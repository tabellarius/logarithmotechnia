package dataframe

import (
	"fmt"
	"logarithmotechnia/vector"
	"testing"
)

func TestDataframe_Filter(t *testing.T) {
	df := New([]vector.Vector{
		vector.IntegerWithNA([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil),
		vector.StringWithNA([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}, nil),
		vector.BooleanWithNA([]bool{true, false, true, false, true, false, true, false, true, false}, nil),
	})

	testData := []struct {
		name         string
		selectorInt  []int
		selectorBool []bool
		dfColumns    []vector.Vector
	}{
		{
			name:        "indices",
			selectorInt: []int{-1, 0, 1, 3, 5, 8, 10, 11, 100},
			dfColumns: []vector.Vector{
				vector.IntegerWithNA([]int{0, 1, 3, 5, 8, 10}, []bool{true, false, false, false, false, false}),
				vector.StringWithNA([]string{"", "1", "3", "5", "8", "10"}, []bool{true, false, false, false, false, false}),
				vector.BooleanWithNA([]bool{false, true, true, true, false, false}, []bool{true, false, false, false, false, false}),
			},
		},
		{
			name:         "boolean full",
			selectorBool: []bool{true, false, true, false, true, false, false, true, false, true},
			dfColumns: []vector.Vector{
				vector.IntegerWithNA([]int{1, 3, 5, 8, 10}, nil),
				vector.StringWithNA([]string{"1", "3", "5", "8", "10"}, nil),
				vector.BooleanWithNA([]bool{true, true, true, false, false}, nil),
			},
		},
		{
			name:         "boolean odd",
			selectorBool: []bool{true, false},
			dfColumns: []vector.Vector{
				vector.IntegerWithNA([]int{1, 3, 5, 7, 9}, nil),
				vector.StringWithNA([]string{"1", "3", "5", "7", "9"}, nil),
				vector.BooleanWithNA([]bool{true, true, true, true, true}, nil),
			},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			var newDf *Dataframe
			if data.selectorInt != nil {
				newDf = df.Filter(data.selectorInt)
			} else if data.selectorBool != nil {
				newDf = df.Filter(data.selectorBool)
			}

			if !vector.CompareVectorArrs(newDf.columns, data.dfColumns) {
				t.Error(fmt.Sprintf("Columns (%v) are not equal to expected (%v)", newDf.columns, data.dfColumns))
			}
		})
	}
}
