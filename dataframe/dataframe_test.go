package dataframe

import (
	"fmt"
	"github.com/dee-ru/logarithmotechnia/vector"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	testData := []struct {
		name      string
		columns   []vector.Vector
		config    []Config
		dfColumns []vector.Vector
		dfConfig  Config
	}{
		{
			name:      "empty",
			columns:   []vector.Vector{},
			config:    []Config{},
			dfColumns: []vector.Vector{},
			dfConfig:  Config{columnNames: []string{}},
		},
		{
			name:      "empty with column names",
			columns:   []vector.Vector{},
			config:    []Config{OptionColumnNames([]string{"one", "two", "three"})},
			dfColumns: []vector.Vector{},
			dfConfig:  Config{columnNames: []string{}},
		},
		{
			name: "normal",
			columns: []vector.Vector{
				vector.Integer([]int{1, 2, 3}, nil),
				vector.String([]string{"1", "2", "3"}, nil),
				vector.Boolean([]bool{true, true, false}, nil),
			},
			config: []Config{},
			dfColumns: []vector.Vector{
				vector.Integer([]int{1, 2, 3}, nil),
				vector.String([]string{"1", "2", "3"}, nil),
				vector.Boolean([]bool{true, true, false}, nil),
			},
			dfConfig: Config{columnNames: []string{"1", "2", "3"}},
		},
		{
			name: "normal with column names",
			columns: []vector.Vector{
				vector.Integer([]int{1, 2, 3}, nil),
				vector.String([]string{"1", "2", "3"}, nil),
				vector.Boolean([]bool{true, true, false}, nil),
			},
			config: []Config{OptionColumnNames([]string{"int", "string", "bool"})},
			dfColumns: []vector.Vector{
				vector.Integer([]int{1, 2, 3}, nil),
				vector.String([]string{"1", "2", "3"}, nil),
				vector.Boolean([]bool{true, true, false}, nil),
			},
			dfConfig: Config{columnNames: []string{"int", "string", "bool"}},
		},
		{
			name: "normal with partial column names",
			columns: []vector.Vector{
				vector.Integer([]int{1, 2, 3}, nil),
				vector.String([]string{"1", "2", "3"}, nil),
				vector.Boolean([]bool{true, true, false}, nil),
			},
			config: []Config{OptionColumnNames([]string{"int", "string"})},
			dfColumns: []vector.Vector{
				vector.Integer([]int{1, 2, 3}, nil),
				vector.String([]string{"1", "2", "3"}, nil),
				vector.Boolean([]bool{true, true, false}, nil),
			},
			dfConfig: Config{columnNames: []string{"int", "string", "3"}},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			df := New(data.columns, data.config...)

			if !reflect.DeepEqual(df.columns, data.dfColumns) {
				t.Error(fmt.Sprintf("Columns (%v) are not equal to expected (%v)", df.columns, data.dfColumns))
			}
			if !reflect.DeepEqual(df.config, data.dfConfig) {
				t.Error(fmt.Sprintf("Config (%v) are not equal to expected (%v)",
					df.config, data.dfConfig))
			}
		})
	}
}
