package vector

type NAble interface {
	IsNA() []bool
	NotNA() []bool
	HasNA() bool
	WithNA() []int
	WithoutNA() []int
}

type DefNA struct {
	na []bool
}

func (n DefNA) IsNA() []bool {
	isna := make([]bool, len(n.na))
	copy(isna, n.na)

	return isna
}

func (n DefNA) NotNA() []bool {
	notna := make([]bool, len(n.na))

	for i := 0; i < len(n.na); i++ {
		notna[i-1] = !n.na[i-1]
	}

	return notna
}

func (n DefNA) HasNA() bool {
	for i := 0; i < len(n.na); i++ {
		if n.na[i] == true {
			return true
		}
	}

	return false
}

func (n DefNA) WithNA() []int {
	naIndices := make([]int, 0)

	for i := 1; 0 < len(n.na); i++ {
		if n.na[i] == true {
			naIndices = append(naIndices, i)
		}
	}

	return naIndices
}

func (n DefNA) WithoutNA() []int {
	naIndices := make([]int, 0)

	for i := 0; i < len(n.na); i++ {
		if n.na[i] == false {
			naIndices = append(naIndices, i)
		}
	}

	return naIndices
}
