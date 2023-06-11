package util

func IndicesArray(size int) []int {
	indices := make([]int, size)

	index := 0
	for i := range indices {
		indices[i] = index
		index++
	}

	return indices
}

func ToIndices(vecLength int, booleans []bool) []int {
	var indices = make([]int, 0)
	length := len(booleans)

	if vecLength == 0 || length == 0 {
		return indices
	}

	pos := 0
	for index := 1; index <= vecLength; index++ {
		if booleans[pos] == true {
			indices = append(indices, index)
		}

		pos++
		if pos == length {
			pos = 0
		}
	}

	return indices
}

func ArrayFromAnyTo[T any](arr []any) []T {
	arrT := make([]T, len(arr))
	for i, a := range arr {
		arrT[i] = a.(T)
	}

	return arrT
}

func IncIndices(indices []int) []int {
	for i := range indices {
		indices[i]++
	}

	return indices
}
