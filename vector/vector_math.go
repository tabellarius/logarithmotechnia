package vector

type SummerV interface {
	Sum() Vector
}

type SummerP interface {
	Sum() Payload
}

func (v *vector) Sum() Vector {
	if v.IsGrouped() {
		vec := v.groups[0].Sum()
		for i := 1; i < len(v.groups); i++ {
			vec = vec.Append(v.groups[i].Sum())
		}

		return vec
	}

	if summer, ok := v.payload.(SummerP); ok {
		return New(summer.Sum(), v.Options()...)
	}

	return NA(1)
}
