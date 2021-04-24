package vector

type NAble interface {
	NA() []bool
	IsNA(idx int) bool
	HasNA() bool
	OnlyNA() []int
	WithoutNA() []int
}

type DefNAble struct {
	vec Vector
	na  []bool
}

func (n *DefNAble) Refresh() {
	if n.vec.Length() == 0 {
		return
	}

	na := make([]bool, len(n.na))
	copy(na, n.na)

	n.na = na
}

func (n *DefNAble) NA() []bool {
	length := len(n.na) - 1
	na := make([]bool, length)
	copy(na, n.na[1:])
	return na
}

func (n *DefNAble) IsNA(idx int) bool {
	if idx >= 1 && idx <= len(n.na) {
		return n.na[idx]
	}

	return false
}

func (n *DefNAble) HasNA() bool {
	for i := 1; i <= n.vec.Length(); i++ {
		if n.na[i] == true {
			return true
		}
	}

	return false
}

func (n *DefNAble) OnlyNA() []int {
	naIndices := make([]int, 0)

	for i := 1; i <= n.vec.Length(); i++ {
		if n.na[i] == true {
			naIndices = append(naIndices, i)
		}
	}

	return naIndices
}

func (n *DefNAble) WithoutNA() []int {
	naIndices := make([]int, 0)

	for i := 1; i <= n.vec.Length(); i++ {
		if n.na[i] == false {
			naIndices = append(naIndices, i)
		}
	}

	return naIndices
}
