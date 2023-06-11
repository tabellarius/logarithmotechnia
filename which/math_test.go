package which

import (
	"logarithmotechnia/vector"
	"math"
	"reflect"
	"testing"
)

func TestIsInf(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  []bool
	}{
		{
			name: "float64",
			in:   vector.Float([]float64{1, 2, 3, math.Inf(1), math.Inf(-1), 4, 5, 6}),
			out:  []bool{false, false, false, true, true, false, false, false},
		},
		{
			name: "complex128",
			in:   vector.Complex([]complex128{1, 2, 3, complex(math.Inf(1), 0), complex(math.Inf(-1), 0), 4, 5, 6}),
			out:  []bool{false, false, false, true, true, false, false, false},
		},
		{
			name: "int",
			in:   vector.Integer([]int{1, 2, 3, 4, 5, 6}),
			out:  []bool{false, false, false, false, false, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := IsInf(data.in)
			if !reflect.DeepEqual(out, data.out) {
				t.Errorf("Expected %v, got %v", data.out, out)
			}
		})
	}

}

func TestIsNaN(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  []bool
	}{
		{
			name: "float64",
			in:   vector.Float([]float64{1, 2, 3, math.NaN(), 4, 5, 6}),
			out:  []bool{false, false, false, true, false, false, false},
		},
		{
			name: "complex128",
			in:   vector.Complex([]complex128{1, 2, 3, complex(math.NaN(), 0), 4, 5, 6}),
			out:  []bool{false, false, false, true, false, false, false},
		},
		{
			name: "int",
			in:   vector.Integer([]int{1, 2, 3, 4, 5, 6}),
			out:  []bool{false, false, false, false, false, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := IsNaN(data.in)
			if !reflect.DeepEqual(out, data.out) {
				t.Errorf("Expected %v, got %v", data.out, out)
			}
		})
	}
}

func TestSignbit(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  []bool
	}{
		{
			name: "float64",
			in:   vector.Float([]float64{1, 2, 3, -4, 5, 6}),
			out:  []bool{false, false, false, true, false, false},
		},
		{
			name: "int",
			in:   vector.Integer([]int{1, 2, -3, -4, 5, 6}),
			out:  []bool{false, false, true, true, false, false},
		},
		{
			name: "string",
			in:   vector.String([]string{"1", "2", "-3", "-4", "5", "6"}),
			out:  []bool{false, false, false, false, false, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := Signbit(data.in)
			if !reflect.DeepEqual(out, data.out) {
				t.Errorf("Expected %v, got %v", data.out, out)
			}
		})
	}
}
