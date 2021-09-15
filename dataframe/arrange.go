package dataframe

import (
	"logarithmotechnia/vector"
)

func (df *Dataframe) Arrange(args ...interface{}) *Dataframe {
	potentialColumns := []string{}
	options := []vector.Option{}

	for _, arg := range args {
		switch arg.(type) {
		case string:
			potentialColumns = append(potentialColumns, arg.(string))
		case []string:
			potentialColumns = append(potentialColumns, arg.([]string)...)
		case vector.Option:
			options = append(options, arg.(vector.Option))
		}
	}

	columns := []string{}
	for _, column := range potentialColumns {
		if df.HasColumn(column) {
			columns = append(columns, column)
		}
	}

	if len(columns) == 0 {
		return df
	}

	arrangeBy := columns[0]

	var indices []int
	if len(columns) == 1 || df.rowNum <= 1 {
		indices = df.C(arrangeBy).SortedIndices()
	} else {
		indices = df.arrangeByMultiple(arrangeBy, columns[1:])
	}

	conf := vector.MergeOptions(options)
	if conf.HasOption(vector.KeyOptionArrangeReverse) && conf.Value(vector.KeyOptionArrangeReverse).(bool) {
		newIndices := make([]int, len(indices))
		idx := 0
		for i := len(indices) - 1; i >= 0; i-- {
			newIndices[idx] = indices[i]
			idx++
		}
		indices = newIndices
	}

	return df.ByIndices(indices)
}

func (df *Dataframe) arrangeByMultiple(arrangeBy string, additional []string) []int {
	if len(additional) == 0 {
		return df.Cn(arrangeBy).SortedIndices()
	}

	indices, ranks := df.C(arrangeBy).SortedIndicesWithRanks()

	subIndices := []int{indices[0]}
	ranks = append(ranks, -1)
	for i := 1; i < len(ranks); i++ {
		if ranks[i] != ranks[i-1] {
			if len(subIndices) > 1 {
				subDf := df.Select(additional).ByIndices(subIndices)
				reIndices := subDf.arrangeByMultiple(additional[0], additional[1:])

				newSubIndices := make([]int, len(subIndices))
				for j, reidx := range reIndices {
					newSubIndices[j] = subIndices[reidx-1]
				}

				z := 0
				for j := i - len(subIndices); j < i; j++ {
					indices[j] = newSubIndices[z]
					z++
				}
			}

			if ranks[i] != -1 {
				subIndices = []int{indices[i]}
			}
		} else {
			subIndices = append(subIndices, indices[i])
		}
	}

	return indices
}
