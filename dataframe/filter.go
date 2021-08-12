package dataframe

import (
	"github.com/dee-ru/logarithmotechnia/util"
	"github.com/dee-ru/logarithmotechnia/vector"
)

func (df *Dataframe) Filter(filter interface{}) *Dataframe {
	if indices, ok := filter.([]int); ok {
		return df.ByIndices(indices)
	}

	if which, ok := filter.([]bool); ok {
		indices := util.ToIndices(df.rowNum, which)
		return df.ByIndices(indices)
	}

	return New([]vector.Vector{})
}
