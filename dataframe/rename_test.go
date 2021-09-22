package dataframe

import (
	"fmt"
	"logarithmotechnia/vector"
	"reflect"
	"testing"
)

func TestDataframe_Rename(t *testing.T) {
	df := getTestDataFrame().Select("name", "age", "gender")

	testData := []struct {
		name        string
		renames     []Rename
		columnNames []string
		columns     []vector.Vector
	}{
		{
			name:        "normal",
			renames:     []Rename{{"name", "nickname"}},
			columnNames: []string{"nickname", "age", "gender"},
			columns: []vector.Vector{
				vector.String([]string{"Jim", "SPARC-001", "Anna", "Lucius", "Maria"}, nil),
				vector.Integer([]int{31, 3, 24, 41, 33}, nil),
				vector.String([]string{"m", "", "f", "m", "f"}, []bool{false, true, false, false, false}),
			},
		},
		{
			name:        "normal",
			renames:     []Rename{{"name", "nickname"}, {"age", "nickname"}},
			columnNames: []string{"nickname", "nickname", "gender"},
			columns: []vector.Vector{
				vector.String([]string{"Jim", "SPARC-001", "Anna", "Lucius", "Maria"}, nil),
				vector.Integer([]int{31, 3, 24, 41, 33}, nil),
				vector.String([]string{"m", "", "f", "m", "f"}, []bool{false, true, false, false, false}),
			},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			newDf := df.Rename(data.renames...)

			if !reflect.DeepEqual(newDf.columnNames, data.columnNames) {
				t.Error(fmt.Sprintf("Dataframe column names (%v) are not equal to expected (%v)",
					newDf.columnNames, data.columnNames))
			}

			if !vector.CompareVectorArrs(newDf.columns, data.columns) {
				t.Error(fmt.Sprintf("Columns (%v) are not equal to expected (%v)",
					newDf.columns, data.columns))
			}
		})
	}
}
