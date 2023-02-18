package dataframe

import (
	"fmt"
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
	groupIndex        vector.GroupIndex
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

func (df *Dataframe) C(selector any) vector.Vector {
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

func (df *Dataframe) Pick(idx int) map[string]any {
	rowMap := map[string]any{}
	if idx < 1 || idx > df.rowNum {
		return rowMap
	}

	for _, column := range df.columns {
		rowMap[column.Name()] = column.Pick(idx)
	}

	return rowMap
}

func (df *Dataframe) Traverse(traverser any) {
	if fn, ok := traverser.(func(int, map[string]any)); ok {
		for i := 1; i <= df.rowNum; i++ {
			fn(i, df.Pick(i))
		}
	}

	if fn, ok := traverser.(func(map[string]any)); ok {
		for i := 1; i <= df.rowNum; i++ {
			fn(df.Pick(i))
		}
	}
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

func (df *Dataframe) String() string {
	var str string

	for i, column := range df.columns {
		str += fmt.Sprintf("%s: %v\n", df.columnNames[i], column)
	}

	return str
}

func (df *Dataframe) ToMap() map[string][]any {
	dataMap := map[string][]any{}

	for _, columnName := range df.columnNames {
		data, _ := df.Cn(columnName).Anies()
		dataMap[columnName] = data
	}

	return dataMap
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
	return append(df.Options(), OptionColumnNames(df.columnNames))
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

func New(data any, options ...vector.Option) *Dataframe {
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

	options = append(options, OptionColumnNames(names))

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
			vectors[i] = v.Clone()
		}
	}

	colNum := len(vectors)

	conf := vector.MergeOptions(options)

	columnNames := generateColumnNames(colNum)
	if conf.HasOption(KeyOptionColumnNames) {
		names := conf.Value(KeyOptionColumnNames).([]string)
		names = renameDuplicateColumns(names)
		if colNum >= len(names) {
			copy(columnNames, names)
		} else {
			copy(columnNames, names[0:colNum])
		}
	}

	vectorOptions := []vector.Option{}
	if conf.HasOption(KeyOptionVectorOptions) {
		vectorOptions = conf.Value(KeyOptionVectorOptions).([]vector.Option)
	}

	for i, columnName := range columnNames {
		vectors[i].SetName(columnName)
		if len(vectorOptions) > 0 {
			for _, option := range vectorOptions {
				vectors[i].SetOption(option)
			}
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

func renameDuplicateColumns(names []string) []string {
	if len(names) == 0 {
		return names
	}

	uniqueNames := make([]string, len(names))
	uniqueNames[0] = names[0]
	for i := 1; i < len(names); i++ {
		id := 1
		name := names[i]

		for {
			duplicate := false
			for j := 0; j < i; j++ {
				if uniqueNames[j] == name {
					duplicate = true
					break
				}
			}

			if !duplicate {
				break
			}

			name = names[i] + "_" + strconv.Itoa(id)
			id++
		}

		uniqueNames[i] = name
	}

	return uniqueNames
}
