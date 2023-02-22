package vector

// NAble is an interface a payload has to satisfy in order to support NA-values.
type NAble interface {
	IsNA() []bool
	NotNA() []bool
	HasNA() bool
	WithNA() []int
	WithoutNA() []int
}

// DefNAble is an easy to embed implementation of NAble interface.
type DefNAble struct {
	na []bool
}

func (n *DefNAble) IsNA() []bool {
	isna := make([]bool, len(n.na))
	copy(isna, n.na)

	return isna
}

func (n *DefNAble) NotNA() []bool {
	notna := make([]bool, len(n.na))

	for i := 0; i < len(n.na); i++ {
		notna[i] = !n.na[i]
	}

	return notna
}

func (n *DefNAble) HasNA() bool {
	for i := 0; i < len(n.na); i++ {
		if n.na[i] == true {
			return true
		}
	}

	return false
}

func (n *DefNAble) WithNA() []int {
	naIndices := make([]int, 0)

	for i := 0; i < len(n.na); i++ {
		if n.na[i] == true {
			naIndices = append(naIndices, i+1)
		}
	}

	return naIndices
}

func (n *DefNAble) WithoutNA() []int {
	naIndices := make([]int, 0)

	for i := 0; i < len(n.na); i++ {
		if n.na[i] == false {
			naIndices = append(naIndices, i+1)
		}
	}

	return naIndices
}
