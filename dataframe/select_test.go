package dataframe

import (
	"fmt"
	"logarithmotechnia/vector"
	"reflect"
	"testing"
)

func getTestDataFrame() *Dataframe {
	return New([]vector.Vector{
		vector.StringWithNA([]string{"Jim", "SPARC-001", "Anna", "Lucius", "Maria"}, nil),
		vector.IntegerWithNA([]int{31, 3, 24, 41, 33}, nil),
		vector.StringWithNA([]string{"m", "", "f", "m", "f"}, []bool{false, true, false, false, false}),
		vector.IntegerWithNA([]int{110000, 0, 50000, 120000, 80000}, nil),
		vector.BooleanWithNA([]bool{true, true, true, false, true}, nil),
		vector.StringWithNA([]string{"damage", "heavy", "support", "damage", "healer"}, nil),
	}, OptionColumnNames([]string{"name", "age", "gender", "salary", "active", "-class"}))

}

func TestDataframe_Select(t *testing.T) {
	df := getTestDataFrame()

	testData := []struct {
		name        string
		selectors   []any
		columns     []vector.Vector
		columnNames []string
	}{
		{
			name:      "zero selectors",
			selectors: []any{},
			columns: []vector.Vector{df.columns[0], df.columns[1], df.columns[2], df.columns[3],
				df.columns[4], df.columns[5]},
			columnNames: []string{"name", "age", "gender", "salary", "active", "-class"},
		},
		{
			name:        "one column with string selector",
			selectors:   []any{"name"},
			columns:     []vector.Vector{df.columns[0]},
			columnNames: []string{"name"},
		},
		{
			name:        "four columns with string selectors",
			selectors:   []any{"name", "salary", "gender", "-class"},
			columns:     []vector.Vector{df.columns[0], df.columns[3], df.columns[2], df.columns[5]},
			columnNames: []string{"name", "salary", "gender", "-class"},
		},
		{
			name:        "four columns with duplicate, string selectors",
			selectors:   []any{"name", "salary", "name", "gender", "-class", "salary"},
			columns:     []vector.Vector{df.columns[0], df.columns[3], df.columns[2], df.columns[5]},
			columnNames: []string{"name", "salary", "gender", "-class"},
		},
		{
			name:        "remove with string selector",
			selectors:   []any{"-age"},
			columns:     []vector.Vector{df.columns[0], df.columns[2], df.columns[3], df.columns[4], df.columns[5]},
			columnNames: []string{"name", "gender", "salary", "active", "-class"},
		},
		{
			name:        "three removes with string selector",
			selectors:   []any{"-age", "-gender", "--class"},
			columns:     []vector.Vector{df.columns[0], df.columns[3], df.columns[4]},
			columnNames: []string{"name", "salary", "active"},
		},
		{
			name:        "normal and removal string selectors",
			selectors:   []any{"name", "age", "salary", "gender", "-age"},
			columns:     []vector.Vector{df.columns[0], df.columns[3], df.columns[2]},
			columnNames: []string{"name", "salary", "gender"},
		},
		{
			name:        "non-existent string selector",
			selectors:   []any{"exp"},
			columns:     []vector.Vector{},
			columnNames: []string{},
		},
		{
			name:        "non-existent removal string selector",
			selectors:   []any{"-exp"},
			columns:     []vector.Vector{},
			columnNames: []string{},
		},
		{
			name:        "string array selector",
			selectors:   []any{[]string{"gender", "name", "-class", "age"}},
			columns:     []vector.Vector{df.columns[2], df.columns[0], df.columns[5], df.columns[1]},
			columnNames: []string{"gender", "name", "-class", "age"},
		},
		{
			name:        "string array selector three removals",
			selectors:   []any{[]string{"-age", "-gender", "--class"}},
			columns:     []vector.Vector{df.columns[0], df.columns[3], df.columns[4]},
			columnNames: []string{"name", "salary", "active"},
		},
		{
			name:        "index selector",
			selectors:   []any{1},
			columns:     []vector.Vector{df.columns[0]},
			columnNames: []string{"name"},
		},
		{
			name:        "non-existent index selector",
			selectors:   []any{-1},
			columns:     []vector.Vector{},
			columnNames: []string{},
		},
		{
			name:        "duplicate index selector",
			selectors:   []any{1, 1, 1},
			columns:     []vector.Vector{df.columns[0]},
			columnNames: []string{"name"},
		},
		{
			name:        "multiple index selectors",
			selectors:   []any{1, 2, 5, 6},
			columns:     []vector.Vector{df.columns[0], df.columns[1], df.columns[4], df.columns[5]},
			columnNames: []string{"name", "age", "active", "-class"},
		},
		{
			name:        "multiple index selectors with non-existent and duplicate",
			selectors:   []any{1, -1, 2, 0, 5, 10, 1, 6, 5},
			columns:     []vector.Vector{df.columns[0], df.columns[1], df.columns[4], df.columns[5]},
			columnNames: []string{"name", "age", "active", "-class"},
		},
		{
			name:        "string and index selectors combined",
			selectors:   []any{[]string{"gender", "name"}, 6, "age"},
			columns:     []vector.Vector{df.columns[2], df.columns[0], df.columns[5], df.columns[1]},
			columnNames: []string{"gender", "name", "-class", "age"},
		},
		{
			name:        "boolean selector - full",
			selectors:   []any{[]bool{true, true, false, true, false, false}},
			columns:     []vector.Vector{df.columns[0], df.columns[1], df.columns[3]},
			columnNames: []string{"name", "age", "salary"},
		},
		{
			name:        "FromTo regular",
			selectors:   []any{FromToColNames{"name", "salary"}},
			columns:     []vector.Vector{df.columns[0], df.columns[1], df.columns[2], df.columns[3]},
			columnNames: []string{"name", "age", "gender", "salary"},
		},
		{
			name:        "FromTo names reverse",
			selectors:   []any{FromToColNames{"salary", "name"}},
			columns:     []vector.Vector{df.columns[3], df.columns[2], df.columns[1], df.columns[0]},
			columnNames: []string{"salary", "gender", "age", "name"},
		},
		{
			name:        "FromTo names incorrect from",
			selectors:   []any{FromToColNames{"nam", "salary"}},
			columns:     []vector.Vector{},
			columnNames: []string{},
		},
		{
			name:        "FromTo names incorrect to",
			selectors:   []any{FromToColNames{"name", "salar"}},
			columns:     []vector.Vector{},
			columnNames: []string{},
		},
		{
			name:        "FromTo indices regular",
			selectors:   []any{FromToColIndices{1, 4}},
			columns:     []vector.Vector{df.columns[0], df.columns[1], df.columns[2], df.columns[3]},
			columnNames: []string{"name", "age", "gender", "salary"},
		},
		{
			name:        "FromTo indices reverse",
			selectors:   []any{FromToColIndices{4, 1}},
			columns:     []vector.Vector{df.columns[3], df.columns[2], df.columns[1], df.columns[0]},
			columnNames: []string{"salary", "gender", "age", "name"},
		},
		{
			name:        "FromTo indices incorrect to",
			selectors:   []any{FromToColIndices{0, 4}},
			columns:     []vector.Vector{},
			columnNames: []string{},
		},
		{
			name:        "FromTo indices incorrect from",
			selectors:   []any{FromToColIndices{1, 7}},
			columns:     []vector.Vector{},
			columnNames: []string{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			newDf := df.Select(data.selectors...)

			if !reflect.DeepEqual(newDf.columns, data.columns) {
				t.Error(fmt.Sprintf("Columns (%v) are not equal to expected (%v)",
					newDf.columns, data.columns))
			}

			if !reflect.DeepEqual(newDf.columnNames, data.columnNames) {
				t.Error(fmt.Sprintf("Column names (%v) are not equal to expected (%v)",
					newDf.columnNames, data.columnNames))
			}
		})
	}

}
