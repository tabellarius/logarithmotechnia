package dataframe

import (
	"logarithmotechnia.com/logarithmotechnia/vector"
)

type Dataframe interface {
	RowNum() int
	ColNum() int
	Clone() Dataframe

	ByIndices(indices []int) Dataframe

	Filter(filter interface{}) Dataframe
	SupportsFilter(filter interface{}) bool

	Select(selectors ...interface{})
	Mutate(params ...interface{})
	Transmute(params ...interface{})
	Relocate(params ...interface{})

	Vectors() []vector.Vector

	IsEmpty() bool
}

type dataframe struct {
	rowNum   int
	colNum   int
	columns  []vector.Vector
	colNames []string
}

func (df *dataframe) Clone() Dataframe {
	panic("implement me")
}

func (df *dataframe) ByIndices(indices []int) Dataframe {
	newColumns := make([]vector.Vector, df.colNum)

	for i, column := range df.columns {
		newColumns[i] = column.ByIndices(indices)
	}

	return New(newColumns)
}

func (df *dataframe) Filter(filter interface{}) Dataframe {
	panic("implement me")
}

func (df *dataframe) SupportsFilter(filter interface{}) bool {
	panic("implement me")
}

func (df *dataframe) Select(selectors ...interface{}) {
	panic("implement me")
}

func (df *dataframe) Mutate(params ...interface{}) {
	panic("implement me")
}

func (df *dataframe) Transmute(params ...interface{}) {
	panic("implement me")
}

func (df *dataframe) Relocate(params ...interface{}) {
	panic("implement me")
}

func (df *dataframe) Vectors() []vector.Vector {
	panic("implement me")
}

func (df *dataframe) IsEmpty() bool {
	panic("implement me")
}

func (df *dataframe) RowNum() int {
	return df.rowNum
}

func (df *dataframe) ColNum() int {
	return df.colNum
}

func New(data interface{}) Dataframe {
	var df dataframe
	if vectors, ok := data.([]vector.Vector); ok {
		df = dataframeFromVectors(vectors)
	}

	return &df
}

func dataframeFromVectors(vectors []vector.Vector) dataframe {
	maxLen := 0

	for _, v := range vectors {
		if v.Len() > maxLen {
			maxLen = v.Len()
		}
	}

	for i, v := range vectors {
		if v.Len() < maxLen {
			vectors[i] = v.Append(vector.NA(maxLen - v.Len()))
		}
	}

	return dataframe{
		rowNum:   maxLen,
		colNum:   len(vectors),
		columns:  vectors,
		colNames: nil,
	}
}
