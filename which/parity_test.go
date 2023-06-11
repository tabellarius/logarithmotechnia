package which

import (
	"fmt"
	"logarithmotechnia/vector"
	"reflect"
	"testing"
	"time"
)

func TestEven(t *testing.T) {
	testData := []struct {
		name        string
		vec         vector.Vector
		outBooleans []bool
	}{
		{
			name:        "five",
			vec:         vector.Integer([]int{1, 2, 3, 4, 5}),
			outBooleans: []bool{false, true, false, true, false},
		},
		{
			name:        "eight",
			vec:         vector.Float([]float64{1, 2, 3, 4, 5, 6, 7, 8}),
			outBooleans: []bool{false, true, false, true, false, true, false, true},
		},
		{
			name:        "empty",
			vec:         vector.Any([]any{}),
			outBooleans: []bool{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			outBooleans := Even(data.vec)
			if !reflect.DeepEqual(outBooleans, data.outBooleans) {
				t.Error(fmt.Sprintf("vec.Even() (%v) does not match data.outBooleans (%v)", outBooleans, data.outBooleans))
			}
		})
	}

}

func TestOdd(t *testing.T) {
	testData := []struct {
		name        string
		vec         vector.Vector
		outBooleans []bool
	}{
		{
			name:        "five",
			vec:         vector.Integer([]int{1, 2, 3, 4, 5}),
			outBooleans: []bool{true, false, true, false, true},
		},
		{
			name:        "eight",
			vec:         vector.Float([]float64{1, 2, 3, 4, 5, 6, 7, 8}),
			outBooleans: []bool{true, false, true, false, true, false, true, false},
		},
		{
			name:        "empty",
			vec:         vector.Time([]time.Time{}),
			outBooleans: []bool{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			outBooleans := Odd(data.vec)
			if !reflect.DeepEqual(outBooleans, data.outBooleans) {
				t.Error(fmt.Sprintf("vec.Even() (%v) does not match data.outBooleans (%v)", outBooleans, data.outBooleans))
			}
		})
	}
}

func TestNth(t *testing.T) {
	testData := []struct {
		name        string
		nth         int
		vec         vector.Vector
		outBooleans []bool
	}{
		{
			name:        "every 3",
			nth:         3,
			vec:         vector.Integer([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			outBooleans: []bool{false, false, true, false, false, true, false, false, true, false},
		},
		{
			name:        "every 5",
			nth:         5,
			vec:         vector.Integer([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			outBooleans: []bool{false, false, false, false, true, false, false, false, false, true},
		},
		{
			name:        "every 1",
			nth:         1,
			vec:         vector.Integer([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			outBooleans: []bool{true, true, true, true, true, true, true, true, true, true},
		},
		{
			name:        "empty",
			nth:         3,
			vec:         vector.Time([]time.Time{}),
			outBooleans: []bool{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			outBooleans := Nth(data.vec, data.nth)
			if !reflect.DeepEqual(outBooleans, data.outBooleans) {
				t.Error(fmt.Sprintf("vec.Even() (%v) does not match data.outBooleans (%v)", outBooleans, data.outBooleans))
			}
		})
	}

}
