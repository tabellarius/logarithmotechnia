package dataframe

import (
	"logarithmotechnia/vector"
	"strconv"
)

type Column struct {
	name   string
	vector vector.Vector
}

type Dataframe struct {
	rowNum            int
	colNum            int
	columns           []vector.Vector
	columnNames       []string
	columnNamesVector vector.Vector
	groupedBy         []string
}

func (df *Dataframe) RowNum() int {
	return df.rowNum
}

func (df *Dataframe) ColNum() int {
	return df.colNum
}

func (df *Dataframe) Clone() *Dataframe {
	return New(df.columns, df.OptionsWithNames()...)
}

func (df *Dataframe) Cn(name string) vector.Vector {
	index := df.columnIndexByName(name)

	if index > 0 {
		return df.columns[index-1]
	}

	return nil
}

func (df *Dataframe) C(selector interface{}) vector.Vector {
	if index, ok := selector.(int); ok {
		return df.Ci(index)
	}

	if name, ok := selector.(string); ok {
		return df.Cn(name)
	}

	return nil
}

func (df *Dataframe) Ci(index int) vector.Vector {
	if df.IsValidColumnIndex(index) {
		return df.columns[index-1]
	}

	return nil
}

func (df *Dataframe) SetColumnName(index int, name string) *Dataframe {
	df.setColumnName(index, name)
	df.columnNamesVector = vector.StringWithNA(df.columnNames, nil)

	return df
}

func (df *Dataframe) setColumnName(index int, name string) {
	if df.IsValidColumnIndex(index) {
		if !df.HasColumn(name) {
			df.columnNames[index-1] = name
		}
	}
}

func (df *Dataframe) SetColumnNames(names []string) *Dataframe {
	index := 1
	for _, name := range names {
		df.setColumnName(index, name)
		index++
		if index > df.colNum {
			break
		}
	}
	df.columnNamesVector = vector.StringWithNA(df.columnNames, nil)

	return df
}

func (df *Dataframe) Names() vector.Vector {
	return df.columnNamesVector
}

func (df *Dataframe) NamesAsStrings() []string {
	names := make([]string, df.colNum)
	copy(names, df.columnNames)

	return names
}

func (df *Dataframe) ByIndices(indices []int) *Dataframe {
	newColumns := make([]vector.Vector, df.colNum)

	for i, column := range df.columns {
		newColumns[i] = column.ByIndices(indices)
	}

	return New(newColumns, df.OptionsWithNames()...)
}

func (df *Dataframe) Columns() []vector.Vector {
	return df.columns
}

func (df *Dataframe) IsEmpty() bool {
	return df.colNum == 0
}

func (df *Dataframe) IsValidColumnIndex(index int) bool {
	if index >= 1 && index <= df.colNum {
		return true
	}

	return false
}

func (df *Dataframe) HasColumn(name string) bool {
	return strPosInSlice(df.columnNames, name) != -1
}

func (df *Dataframe) GroupBy(selectors ...interface{}) *Dataframe {
	columns := []string{}
	for _, selector := range selectors {
		switch selector.(type) {
		case string:
			columns = append(columns, selector.(string))
		case []string:
			columns = append(columns, selector.([]string)...)
		}
	}

	groupByColumns := []string{}
	for _, column := range columns {
		if df.Names().Has(column) {
			groupByColumns = append(groupByColumns, column)
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

	newColumns := make([]vector.Vector, df.colNum)
	for i, column := range df.columns {
		newColumns[i] = column.GroupByIndices(groups)
	}
	newDf := New(newColumns, df.OptionsWithNames()...)
	newDf.groupedBy = groupByColumns

	return newDf
}

func (df *Dataframe) groupByColumn(groupBy string, curGroups [][]int) [][]int {
	if len(curGroups) == 0 {
		return df.Cn(groupBy).Groups()
	}

	newIndices := [][]int{}
	for _, indices := range curGroups {
		if len(indices) == 1 {
			newIndices = append(newIndices, indices)
			continue
		}

		subGroups := df.Cn(groupBy).ByIndices(indices).Groups()
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

func (df *Dataframe) IsGrouped() bool {
	return len(df.groupedBy) > 0
}

func (df *Dataframe) GroupedBy() []string {
	groupedBy := make([]string, len(df.groupedBy))
	copy(groupedBy, df.groupedBy)

	return groupedBy
}

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

func (df *Dataframe) columnIndexByName(name string) int {
	index := 1

	for _, columnName := range df.columnNames {
		if columnName == name {
			return index
		}
		index++
	}

	return 0
}

func (df *Dataframe) OptionsWithNames() []vector.Option {
	return append(df.Options(), vector.OptionColumnNames(df.columnNames))
}

func (df *Dataframe) Options() []vector.Option {
	return []vector.Option{}
}

func generateColumnNames(length int) []string {
	names := make([]string, length)

	for i := 1; i <= length; i++ {
		names[i-1] = strconv.Itoa(i)
	}

	return names
}

func New(data interface{}, options ...vector.Option) *Dataframe {
	var df *Dataframe
	switch data.(type) {
	case []vector.Vector:
		df = dataframeFromVectors(data.([]vector.Vector), options...)
	case []Column:
		df = dataframeFromColumns(data.([]Column), options...)
	default:
		df = dataframeFromVectors([]vector.Vector{})
	}

	return df
}

func dataframeFromColumns(columns []Column, options ...vector.Option) *Dataframe {
	vectors := []vector.Vector{}
	names := []string{}

	for _, column := range columns {
		vectors = append(vectors, column.vector)
		names = append(names, column.name)
	}

	options = append(options, vector.OptionColumnNames(names))

	return dataframeFromVectors(vectors, options...)
}

func dataframeFromVectors(vectors []vector.Vector, options ...vector.Option) *Dataframe {
	maxLen := 0

	for _, v := range vectors {
		if v.Len() > maxLen {
			maxLen = v.Len()
		}
	}

	for i, v := range vectors {
		if v.Len() < maxLen {
			vectors[i] = v.Append(vector.NA(maxLen - v.Len()))
		} else {
			vectors[i] = v
		}
	}

	colNum := len(vectors)

	conf := vector.MergeOptions(options)

	columnNames := generateColumnNames(colNum)
	if conf.HasOption(vector.KeyOptionColumnNames) {
		names := conf.Value(vector.KeyOptionColumnNames).([]string)
		if colNum >= len(names) {
			copy(columnNames, names)
		} else {
			copy(columnNames, names[0:colNum])
		}
	}

	return &Dataframe{
		rowNum:            maxLen,
		colNum:            colNum,
		columns:           vectors,
		columnNames:       columnNames,
		columnNamesVector: vector.StringWithNA(columnNames, nil),
	}
}
