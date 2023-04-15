package util

func FromTo(from, to, length int) []int {
	/* from and to have different signs */
	if from*to < 0 {
		return []int{}
	}

	var indices []int
	if from == 0 && to == 0 {
		indices = []int{}
	} else if from > 0 && from > to {
		indices = byFromToReverse(to, from, length)
	} else if from <= 0 && to <= 0 {
		from *= -1
		to *= -1
		if from > to {
			from, to = to, from
		}
		indices = byFromToWithRemove(from, to, length)
	} else {
		indices = byFromToRegular(from, to, length)
	}

	return indices
}

func byFromToRegular(from, to, length int) []int {
	from, to = normalizeFromTo(from, to, length)

	indices := make([]int, to-from+1)
	index := 0
	for idx := from; idx <= to; idx++ {
		indices[index] = idx
		index++
	}

	return indices
}

func byFromToReverse(from, to, length int) []int {
	from, to = normalizeFromTo(from, to, length)

	indices := make([]int, to-from+1)
	index := 0
	for idx := to; idx >= from; idx-- {
		indices[index] = idx
		index++
	}

	return indices
}

func byFromToWithRemove(from, to, length int) []int {
	from, to = normalizeFromTo(from, to, length)

	indices := make([]int, from-1+length-to)
	index := 0
	for idx := 1; idx < from; idx++ {
		indices[index] = idx
		index++
	}
	for idx := to + 1; idx <= length; idx++ {
		indices[index] = idx
		index++
	}

	return indices
}

func normalizeFromTo(from, to, length int) (int, int) {
	if to > length {
		to = length
	}
	if from < 1 {
		from = 1
	}

	return from, to
}
