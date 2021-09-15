package vector

func indicesArray(size int) []int {
	indices := make([]int, size)

	index := 0
	for i := range indices {
		indices[i] = index
		index++
	}

	return indices
}

func incIndices(indices []int) []int {
	for i := range indices {
		indices[i]++
	}

	return indices
}
