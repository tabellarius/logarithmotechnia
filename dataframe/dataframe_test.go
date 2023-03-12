package dataframe

import (
	"fmt"
	"logarithmotechnia/vector"
	"reflect"
	"strconv"
	"testing"
)

func TestNew(t *testing.T) {
	testData := []struct {
		name          string
		columns       any
		columnNames   []string
		dfColumns     []vector.Vector
		dfColumnNames []string
	}{
		{
			name:          "empty",
			columns:       []vector.Vector{},
			columnNames:   []string{},
			dfColumns:     []vector.Vector{},
			dfColumnNames: []string{},
		},
		{
			name:          "empty with column names",
			columns:       []vector.Vector{},
			columnNames:   []string{"one", "two", "three"},
			dfColumns:     []vector.Vector{},
			dfColumnNames: []string{},
		},
		{
			name:          "incorrect data type",
			columns:       []int{1, 2, 3},
			columnNames:   []string{},
			dfColumns:     []vector.Vector{},
			dfColumnNames: []string{},
		},
		{
			name: "normal",
			columns: []vector.Vector{
				vector.IntegerWithNA([]int{1, 2, 3}, nil),
				vector.StringWithNA([]string{"1", "2", "3"}, nil),
				vector.BooleanWithNA([]bool{true, true, false}, nil),
			},
			dfColumns: []vector.Vector{
				vector.IntegerWithNA([]int{1, 2, 3}, nil),
				vector.StringWithNA([]string{"1", "2", "3"}, nil),
				vector.BooleanWithNA([]bool{true, true, false}, nil),
			},
			dfColumnNames: []string{"1", "2", "3"},
		},
		{
			name: "normal with column names",
			columns: []vector.Vector{
				vector.IntegerWithNA([]int{1, 2, 3}, nil),
				vector.StringWithNA([]string{"1", "2", "3"}, nil),
				vector.BooleanWithNA([]bool{true, true, false}, nil),
			},
			columnNames: []string{"int", "string", "bool"},
			dfColumns: []vector.Vector{
				vector.IntegerWithNA([]int{1, 2, 3}, nil),
				vector.StringWithNA([]string{"1", "2", "3"}, nil),
				vector.BooleanWithNA([]bool{true, true, false}, nil),
			},
			dfColumnNames: []string{"int", "string", "bool"},
		},
		{
			name: "normal with partial column names",
			columns: []vector.Vector{
				vector.IntegerWithNA([]int{1, 2, 3}, nil),
				vector.StringWithNA([]string{"1", "2", "3"}, nil),
				vector.BooleanWithNA([]bool{true, true, false}, nil),
			},
			columnNames: []string{"int", "string"},
			dfColumns: []vector.Vector{
				vector.IntegerWithNA([]int{1, 2, 3}, nil),
				vector.StringWithNA([]string{"1", "2", "3"}, nil),
				vector.BooleanWithNA([]bool{true, true, false}, nil),
			},
			dfColumnNames: []string{"int", "string", "3"},
		},
		{
			name: "different columns' length",
			columns: []vector.Vector{
				vector.IntegerWithNA([]int{1, 2, 3, 4, 5}, nil),
				vector.StringWithNA([]string{"1", "2", "3"}, nil),
				vector.BooleanWithNA([]bool{true, true}, nil),
			},
			dfColumns: []vector.Vector{
				vector.IntegerWithNA([]int{1, 2, 3, 4, 5}, nil),
				vector.StringWithNA([]string{"1", "2", "3", "", ""}, []bool{false, false, false, true, true}),
				vector.BooleanWithNA([]bool{true, true, false, false, false}, []bool{false, false, true, true, true}),
			},
			dfColumnNames: []string{"1", "2", "3"},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			var df *Dataframe
			if data.columnNames != nil {
				df = New(data.columns, OptionColumnNames(data.columnNames))
			} else {
				df = New(data.columns)
			}

			if !vector.CompareVectorArrs(df.columns, data.dfColumns) {
				t.Error(fmt.Sprintf("Columns (%v) are not equal to expected (%v)", df.columns, data.dfColumns))
			}

			if !reflect.DeepEqual(df.columnNames, data.dfColumnNames) {
				t.Error(fmt.Sprintf("Column names (%v) are not equal to expected (%v)",
					df.columnNames, data.dfColumnNames))
			}

		})
	}
}

