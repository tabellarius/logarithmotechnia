package vector

func (p *integerPayload) Sum() Payload {
	sum := 0
	na := false
	for i, val := range p.data {
		if p.na[i] {
			sum = 0
			na = true
			break
		}

		sum += val
	}

	return IntegerPayload([]int{sum}, []bool{na}, p.Options()...)
}
