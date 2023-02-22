package vector

func (p *anyPayload) Max() Payload {
	if p.length == 0 || p.fn.Eq == nil || p.fn.Lt == nil {
		return AnyPayload([]any{nil}, []bool{true}, p.Options()...)
	}

	max := p.data[0]
	for i := 1; i < p.length; i++ {
		if p.na[i] {
			return AnyPayload([]any{nil}, []bool{true}, p.Options()...)
		}

		if p.fn.Lt(max, p.data[i]) {
			max = p.data[i]
		}
	}

	return AnyPayload([]any{max}, []bool{false}, p.Options()...)

}

func (p *anyPayload) Min() Payload {
	if p.length == 0 || p.fn.Eq == nil || p.fn.Lt == nil {
		return AnyPayload([]any{nil}, []bool{true}, p.Options()...)
	}

	min := p.data[0]
	for i := 1; i < p.length; i++ {
		if p.na[i] {
			return AnyPayload([]any{nil}, []bool{true}, p.Options()...)
		}

		if p.fn.Lt(p.data[i], min) {
			min = p.data[i]
		}
	}

	return AnyPayload([]any{min}, []bool{false}, p.Options()...)

}
