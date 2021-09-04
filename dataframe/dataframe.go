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
}

func (df *Dataframe) RowNum() int {
	return df.rowNum
}

func (df *Dataframe) ColNum() int {
	return df.colNum
}

func (df *Dataframe) Clone() *Dataframe {
	return New(df.columns, df.Options()...)
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
	df.columnNamesVector = vector.String(df.columnNames, nil)

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
	df.columnNamesVector = vector.String(df.columnNames, nil)

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

	return New(newColumns, df.Options()...)
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

func New(data interface{}, options ...vector.Option) *Dataframe {
	var df *Dataframe
	switch data.(type) {
	case []vector.Vector:
		df = dataframeFromVectors(data.([]vector.Vector), options...)
	case []Column:
		df = dateframeFromColumns(data.([]Column), options...)
	default:
		df = dataframeFromVectors([]vector.Vector{})
	}

	return df
}

func dateframeFromColumns(columns []Column, options ...vector.Option) *Dataframe {
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
		columnNamesVector: vector.String(columnNames, nil),
	}
}

func (df *Dataframe) Options() []vector.Option {
	return []vector.Option{
		vector.OptionColumnNames(df.columnNames),
	}
}

func generateColumnNames(length int) []string {
	names := make([]string, length)

	for i := 1; i <= length; i++ {
		names[i-1] = strconv.Itoa(i)
	}

	return names
}
