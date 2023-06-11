package embed

// NAble is an easy to embed implementation of NAble interface.
type NAble struct {
	NA []bool
}

func (n *NAble) IsNA() []bool {
	isna := make([]bool, len(n.NA))
	copy(isna, n.NA)

	return isna
}

func (n *NAble) NotNA() []bool {
	notna := make([]bool, len(n.NA))

	for i := 0; i < len(n.NA); i++ {
		notna[i] = !n.NA[i]
	}

	return notna
}

func (n *NAble) HasNA() bool {
	for i := 0; i < len(n.NA); i++ {
		if n.NA[i] == true {
			return true
		}
	}

	return false
}

func (n *NAble) WithNA() []int {
	naIndices := make([]int, 0)

	for i := 0; i < len(n.NA); i++ {
		if n.NA[i] == true {
			naIndices = append(naIndices, i+1)
		}
	}

	return naIndices
}

func (n *NAble) WithoutNA() []int {
	naIndices := make([]int, 0)

	for i := 0; i < len(n.NA); i++ {
		if n.NA[i] == false {
			naIndices = append(naIndices, i+1)
		}
	}

	return naIndices
}
