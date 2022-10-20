package dataframe

import (
	"logarithmotechnia/vector"
)

func (df *Dataframe) Arrange(args ...any) *Dataframe {
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

	conf := vector.MergeOptions(options)
	reverseColumns := []string{}
	if conf.HasOption(vector.KeyOptionArrangeReverseColumns) {
		reverseColumns = conf.Value(vector.KeyOptionArrangeReverseColumns).([]string)
	}
	rcVec := vector.StringWithNA(reverseColumns, nil)

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
		if rcVec.Has(arrangeBy) {
			indices = reverseIndices(indices)
		}
	} else {
		indices = df.arrangeByMultiple(arrangeBy, columns[1:], rcVec)
	}

	if conf.HasOption(vector.KeyOptionArrangeReverse) && conf.Value(vector.KeyOptionArrangeReverse).(bool) {
		indices = reverseIndices(indices)
	}

	return df.ByIndices(indices)
}

func (df *Dataframe) arrangeByMultiple(arrangeBy string, additional []string, reverse vector.Vector) []int {
	if len(additional) == 0 {
		indices := df.Cn(arrangeBy).SortedIndices()
		if reverse.Has(arrangeBy) {
			indices = reverseIndices(indices)
		}
		return indices
	}

	indices, ranks := df.C(arrangeBy).SortedIndicesWithRanks()

	subIndices := []int{indices[0]}
	ranks = append(ranks, -1)
	for i := 1; i < len(ranks); i++ {
		if ranks[i] != ranks[i-1] {
			if len(subIndices) > 1 {
				subDf := df.Select(additional).ByIndices(subIndices)
				reIndices := subDf.arrangeByMultiple(additional[0], additional[1:], reverse)

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

	if reverse.Has(arrangeBy) {
		indices = reverseIndices(indices)
	}

	return indices
}
