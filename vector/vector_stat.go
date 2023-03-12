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

		return Combine(outValues...).SetName(v.Name() + "_sum")
	}

	vec := NA(1)
	if summer, ok := v.payload.(Summer); ok {
		vec = New(summer.Sum(), v.Options()...)
	}
	vec.SetName(v.Name() + "_sum")

	return vec
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

		return Combine(outValues...).SetName(v.Name() + "_max")
	}

	vec := NA(1)
	if maxxer, ok := v.payload.(Maxxer); ok {
		return New(maxxer.Max(), v.Options()...)
	}
	vec.SetName(v.Name() + "_max")

	return vec
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

		return Combine(outValues...).SetName(v.Name() + "_min")
	}

	vec := NA(1)
	if minner, ok := v.payload.(Minner); ok {
		vec = New(minner.Min(), v.Options()...)
	}
	vec.SetName(v.Name() + "_min")

	return vec
}
