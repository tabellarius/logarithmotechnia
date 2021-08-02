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

	IsEmpty() bool
}

type dataframe struct {
	rowNum   int
	colNum   int
	vectors  []vector.Vector
	colNames []string
}

func Df(data interface{}) (Dataframe, error) {
	return nil, nil
}
