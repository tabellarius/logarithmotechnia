package dataframe

import (
	"fmt"
	"logarithmotechnia/vector"
	"math"
	"reflect"
	"testing"
)

func TestFromCSVFile(t *testing.T) {
	df, _ := FromCSVFile("./test_data/persons.csv",
		CSVOptionSeparator(';'),
		CSVOptionSkipFirstLine(true))

	expectedColumnNames := []string{"Name", "DepType", "Salary", "KPI", "Group", "Active"}
	expectedColumns := []vector.Vector{
		vector.String([]string{"John", "Jane", "Jack", "Robert", "Marcius", "Catullus", "Marcia", "Gera", "Zeus", "Hephaestus", "Hades"}),
		vector.String([]string{"research", "research", "production", "research", "production", "logistics", "production", "sales", "sales", "factory", ""}),
		vector.IntegerWithNA([]int{120000, 0, 80000, 140000, 0, 100000, 60000, 150000, 225000, 150000, 175000},
			[]bool{false, true, false, false, true, false, false, false, false, false, false}),
		vector.FloatWithNA([]float64{1.45, 2.3, 3, 1, 0.67, math.NaN(), 1.44, 1.8, 1.125, 1.4, math.NaN()},
			[]bool{false, false, false, false, false, true, false, false, false, false, true}),
		vector.String([]string{"A", "A", "B", "B", "A", "A", "B", "A", "A", "B", "A"}),
		vector.BooleanWithNA([]bool{true, true, false, false, true, true, false, false, true, false, false},
			[]bool{false, false, false, false, false, false, false, false, false, false, true}),
	}

	if !reflect.DeepEqual(df.columnNames, expectedColumnNames) {
		t.Error(fmt.Sprintf("Column names (%v) are not equal to expected (%v)", df.columnNames, expectedColumnNames))
	}

	if !vector.CompareVectorArrs(df.columns, expectedColumns) {
		t.Error(fmt.Sprintf("Dataframe columns (%v) are not equal to expected (%v)",
			df.columns, expectedColumns))
	}
}

func TestCSVOptions(t *testing.T) {
	testData := []struct {
		name      string
		result    Option
		reference Option
	}{
		{
			name:      "CSVOptionSkipFirstLine",
			result:    CSVOptionSkipFirstLine(true),
			reference: ConfOption{optionCSVSkipFirstLine, true},
		},
		{
			name:      "CSVOptionSeparator",
			result:    CSVOptionSeparator(';'),
			reference: ConfOption{optionCSVSeparator, ';'},
		},
		{
			name:   "CSVOptionDataframeOptions",
			result: CSVOptionDataframeOptions(OptionColumnNames([]string{"id", "price"})),
			reference: ConfOption{optionCSVDataframeOptions,
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
