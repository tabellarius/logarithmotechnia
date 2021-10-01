package vector

func (p *complexPayload) Sum() Payload {
	sum := 0 + 0i
	na := false
	for i, val := range p.data {
		if p.na[i] {
			sum = 0
			na = true
			break
		}

		sum += val
	}

	return ComplexPayload([]complex128{sum}, []bool{na}, p.Options()...)
}
