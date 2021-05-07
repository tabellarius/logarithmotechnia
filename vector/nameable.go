package vector

type Nameable interface {
	Names() []string
	NamesMap() map[string]int
	InvertedNamesMap() map[int]string
	Name(idx int) string
	Index(name string) int
	NamesForIndices(indices []int) map[string]int
	SetName(idx int, name string) Vector
	SetNames(names []string) Vector
	SetNamesMap(names map[string]int) Vector
	HasName(name string) bool
	HasNameFor(idx int) bool
	ByNames(names []string) Vector
}

type DefNameable struct {
	vec   Vector
	names map[string]int
	marked bool
}

func (n *DefNameable) HasName(name string) bool {
	_, exists := n.names[name]
	return exists
}

func (n *DefNameable) HasNameFor(idx int) bool {
	if idx >= 1 && idx <= n.vec.Length() {
		for _, index := range n.names {
			if idx == index {
				return true
			}
		}
	}

	return false
}

func (n *DefNameable) Clone() Nameable {
	n.marked = true

	return &DefNameable{
		vec:    n.vec,
		names:  n.names,
		marked: true,
	}
}

func (n *DefNameable) Refresh() {
	names := map[string]int{}

	for name, idx := range n.names {
		names[name] = idx
	}

	n.names = names
}

func (n *DefNameable) Names() []string {
	names := make([]string, n.vec.Length())

	for name, idx := range n.names {
		names[idx-1] = name
	}

	return names
}

func (n *DefNameable) NamesMap() map[string]int {
	names := make(map[string]int)

	for name, idx := range n.names {
		names[name] = idx
	}

	return names
}

func (n *DefNameable) InvertedNamesMap() map[int]string {
	inverted := make(map[int]string)

	for name, idx := range n.names {
		inverted[idx] = name
	}

	return inverted
}

func (n *DefNameable) Name(index int) string {
	if index >= 1 && index <= n.vec.Length() {
		for name, idx := range n.names {
			if index == idx {
				return name
			}
		}
	}

	return ""
}

func (n *DefNameable) Index(name string) int {
	idx, ok := n.names[name]
	if ok {
		return idx
	}
	return 0
}

func (n *DefNameable) NamesForIndices(indices []int) map[string]int {
	inverted := n.InvertedNamesMap()
	names := map[string]int{}

	for _, idx := indices {

	}

}

func (n *DefNameable) SetName(idx int, name string) Vector {
	if n.marked {
		n.Refresh()
	}

	if name != "" && idx >= 1 && idx <= n.vec.Length() {
		n.names[name] = idx
	}

	return n.vec
}

func (n *DefNameable) SetNames(names []string) Vector {
	length := len(names)
	if length > n.vec.Length() {
		length = n.vec.Length()
		if rep, ok := n.vec.(Reporter); ok {
			rep.Report().AddWarning("SetNames(): names []string is longer than vectors length")
		}
	}

	for i := 1; i <= length; i++ {
		n.SetName(i, names[i-1])
	}

	return n.vec
}

func (n *DefNameable) SetNamesMap(names map[string]int) Vector {
	n.names = make(map[string]int)
	for name, idx := range names {
		n.SetName(idx, name)
	}

	return n.vec
}

func (n *DefNameable) ByNames(names []string) Vector {
	indices := make([]int, 0)

	for _, name := range names {
		if idx, ok := n.names[name]; ok {
			indices = append(indices, idx)
		}
	}

	return n.vec.ByIndices(indices)
}

func newDefNameable(vec Vector) DefNameable {
	nameable := DefNameable{
		vec:   vec,
		names: map[string]int{},
		marked: false,
	}

	return nameable
}
