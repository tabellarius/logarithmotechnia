package dataframe

import "logarithmotechnia.com/logarithmotechnia/vector"

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
	vectors  []vector.Vector
	colNames []string
}

func (df *dataframe) Clone() Dataframe {
	panic("implement me")
}

func (df *dataframe) ByIndices(indices []int) Dataframe {
	panic("implement me")
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

func Df(data interface{}) (Dataframe, error) {
	var df dataframe
	if vectors, ok := data.([]vector.Vector); ok {
		df = dataframeFromVectors(vectors)
	}

	return &df, nil
}

func dataframeFromVectors(vectors []vector.Vector) dataframe {

	return dataframe{
		rowNum:   0,
		colNum:   0,
		vectors:  vectors,
		colNames: nil,
	}
}
