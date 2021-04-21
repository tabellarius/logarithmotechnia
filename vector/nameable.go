package vector

type Nameable interface {
	Names() []string
	NamesMap() map[string]int
	InvertedNamesMap() map[int]string
	Name(idx int) string
	Index(name string) int
	SetName(name string, idx int) Vector
	SetNames(names []string) Vector
	SetNamesMap(names map[string]int) Vector
	ByNames(names []string) Vector
}

type DefaultNameable struct {
	vec   Vector
	names map[string]int
}

func (n *DefaultNameable) Names() []string {
	names := make([]string, n.vec.Length())

	for name, idx := range n.names {
		names[idx] = name
	}

	return names
}

func (n *DefaultNameable) NamesMap() map[string]int {
	names := make(map[string]int)

	for name, idx := range n.names {
		names[name] = idx
	}

	return names
}

func (n *DefaultNameable) InvertedNamesMap() map[int]string {
	inverted := make(map[int]string)

	for name, idx := range n.names {
		inverted[idx] = name
	}

	return inverted
}

func (n *DefaultNameable) Name(index int) string {
	if index >= 1 && index <= n.vec.Length() {
		for name, idx := range n.names {
			if index == idx {
				return name
			}
		}
	}

	return ""
}

func (n *DefaultNameable) Index(name string) int {
	idx, ok := n.names[name]
	if ok {
		return idx
	}
	return 0
}

func (n *DefaultNameable) SetName(name string, idx int) Vector {
	if name != "" && idx >= 1 && idx <= n.vec.Length() {
		n.names[name] = idx
	}

	return n.vec
}

func (n *DefaultNameable) SetNames(names []string) Vector {
	length := len(names)
	if length > n.vec.Length() {
		length = n.vec.Length()
	}

	for i := 1; i <= length; i++ {
		n.SetName(names[i], i)
	}

	return n.vec
}

func (n *DefaultNameable) SetNamesMap(names map[string]int) Vector {
	n.names = make(map[string]int)
	for name, idx := range names {
		n.SetName(name, idx)
	}

	return n.vec
}

func (n *DefaultNameable) ByNames(names []string) Vector {
	indices := make([]int, 0)

	for _, name := range names {
		if idx, ok := n.names[name]; ok {
			indices = append(indices, idx)
		}
	}

	return n.vec.ByIndex(indices)
}
