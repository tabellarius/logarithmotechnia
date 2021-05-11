package vector

type Nameable interface {
	Names() []string
	NamesMap() map[string]int
	InvertedNamesMap() map[int]string
	Name(idx int) string
	Index(name string) int
	NamesForIndices(indices []int) map[string]int
	SetName(idx int, name string)
	SetNamesMap(names map[string]int)
	SetNames(names []string)
	HasName(name string) bool
	HasNameFor(idx int) bool
}

type DefNameable struct {
	length int
	names  map[string]int
}

func (v *DefNameable) Names() []string {
	names := make([]string, v.length)

	for name, idx := range v.names {
		names[idx-1] = name
	}

	return names
}

func (v *DefNameable) NamesMap() map[string]int {
	names := make(map[string]int)

	for name, idx := range v.names {
		names[name] = idx
	}

	return names
}

func (v *DefNameable) InvertedNamesMap() map[int]string {
	inverted := make(map[int]string)

	for name, idx := range v.names {
		inverted[idx] = name
	}

	return inverted
}

func (v *DefNameable) Name(index int) string {
	if index >= 1 && index <= v.length {
		for name, idx := range v.names {
			if index == idx {
				return name
			}
		}
	}

	return ""
}

func (v *DefNameable) Index(name string) int {
	idx, ok := v.names[name]
	if ok {
		return idx
	}
	return 0
}

func (v *DefNameable) NamesForIndices(indices []int) map[string]int {
	inverted := v.InvertedNamesMap()
	names := map[string]int{}

	for _, idx := range indices {
		if name, ok := inverted[idx]; ok {
			names[name] = idx
		}
	}

	return names
}

func (v *DefNameable) SetName(idx int, name string) {
	if name != "" && idx >= 1 && idx <= v.length {
		v.names[name] = idx
	}
}

func (v *DefNameable) SetNames(names []string) {
	length := len(names)

	/*
		if length != v.length {
			v.report.AddWarning("SetNames(): names []string is not equal to vector's length")
		}
	*/

	if length > v.length {
		length = v.length
	}

	for i := 1; i <= length; i++ {
		v.SetName(i, names[i-1])
	}
}

func (v *DefNameable) SetNamesMap(names map[string]int) {
	v.names = map[string]int{}

	for name, idx := range names {
		v.SetName(idx, name)
	}
}

func (v *DefNameable) HasName(name string) bool {
	_, exists := v.names[name]
	return exists
}

func (v *DefNameable) HasNameFor(idx int) bool {
	if idx >= 1 && idx <= v.length {
		for _, index := range v.names {
			if idx == index {
				return true
			}
		}
	}

	return false
}
