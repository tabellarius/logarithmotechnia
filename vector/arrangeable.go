package vector

import "sort"

type DefArrangeable struct {
	length int
	DefNAble
	fnLess  func(i, j int) bool
	fnEqual func(i, j int) bool
}

func (ar *DefArrangeable) sortedIndices() []int {
	indices := indicesArray(ar.length)

	var fn func(i, j int) bool
	if ar.HasNA() {
		fn = func(i, j int) bool {
			if ar.na[indices[i]] && ar.na[indices[j]] {
				return i < j
			}

			if ar.na[indices[i]] {
				return true
			}

			if ar.na[indices[j]] {
				return false
			}

			return ar.fnLess(indices[i], indices[j])
		}
	} else {
		fn = func(i, j int) bool {
			return ar.fnLess(indices[i], indices[j])
		}
	}

	sort.Slice(indices, fn)

	return indices
}

func (ar *DefArrangeable) SortedIndices() []int {
	return incIndices(ar.sortedIndices())
}

func (ar *DefArrangeable) SortedIndicesWithRanks() ([]int, []int) {
	indices := ar.sortedIndices()

	if len(indices) == 0 {
		return indices, []int{}
	}

	if len(indices) == 1 {
		return indices, []int{1}
	}

	rank := 1
	ranks := make([]int, ar.length)
	if ar.na[0] {
		rank = 0
	}
	ranks[0] = rank
	for i := 1; i < ar.length; i++ {
		if ar.na[indices[i]] != ar.na[indices[i-1]] || !ar.fnEqual(indices[i], indices[i-1]) {
			rank++
			ranks[i] = rank
		} else {
			ranks[i] = rank
		}
	}

	return incIndices(indices), ranks
}
