package vector

type NAble interface {
	NA() []bool
	IsNA(idx int) bool
	HasNA() bool
	OnlyNA() []int
	WithoutNA() []int
}

type DefaultNAble struct {
	vec Vector
	na  []bool
}

func (n *DefaultNAble) Refresh() {
	if len(n.na) == 0 {
		return
	}

	na := make([]bool, len(n.na))
	copy(na, n.na)

	n.na = na
}

func (n *DefaultNAble) NA() []bool {
	length := len(n.na)
	na := make([]bool, length)
	copy(na, n.na)
	return na
}

func (n *DefaultNAble) IsNA(idx int) bool {
	if idx >= 1 && idx <= len(n.na) {
		return n.na[idx]
	}

	return false
}

func (n *DefaultNAble) HasNA() bool {
	for i := 1; i <= len(n.na); i++ {
		if n.na[i] == true {
			return true
		}
	}

	return false
}

func (n *DefaultNAble) OnlyNA() []int {
	naIndices := make([]int, 0)

	for i := 1; i <= len(n.na); i++ {
		if n.na[i] == true {
			naIndices = append(naIndices, i)
		}
	}

	return naIndices
}

func (n *DefaultNAble) WithoutNA() []int {
	naIndices := make([]int, 0)

	for i := 1; i <= len(n.na); i++ {
		if n.na[i] == false {
			naIndices = append(naIndices, i)
		}
	}

	return naIndices
}
