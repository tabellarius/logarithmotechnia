package embed

// DefNAble is an easy to embed implementation of NAble interface.
type DefNAble struct {
	NA []bool
}

func (n *DefNAble) IsNA() []bool {
	isna := make([]bool, len(n.NA))
	copy(isna, n.NA)

	return isna
}

func (n *DefNAble) NotNA() []bool {
	notna := make([]bool, len(n.NA))

	for i := 0; i < len(n.NA); i++ {
		notna[i] = !n.NA[i]
	}

	return notna
}

func (n *DefNAble) HasNA() bool {
	for i := 0; i < len(n.NA); i++ {
		if n.NA[i] == true {
			return true
		}
	}

	return false
}

func (n *DefNAble) WithNA() []int {
	naIndices := make([]int, 0)

	for i := 0; i < len(n.NA); i++ {
		if n.NA[i] == true {
			naIndices = append(naIndices, i+1)
		}
	}

	return naIndices
}

func (n *DefNAble) WithoutNA() []int {
	naIndices := make([]int, 0)

	for i := 0; i < len(n.NA); i++ {
		if n.NA[i] == false {
			naIndices = append(naIndices, i+1)
		}
	}

	return naIndices
}
