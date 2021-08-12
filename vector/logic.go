package vector

func And(booleans ...[]bool) []bool {
	if len(booleans) == 0 {
		return []bool{}
	}

	return nil
}

func Or(booleans ...[]bool) []bool {
	if len(booleans) == 0 {
		return []bool{}
	}

	return nil
}

func Xor(booleans ...[]bool) []bool {
	if len(booleans) == 0 {
		return []bool{}
	}

	return nil
}

func Not(in []bool) []bool {
	out := make([]bool, len(in))

	for i, v := range in {
		if v {
			out[i] = false
		} else {
			out[i] = true
		}
	}

	return out
}
