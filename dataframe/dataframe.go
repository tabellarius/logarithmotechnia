package dataframe

import (
	"logarithmotechnia.com/logarithmotechnia/vector"
)

/*
type Dataframe interface {
	RowNum() int
	ColNum() int
	Clone() Dataframe

	ColNames()
	SetColname(index int, name string)
	SetColNames([]string)
	Columns() []vector.Vector

	ByIndices(indices []int) Dataframe

	Filter(filter interface{}) Dataframe
	SupportsFilter(filter interface{}) bool

	Select(selectors ...interface{})
	Mutate(params ...interface{})
	Transmute(params ...interface{})
	Relocate(params ...interface{})

	IsEmpty() bool
}
*/

type Dataframe struct {
	rowNum   int
	colNum   int
	columns  []vector.Vector
	colNames []string
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

func (df *Dataframe) ColNames() {
	panic("implement me")
}

func (df *Dataframe) SetColname(index int, name string) {
	panic("implement me")
}

func (df *Dataframe) SetColNames(strings []string) {
	panic("implement me")
}

func (df *Dataframe) ByIndices(indices []int) *Dataframe {
	newColumns := make([]vector.Vector, df.colNum)

	for i, column := range df.columns {
		newColumns[i] = column.ByIndices(indices)
	}

	return New(newColumns)
}

func (df *Dataframe) Filter(filter interface{}) Dataframe {
	panic("implement me")
}

func (df *Dataframe) SupportsFilter(filter interface{}) bool {
	panic("implement me")
}

func (df *Dataframe) Select(selectors ...interface{}) {
	panic("implement me")
}

func (df *Dataframe) Mutate(params ...interface{}) {
	panic("implement me")
}

func (df *Dataframe) Transmute(params ...interface{}) {
	panic("implement me")
}

func (df *Dataframe) Relocate(params ...interface{}) {
	panic("implement me")
}

func (df *Dataframe) Columns() []vector.Vector {
	panic("implement me")
}

func (df *Dataframe) IsEmpty() bool {
	panic("implement me")
}

func New(data interface{}) *Dataframe {
	var df *Dataframe
	if vectors, ok := data.([]vector.Vector); ok {
		df = dataframeFromVectors(vectors)
	}

	return df
}

func dataframeFromVectors(vectors []vector.Vector) *Dataframe {
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

	return &Dataframe{
		rowNum:   maxLen,
		colNum:   len(vectors),
		columns:  vectors,
		colNames: nil,
	}
}
