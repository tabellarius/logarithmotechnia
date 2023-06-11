package vector

func (p *floatPayload) Add(p2 Payload) Payload {
	numbers := make([]float64, p.length)
	na := make([]bool, p.length)

	if p.length != p2.Len() {
		p2 = p2.Adjust(p.length)
	}

	var opNumbers []float64
	var opNA []bool
	if pType, ok := p2.(*floatPayload); ok {
		opNumbers = pType.data
		opNA = pType.NA
	} else if floatable, ok := p2.(Floatable); ok {
		opNumbers, opNA = floatable.Floats()
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

	return FloatPayload(numbers, na, p.Options()...)
}

func (p *floatPayload) Sub(p2 Payload) Payload {
	numbers := make([]float64, p.length)
	na := make([]bool, p.length)

	if p.length != p2.Len() {
		p2 = p2.Adjust(p.length)
	}

	var opNumbers []float64
	var subNA []bool
	if pType, ok := p2.(*floatPayload); ok {
		opNumbers = pType.data
		subNA = pType.NA
	} else if floatable, ok := p2.(Floatable); ok {
		opNumbers, subNA = floatable.Floats()
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

	return FloatPayload(numbers, na, p.Options()...)
}

func (p *floatPayload) Mul(p2 Payload) Payload {
	numbers := make([]float64, p.length)
	na := make([]bool, p.length)

	if p.length != p2.Len() {
		p2 = p2.Adjust(p.length)
	}

	var opNumbers []float64
	var opNA []bool
	if pType, ok := p2.(*floatPayload); ok {
		opNumbers = pType.data
		opNA = pType.NA
	} else if floatable, ok := p2.(Floatable); ok {
		opNumbers, opNA = floatable.Floats()
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

	return FloatPayload(numbers, na, p.Options()...)
}

func (p *floatPayload) Div(p2 Payload) Payload {
	numbers := make([]float64, p.length)
	na := make([]bool, p.length)

	if p.length != p2.Len() {
		p2 = p2.Adjust(p.length)
	}

	var opNumbers []float64
	var opNA []bool
	if pType, ok := p2.(*floatPayload); ok {
		opNumbers = pType.data
		opNA = pType.NA
	} else if floatable, ok := p2.(Floatable); ok {
		opNumbers, opNA = floatable.Floats()
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

	return FloatPayload(numbers, na, p.Options()...)
}
