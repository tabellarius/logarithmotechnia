package dataframe

import (
	"fmt"
	"logarithmotechnia/vector"
	"testing"
)

func TestDataframe_Relocate(t *testing.T) {
	df := getTestDataFrame()

	testData := []struct {
		name      string
		selectors []interface{}
		options   []interface{}
		columns   []vector.Vector
	}{
		{
			name:      "string",
			selectors: []interface{}{"name"},
			columns: []vector.Vector{
				vector.IntegerWithNA([]int{31, 3, 24, 41, 33}, nil),
				vector.StringWithNA([]string{"m", "", "f", "m", "f"}, []bool{false, true, false, false, false}),
				vector.IntegerWithNA([]int{110000, 0, 50000, 120000, 80000}, nil),
				vector.BooleanWithNA([]bool{true, true, true, false, true}, nil),
				vector.StringWithNA([]string{"damage", "heavy", "support", "damage", "healer"}, nil),
				vector.StringWithNA([]string{"Jim", "SPARC-001", "Anna", "Lucius", "Maria"}, nil),
			},
		},
		{
			name:      "[]string",
			selectors: []interface{}{[]string{"name", "gender"}},
			columns: []vector.Vector{
				vector.IntegerWithNA([]int{31, 3, 24, 41, 33}, nil),
				vector.IntegerWithNA([]int{110000, 0, 50000, 120000, 80000}, nil),
				vector.BooleanWithNA([]bool{true, true, true, false, true}, nil),
				vector.StringWithNA([]string{"damage", "heavy", "support", "damage", "healer"}, nil),
				vector.StringWithNA([]string{"Jim", "SPARC-001", "Anna", "Lucius", "Maria"}, nil),
				vector.StringWithNA([]string{"m", "", "f", "m", "f"}, []bool{false, true, false, false, false}),
			},
		},
		{
			name:      "string + []string",
			selectors: []interface{}{"salary", []string{"name", "gender"}},
			columns: []vector.Vector{
				vector.IntegerWithNA([]int{31, 3, 24, 41, 33}, nil),
				vector.BooleanWithNA([]bool{true, true, true, false, true}, nil),
				vector.StringWithNA([]string{"damage", "heavy", "support", "damage", "healer"}, nil),
				vector.IntegerWithNA([]int{110000, 0, 50000, 120000, 80000}, nil),
				vector.StringWithNA([]string{"Jim", "SPARC-001", "Anna", "Lucius", "Maria"}, nil),
				vector.StringWithNA([]string{"m", "", "f", "m", "f"}, []bool{false, true, false, false, false}),
			},
		},
		{
			name:      "string + []string + after column",
			selectors: []interface{}{"salary", []string{"name", "gender"}},
			options:   []interface{}{vector.OptionAfterColumn("active")},
			columns: []vector.Vector{
				vector.IntegerWithNA([]int{31, 3, 24, 41, 33}, nil),
				vector.BooleanWithNA([]bool{true, true, true, false, true}, nil),
				vector.IntegerWithNA([]int{110000, 0, 50000, 120000, 80000}, nil),
				vector.StringWithNA([]string{"Jim", "SPARC-001", "Anna", "Lucius", "Maria"}, nil),
				vector.StringWithNA([]string{"m", "", "f", "m", "f"}, []bool{false, true, false, false, false}),
				vector.StringWithNA([]string{"damage", "heavy", "support", "damage", "healer"}, nil),
			},
		},
		{
			name:      "string + []string + before column",
			selectors: []interface{}{"salary", []string{"name", "gender"}},
			options:   []interface{}{vector.OptionBeforeColumn("active")},
			columns: []vector.Vector{
				vector.IntegerWithNA([]int{31, 3, 24, 41, 33}, nil),
				vector.IntegerWithNA([]int{110000, 0, 50000, 120000, 80000}, nil),
				vector.StringWithNA([]string{"Jim", "SPARC-001", "Anna", "Lucius", "Maria"}, nil),
				vector.StringWithNA([]string{"m", "", "f", "m", "f"}, []bool{false, true, false, false, false}),
				vector.BooleanWithNA([]bool{true, true, true, false, true}, nil),
				vector.StringWithNA([]string{"damage", "heavy", "support", "damage", "healer"}, nil),
			},
		},
		{
			name:      "string",
			selectors: []interface{}{"name"},
			columns: []vector.Vector{
				vector.IntegerWithNA([]int{31, 3, 24, 41, 33}, nil),
				vector.StringWithNA([]string{"m", "", "f", "m", "f"}, []bool{false, true, false, false, false}),
				vector.IntegerWithNA([]int{110000, 0, 50000, 120000, 80000}, nil),
				vector.BooleanWithNA([]bool{true, true, true, false, true}, nil),
				vector.StringWithNA([]string{"damage", "heavy", "support", "damage", "healer"}, nil),
				vector.StringWithNA([]string{"Jim", "SPARC-001", "Anna", "Lucius", "Maria"}, nil),
			},
		},
		{
			name:      "boolean full",
			selectors: []interface{}{[]bool{true, false, true, false, true, false}},
			columns: []vector.Vector{
				vector.IntegerWithNA([]int{31, 3, 24, 41, 33}, nil),
				vector.IntegerWithNA([]int{110000, 0, 50000, 120000, 80000}, nil),
				vector.StringWithNA([]string{"damage", "heavy", "support", "damage", "healer"}, nil),
				vector.StringWithNA([]string{"Jim", "SPARC-001", "Anna", "Lucius", "Maria"}, nil),
				vector.StringWithNA([]string{"m", "", "f", "m", "f"}, []bool{false, true, false, false, false}),
				vector.BooleanWithNA([]bool{true, true, true, false, true}, nil),
			},
		},
		{
			name:      "boolean partial",
			selectors: []interface{}{[]bool{true, false}},
			columns: []vector.Vector{
				vector.IntegerWithNA([]int{31, 3, 24, 41, 33}, nil),
				vector.IntegerWithNA([]int{110000, 0, 50000, 120000, 80000}, nil),
				vector.StringWithNA([]string{"damage", "heavy", "support", "damage", "healer"}, nil),
				vector.StringWithNA([]string{"Jim", "SPARC-001", "Anna", "Lucius", "Maria"}, nil),
				vector.StringWithNA([]string{"m", "", "f", "m", "f"}, []bool{false, true, false, false, false}),
				vector.BooleanWithNA([]bool{true, true, true, false, true}, nil),
			},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			parameters := append(data.selectors, data.options...)
			newDf := df.Relocate(parameters...)

			if !vector.CompareVectorArrs(newDf.columns, data.columns) {
				t.Error(fmt.Sprintf("Columns (%v) are not equal to expected (%v)",
					newDf.columns, data.columns))
			}
		})
	}
}
