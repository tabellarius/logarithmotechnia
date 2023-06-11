package vector

func (p *complexPayload) Add(p2 Payload) Payload {
	numbers := make([]complex128, p.length)
	na := make([]bool, p.length)

	if p.length != p2.Len() {
		p2 = p2.Adjust(p.length)
	}

	var opNumbers []complex128
	var opNA []bool
	if pType, ok := p2.(*complexPayload); ok {
		opNumbers = pType.data
		opNA = pType.NA
	} else if complexable, ok := p2.(Complexable); ok {
		opNumbers, opNA = complexable.Complexes()
	} else {
		return NAPayload(p.length)
	}

	for i := 0; i < p.length; i++ {
		if p.NA[i] || opNA[i] {
			na[i] = true
		} else {
			numbers[i] = p.data[i] + opNumbers[i]
		}
	}

	return ComplexPayload(numbers, na, p.Options()...)
}

func (p *complexPayload) Sub(p2 Payload) Payload {
	numbers := make([]complex128, p.length)
	na := make([]bool, p.length)

	if p.length != p2.Len() {
		p2 = p2.Adjust(p.length)
	}

	var opNumbers []complex128
	var subNA []bool
	if pType, ok := p2.(*complexPayload); ok {
		opNumbers = pType.data
		subNA = pType.NA
	} else if complexable, ok := p2.(Complexable); ok {
		opNumbers, subNA = complexable.Complexes()
	} else {
		return NAPayload(p.length)
	}

	for i := 0; i < p.length; i++ {
		if p.NA[i] || subNA[i] {
			na[i] = true
		} else {
			numbers[i] = p.data[i] - opNumbers[i]
		}
	}

	return ComplexPayload(numbers, na, p.Options()...)
}

func (p *complexPayload) Mul(p2 Payload) Payload {
	numbers := make([]complex128, p.length)
	na := make([]bool, p.length)

	if p.length != p2.Len() {
		p2 = p2.Adjust(p.length)
	}

	var opNumbers []complex128
	var opNA []bool
	if pType, ok := p2.(*complexPayload); ok {
		opNumbers = pType.data
		opNA = pType.NA
	} else if complexable, ok := p2.(Complexable); ok {
		opNumbers, opNA = complexable.Complexes()
	} else {
		return NAPayload(p.length)
	}

	for i := 0; i < p.length; i++ {
		if p.NA[i] || opNA[i] {
			na[i] = true
		} else {
			numbers[i] = p.data[i] * opNumbers[i]
		}
	}

	return ComplexPayload(numbers, na, p.Options()...)
}

func (p *complexPayload) Div(p2 Payload) Payload {
	numbers := make([]complex128, p.length)
	na := make([]bool, p.length)

	if p.length != p2.Len() {
		p2 = p2.Adjust(p.length)
	}

	var opNumbers []complex128
	var opNA []bool
	if pType, ok := p2.(*complexPayload); ok {
		opNumbers = pType.data
		opNA = pType.NA
	} else if complexable, ok := p2.(Complexable); ok {
		opNumbers, opNA = complexable.Complexes()
	} else {
		return NAPayload(p.length)
	}

	for i := 0; i < p.length; i++ {
		if p.NA[i] || opNA[i] || opNumbers[i] == 0 {
			na[i] = true
		} else {
			numbers[i] = p.data[i] / opNumbers[i]
		}
	}

	return ComplexPayload(numbers, na, p.Options()...)
}