func TestDataframe_ByIndices(t *testing.T) {
	df := New([]vector.Vector{
		vector.IntegerWithNA([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil),
		vector.StringWithNA([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}, nil),
		vector.BooleanWithNA([]bool{true, false, true, false, true, false, true, false, true, false}, nil),
	})

	testData := []struct {
		name      string
		indices   []int
		dfColumns []vector.Vector
	}{
		{
			name:    "normal",
			indices: []int{1, 3, 5, 8, 10},
			dfColumns: []vector.Vector{
				vector.IntegerWithNA([]int{1, 3, 5, 8, 10}, nil),
				vector.StringWithNA([]string{"1", "3", "5", "8", "10"}, nil),
				vector.BooleanWithNA([]bool{true, true, true, false, false}, nil),
			},
		},
		{
			name:    "with invalid",
			indices: []int{-1, 0, 1, 3, 5, 8, 10, 11, 100},
			dfColumns: []vector.Vector{
				vector.IntegerWithNA([]int{0, 1, 3, 5, 8, 10}, []bool{true, false, false, false, false, false}),
				vector.StringWithNA([]string{"", "1", "3", "5", "8", "10"}, []bool{true, false, false, false, false, false}),
				vector.BooleanWithNA([]bool{false, true, true, true, false, false}, []bool{true, false, false, false, false, false}),
			},
		},
		{
			name:    "reverse",
			indices: []int{10, 8, 5, 3, 1},
			dfColumns: []vector.Vector{
				vector.IntegerWithNA([]int{10, 8, 5, 3, 1}, nil),
				vector.StringWithNA([]string{"10", "8", "5", "3", "1"}, nil),
				vector.BooleanWithNA([]bool{false, false, true, true, true}, nil),
			},
		},
		{
			name:    "empty",
			indices: []int{},
			dfColumns: []vector.Vector{
				vector.IntegerWithNA([]int{}, nil),
				vector.StringWithNA([]string{}, nil),
				vector.BooleanWithNA([]bool{}, nil),
			},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			newDf := df.ByIndices(data.indices)

			if !vector.CompareVectorArrs(newDf.columns, data.dfColumns) {
				t.Error(fmt.Sprintf("Columns (%v) are not equal to expected (%v)", newDf.columns, data.dfColumns))
			}
		})
	}
}

func TestDataframe_ColNum(t *testing.T) {
	df := New([]vector.Vector{
		vector.IntegerWithNA([]int{1, 2, 3, 4, 5}, nil),
		vector.StringWithNA([]string{"1", "2", "3"}, nil),
		vector.BooleanWithNA([]bool{true, true}, nil),
	})

	if df.ColNum() != 3 {
		t.Error("Column number is incorrect!")
	}
}

func TestDataframe_RowNum(t *testing.T) {
	df := New([]vector.Vector{
		vector.IntegerWithNA([]int{1, 2, 3, 4, 5}, nil),
		vector.StringWithNA([]string{"1", "2", "3"}, nil),
		vector.BooleanWithNA([]bool{true, true}, nil),
	})

	if df.RowNum() != 5 {
		t.Error("Row number is incorrect!")
	}
}

func TestDataframe_Ci(t *testing.T) {
	df := New([]vector.Vector{
		vector.IntegerWithNA([]int{1, 2, 3, 4, 5}, nil),
		vector.StringWithNA([]string{"1", "2", "3"}, nil),
		vector.BooleanWithNA([]bool{true, true}, nil),
	}, OptionColumnNames([]string{"int", "string", "bool"}))

	testData := []struct {
		index  int
		column vector.Vector
	}{
		{1, vector.IntegerWithNA([]int{1, 2, 3, 4, 5}, nil)},
		{2, vector.StringWithNA([]string{"1", "2", "3", "", ""}, []bool{false, false, false, true, true})},
		{3, vector.BooleanWithNA([]bool{true, true, false, false, false}, []bool{false, false, true, true, true})},
		{0, nil},
		{-1, nil},
		{4, nil},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			column := df.Ci(data.index)

			if !vector.CompareVectorsForTest(column, data.column) {
				t.Error(fmt.Sprintf("Columns (%v) are not equal to expected (%v)", column, data.column))
			}
		})
	}
}

