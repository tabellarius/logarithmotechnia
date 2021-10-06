package dataframe

import (
	"fmt"
	"logarithmotechnia/vector"
	"reflect"
	"testing"
)

func TestDataframe_BindColumns(t *testing.T) {
	df1 := New([]Column{
		{"A", vector.IntegerWithNA([]int{10, 20, 30}, []bool{false, true, true})},
		{"B", vector.Boolean([]bool{true, false, true})},
		{"C", vector.String([]string{"1", "2", "3"})},
	})

	df2 := New([]Column{
		{"E", vector.Boolean([]bool{false, false, true})},
		{"A", vector.IntegerWithNA([]int{1, 2, 3}, []bool{false, false, true})},
	})

	df3 := New([]Column{
		{"C", vector.String([]string{"1", "2"})},
		{"D", vector.Integer([]int{1, 2})},
	})

	testData := []struct {
		name            string
		df              *Dataframe
		expectedColumns []vector.Vector
		columnNames     []string
	}{
		{
			name: "bindColumns empty",
			df:   df1.BindColumns(),
			expectedColumns: []vector.Vector{
				vector.IntegerWithNA([]int{10, 20, 30}, []bool{false, true, true}),
				vector.Boolean([]bool{true, false, true}),
				vector.String([]string{"1", "2", "3"}),
			},
			columnNames: []string{"A", "B", "C"},
		},
		{
			name: "bindColumns",
			df:   df1.BindColumns(df2),
			expectedColumns: []vector.Vector{
				vector.IntegerWithNA([]int{10, 20, 30}, []bool{false, true, true}),
				vector.Boolean([]bool{true, false, true}),
				vector.String([]string{"1", "2", "3"}),
				vector.Boolean([]bool{false, false, true}),
				vector.IntegerWithNA([]int{1, 2, 3}, []bool{false, false, true}),
			},
			columnNames: []string{"A", "B", "C", "E", "A_1"},
		},
		{
			name: "bindColumns 2 args",
			df:   df1.BindColumns(df2, df3),
			expectedColumns: []vector.Vector{
				vector.IntegerWithNA([]int{10, 20, 30}, []bool{false, true, true}),
				vector.Boolean([]bool{true, false, true}),
				vector.String([]string{"1", "2", "3"}),
				vector.Boolean([]bool{false, false, true}),
				vector.IntegerWithNA([]int{1, 2, 3}, []bool{false, false, true}),
				vector.StringWithNA([]string{"1", "2", ""}, []bool{false, false, true}),
				vector.IntegerWithNA([]int{1, 2, 0}, []bool{false, false, true}),
			},
			columnNames: []string{"A", "B", "C", "E", "A_1", "C_1", "D"},
		},
		{
			name: "bindColumns arr",
			df:   df1.BindColumns([]*Dataframe{df2, df3}),
			expectedColumns: []vector.Vector{
				vector.IntegerWithNA([]int{10, 20, 30}, []bool{false, true, true}),
				vector.Boolean([]bool{true, false, true}),
				vector.String([]string{"1", "2", "3"}),
				vector.Boolean([]bool{false, false, true}),
				vector.IntegerWithNA([]int{1, 2, 3}, []bool{false, false, true}),
				vector.StringWithNA([]string{"1", "2", ""}, []bool{false, false, true}),
				vector.IntegerWithNA([]int{1, 2, 0}, []bool{false, false, true}),
			},
			columnNames: []string{"A", "B", "C", "E", "A_1", "C_1", "D"},
		},
		{
			name: "bindColumns to itself",
			df:   df1.BindColumns(df1),
			expectedColumns: []vector.Vector{
				vector.IntegerWithNA([]int{10, 20, 30}, []bool{false, true, true}),
				vector.Boolean([]bool{true, false, true}),
				vector.String([]string{"1", "2", "3"}),
				vector.IntegerWithNA([]int{10, 20, 30}, []bool{false, true, true}),
				vector.Boolean([]bool{true, false, true}),
				vector.String([]string{"1", "2", "3"}),
			},
			columnNames: []string{"A", "B", "C", "A_1", "B_1", "C_1"},
		},
		{
			name: "bindColumns to itself + other",
			df:   df1.BindColumns(df1, df2),
			expectedColumns: []vector.Vector{
				vector.IntegerWithNA([]int{10, 20, 30}, []bool{false, true, true}),
				vector.Boolean([]bool{true, false, true}),
				vector.String([]string{"1", "2", "3"}),
				vector.IntegerWithNA([]int{10, 20, 30}, []bool{false, true, true}),
				vector.Boolean([]bool{true, false, true}),
				vector.String([]string{"1", "2", "3"}),
				vector.Boolean([]bool{false, false, true}),
				vector.IntegerWithNA([]int{1, 2, 3}, []bool{false, false, true}),
			},
			columnNames: []string{"A", "B", "C", "A_1", "B_1", "C_1", "E", "A_2"},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if !vector.CompareVectorArrs(data.df.columns, data.expectedColumns) {
				t.Error(fmt.Sprintf("Dataframe columns (%v) are not equal to expected (%v)",
					data.df.columns, data.expectedColumns))
			}

			if !reflect.DeepEqual(data.df.columnNames, data.columnNames) {
				t.Error(fmt.Sprintf("Dataframe column names (%v) are not equal to expected (%v)",
					data.df.columnNames, data.columnNames))
			}
		})
	}
}

func TestDataframe_BindRows(t *testing.T) {
	df1 := New([]Column{
		{"A", vector.IntegerWithNA([]int{10, 20, 30}, []bool{false, true, true})},
		{"B", vector.Boolean([]bool{true, false, true})},
		{"C", vector.String([]string{"1", "2", "3"})},
	})

	df2 := New([]Column{
		{"E", vector.Boolean([]bool{false, false, true})},
		{"C", vector.String([]string{"4", "5"})},
		{"A", vector.IntegerWithNA([]int{1, 2, 3}, []bool{false, false, true})},
	})

	df3 := New([]Column{
		{"C", vector.String([]string{"alpha", "omega"})},
		{"D", vector.Integer([]int{1, 2})},
	})

	testData := []struct {
		name            string
		df              *Dataframe
		expectedColumns []vector.Vector
	}{
		{
			name: "bindRows empty",
			df:   df1.BindColumns(),
			expectedColumns: []vector.Vector{
				vector.IntegerWithNA([]int{10, 20, 30}, []bool{false, true, true}),
				vector.Boolean([]bool{true, false, true}),
				vector.String([]string{"1", "2", "3"}),
			},
		},
		{
			name: "bindRows 1 arg",
			df:   df1.BindRows(df2),
			expectedColumns: []vector.Vector{
				vector.IntegerWithNA([]int{10, 0, 0, 1, 2, 0}, []bool{false, true, true, false, false, true}),
				vector.BooleanWithNA([]bool{true, false, true, false, false, false},
					[]bool{false, false, false, true, true, true}),
				vector.StringWithNA([]string{"1", "2", "3", "4", "5", ""}, []bool{false, false, false, false, false, true}),
			},
		},
		{
			name: "bindRows 2 args",
			df:   df1.BindRows(df2, df3),
			expectedColumns: []vector.Vector{
				vector.IntegerWithNA([]int{10, 0, 0, 1, 2, 0, 0, 0},
					[]bool{false, true, true, false, false, true, true, true}),
				vector.BooleanWithNA([]bool{true, false, true, false, false, false, false, false},
					[]bool{false, false, false, true, true, true, true, true}),
				vector.StringWithNA([]string{"1", "2", "3", "4", "5", "", "alpha", "omega"},
					[]bool{false, false, false, false, false, true, false, false}),
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
