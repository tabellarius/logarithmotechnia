package dataframe

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"logarithmotechnia/vector"
	"reflect"
	"testing"
)

func TestFromSQL(t *testing.T) {
	db, err := sql.Open("sqlite3", "./test_data/items.sqlite")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}

	df, err := FromSQL(tx, "SELECT * FROM sku", []any{})

	columnNames := []string{"id", "title", "vendor_id", "price"}
	expectedColumns := []vector.Vector{
		vector.Integer([]int{1, 2, 3, 4, 5}),
		vector.StringWithNA([]string{"Item 1", "Item 2", "Item 3", "", "Item 5"},
			[]bool{false, false, false, true, false}),
		vector.StringWithNA([]string{"VND001-001", "VND001-002", "VND002-001", "VND002-002", ""},
			[]bool{false, false, false, false, true}),
		vector.FloatWithNA([]float64{3050.000, 249.990, 0.000, 1101.100, 150.000},
			[]bool{false, false, true, false, false}),
	}

	if !reflect.DeepEqual(df.columnNames, columnNames) {
		t.Error(fmt.Sprintf("Dataframe column names (%v) are not equal to expected (%v)",
			df.columnNames, columnNames))
	}

	if !vector.CompareVectorArrs(df.columns, expectedColumns) {
		t.Error(fmt.Sprintf("Dataframe columns (%v) are not equal to expected (%v)",
			df.columns, expectedColumns))
	}
}

func TestSQLOptions(t *testing.T) {
	testData := []struct {
		name      string
		result    Option
		reference Option
	}{
		{
			name:   "SQLOptionDataframeOptions",
			result: SQLOptionDataframeOptions(OptionColumnNames([]string{"id", "price"})),
			reference: ConfOption{optionSQLDataframeOptions,
				[]Option{OptionColumnNames([]string{"id", "price"})}},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if !reflect.DeepEqual(data.result, data.reference) {
				t.Error(fmt.Sprintf("Resulting conf option (%v) does not match reference (%v)",
					data.result, data.reference))
			}
		})
	}
}

func TestSQLOptionTransformers(t *testing.T) {
	transformerFn := func(vec vector.Vector) vector.Vector {
		return vec
	}
	option := SQLOptionTransformers(map[string]transformerFunc{"stubFn": transformerFn})
	val := option.Value().(map[string]transformerFunc)

	if option.Key() != optionSQLDataframeTransformers ||
		reflect.ValueOf(val["stubFn"]).Pointer() != reflect.ValueOf(transformerFn).Pointer() {
		t.Error("SQLOptionTransformers() failed")
	}
}