func TestDataframe_Cn(t *testing.T) {
	df := New([]vector.Vector{
		vector.IntegerWithNA([]int{1, 2, 3, 4, 5}, nil),
		vector.StringWithNA([]string{"1", "2", "3"}, nil),
		vector.BooleanWithNA([]bool{true, true}, nil),
	}, OptionColumnNames([]string{"int", "string", "bool"}))

	testData := []struct {
		name   string
		column vector.Vector
	}{
		{"int", vector.IntegerWithNA([]int{1, 2, 3, 4, 5}, nil)},
		{"string", vector.StringWithNA([]string{"1", "2", "3", "", ""}, []bool{false, false, false, true, true})},
		{"bool", vector.BooleanWithNA([]bool{true, true, false, false, false}, []bool{false, false, true, true, true})},
		{"", nil},
		{"some", nil},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			column := df.Cn(data.name)

			if !vector.CompareVectorsForTest(column, data.column) {
				t.Error(fmt.Sprintf("Columns (%v) are not equal to expected (%v)", column, data.column))
			}
		})
	}
}

func TestDataframe_C(t *testing.T) {
	df := New([]vector.Vector{
		vector.IntegerWithNA([]int{1, 2, 3}, nil),
		vector.StringWithNA([]string{"1", "2", "3"}, nil),
		vector.BooleanWithNA([]bool{true, false, true}, nil),
	}, OptionColumnNames([]string{"int", "string", "bool"}))

	testData := []struct {
		selector any
		column   vector.Vector
	}{
		{"int", vector.IntegerWithNA([]int{1, 2, 3}, nil)},
		{2, vector.StringWithNA([]string{"1", "2", "3"}, nil)},
		{0, nil},
		{4, nil},
		{"some", nil},
		{2 + 2i, nil},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			column := df.C(data.selector)

			if !vector.CompareVectorsForTest(column, data.column) {
				t.Error(fmt.Sprintf("Columns (%v) are not equal to expected (%v)", column, data.column))
			}
		})
	}
}

func TestDataframe_NamesAsStrings(t *testing.T) {
	testData := []struct {
		name        string
		dataframe   *Dataframe
		columnNames []string
	}{
		{
			name: "with names",
			dataframe: New([]vector.Vector{
				vector.IntegerWithNA([]int{1, 2, 3}, nil),
				vector.StringWithNA([]string{"1", "2", "3"}, nil),
				vector.BooleanWithNA([]bool{true, false, true}, nil),
			}, OptionColumnNames([]string{"int", "string", "bool"})),
			columnNames: []string{"int", "string", "bool"},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			columnNames := data.dataframe.NamesAsStrings()

			if !reflect.DeepEqual(columnNames, data.columnNames) {
				t.Error(fmt.Sprintf("Columns names (%v) are not equal to expected (%v)",
					columnNames, data.columnNames))
			}
		})
	}
}

func TestDataframe_IsEmpty(t *testing.T) {
	testData := []struct {
		name      string
		dataframe *Dataframe
		isEmpty   bool
	}{
		{
			name:      "empty",
			dataframe: New([]vector.Vector{}),
			isEmpty:   true,
		},
		{
			name: "non-empty",
			dataframe: New([]vector.Vector{
				vector.IntegerWithNA([]int{1, 2, 3}, nil),
				vector.StringWithNA([]string{"1", "2", "3"}, nil),
				vector.BooleanWithNA([]bool{true, false, true}, nil),
			}, OptionColumnNames([]string{"int", "string", "bool"})),
			isEmpty: false,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			isEmpty := data.dataframe.IsEmpty()

			if isEmpty != data.isEmpty {
				t.Error(fmt.Sprintf("IsEmpty (%t) is not equal to expected (%t)",
					isEmpty, data.isEmpty))
			}
		})
	}
}

