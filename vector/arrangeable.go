package vector

import (
	"sort"
)

type DefArrangeable struct {
	Length int
	DefNAble
	FnLess  func(i, j int) bool
	FnEqual func(i, j int) bool
}

func (ar *DefArrangeable) sortedIndices() []int {
	indices := indicesArray(ar.Length)

	var fn func(i, j int) bool
	if ar.HasNA() {
		fn = func(i, j int) bool {
			if ar.na[indices[i]] && ar.na[indices[j]] {
				return i < j
			}

			if ar.na[indices[i]] {
				return false
			}

			if ar.na[indices[j]] {
				return true
			}

			return ar.FnLess(indices[i], indices[j])
		}
	} else {
		fn = func(i, j int) bool {
			return ar.FnLess(indices[i], indices[j])
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
	ranks := make([]int, ar.Length)
	if ar.na[0] {
		rank = 0
	}
	ranks[0] = rank
	for i := 1; i < ar.Length; i++ {
		if ar.na[indices[i]] != ar.na[indices[i-1]] || !ar.FnEqual(indices[i], indices[i-1]) {
			rank++
			ranks[i] = rank
		} else {
			ranks[i] = rank
		}
	}

	return incIndices(indices), ranks
}
