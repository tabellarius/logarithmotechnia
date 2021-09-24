package dataframe

import (
	"fmt"
	"logarithmotechnia/vector"
	"reflect"
	"testing"
)

func TestDataframe_Mutate(t *testing.T) {
	df := New([]Column{
		{"name", vector.StringWithNA([]string{"Jim", "Lucius", "Alice", "Marcus", "Leticia"}, nil)},
		{"salary", vector.IntegerWithNA([]int{100000, 120000, 80000, 70000, 90000}, nil)},
		{"active", vector.BooleanWithNA([]bool{true, false, true, false, true}, nil)},
	})

	mutateColumns := []Column{
		{"city", vector.StringWithNA([]string{"London", "Rome", "Paris", "New York", "Tokyo"}, nil)},
		{"missions", vector.IntegerWithNA([]int{10, 27, 4, 6, 8}, nil)},
	}

	testData := []struct {
		name     string
		df       *Dataframe
		columns  []Column
		options  []vector.Option
		expected *Dataframe
	}{
		{
			name:    "simple",
			df:      df,
			columns: mutateColumns,
			expected: New([]Column{
				{"name", vector.StringWithNA([]string{"Jim", "Lucius", "Alice", "Marcus", "Leticia"}, nil)},
				{"salary", vector.IntegerWithNA([]int{100000, 120000, 80000, 70000, 90000}, nil)},
				{"active", vector.BooleanWithNA([]bool{true, false, true, false, true}, nil)},
				{"city", vector.StringWithNA([]string{"London", "Rome", "Paris", "New York", "Tokyo"}, nil)},
				{"missions", vector.IntegerWithNA([]int{10, 27, 4, 6, 8}, nil)},
			}),
		},
		{
			name:    "after column",
			df:      df,
			columns: mutateColumns,
			options: []vector.Option{vector.OptionAfterColumn("name")},
			expected: New([]Column{
				{"name", vector.StringWithNA([]string{"Jim", "Lucius", "Alice", "Marcus", "Leticia"}, nil)},
				{"city", vector.StringWithNA([]string{"London", "Rome", "Paris", "New York", "Tokyo"}, nil)},
				{"missions", vector.IntegerWithNA([]int{10, 27, 4, 6, 8}, nil)},
				{"salary", vector.IntegerWithNA([]int{100000, 120000, 80000, 70000, 90000}, nil)},
				{"active", vector.BooleanWithNA([]bool{true, false, true, false, true}, nil)},
			}),
		},
		{
			name:    "before column",
			df:      df,
			columns: mutateColumns,
			options: []vector.Option{vector.OptionBeforeColumn("salary")},
			expected: New([]Column{
				{"name", vector.StringWithNA([]string{"Jim", "Lucius", "Alice", "Marcus", "Leticia"}, nil)},
				{"city", vector.StringWithNA([]string{"London", "Rome", "Paris", "New York", "Tokyo"}, nil)},
				{"missions", vector.IntegerWithNA([]int{10, 27, 4, 6, 8}, nil)},
				{"salary", vector.IntegerWithNA([]int{100000, 120000, 80000, 70000, 90000}, nil)},
				{"active", vector.BooleanWithNA([]bool{true, false, true, false, true}, nil)},
			}),
		},
		{
			name:    "before first column",
			df:      df,
			columns: mutateColumns,
			options: []vector.Option{vector.OptionBeforeColumn("name")},
			expected: New([]Column{
				{"city", vector.StringWithNA([]string{"London", "Rome", "Paris", "New York", "Tokyo"}, nil)},
				{"missions", vector.IntegerWithNA([]int{10, 27, 4, 6, 8}, nil)},
				{"name", vector.StringWithNA([]string{"Jim", "Lucius", "Alice", "Marcus", "Leticia"}, nil)},
				{"salary", vector.IntegerWithNA([]int{100000, 120000, 80000, 70000, 90000}, nil)},
				{"active", vector.BooleanWithNA([]bool{true, false, true, false, true}, nil)},
			}),
		},
		{
			name:    "after non-existant column",
			df:      df,
			columns: mutateColumns,
			options: []vector.Option{vector.OptionAfterColumn("pods")},
			expected: New([]Column{
				{"name", vector.StringWithNA([]string{"Jim", "Lucius", "Alice", "Marcus", "Leticia"}, nil)},
				{"salary", vector.IntegerWithNA([]int{100000, 120000, 80000, 70000, 90000}, nil)},
				{"active", vector.BooleanWithNA([]bool{true, false, true, false, true}, nil)},
				{"city", vector.StringWithNA([]string{"London", "Rome", "Paris", "New York", "Tokyo"}, nil)},
				{"missions", vector.IntegerWithNA([]int{10, 27, 4, 6, 8}, nil)},
			}),
		},
		{
			name:    "before non-existant column",
			df:      df,
			columns: mutateColumns,
			options: []vector.Option{vector.OptionBeforeColumn("pods")},
			expected: New([]Column{
				{"name", vector.StringWithNA([]string{"Jim", "Lucius", "Alice", "Marcus", "Leticia"}, nil)},
				{"salary", vector.IntegerWithNA([]int{100000, 120000, 80000, 70000, 90000}, nil)},
				{"active", vector.BooleanWithNA([]bool{true, false, true, false, true}, nil)},
				{"city", vector.StringWithNA([]string{"London", "Rome", "Paris", "New York", "Tokyo"}, nil)},
				{"missions", vector.IntegerWithNA([]int{10, 27, 4, 6, 8}, nil)},
			}),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			newDf := data.df.Mutate(data.columns, data.options...)

			if !vector.CompareVectorArrs(newDf.columns, data.expected.columns) {
				t.Error(fmt.Sprintf("Columns (%v) are not equal to expected (%v)",
					newDf.columns, data.expected.columns))
			}

			if !reflect.DeepEqual(newDf.columnNames, data.expected.columnNames) {
				t.Error(fmt.Sprintf("Column names (%v) are not equal to expected (%v)",
					newDf.columns, data.expected.columns))
			}
		})
	}
}
