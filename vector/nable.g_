package vector

type NAble interface {
	NA() []bool
	IsNA(idx int) bool
	SetNA(na []bool) Vector
	HasNA() bool
	OnlyNA() []int
	WithoutNA() []int
}

type DefNAble struct {
	vec    Vector
	na     []bool
	marked bool
}

func (n *DefNAble) Clone() NAble {
	n.marked = true

	return &DefNAble{
		vec:    n.vec,
		na:     n.na,
		marked: true,
	}
}

func (n *DefNAble) Refresh() {
	if n.vec.Length() == 0 {
		return
	}

	na := make([]bool, len(n.na))
	copy(na, n.na)

	n.na = na
	n.marked = false
}

func (n *DefNAble) NA() []bool {
	length := len(n.na) - 1
	na := make([]bool, length)
	copy(na, n.na[1:])
	return na
}

func (n *DefNAble) IsNA(idx int) bool {
	if idx >= 1 && idx < len(n.na) {
		return n.na[idx]
	}

	return false
}

func (n *DefNAble) SetNA(na []bool) Vector {
	if len(na) != n.vec.Length() {
		if rep, ok := n.vec.(Reporter); ok {
			rep.Report().AddWarning("SetNA(na []bool): length of na is not equal to vector's length")
		}
		return n.vec
	}

	n.marked = false
	n.na = make([]bool, n.vec.Length()+1)
	copy(n.na[1:], na)
	return n.vec
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

func newDefaultNAble(vec Vector) DefNAble {
	nable := DefNAble{
		vec:    vec,
		na:     make([]bool, vec.Length()+1),
		marked: false,
	}

	return nable
}
