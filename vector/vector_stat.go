package vector

func invokeGroupFunction(
	v *vector,
	checkFn func(*vector) bool,
	actionFn func(*vector) Payload,
	columnPostfix string,
) Vector {
	if v.IsGrouped() {
		vectors := v.GroupVectors()
		outValues := make([]Vector, len(vectors))
		for i := 0; i < len(vectors); i++ {
			outValues[i] = invokeGroupFunction(vectors[i].(*vector), checkFn, actionFn, columnPostfix)
		}

		return Combine(outValues...).SetName(v.Name() + columnPostfix)
	}

	vec := NA(1)
	if checkFn(v) {
		vec = New(actionFn(v), v.Options()...)
	}
	vec.SetName(v.Name() + "_sum")

	return vec
}

func invokeFunction(
	v *vector,
	checkFn func(*vector) bool,
	actionFn func(*vector) Payload,
	columnPostfix string,
) Vector {
	vec := NA(v.length)
	if checkFn(v) {
		vec = New(actionFn(v), v.Options()...)
	}
	vec.SetName(v.Name() + columnPostfix)

	return vec

}

type Statistics interface {
	Sum() Vector
	Max() Vector
	Min() Vector
	Mean() Vector
	Median() Vector
	Prod() Vector
	CumSum() Vector
	CumProd() Vector
	CumMax() Vector
	CumMin() Vector
}

// Summer has to be implemented by the payload to be able to calculate the sum of it.
type Summer interface {
	Sum() Payload
}

// Sum returns the sum of the vector.
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
		"_sum",
	)
}

// Proder has to be implemented by the payload to be able to calculate the product of it.
type Proder interface {
	Prod() Payload
}

// Prod returns the product of the vector.
func (v *vector) Prod() Vector {
	return invokeGroupFunction(
		v,
		func(v *vector) bool {
			_, ok := v.payload.(Proder)
			return ok
		},
		func(v *vector) Payload {
			return v.payload.(Proder).Prod()
		},
		"_prod",
	)
}

// Maxxer has to be implemented by the payload to be able to calculate the maximum of it.
type Maxxer interface {
	Max() Payload
}

// Max returns the maximum of the vector.
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
		"_max",
	)
}

// Minner has to be implemented by the payload to be able to calculate the minimum of it.
type Minner interface {
	Min() Payload
}

// Min returns the minimum of the vector.
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
		"_min",
	)
}

// Meaner has to be implemented by the payload to be able to calculate the mean of it.
type Meaner interface {
	Mean() Payload
}

// Mean returns the mean of the vector.
func (v *vector) Mean() Vector {
	return invokeGroupFunction(
		v,
		func(v *vector) bool {
			_, ok := v.payload.(Meaner)
			return ok
		},
		func(v *vector) Payload {
			return v.payload.(Meaner).Mean()
		},
		"_mean",
	)
}

// Medianer has to be implemented by the payload to be able to calculate the median of it.
type Medianer interface {
	Median() Payload
}

// Median returns the median of the vector.
func (v *vector) Median() Vector {
	return invokeGroupFunction(
		v,
		func(v *vector) bool {
			_, ok := v.payload.(Medianer)
			return ok
		},
		func(v *vector) Payload {
			return v.payload.(Medianer).Median()
		},
		"_median",
	)
}

// CumSummer has to be implemented by the payload to be able to calculate the cumulative sum of it.
type CumSummer interface {
	CumSum() Payload
}

// CumSum returns the cumulative sum of the vector.
// Non-groupable for now.
func (v *vector) CumSum() Vector {
	return invokeFunction(
		v,
		func(v *vector) bool {
			_, ok := v.payload.(CumSummer)
			return ok
		},
		func(v *vector) Payload {
			return v.payload.(CumSummer).CumSum()
		},
		"_cumsum",
	)
}

// CumProder has to be implemented by the payload to be able to calculate the cumulative product of it.
type CumProder interface {
	CumProd() Payload
}

// CumProd returns the cumulative product of the vector.
// Non-groupable for now.
func (v *vector) CumProd() Vector {
	return invokeFunction(
		v,
		func(v *vector) bool {
			_, ok := v.payload.(CumProder)
			return ok
		},
		func(v *vector) Payload {
			return v.payload.(CumProder).CumProd()
		},
		"_cumprod",
	)
}

// CumMaxer has to be implemented by the payload to be able to calculate the cumulative maximum of it.
type CumMaxer interface {
	CumMax() Payload
}

// CumMax returns the cumulative maximum of the vector.
// Non-groupable for now.
func (v *vector) CumMax() Vector {
	return invokeFunction(
		v,
		func(v *vector) bool {
			_, ok := v.payload.(CumMaxer)
			return ok
		},
		func(v *vector) Payload {
			return v.payload.(CumMaxer).CumMax()
		},
		"_cummax",
	)
}

// CumMinner has to be implemented by the payload to be able to calculate the cumulative minimum of it.
type CumMinner interface {
	CumMin() Payload
}

// CumMin returns the cumulative minimum of the vector.
// Non-groupable for now.
func (v *vector) CumMin() Vector {
	return invokeFunction(
		v,
		func(v *vector) bool {
			_, ok := v.payload.(CumMinner)
			return ok
		},
		func(v *vector) Payload {
			return v.payload.(CumMinner).CumMin()
		},
		"_cummin",
	)
}