func TestDataframe_Clone(t *testing.T) {
	testData := []struct {
		name      string
		dataframe *Dataframe
		isEmpty   bool
	}{
		{
			name:      "empty",
			dataframe: New([]vector.Vector{}),
			isEmpty:   true,
		},
		{
			name: "non-empty",
			dataframe: New([]vector.Vector{
				vector.IntegerWithNA([]int{1, 2, 3}, nil),
				vector.StringWithNA([]string{"1", "2", "3"}, nil),
				vector.BooleanWithNA([]bool{true, false, true}, nil),
			}, OptionColumnNames([]string{"int", "string", "bool"})),
			isEmpty: false,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			newDf := data.dataframe.Clone()

			if !reflect.DeepEqual(newDf.columns, data.dataframe.columns) {
				t.Error(fmt.Sprintf("Columns (%v) are not equal to expected (%v)",
					newDf.columns, data.dataframe.columns))
			}
			if !reflect.DeepEqual(newDf.columnNames, data.dataframe.columnNames) {
				t.Error(fmt.Sprintf("Column names (%v) is not equal to expected (%v)",
					newDf.columnNames, data.dataframe.columnNames))
			}
		})
	}
}

func TestDataframe_Columns(t *testing.T) {
	columns := []vector.Vector{
		vector.IntegerWithNA([]int{1, 2, 3}, nil),
		vector.StringWithNA([]string{"1", "2", "3"}, nil),
		vector.BooleanWithNA([]bool{true, false, true}, nil),
	}

	df := New(columns)

	dfColumns := df.Columns()

	if !reflect.DeepEqual(columns, dfColumns) {
		t.Error(fmt.Sprintf("Columns (%v) are not equal to expected (%v)",
			columns, dfColumns))
	}

}

func TestDataframe_HasColumn(t *testing.T) {
	df := New([]vector.Vector{
		vector.IntegerWithNA([]int{1, 2, 3}, nil),
		vector.StringWithNA([]string{"1", "2", "3"}, nil),
		vector.BooleanWithNA([]bool{true, false, true}, nil),
	}, OptionColumnNames([]string{"int", "string", "bool"}))

	testData := []struct {
		name      string
		column    string
		hasColumn bool
	}{
		{
			name:      "exists 1",
			column:    "int",
			hasColumn: true,
		},
		{
			name:      "exists 2",
			column:    "string",
			hasColumn: true,
		},
		{
			name:      "not exists",
			column:    "float",
			hasColumn: false,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			hasColumn := df.HasColumn(data.column)

			if hasColumn != data.hasColumn {
				t.Error(fmt.Sprintf("hasColumn (%v) are not equal to expected (%v)",
					hasColumn, data.hasColumn))
			}
		})
	}
}

func TestDataframe_IsValidColumnIndex(t *testing.T) {
	df := New([]vector.Vector{
		vector.IntegerWithNA([]int{1, 2, 3}, nil),
		vector.StringWithNA([]string{"1", "2", "3"}, nil),
		vector.BooleanWithNA([]bool{true, false, true}, nil),
	}, OptionColumnNames([]string{"int", "string", "bool"}))

	testData := []struct {
		name  string
		index int
		valid bool
	}{
		{"minimum", 1, true},
		{"maximum", 3, true},
		{"zero", 0, false},
		{"overflow", 4, false},
		{"negative", -1, false},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			valid := df.IsValidColumnIndex(data.index)

			if valid != data.valid {
				t.Error(fmt.Sprintf("hasColumn (%v) are not equal to expected (%v)",
					valid, data.valid))
			}
		})
	}
}

