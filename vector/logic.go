package vector

func And(booleans ...[]bool) []bool {
	if len(booleans) == 0 {
		return []bool{}
	}

	if len(booleans) == 1 {
		return booleans[0]
	}

	src := booleans[0]
	srcLen := len(src)
	for i := 1; i < len(booleans); i++ {
		cmp := fitCmpToSrc(src, booleans[i])

		for j := 0; j < srcLen; j++ {
			src[j] = src[j] && cmp[j]
		}
	}

	return src
}

func Or(booleans ...[]bool) []bool {
	if len(booleans) == 0 {
		return []bool{}
	}

	if len(booleans) == 1 {
		return booleans[0]
	}

	src := booleans[0]
	srcLen := len(src)
	for i := 1; i < len(booleans); i++ {
		cmp := fitCmpToSrc(src, booleans[i])

		for j := 0; j < srcLen; j++ {
			src[j] = src[j] || cmp[j]
		}
	}

	return src
}

func Xor(booleans ...[]bool) []bool {
	if len(booleans) == 0 {
		return []bool{}
	}

	if len(booleans) == 1 {
		return booleans[0]
	}

	src := booleans[0]
	srcLen := len(src)
	for i := 1; i < len(booleans); i++ {
		cmp := fitCmpToSrc(src, booleans[i])

		for j := 0; j < srcLen; j++ {
			src[j] = src[j] != cmp[j]
		}
	}

	return src
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

func fitCmpToSrc(src, cmp []bool) []bool {
	srcLen := len(src)
	cmpLen := len(cmp)

	if srcLen != cmpLen {
		bvec := Boolean(cmp, nil)
		booleans, _ := bvec.Adjust(srcLen).Booleans()
		return booleans
	}

	return cmp
}
