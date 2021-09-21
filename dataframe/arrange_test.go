package dataframe

import (
	"fmt"
	"logarithmotechnia/vector"
	"testing"
)

func TestDataframe_Arrange(t *testing.T) {
	df := New([]Column{
		{"activity", vector.Boolean([]bool{true, false, true, false, true, false, true, false, true,
			false, true, false, true, false, true}, nil)},
		{"type", vector.Integer([]int{1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3}, nil)},
		{"salary", vector.Integer([]int{100, 300, 50, 50, 350, 45, 120, 225, 60, 30, 220, 70, 180, 35, 110}, nil)},
	})

	testData := []struct {
		name            string
		df              *Dataframe
		expectedColumns []vector.Vector
	}{
		{
			name: "one arrangeBy",
			df:   df.Arrange("salary"),
			expectedColumns: []vector.Vector{
				vector.Boolean([]bool{false, false, false, true, false, true, false, true, true, true, true, true, false, false, true}, nil),
				vector.Integer([]int{1, 2, 3, 3, 1, 3, 3, 1, 3, 1, 1, 2, 2, 2, 2}, nil),
				vector.Integer([]int{30, 35, 45, 50, 50, 60, 70, 100, 110, 120, 180, 220, 225, 300, 350}, nil),
			},
		},
		{
			name: "one arrangeBy with false reverse",
			df:   df.Arrange("salary", vector.OptionArrangeReverse(false)),
			expectedColumns: []vector.Vector{
				vector.Boolean([]bool{false, false, false, true, false, true, false, true, true, true, true, true, false, false, true}, nil),
				vector.Integer([]int{1, 2, 3, 3, 1, 3, 3, 1, 3, 1, 1, 2, 2, 2, 2}, nil),
				vector.Integer([]int{30, 35, 45, 50, 50, 60, 70, 100, 110, 120, 180, 220, 225, 300, 350}, nil),
			},
		},
		{
			name: "one arrangeBy with reverse",
			df:   df.Arrange("salary", vector.OptionArrangeReverse(true)),
			expectedColumns: []vector.Vector{
				vector.Boolean([]bool{true, false, false, true, true, true, true, true, false, true, false, true, false, false, false}, nil),
				vector.Integer([]int{2, 2, 2, 2, 1, 1, 3, 1, 3, 3, 1, 3, 3, 2, 1}, nil),
				vector.Integer([]int{350, 300, 225, 220, 180, 120, 110, 100, 70, 60, 50, 50, 45, 35, 30}, nil),
			},
		},
		{
			name: "two arrangeBy",
			df:   df.Arrange("type", "salary"),
			expectedColumns: []vector.Vector{
				vector.Boolean([]bool{false, false, true, true, true, false, true, false, false, true, false, true, true, false, true}, nil),
				vector.Integer([]int{1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3}, nil),
				vector.Integer([]int{30, 50, 100, 120, 180, 35, 220, 225, 300, 350, 45, 50, 60, 70, 110}, nil),
			},
		},
		{
			name: "two arrangeBy with false reverse",
			df:   df.Arrange("type", "salary", vector.OptionArrangeReverse(false)),
			expectedColumns: []vector.Vector{
				vector.Boolean([]bool{false, false, true, true, true, false, true, false, false, true, false, true, true, false, true}, nil),
				vector.Integer([]int{1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3}, nil),
				vector.Integer([]int{30, 50, 100, 120, 180, 35, 220, 225, 300, 350, 45, 50, 60, 70, 110}, nil),
			},
		},
		{
			name: "three arrangeBy",
			df:   df.Arrange("activity", "type", "salary"),
			expectedColumns: []vector.Vector{
				vector.Boolean([]bool{false, false, false, false, false, false, false, true, true, true, true, true, true, true, true}, nil),
				vector.Integer([]int{1, 1, 2, 2, 2, 3, 3, 1, 1, 1, 2, 2, 3, 3, 3}, nil),
				vector.Integer([]int{30, 50, 35, 225, 300, 45, 70, 100, 120, 180, 220, 350, 50, 60, 110}, nil),
			},
		},
		{
			name: "three arrangeBy with reverse",
			df:   df.Arrange("activity", "type", "salary", vector.OptionArrangeReverse(true)),
			expectedColumns: []vector.Vector{
				vector.Boolean([]bool{true, true, true, true, true, true, true, true, false, false, false, false, false, false, false}, nil),
				vector.Integer([]int{3, 3, 3, 2, 2, 1, 1, 1, 3, 3, 2, 2, 2, 1, 1}, nil),
				vector.Integer([]int{110, 60, 50, 350, 220, 180, 120, 100, 70, 45, 300, 225, 35, 50, 30}, nil),
			},
		},
		{
			name: "three arrangeBy with one column reversed",
			df:   df.Arrange("activity", "type", "salary", vector.OptionArrangeReverseColumns("salary")),
			expectedColumns: []vector.Vector{
				vector.Boolean([]bool{false, false, false, false, false, false, false, true, true, true, true, true, true, true, true}, nil),
				vector.Integer([]int{1, 1, 2, 2, 2, 3, 3, 1, 1, 1, 2, 2, 3, 3, 3}, nil),
				vector.Integer([]int{50, 30, 300, 225, 35, 70, 45, 180, 120, 100, 350, 220, 110, 60, 50}, nil),
			},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if !vector.CompareVectorArrs(data.df.columns, data.expectedColumns) {
				t.Error(fmt.Sprintf("Dataframe columns (%v) are not equal to expected (%v)",
					data.df.columns, data.expectedColumns))
			}
		})
	}
}
