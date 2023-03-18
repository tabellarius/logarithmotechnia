package vector

import (
	"fmt"
	"testing"
)

func TestVector_Sum(t *testing.T) {
	testData := []struct {
		name   string
		vec    Vector
		sumVec Vector
	}{
		{
			name:   "normal summer",
			vec:    Integer([]int{10, 2, 8, 12, 18}),
			sumVec: Integer([]int{50}),
		},
		{
			name:   "normal non-summer",
			vec:    String([]string{"one", "two", "8", "12", "18"}),
			sumVec: NA(1),
		},
		{
			name:   "normal grouped summer",
			vec:    Integer([]int{10, 2, 10, 12, 0, 10, 2}).GroupByIndices([][]int{{1, 3, 6}, {2, 4, 7}, {5}}),
			sumVec: Integer([]int{30, 16, 0}),
		},
		{
			name:   "normal grouped non-summer",
			vec:    String([]string{"one", "two", "8", "12", "18"}).GroupByIndices([][]int{{1, 3}, {2, 5}, {4}}),
			sumVec: NA(3),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			sumVec := data.vec.Sum()

			if !CompareVectorsForTest(sumVec, data.sumVec) {
				t.Error(fmt.Sprintf("Sum vector (%v) does not match expected (%v)",
					sumVec, data.sumVec))
			}
		})
	}
}

func TestVector_Max(t *testing.T) {
	testData := []struct {
		name   string
		vec    Vector
		valVec Vector
	}{
		{
			name:   "normal maxxer",
			vec:    Integer([]int{10, 2, 8, 12, 18}),
			valVec: Integer([]int{18}),
		},
		{
			name:   "normal non-maxxer",
			vec:    Complex([]complex128{10, 2, 8, 12, 18}),
			valVec: NA(1),
		},
		{
			name:   "normal grouped maxxer",
			vec:    Integer([]int{10, 2, 10, 12, 0, 10, 2}).GroupByIndices([][]int{{1, 3, 6}, {2, 4, 7}, {5}}),
			valVec: Integer([]int{10, 12, 0}),
		},
		{
			name:   "normal grouped non-maxxer",
			vec:    Complex([]complex128{1, 2, 8, 12, 18}).GroupByIndices([][]int{{1, 3}, {2, 5}, {4}}),
			valVec: NA(3),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			valVec := data.vec.Max()

			if !CompareVectorsForTest(valVec, data.valVec) {
				t.Error(fmt.Sprintf("Max vector (%v) does not match expected (%v)",
					valVec, data.valVec))
			}
		})
	}
}

func TestVector_Min(t *testing.T) {
	testData := []struct {
		name   string
		vec    Vector
		valVec Vector
	}{
		{
			name:   "normal minner",
			vec:    Integer([]int{10, 2, 8, 12, 18}),
			valVec: Integer([]int{2}),
		},
		{
			name:   "normal non-minner",
			vec:    Complex([]complex128{10, 2, 8, 12, 18}),
			valVec: NA(1),
		},
		{
			name:   "normal grouped minner",
			vec:    Integer([]int{10, 2, 10, 12, 0, 10, 2}).GroupByIndices([][]int{{1, 3, 6}, {2, 4, 7}, {5}}),
			valVec: Integer([]int{10, 2, 0}),
		},
		{
			name:   "normal grouped non-minner",
			vec:    Complex([]complex128{1, 2, 8, 12, 18}).GroupByIndices([][]int{{1, 3}, {2, 5}, {4}}),
			valVec: NA(3),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			valVec := data.vec.Min()

			if !CompareVectorsForTest(valVec, data.valVec) {
				t.Error(fmt.Sprintf("Min vector (%v) does not match expected (%v)",
					valVec, data.valVec))
			}
		})
	}
}

func TestVector_Mean(t *testing.T) {
	testData := []struct {
		name   string
		vec    Vector
		valVec Vector
	}{
		{
			name:   "normal meaner",
			vec:    Integer([]int{10, 2, 8, 12, 18}),
			valVec: Float([]float64{10}),
		},
		{
			name:   "normal non-meaner",
			vec:    String([]string{"10", "2", "8", "12", "18"}),
			valVec: NA(1),
		},
		{
			name:   "normal grouped minner",
			vec:    Integer([]int{10, 2, 10, 12, 0, 10, 4}).GroupByIndices([][]int{{1, 3, 6}, {2, 4, 7}, {5}}),
			valVec: Float([]float64{10, 6, 0}),
		},
		{
			name:   "normal grouped non-minner",
			vec:    String([]string{"1", "2", "8", "12", "18"}).GroupByIndices([][]int{{1, 3}, {2, 5}, {4}}),
			valVec: NA(3),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			valVec := data.vec.Mean()

			if !CompareVectorsForTest(valVec, data.valVec) {
				t.Error(fmt.Sprintf("Mean vector (%v) does not match expected (%v)",
					valVec, data.valVec))
			}
		})
	}
}
