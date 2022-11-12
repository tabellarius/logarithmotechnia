package dataframe

import (
	"logarithmotechnia/util"
	"logarithmotechnia/vector"
)

func (df *Dataframe) Filter(filter any) *Dataframe {
	switch filter.(type) {
	case []int:
		return df.ByIndices(filter.([]int))
	case []bool:
		indices := util.ToIndices(df.rowNum, filter.([]bool))
		return df.ByIndices(indices)
	case func(int, map[string]any) bool:
		fn := filter.(func(int, map[string]any) bool)
		indices := util.ToIndices(df.rowNum, df.filterByFunc(fn))
		return df.ByIndices(indices)
	case func(map[string]any) bool:
		fn := filter.(func(map[string]any) bool)
		indices := util.ToIndices(df.rowNum, df.filterByCompactFunc(fn))
		return df.ByIndices(indices)
	}

	return New([]vector.Vector{}, df.Options()...)
}

func (df *Dataframe) filterByFunc(fn func(int, map[string]any) bool) []bool {
	booleans := make([]bool, df.rowNum)

	for i := 1; i <= df.rowNum; i++ {
		booleans[i-1] = fn(i, df.Pick(i))
	}

	return booleans
}

func (df *Dataframe) filterByCompactFunc(fn func(map[string]any) bool) []bool {
	booleans := make([]bool, df.rowNum)

	for i := 1; i <= df.rowNum; i++ {
		booleans[i-1] = fn(df.Pick(i))
	}

	return booleans
}
