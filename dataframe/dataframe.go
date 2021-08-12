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

func (df *Dataframe) Clone() Dataframe {
	panic("implement me")
}

func (df *Dataframe) ColumnNames() {
	panic("implement me")
}

func (df *Dataframe) Cn(name string) vector.Vector {
	index := df.columnIndexByName(name)

	if index > 0 {
		return df.columns[index]
	}

	return nil
}

func (df *Dataframe) Ci(index int) vector.Vector {
	if df.isValidIndex(index) {
		return df.columns[index]
	}

	return nil
}

func (df *Dataframe) SetColumnName(index int, name string) *Dataframe {
	if df.isValidIndex(index) {
		df.config.columnNames[index] = name
	}

	return df
}

func (df *Dataframe) SetColumnNames(strings []string) *Dataframe {

	return df
}

func (df *Dataframe) GetColumnNames() []string {
	names := make([]string, df.colNum)
	copy(names, df.config.columnNames)

	return names
}

func (df *Dataframe) ByIndices(indices []int) *Dataframe {
	newColumns := make([]vector.Vector, df.colNum)

	for i, column := range df.columns {
		newColumns[i] = column.ByIndices(indices)
	}

	return New(newColumns)
}

func (df *Dataframe) Columns() []vector.Vector {
	return df.columns
}

func (df *Dataframe) IsEmpty() bool {
	panic("implement me")
}

func (df *Dataframe) isValidIndex(index int) bool {
	if index >= 1 && index <= df.colNum {
		return true
	}

	return false
}

func (df *Dataframe) columnIndexByName(name string) int {
	index := 0

	for _, columnName := range df.config.columnNames {
		index++
		if columnName == name {
			break
		}
	}

	return index
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
