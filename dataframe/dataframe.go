package dataframe

import (
	"github.com/dee-ru/logarithmotechnia/vector"
	"strconv"
)

type Dataframe struct {
	rowNum  int
	colNum  int
	columns []vector.Vector
	config  Config
}

func (df *Dataframe) RowNum() int {
	return df.rowNum
}

func (df *Dataframe) ColNum() int {
	return df.colNum
}

func (df *Dataframe) Clone() *Dataframe {
	return New(df.columns, df.config)
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
	df.config.columnNamesVector = vector.String(df.config.columnNames, nil)

	return df
}

func (df *Dataframe) setColumnName(index int, name string) {
	if df.IsValidColumnIndex(index) {
		if !df.HasColumn(name) {
			df.config.columnNames[index-1] = name
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
	df.config.columnNamesVector = vector.String(df.config.columnNames, nil)

	return df
}

func (df *Dataframe) Names() vector.Vector {
	return df.config.columnNamesVector
}

func (df *Dataframe) NamesAsStrings() []string {
	names := make([]string, df.colNum)
	copy(names, df.config.columnNames)

	return names
}

func (df *Dataframe) ByIndices(indices []int) *Dataframe {
	newColumns := make([]vector.Vector, df.colNum)

	for i, column := range df.columns {
		newColumns[i] = column.ByIndices(indices)
	}

	return New(newColumns, df.config)
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
	return strPosInSlice(df.config.columnNames, name) != -1
}

func (df *Dataframe) columnIndexByName(name string) int {
	index := 1

	for _, columnName := range df.config.columnNames {
		if columnName == name {
			return index
		}
		index++
	}

	return 0
}

func New(data interface{}, options ...Config) *Dataframe {
	var df *Dataframe
	if vectors, ok := data.([]vector.Vector); ok {
		df = dataframeFromVectors(vectors, options...)
	} else {
		df = dataframeFromVectors([]vector.Vector{})
	}

	return df
}

func dataframeFromVectors(vectors []vector.Vector, options ...Config) *Dataframe {
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

	config := mergeConfigs(options)

	columnNames := generateColumnNames(colNum)
	if colNum >= len(config.columnNames) {
		copy(columnNames, config.columnNames)
	} else {
		copy(columnNames, config.columnNames[0:colNum])
	}
	config.columnNames = columnNames

	return &Dataframe{
		rowNum:  maxLen,
		colNum:  colNum,
		columns: vectors,
		config:  config,
	}
}

func generateColumnNames(length int) []string {
	names := make([]string, length)

	for i := 1; i <= length; i++ {
		names[i-1] = strconv.Itoa(i)
	}

	return names
}
