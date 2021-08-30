package dataframe

import (
	"fmt"
	"logarithmotechnia/vector"
	"reflect"
	"testing"
)

func getTestDataFrame() *Dataframe {
	return New([]vector.Vector{
		vector.String([]string{"Jim", "SPARC-001", "Anna", "Lucius", "Maria"}, nil),
		vector.Integer([]int{31, 3, 24, 41, 33}, nil),
		vector.String([]string{"m", "", "f", "m", "f"}, []bool{false, true, false, false, false}),
		vector.Integer([]int{110000, 0, 50000, 120000, 80000}, nil),
		vector.Boolean([]bool{true, true, true, false, true}, nil),
		vector.String([]string{"damage", "heavy", "support", "damage", "healer"}, nil),
	}, OptionColumnNames([]string{"name", "age", "gender", "salary", "active", "-class"}))

}

func TestDataframe_Select(t *testing.T) {
	df := getTestDataFrame()

	testData := []struct {
		name      string
		selectors []interface{}
		columns   []vector.Vector
	}{
		{
			name:      "zero selectors",
			selectors: []interface{}{},
			columns: []vector.Vector{df.columns[0], df.columns[1], df.columns[2], df.columns[3],
				df.columns[4], df.columns[5]},
		},
		{
			name:      "one column with string selector",
			selectors: []interface{}{"name"},
			columns:   []vector.Vector{df.columns[0]},
		},
		{
			name:      "four columns with string selectors",
			selectors: []interface{}{"name", "salary", "gender", "-class"},
			columns:   []vector.Vector{df.columns[0], df.columns[3], df.columns[2], df.columns[5]},
		},
		{
			name:      "four columns with duplicate, string selectors",
			selectors: []interface{}{"name", "salary", "name", "gender", "-class", "salary"},
			columns:   []vector.Vector{df.columns[0], df.columns[3], df.columns[2], df.columns[5]},
		},
		{
			name:      "remove with string selector",
			selectors: []interface{}{"-age"},
			columns:   []vector.Vector{df.columns[0], df.columns[2], df.columns[3], df.columns[4], df.columns[5]},
		},
		{
			name:      "three removes with string selector",
			selectors: []interface{}{"-age", "-gender", "--class"},
			columns:   []vector.Vector{df.columns[0], df.columns[3], df.columns[4]},
		},
		{
			name:      "normal and removal string selectors",
			selectors: []interface{}{"name", "age", "salary", "gender", "-age"},
			columns:   []vector.Vector{df.columns[0], df.columns[3], df.columns[2]},
		},
		{
			name:      "non-existent string selector",
			selectors: []interface{}{"exp"},
			columns:   []vector.Vector{},
		},
		{
			name:      "non-existent removal string selector",
			selectors: []interface{}{"-exp"},
			columns:   []vector.Vector{},
		},
		{
			name:      "string array selector",
			selectors: []interface{}{[]string{"gender", "name", "-class", "age"}},
			columns:   []vector.Vector{df.columns[2], df.columns[0], df.columns[5], df.columns[1]},
		},
		{
			name:      "string array selector three removals",
			selectors: []interface{}{[]string{"-age", "-gender", "--class"}},
			columns:   []vector.Vector{df.columns[0], df.columns[3], df.columns[4]},
		},
		{
			name:      "index selector",
			selectors: []interface{}{1},
			columns:   []vector.Vector{df.columns[0]},
		},
		{
			name:      "non-existent index selector",
			selectors: []interface{}{-1},
			columns:   []vector.Vector{},
		},
		{
			name:      "duplicate index selector",
			selectors: []interface{}{1, 1, 1},
			columns:   []vector.Vector{df.columns[0]},
		},
		{
			name:      "multiple index selectors",
			selectors: []interface{}{1, 2, 5, 6},
			columns:   []vector.Vector{df.columns[0], df.columns[1], df.columns[4], df.columns[5]},
		},
		{
			name:      "multiple index selectors with non-existent and duplicate",
			selectors: []interface{}{1, -1, 2, 0, 5, 10, 1, 6, 5},
			columns:   []vector.Vector{df.columns[0], df.columns[1], df.columns[4], df.columns[5]},
		},
		{
			name:      "string and index selectors combined",
			selectors: []interface{}{[]string{"gender", "name"}, 6, "age"},
			columns:   []vector.Vector{df.columns[2], df.columns[0], df.columns[5], df.columns[1]},
		},
		{
			name:      "boolean selector - full",
			selectors: []interface{}{[]bool{true, true, false, true, false, false}},
			columns:   []vector.Vector{df.columns[0], df.columns[1], df.columns[3]},
		},
		{
			name:      "FromTo regular",
			selectors: []interface{}{FromToColNames{"name", "salary"}},
			columns:   []vector.Vector{df.columns[0], df.columns[1], df.columns[2], df.columns[3]},
		},
		{
			name:      "FromTo names reverse",
			selectors: []interface{}{FromToColNames{"salary", "name"}},
			columns:   []vector.Vector{df.columns[3], df.columns[2], df.columns[1], df.columns[0]},
		},
		{
			name:      "FromTo names incorrect from",
			selectors: []interface{}{FromToColNames{"nam", "salary"}},
			columns:   []vector.Vector{},
		},
		{
			name:      "FromTo names incorrect to",
			selectors: []interface{}{FromToColNames{"name", "salar"}},
			columns:   []vector.Vector{},
		},
		{
			name:      "FromTo indices regular",
			selectors: []interface{}{FromToColIndices{1, 4}},
			columns:   []vector.Vector{df.columns[0], df.columns[1], df.columns[2], df.columns[3]},
		},
		{
			name:      "FromTo indices reverse",
			selectors: []interface{}{FromToColIndices{4, 1}},
			columns:   []vector.Vector{df.columns[3], df.columns[2], df.columns[1], df.columns[0]},
		},
		{
			name:      "FromTo indices incorrect to",
			selectors: []interface{}{FromToColIndices{0, 4}},
			columns:   []vector.Vector{},
		},
		{
			name:      "FromTo indices incorrect from",
			selectors: []interface{}{FromToColIndices{1, 7}},
			columns:   []vector.Vector{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			newDf := df.Select(data.selectors...)

			if !reflect.DeepEqual(newDf.columns, data.columns) {
				t.Error(fmt.Sprintf("Columns (%v) are not equal to expected (%v)",
					newDf.columns, data.columns))
			}
		})
	}

}
