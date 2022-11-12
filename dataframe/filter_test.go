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
	}, OptionColumnNames([]string{"int", "string", "bool"}))

	testData := []struct {
		name      string
		selector  any
		dfColumns []vector.Vector
	}{
		{
			name:     "indices",
			selector: []int{-1, 0, 1, 3, 5, 8, 10, 11, 100},
			dfColumns: []vector.Vector{
				vector.IntegerWithNA([]int{0, 1, 3, 5, 8, 10}, []bool{true, false, false, false, false, false}),
				vector.StringWithNA([]string{"", "1", "3", "5", "8", "10"}, []bool{true, false, false, false, false, false}),
				vector.BooleanWithNA([]bool{false, true, true, true, false, false}, []bool{true, false, false, false, false, false}),
			},
		},
		{
			name:     "boolean full",
			selector: []bool{true, false, true, false, true, false, false, true, false, true},
			dfColumns: []vector.Vector{
				vector.IntegerWithNA([]int{1, 3, 5, 8, 10}, nil),
				vector.StringWithNA([]string{"1", "3", "5", "8", "10"}, nil),
				vector.BooleanWithNA([]bool{true, true, true, false, false}, nil),
			},
		},
		{
			name:     "boolean odd",
			selector: []bool{true, false},
			dfColumns: []vector.Vector{
				vector.IntegerWithNA([]int{1, 3, 5, 7, 9}, nil),
				vector.StringWithNA([]string{"1", "3", "5", "7", "9"}, nil),
				vector.BooleanWithNA([]bool{true, true, true, true, true}, nil),
			},
		},
		{
			name: "function",
			selector: func(i int, row map[string]any) bool {
				if i <= 5 && row["bool"].(bool) {
					return true
				}

				return false
			},
			dfColumns: []vector.Vector{
				vector.IntegerWithNA([]int{1, 3, 5}, nil),
				vector.StringWithNA([]string{"1", "3", "5"}, nil),
				vector.BooleanWithNA([]bool{true, true, true}, nil),
			},
		},
		{
			name: "compact function",
			selector: func(row map[string]any) bool {
				return row["bool"].(bool)
			},
			dfColumns: []vector.Vector{
				vector.IntegerWithNA([]int{1, 3, 5, 7, 9}, nil),
				vector.StringWithNA([]string{"1", "3", "5", "7", "9"}, nil),
				vector.BooleanWithNA([]bool{true, true, true, true, true}, nil),
			},
		},
		{
			name:      "invalid selector",
			selector:  []complex128{0 + 0i, 1 + 1i},
			dfColumns: []vector.Vector{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			newDf := df.Filter(data.selector)

			if !vector.CompareVectorArrs(newDf.columns, data.dfColumns) {
				t.Error(fmt.Sprintf("Columns (%v) are not equal to expected (%v)", newDf.columns, data.dfColumns))
			}
		})
	}
}
