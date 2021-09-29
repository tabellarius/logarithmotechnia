package dataframe

import (
	"fmt"
	"logarithmotechnia/vector"
	"testing"
)

func TestDataframe_Sum(t *testing.T) {
	df := New([]Column{
		{"A", vector.Integer([]int{100, 200, 200, 30, 30})},
		{"B", vector.IntegerWithNA([]int{100, 100, 40, 30, 40}, []bool{false, true, true, true, false})},
		{"C", vector.Boolean([]bool{true, false, true, false, true})},
		{"D", vector.String([]string{"1", "2", "3", "4", "5"})},
	})

	testData := []struct {
		name    string
		df      *Dataframe
		columns []vector.Vector
	}{
		{
			name: "normal",
			df:   df.Sum(),
			columns: []vector.Vector{
				vector.Integer([]int{560}),
				vector.IntegerWithNA([]int{0}, []bool{true}),
				vector.Integer([]int{3}),
				vector.NA(1),
			},
		},
		{
			name: "grouped",
			df:   df.GroupBy("A").Sum(),
			columns: []vector.Vector{
				vector.Integer([]int{100, 400, 60}),
				vector.IntegerWithNA([]int{0}, []bool{true}),
				vector.Integer([]int{3}),
				vector.NA(1),
			},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if !vector.CompareVectorArrs(data.df.columns, data.columns) {
				t.Error(fmt.Sprintf("Columns (%v) are not equal to expected (%v)", data.df.columns, data.columns))
			}
		})
	}
}
