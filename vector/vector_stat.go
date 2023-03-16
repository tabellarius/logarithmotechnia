package vector

func invokeGroupFunction(v *vector, checkFn func(*vector) bool, actionFn func(*vector) Payload) Vector {
	if v.IsGrouped() {
		vectors := v.GroupVectors()
		outValues := make([]Vector, len(vectors))
		for i := 0; i < len(vectors); i++ {
			outValues[i] = invokeGroupFunction(vectors[i].(*vector), checkFn, actionFn)
		}

		return Combine(outValues...).SetName(v.Name() + "_sum")
	}

	vec := NA(1)
	if checkFn(v) {
		vec = New(actionFn(v), v.Options()...)
	}
	vec.SetName(v.Name() + "_sum")

	return vec
}

type Statistics interface {
	Sum() Vector
	Max() Vector
	Min() Vector
}

type Summer interface {
	Sum() Payload
}

func (v *vector) Sum() Vector {
	return invokeGroupFunction(
		v,
		func(v *vector) bool {
			_, ok := v.payload.(Summer)
			return ok
		},
		func(v *vector) Payload {
			return v.payload.(Summer).Sum()
		},
	)
}

type Maxxer interface {
	Max() Payload
}

func (v *vector) Max() Vector {
	return invokeGroupFunction(
		v,
		func(v *vector) bool {
			_, ok := v.payload.(Maxxer)
			return ok
		},
		func(v *vector) Payload {
			return v.payload.(Maxxer).Max()
		},
	)
}

type Minner interface {
	Min() Payload
}

func (v *vector) Min() Vector {
	return invokeGroupFunction(
		v,
		func(v *vector) bool {
			_, ok := v.payload.(Minner)
			return ok
		},
		func(v *vector) Payload {
			return v.payload.(Minner).Min()
		},
	)
}
