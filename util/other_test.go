package util

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToIndexes(t *testing.T) {
	vecLength := 10

	testData := []struct {
		name    string
		which   []bool
		indices []int
	}{
		{
			name:    "random",
			which:   []bool{false, false, true, true, true, false, false, true, false, true},
			indices: []int{3, 4, 5, 8, 10},
		},
		{
			name:    "multiple",
			which:   []bool{true, false, true, false, false},
			indices: []int{1, 3, 6, 8},
		},
		{
			name:    "non-multiple",
			which:   []bool{false, true, true, false},
			indices: []int{2, 3, 6, 7, 10},
		},
		{
			name:    "one of two",
			which:   []bool{true, false},
			indices: []int{1, 3, 5, 7, 9},
		},
		{
			name:    "one of three",
			which:   []bool{true, false, false},
			indices: []int{1, 4, 7, 10},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			indices := ToIndices(vecLength, data.which)
			if !reflect.DeepEqual(indices, data.indices) {
				t.Error(fmt.Sprintf("Indices (%v) is not equal to expected (%v)", indices, data.indices))
			}
		})
	}
}
