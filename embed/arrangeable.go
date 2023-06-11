package embed

import (
	"logarithmotechnia/internal/util"
	"sort"
)

// Arrangeable can be embedded into a payload for easy implementation of Arrangeable interface to make the payload
// sortable. It uses NAble to support NA-values.
type Arrangeable struct {
	Length int
	NAble
	FnLess      func(i, j int) bool
	FnEqual     func(i, j int) bool
	sortedCache []int
}

func (ar *Arrangeable) SortedIndicesZeroBased() []int {
	if ar.sortedCache != nil {
		cached := make([]int, len(ar.sortedCache))
		copy(cached, ar.sortedCache)

		return ar.sortedCache
	}

	indices := util.IndicesArray(ar.Length)

	var fn func(i, j int) bool
	if ar.HasNA() {
		fn = func(i, j int) bool {
			if ar.NA[indices[i]] && ar.NA[indices[j]] {
				return i < j
			}

			if ar.NA[indices[i]] {
				return false
			}

			if ar.NA[indices[j]] {
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

func (ar *Arrangeable) SortedIndices() []int {
	return util.IncIndices(ar.SortedIndicesZeroBased())
}

func (ar *Arrangeable) SortedIndicesWithRanks() ([]int, []int) {
	indices := ar.SortedIndicesZeroBased()

	if len(indices) == 0 {
		return indices, []int{}
	}

	if len(indices) == 1 {
		return indices, []int{1}
	}

	rank := 1
	ranks := make([]int, ar.Length)
	if ar.NA[0] {
		rank = 0
	}
	ranks[0] = rank
	for i := 1; i < ar.Length; i++ {
		if ar.NA[indices[i]] != ar.NA[indices[i-1]] || !ar.FnEqual(indices[i], indices[i-1]) {
			rank++
			ranks[i] = rank
		} else {
			ranks[i] = rank
		}
	}

	return util.IncIndices(indices), ranks
}
