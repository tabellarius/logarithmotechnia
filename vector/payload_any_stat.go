package vector

func (p *anyPayload) Max() Payload {
	if p.length == 0 || p.fn.Eq == nil || p.fn.Lt == nil {
		return AnyPayload([]any{nil}, []bool{true}, p.Options()...)
	}

	max := p.data[0]
	for i := 1; i < p.length; i++ {
		if p.NA[i] {
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
		if p.NA[i] {
			return AnyPayload([]any{nil}, []bool{true}, p.Options()...)
		}

		if p.fn.Lt(p.data[i], min) {
			min = p.data[i]
		}
	}

	return AnyPayload([]any{min}, []bool{false}, p.Options()...)

}

// CumMax returns the cumulative maximum of the payload. The payload must have Eq and Lt callbacks set.
func (p *anyPayload) CumMax() Payload {
	if p.length == 0 || p.fn.Eq == nil || p.fn.Lt == nil {
		return AnyPayload(make([]any, p.length), trueBooleanArr(p.length))
	}

	data := make([]any, p.length)
	na := make([]bool, p.length)

	max := p.data[0]
	isNA := false
	for i := 0; i < p.length; i++ {
		if isNA {
			na[i] = true
			continue
		}

		if p.NA[i] {
			na[i] = true
			isNA = true
			continue
		}

		if p.fn.Lt(max, p.data[i]) {
			max = p.data[i]
		}

		data[i] = max
	}

	return AnyPayload(data, na, p.Options()...)

}

// CumMin returns the cumulative minimum of the payload. The payload must have Eq and Lt callbacks set.
func (p *anyPayload) CumMin() Payload {
	if p.length == 0 || p.fn.Eq == nil || p.fn.Lt == nil {
		return AnyPayload(make([]any, p.length), trueBooleanArr(p.length))
	}

	data := make([]any, p.length)
	na := make([]bool, p.length)

	min := p.data[0]
	isNA := false
	for i := 0; i < p.length; i++ {
		if isNA {
			na[i] = true
			continue
		}

		if p.NA[i] {
			na[i] = true
			isNA = true
			continue
		}

		if p.fn.Lt(p.data[i], min) {
			min = p.data[i]
		}

		data[i] = min
	}

	return AnyPayload(data, na, p.Options()...)

}
