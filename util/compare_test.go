package util

import (
	"fmt"
	"math"
	"math/cmplx"
	"testing"
)

func TestEqualFloatArrays(t *testing.T) {
	testData := []struct {
		name   string
		arr1   []float64
		arr2   []float64
		result bool
	}{
		{
			name:   "equal",
			arr1:   []float64{1.1, 2.2, math.NaN()},
			arr2:   []float64{1.1, 2.2, math.NaN()},
			result: true,
		},
		{
			name:   "not equal",
			arr1:   []float64{1.1, 2.2, math.NaN()},
			arr2:   []float64{3.0, 2.2, math.NaN()},
			result: false,
		},
		{
			name:   "not equal nan",
			arr1:   []float64{1.1, 2.2, math.NaN()},
			arr2:   []float64{1.1, 2.2, 3.3},
			result: false,
		},
		{
			name:   "different length",
			arr1:   []float64{1.1, 2.2, math.NaN()},
			arr2:   []float64{1.1, 2.2},
			result: false,
		},
		{
			name:   "zero length",
			arr1:   []float64{},
			arr2:   []float64{},
			result: true,
		},
		{
			name:   "both nil",
			arr1:   nil,
			arr2:   nil,
			result: true,
		},
		{
			name:   "arr1 is nil",
			arr1:   nil,
			arr2:   []float64{},
			result: false,
		},
		{
			name:   "arr2 is nil",
			arr1:   []float64{},
			arr2:   nil,
			result: false,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			result := EqualFloatArrays(data.arr1, data.arr2)
			if result != data.result {
				t.Error(fmt.Sprintf("Result (%t) is not equal to expected (%t)", result, data.result))
			}
		})
	}
}

func TestEqualComplexArrays(t *testing.T) {
	testData := []struct {
		name   string
		arr1   []complex128
		arr2   []complex128
		result bool
	}{
		{
			name:   "equal",
			arr1:   []complex128{1.1 + 1.1i, 2.2 + 2.2i, cmplx.NaN()},
			arr2:   []complex128{1.1 + 1.1i, 2.2 + 2.2i, cmplx.NaN()},
			result: true,
		},
		{
			name:   "not equal",
			arr1:   []complex128{1.1 + 1.1i, 2.2 + 2.2i, cmplx.NaN()},
			arr2:   []complex128{3.0 + 1.0i, 2.2 + 2.2i, cmplx.NaN()},
			result: false,
		},
		{
			name:   "not equal nan",
			arr1:   []complex128{1.1 + 1.1i, 2.2 + 2.2i, cmplx.NaN()},
			arr2:   []complex128{1.1 + 1.1i, 2.2 + 2.2i, 3.3 + 3.3i},
			result: false,
		},
		{
			name:   "different length",
			arr1:   []complex128{1.1 + 1.1i, 2.2 + 2.2i, cmplx.NaN()},
			arr2:   []complex128{3.0 + 1.0i, 2.2 + 2.2i},
			result: false,
		},
		{
			name:   "zero length",
			arr1:   []complex128{},
			arr2:   []complex128{},
			result: true,
		},
		{
			name:   "both nil",
			arr1:   nil,
			arr2:   nil,
			result: true,
		},
		{
			name:   "arr1 is nil",
			arr1:   nil,
			arr2:   []complex128{},
			result: false,
		},
		{
			name:   "arr2 is nil",
			arr1:   []complex128{},
			arr2:   nil,
			result: false,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			result := EqualComplexArrays(data.arr1, data.arr2)
			if result != data.result {
				t.Error(fmt.Sprintf("Result (%t) is not equal to expected (%t)", result, data.result))
			}
		})
	}
}
