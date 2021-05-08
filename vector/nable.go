package vector

type NAble interface {
	IsNA() []bool
	NotNA() []bool
	HasNA() bool
	WithNA() []int
	WithoutNA() []int
}

type DefNA struct {
	length int
	na     []bool
}

func (n DefNA) IsNA() []bool {
	isna := make([]bool, n.length)
	copy(isna, n.na[1:])

	return isna
}

func (n DefNA) NotNA() []bool {
	notna := make([]bool, n.length)

	for i := 1; i < n.length; i++ {
		notna[i-1] = !n.na[i]
	}

	return notna
}

func (n DefNA) HasNA() bool {
	if n.na[0] {
		return true
	}

	length := len(n.na) - 1
	for i := 1; i <= length; i++ {
		if n.na[i] == true {
			n.na[0] = true
			return true
		}
	}

	return false
}

func (n DefNA) WithNA() []int {
	naIndices := make([]int, 0)

	for i := 1; i <= n.length; i++ {
		if n.na[i] == true {
			naIndices = append(naIndices, i)
		}
	}

	return naIndices
}

func (n DefNA) WithoutNA() []int {
	naIndices := make([]int, 0)

	for i := 1; i <= n.length; i++ {
		if n.na[i] == false {
			naIndices = append(naIndices, i)
		}
	}

	return naIndices
}
