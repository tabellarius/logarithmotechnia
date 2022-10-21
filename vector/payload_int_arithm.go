package vector

func (p *integerPayload) Add(p2 Payload) Payload {
	integers := make([]int, p.length)
	na := make([]bool, p.length)

	if p.length != p2.Len() {
		p2 = p2.Adjust(p.length)
	}

	var addIntegers []int
	var addNA []bool
	if pType, ok := p2.(*integerPayload); ok {
		addIntegers = pType.data
		addNA = pType.na
	} else if intable, ok := p2.(Intable); ok {
		addIntegers, addNA = intable.Integers()
	} else {
		return NAPayload(p.length)
	}

	for i := 0; i < p.length; i++ {
		if p.na[i] || addNA[i] {
			na[i] = true
		} else {
			integers[i] = p.data[i] + addIntegers[i]
		}
	}

	return IntegerPayload(integers, na, p.Options()...)
}

func (p *integerPayload) Sub(p2 Payload) Payload {
	integers := make([]int, p.length)
	na := make([]bool, p.length)

	if p.length != p2.Len() {
		p2 = p2.Adjust(p.length)
	}

	var subIntegers []int
	var subNA []bool
	if pType, ok := p2.(*integerPayload); ok {
		subIntegers = pType.data
		subNA = pType.na
	} else if intable, ok := p2.(Intable); ok {
		subIntegers, subNA = intable.Integers()
	} else {
		return NAPayload(p.length)
	}

	for i := 0; i < p.length; i++ {
		if p.na[i] || subNA[i] {
			na[i] = true
		} else {
			integers[i] = p.data[i] - subIntegers[i]
		}
	}

	return IntegerPayload(integers, na, p.Options()...)
}

func (p *integerPayload) Mul(p2 Payload) Payload {
	integers := make([]int, p.length)
	na := make([]bool, p.length)

	if p.length != p2.Len() {
		p2 = p2.Adjust(p.length)
	}

	var mulIntegers []int
	var mulNA []bool
	if pType, ok := p2.(*integerPayload); ok {
		mulIntegers = pType.data
		mulNA = pType.na
	} else if intable, ok := p2.(Intable); ok {
		mulIntegers, mulNA = intable.Integers()
	} else {
		return NAPayload(p.length)
	}

	for i := 0; i < p.length; i++ {
		if p.na[i] || mulNA[i] {
			na[i] = true
		} else {
			integers[i] = p.data[i] * mulIntegers[i]
		}
	}

	return IntegerPayload(integers, na, p.Options()...)
}

func (p *integerPayload) Div(p2 Payload) Payload {
	integers := make([]int, p.length)
	na := make([]bool, p.length)

	if p.length != p2.Len() {
		p2 = p2.Adjust(p.length)
	}

	var divIntegers []int
	var divNA []bool
	if pType, ok := p2.(*integerPayload); ok {
		divIntegers = pType.data
		divNA = pType.na
	} else if intable, ok := p2.(Intable); ok {
		divIntegers, divNA = intable.Integers()
	} else {
		return NAPayload(p.length)
	}

	for i := 0; i < p.length; i++ {
		if p.na[i] || divNA[i] || divIntegers[i] == 0 {
			na[i] = true
		} else {
			integers[i] = p.data[i] / divIntegers[i]
		}
	}

	return IntegerPayload(integers, na, p.Options()...)
}
