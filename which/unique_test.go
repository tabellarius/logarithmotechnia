package which

import (
	"fmt"
	"logarithmotechnia/vector"
	"reflect"
	"testing"
)

func TestIsUnique(t *testing.T) {
	testData := []struct {
		name     string
		vector   vector.Vector
		booleans []bool
	}{
		{
			name:     "without NA",
			vector:   vector.Integer([]int{1, 0, 1, 3, 2, 3, 2, 0}),
			booleans: []bool{true, true, false, true, true, false, false, false},
		},
		{
			name: "with NA",
			vector: vector.IntegerWithNA([]int{1, 0, 1, 3, 2, 3, 2, 0},
				[]bool{false, true, true, false, false, false, false, false}),
			booleans: []bool{true, true, false, true, true, false, false, true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			booleans := IsUnique(data.vector)

			if !reflect.DeepEqual(booleans, data.booleans) {
				t.Error(fmt.Sprintf("Result of IsUnique() (%v) do not match expected (%v)",
					booleans, data.booleans))
			}
		})
	}

}

func TestIsNotUnique(t *testing.T) {
	testData := []struct {
		name     string
		vector   vector.Vector
		booleans []bool
	}{
		{
			name:     "without NA",
			vector:   vector.Integer([]int{1, 0, 1, 3, 2, 3, 2, 0}),
			booleans: []bool{false, false, true, false, false, true, true, true},
		},
		{
			name: "with NA",
			vector: vector.IntegerWithNA([]int{1, 0, 1, 3, 2, 3, 2, 0},
				[]bool{false, true, true, false, false, false, false, false}),
			booleans: []bool{false, false, true, false, false, true, true, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			booleans := IsNotUnique(data.vector)

			if !reflect.DeepEqual(booleans, data.booleans) {
				t.Error(fmt.Sprintf("Result of IsUnique() (%v) do not match expected (%v)",
					booleans, data.booleans))
			}
		})
	}

}
