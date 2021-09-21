package dataframe

func strPosInSlice(slice []string, str string) int {
	for i, elem := range slice {
		if str == elem {
			return i
		}
	}

	return -1
}

func reverseIndices(indices []int) []int {
	newIndices := make([]int, len(indices))
	idx := 0

	for i := len(indices) - 1; i >= 0; i-- {
		newIndices[idx] = indices[i]
		idx++
	}

	return newIndices
}
