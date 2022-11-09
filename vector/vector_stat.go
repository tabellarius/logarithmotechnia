package vector

type Statistics interface {
	Sum() Vector
	Max() Vector
	Min() Vector
}

type Summer interface {
	Sum() Payload
}

func (v *vector) Sum() Vector {
	if v.IsGrouped() {
		vectors := v.GroupVectors()
		outValues := make([]Vector, len(vectors))
		for i := 0; i < len(vectors); i++ {
			outValues[i] = vectors[i].Sum()
		}

		return Combine(outValues...)
	}

	if summer, ok := v.payload.(Summer); ok {
		return New(summer.Sum(), v.Options()...)
	}

	return NA(1)
}

type Maxxer interface {
	Max() Payload
}

func (v *vector) Max() Vector {
	if v.IsGrouped() {
		vectors := v.GroupVectors()
		outValues := make([]Vector, len(vectors))
		for i := 0; i < len(vectors); i++ {
			outValues[i] = vectors[i].Max()
		}

		return Combine(outValues...)
	}

	if summer, ok := v.payload.(Maxxer); ok {
		return New(summer.Max(), v.Options()...)
	}

	return NA(1)
}

type Minner interface {
	Min() Payload
}

func (v *vector) Min() Vector {
	if v.IsGrouped() {
		vectors := v.GroupVectors()
		outValues := make([]Vector, len(vectors))
		for i := 0; i < len(vectors); i++ {
			outValues[i] = vectors[i].Min()
		}

		return Combine(outValues...)
	}

	if summer, ok := v.payload.(Minner); ok {
		return New(summer.Min(), v.Options()...)
	}

	return NA(1)
}