func TestDataframe_Pick(t *testing.T) {
	df := New([]vector.Vector{
		vector.IntegerWithNA([]int{1, 2, 3}, nil),
		vector.StringWithNA([]string{"1", "2", "3"}, nil),
		vector.BooleanWithNA([]bool{true, false, true}, nil),
	}, OptionColumnNames([]string{"int", "string", "bool"}))

	testData := []struct {
		index  int
		rowMap map[string]any
	}{
		{
			index:  1,
			rowMap: map[string]any{"int": 1, "string": "1", "bool": true},
		},
		{
			index:  2,
			rowMap: map[string]any{"int": 2, "string": "2", "bool": false},
		},
		{
			index:  3,
			rowMap: map[string]any{"int": 3, "string": "3", "bool": true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			rowMap := df.Pick(data.index)

			if !reflect.DeepEqual(rowMap, data.rowMap) {
				t.Error(fmt.Sprintf("Map (%v) is not equal to expected (%v)",
					rowMap, data.rowMap))
			}
		})
	}
}

func TestDataframe_Traverse(t *testing.T) {
	df := New([]vector.Vector{
		vector.IntegerWithNA([]int{1, 2, 3}, nil),
		vector.StringWithNA([]string{"1", "2", "3"}, nil),
		vector.BooleanWithNA([]bool{true, false, true}, nil),
	}, OptionColumnNames([]string{"int", "string", "bool"}))

	fMapDf := map[int]map[string]any{}
	trF := func(idx int, val map[string]any) {
		fMapDf[idx] = val
	}
	df.Traverse(trF)
	refFMapDf := map[int]map[string]any{
		1: {"int": 1, "string": "1", "bool": true},
		2: {"int": 2, "string": "2", "bool": false},
		3: {"int": 3, "string": "3", "bool": true},
	}

	if !reflect.DeepEqual(fMapDf, refFMapDf) {
		t.Error(fmt.Sprintf("fMapDf %v is not equal to %v", fMapDf, refFMapDf))
	}

	arrMapDf := []map[string]any{}
	trC := func(val map[string]any) {
		arrMapDf = append(arrMapDf, val)
	}
	df.Traverse(trC)
	refArrMapDf := []map[string]any{
		{"int": 1, "string": "1", "bool": true},
		{"int": 2, "string": "2", "bool": false},
		{"int": 3, "string": "3", "bool": true},
	}

	if !reflect.DeepEqual(arrMapDf, refArrMapDf) {
		t.Error(fmt.Sprintf("arrMapDf %v is not equal to %v", arrMapDf, refArrMapDf))
	}
}

func TestDataframe_ToMaps(t *testing.T) {
	testData := []struct {
		name      string
		dataframe *Dataframe
		dataMap   map[string][]any
	}{
		{
			name: "normal",
			dataframe: New([]vector.Vector{
				vector.IntegerWithNA([]int{1, 2, 3}, nil),
				vector.StringWithNA([]string{"1", "2", "3"}, []bool{true, false, true}),
				vector.BooleanWithNA([]bool{true, false, true}, nil),
			}, OptionColumnNames([]string{"int", "string", "bool"})),
			dataMap: map[string][]any{"int": {1, 2, 3}, "string": {nil, "2", nil}, "bool": {true, false, true}},
		},
		{
			name:      "empty",
			dataframe: New([]vector.Vector{}),
			dataMap:   map[string][]any{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			dataMap := data.dataframe.ToMap()

			if !reflect.DeepEqual(dataMap, data.dataMap) {
				t.Error(fmt.Sprintf("Data map (%v) is not equal to expected (%v)",
					dataMap, data.dataMap))
			}
		})
	}
}

func TestDataframe_String(t *testing.T) {
	df := New([]vector.Vector{
		vector.IntegerWithNA([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nil),
		vector.StringWithNA([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}, nil),
		vector.BooleanWithNA([]bool{true, false, true, false, true, false, true, false, true, false}, nil),
	}, OptionColumnNames([]string{"id", "str_id", "is_present"}), OptionVectorOptions([]vector.Option{
		vector.OptionMaxPrintElements(5),
	}))

	vecStr := "# of columns: 3, # of rows: 10\n\nid: [(integer)]1, 2, 3, 4, 5, ...]\nstr_id: [(string)]1, 2, 3, 4, 5, ...]\n" +
		"is_present: [(boolean)]true, false, true, false, true, ...]\n"

	if df.String() != vecStr {
		t.Error("Dataframe String() failed")
	}
}
