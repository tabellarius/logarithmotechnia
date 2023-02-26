package dataframe

import (
	"logarithmotechnia/vector"
)

// GroupBy transforms the dataframe into a grouped one which later can be used for aggregations.
//
// Acceptable selectors are:
//   - name of a column
//   - string slice of column names
//   - index of a column (starting with 1)
//   - array of column indices
func (df *Dataframe) GroupBy(selectors ...any) *Dataframe {
	groupByColumns := []string{}
	for _, selector := range selectors {
		switch val := selector.(type) {
		case string:
			if df.Names().Has(val) {
				groupByColumns = append(groupByColumns, val)
			}
		case []string:
			for _, column := range val {
				if df.Names().Has(column) {
					groupByColumns = append(groupByColumns, column)
				}
			}
		case int:
			if df.IsValidColumnIndex(val) {
				groupByColumns = append(groupByColumns, df.columnNames[val-1])
			}
		case []int:
			for _, index := range val {
				if df.IsValidColumnIndex(index) {
					groupByColumns = append(groupByColumns, df.columnNames[index-1])
				}
			}
		}
	}

	if len(groupByColumns) == 0 {
		return df
	}

	var groups [][]int
	for _, groupBy := range groupByColumns {
		groups = df.groupByColumn(groupBy, groups)
	}

	if len(groups) == 0 {
		return df
	}

	options := append(df.OptionsWithNames(), vector.OptionGroupIndex(groups))

	newColumns := make([]vector.Vector, df.colNum)
	//	fmt.Println(groups)
	for i, column := range df.columns {
		newColumns[i] = column.GroupByIndices(groups)
	}
	newDf := New(newColumns, options...)
	newDf.groupedBy = groupByColumns

	return newDf
}

func (df *Dataframe) groupByColumn(groupBy string, curGroups [][]int) [][]int {
	if len(curGroups) == 0 {
		groups, _ := df.Cn(groupBy).Groups()
		return groups
	}

	newIndices := [][]int{}
	for _, indices := range curGroups {
		if len(indices) == 1 {
			newIndices = append(newIndices, indices)
			continue
		}

		subGroups, _ := df.Cn(groupBy).ByIndices(indices).Groups()
		replaceGroups := make([][]int, len(subGroups))
		for j, subIndices := range subGroups {
			newGroup := make([]int, len(subIndices))
			for k, idx := range subIndices {
				newGroup[k] = indices[idx-1]
			}
			replaceGroups[j] = newGroup
		}

		newIndices = append(newIndices, replaceGroups...)
	}

	return newIndices
}

// IsGrouped returns true if the dataframe is grouped.
func (df *Dataframe) IsGrouped() bool {
	return len(df.groupedBy) > 0
}

// GroupedBy returns string slice with the names of columns by which dataframe was grouped.
func (df *Dataframe) GroupedBy() []string {
	groupedBy := make([]string, len(df.groupedBy))
	copy(groupedBy, df.groupedBy)

	return groupedBy
}

// Ungroup returns ungrouped dataframe.
func (df *Dataframe) Ungroup() *Dataframe {
	if !df.IsGrouped() {
		return df
	}

	columns := make([]vector.Vector, df.colNum)
	for i := 0; i < df.colNum; i++ {
		columns[i] = df.columns[i].Ungroup()
	}

	return New(df.columns, df.OptionsWithNames()...)
}
