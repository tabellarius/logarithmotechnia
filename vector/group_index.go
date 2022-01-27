package vector

type GroupIndex [][]int

func (index GroupIndex) FirstElements() []int {
	elems := make([]int, len(index))

	for i, idx := range index {
		elems[i] = idx[0]
	}

	return elems
}
