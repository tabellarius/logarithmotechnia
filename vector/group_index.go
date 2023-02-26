package vector

// GroupIndex holds element indices for different groups in grouped vector.
type GroupIndex [][]int

// FirstElements returns indices of first element for each group.
func (index GroupIndex) FirstElements() []int {
	elems := make([]int, len(index))

	for i, idx := range index {
		elems[i] = idx[0]
	}

	return elems
}
