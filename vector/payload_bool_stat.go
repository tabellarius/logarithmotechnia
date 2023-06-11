package vector

func (p *booleanPayload) Sum() Payload {
	sum := 0
	na := false
	for i, val := range p.data {
		if p.NA[i] {
			sum = 0
			na = true
			break
		}

		if val {
			sum++
		}
	}

	return IntegerPayload([]int{sum}, []bool{na}, p.Options()...)
}
