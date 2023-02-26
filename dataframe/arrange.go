package dataframe

import (
	"logarithmotechnia/vector"
)

// Arrange orders the rows of a data frame by the values of selected columns.
// Possible parameters: a name of a column, an array of column names, an index of a column, an array of column names.
func (df *Dataframe) Arrange(args ...any) *Dataframe {
	columns := []string{}
	options := []vector.Option{}

	for _, arg := range args {
		switch val := arg.(type) {
		case string:
			if df.HasColumn(val) {
				columns = append(columns, val)
			}
		case []string:
			for _, colName := range val {
				if df.HasColumn(colName) {
					columns = append(columns, colName)
				}
			}
		case int:
			if df.IsValidColumnIndex(val) {
				columns = append(columns, df.columnNames[val-1])
			}
		case []int:
			for _, index := range val {
				if df.IsValidColumnIndex(index) {
					columns = append(columns, df.columnNames[index-1])
				}
			}
		case vector.Option:
			options = append(options, arg.(vector.Option))
		}
	}

	conf := vector.MergeOptions(options)
	reverseColumns := []string{}
	if conf.HasOption(KeyOptionArrangeReverseColumns) {
		reverseColumns = conf.Value(KeyOptionArrangeReverseColumns).([]string)
	}
	rcVec := vector.StringWithNA(reverseColumns, nil)

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

	if conf.HasOption(KeyOptionArrangeReverse) && conf.Value(KeyOptionArrangeReverse).(bool) {
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
