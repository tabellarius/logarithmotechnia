package util

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
