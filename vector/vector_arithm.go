package vector

type Arithmetics interface {
	Add(...Vector) Vector
	Sub(...Vector) Vector
	Mul(...Vector) Vector
	Div(...Vector) Vector
}

type Adder interface {
	Add(Payload) Payload
}

type Subber interface {
	Sub(Payload) Payload
}

type Multiplier interface {
	Mul(Payload) Payload
}

type Divider interface {
	Div(Payload) Payload
}

func (v *vector) Add(vectors ...Vector) Vector {
	if len(vectors) == 0 {
		return v
	}

	adder, ok := v.payload.(Adder)
	if !ok {
		return NA(v.length)
	}

	var payload Payload
	for _, vec := range vectors {
		payload = adder.Add(vec.Payload())
		if adderable, ok := payload.(Adder); ok {
			adder = adderable
		} else {
			return NA(v.length)
		}
	}

	return New(payload, v.Options()...)
}

func (v *vector) Sub(vectors ...Vector) Vector {
	if len(vectors) == 0 {
		return v
	}

	subber, ok := v.payload.(Subber)
	if !ok {
		return NA(v.length)
	}

	var payload Payload
	for _, vec := range vectors {
		payload = subber.Sub(vec.Payload())
		if subberable, ok := payload.(Subber); ok {
			subber = subberable
		} else {
			return NA(v.length)
		}
	}

	return New(payload, v.Options()...)
}

func (v *vector) Mul(vectors ...Vector) Vector {
	if len(vectors) == 0 {
		return v
	}

	multiplier, ok := v.payload.(Multiplier)
	if !ok {
		return NA(v.length)
	}

	var payload Payload
	for _, vec := range vectors {
		payload = multiplier.Mul(vec.Payload())
		if multiplierable, ok := payload.(Multiplier); ok {
			multiplier = multiplierable
		} else {
			return NA(v.length)
		}
	}

	return New(payload, v.Options()...)
}

func (v *vector) Div(vectors ...Vector) Vector {
	if len(vectors) == 0 {
		return v
	}

	divider, ok := v.payload.(Divider)
	if !ok {
		return NA(v.length)
	}

	var payload Payload
	for _, vec := range vectors {
		payload = divider.Div(vec.Payload())
		if dividable, ok := payload.(Divider); ok {
			divider = dividable
		} else {
			return NA(v.length)
		}
	}

	return New(payload, v.Options()...)
}
