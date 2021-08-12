package dataframe

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOptionColumnNames(t *testing.T) {
	testData := []struct {
		name        string
		columnNames []string
		config      Config
	}{
		{
			name:        "normal",
			columnNames: []string{"one", "two", "three"},
			config:      Config{columnNames: []string{"one", "two", "three"}},
		},
		{
			name:        "empty",
			columnNames: []string{},
			config:      Config{columnNames: []string{}},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			config := OptionColumnNames(data.columnNames)

			if !reflect.DeepEqual(config, data.config) {
				t.Error(fmt.Sprintf("Config (%v) is not equal to expected (%v)", config, data.config))
			}
		})
	}
}
