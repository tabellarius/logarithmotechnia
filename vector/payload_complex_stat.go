package vector

func (p *complexPayload) Sum() Payload {
	sum, na := genSum(p.data, p.na)

	return ComplexPayload([]complex128{sum}, []bool{na}, p.Options()...)
}

func (p *complexPayload) Mean() Payload {
	var sum, mean complex128
	var na bool

	if p.length != 0 {
		for i := 0; i < p.length; i++ {
			if p.na[i] {
				mean, na = 0, true
				goto outOfLoop
			}

			sum += p.data[i]
		}

		mean = sum / complex(float64(p.length), 0)
	}
outOfLoop:

	return ComplexPayload([]complex128{mean}, []bool{na}, p.Options()...)
}
