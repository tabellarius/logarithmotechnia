package dataframe

import (
	"logarithmotechnia/internal/util"
	"logarithmotechnia/vector"
)

// Filter filters te dataframe based on the provided filter. Possible filters are:
//   - an integer slice of indices (starting from 1)
//   - a boolean slice of same length as length of the dataframe. All rows of the dataframe, which have same indices
//     as "true" elements of the slice, will be selected. If the slice has more elements than the dataframe, only first
//     N elements will be used (where N is the length of the dataframe). If the slice has fewer elements than the
//     dataframe, the slice will be recycled (f.e. []bool{true, false} will select all odd rows of the dataframe.
//   - a function func(index int, row map[string]any bool which will be called for all rows in the dataframe, but
//     only those, for which the function will have return true, will be selected.
//   - a function func(row map[string]any bool - same as previous but without index argument.
func (df *Dataframe) Filter(filter any) *Dataframe {
	switch f := filter.(type) {
	case []int:
		return df.ByIndices(f)
	case []bool:
		indices := util.ToIndices(df.rowNum, f)
		return df.ByIndices(indices)
	case func(int, map[string]any) bool:
		indices := util.ToIndices(df.rowNum, df.filterByFunc(f))
		return df.ByIndices(indices)
	case func(map[string]any) bool:
		indices := util.ToIndices(df.rowNum, df.filterByCompactFunc(f))
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

func (df *Dataframe) FromTo(from, to int) *Dataframe {
	return df.ByIndices(util.FromTo(from, to, df.RowNum()))
}
