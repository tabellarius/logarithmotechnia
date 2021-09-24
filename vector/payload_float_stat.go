package vector

func (p *floatPayload) Sum() Payload {
	sum := 0.0
	na := false
	for i, val := range p.data {
		if p.na[i] {
			sum = 0
			na = true
			break
		}

		sum += val
	}

	return FloatPayload([]float64{sum}, []bool{na}, p.Options()...)
}
