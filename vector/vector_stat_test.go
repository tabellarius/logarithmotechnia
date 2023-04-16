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
			name:   "normal grouped meaner",
			vec:    Integer([]int{10, 2, 10, 12, 0, 10, 4}).GroupByIndices([][]int{{1, 3, 6}, {2, 4, 7}, {5}}),
			valVec: Float([]float64{10, 6, 0}),
		},
		{
			name:   "normal grouped non-meaner",
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

func TestVector_Median(t *testing.T) {
	testData := []struct {
		name   string
		vec    Vector
		valVec Vector
	}{
		{
			name:   "normal medianer",
			vec:    Integer([]int{-100, 2, 8, 12, 180}),
			valVec: Integer([]int{8}),
		},
		{
			name:   "normal non-medianer",
			vec:    String([]string{"10", "2", "8", "12", "18"}),
			valVec: NA(1),
		},
		{
			name:   "normal grouped medianer",
			vec:    Float([]float64{10, 2, 10, 12, 0, 10, 4}).GroupByIndices([][]int{{1, 3, 6}, {2, 4, 7}, {5}}),
			valVec: Float([]float64{10, 4, 0}),
		},
		{
			name:   "normal grouped non-medianer",
			vec:    String([]string{"1", "2", "8", "12", "18"}).GroupByIndices([][]int{{1, 3}, {2, 5}, {4}}),
			valVec: NA(3),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			valVec := data.vec.Median()

			if !CompareVectorsForTest(valVec, data.valVec) {
				t.Error(fmt.Sprintf("Median vector (%v) does not match expected (%v)",
					valVec, data.valVec))
			}
		})
	}
}

func TestVector_Prod(t *testing.T) {
	testData := []struct {
		name   string
		vec    Vector
		valVec Vector
	}{
		{
			name:   "normal proder",
			vec:    Integer([]int{-100, 2, 8, 12, 180}),
			valVec: Integer([]int{-3456000}),
		},
		{
			name:   "normal non-proder",
			vec:    String([]string{"10", "2", "8", "12", "18"}),
			valVec: NA(1),
		},
		{
			name:   "normal grouped proder",
			vec:    Float([]float64{10, 2, 10, 12, 0, 10, 4}).GroupByIndices([][]int{{1, 3, 6}, {2, 4, 7}, {5}}),
			valVec: Float([]float64{1000, 96, 0}),
		},
		{
			name:   "normal grouped non-proder",
			vec:    String([]string{"1", "2", "8", "12", "18"}).GroupByIndices([][]int{{1, 3}, {2, 5}, {4}}),
			valVec: NA(3),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			valVec := data.vec.Prod()

			if !CompareVectorsForTest(valVec, data.valVec) {
				t.Error(fmt.Sprintf("Prod vector (%v) does not match expected (%v)",
					valVec, data.valVec))
			}
		})
	}
}

func TestVector_CumSum(t *testing.T) {
	testData := []struct {
		name   string
		vec    Vector
		valVec Vector
	}{
		{
			name:   "normal cumsumer",
			vec:    Integer([]int{-100, 2, 8, 12, 180}),
			valVec: Integer([]int{-100, -98, -90, -78, 102}),
		},
		{
			name:   "normal non-cumsumer",
			vec:    String([]string{"10", "2", "8", "12", "18"}),
			valVec: NA(5),
		},
		{
			name:   "normal grouped cumsumer",
			vec:    Float([]float64{10, 2, 10, 12, 0, 10, 4}).GroupByIndices([][]int{{1, 3, 6}, {2, 4, 7}, {5}}),
			valVec: VectorVector([]Vector{Float([]float64{10, 20, 30}), Float([]float64{2, 14, 18}), Float([]float64{0})}),
		},
		{
			name:   "normal grouped non-cumsumer",
			vec:    String([]string{"1", "2", "8", "12", "18"}).GroupByIndices([][]int{{1, 3}, {2, 5}, {4}}),
			valVec: VectorVector([]Vector{NA(2), NA(2), NA(1)}),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			valVec := data.vec.CumSum()

			if !CompareVectorsForTest(valVec, data.valVec) {
				t.Error(fmt.Sprintf("CumSum vector (%v) does not match expected (%v)",
					valVec, data.valVec))
			}
		})
	}
}

func TestVector_CumProd(t *testing.T) {
	testData := []struct {
		name   string
		vec    Vector
		valVec Vector
	}{
		{
			name:   "normal cumproder",
			vec:    Integer([]int{-100, 2, 8, 12, 180}),
			valVec: Integer([]int{-100, -200, -1600, -19200, -3456000}),
		},
		{
			name:   "normal non-cumproder",
			vec:    String([]string{"10", "2", "8", "12", "18"}),
			valVec: NA(5),
		},
		{
			name:   "normal grouped cumproder",
			vec:    Float([]float64{10, 2, 10, 12, 0, 10, 4}).GroupByIndices([][]int{{1, 3, 6}, {2, 4, 7}, {5}}),
			valVec: VectorVector([]Vector{Float([]float64{10, 100, 1000}), Float([]float64{2, 24, 96}), Float([]float64{0})}),
		},
		{
			name:   "normal grouped non-cumproder",
			vec:    String([]string{"1", "2", "8", "12", "18"}).GroupByIndices([][]int{{1, 3}, {2, 5}, {4}}),
			valVec: VectorVector([]Vector{NA(2), NA(2), NA(1)}),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			valVec := data.vec.CumProd()

			if !CompareVectorsForTest(valVec, data.valVec) {
				t.Error(fmt.Sprintf("CumProd vector (%v) does not match expected (%v)",
					valVec, data.valVec))
			}
		})
	}
}

func TestVector_CumMax(t *testing.T) {
	testData := []struct {
		name   string
		vec    Vector
		valVec Vector
	}{
		{
			name:   "normal cummaxer",
			vec:    Integer([]int{-100, 2, 8, 4, 180}),
			valVec: Integer([]int{-100, 2, 8, 8, 180}),
		},
		{
			name:   "normal non-cummaxer",
			vec:    String([]string{"10", "2", "8", "12", "18"}),
			valVec: NA(5),
		},
		{
			name:   "normal grouped cummaxer",
			vec:    Float([]float64{10, 2, 10, 12, 0, 10, 4}).GroupByIndices([][]int{{1, 3, 6}, {2, 4, 7}, {5}}),
			valVec: VectorVector([]Vector{Float([]float64{10, 10, 10}), Float([]float64{2, 12, 12}), Float([]float64{0})}),
		},
		{
			name:   "normal grouped non-cummaxer",
			vec:    String([]string{"1", "2", "8", "12", "18"}).GroupByIndices([][]int{{1, 3}, {2, 5}, {4}}),
			valVec: VectorVector([]Vector{NA(2), NA(2), NA(1)}),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			valVec := data.vec.CumMax()

			if !CompareVectorsForTest(valVec, data.valVec) {
				t.Error(fmt.Sprintf("CumMax vector (%v) does not match expected (%v)",
					valVec, data.valVec))
			}
		})
	}
}

func TestVector_CumMin(t *testing.T) {
	testData := []struct {
		name   string
		vec    Vector
		valVec Vector
	}{
		{
			name:   "normal cumminer",
			vec:    Integer([]int{8, 2, 4, -10, 180}),
			valVec: Integer([]int{8, 2, 2, -10, -10}),
		},
		{
			name:   "normal non-cumminer",
			vec:    String([]string{"10", "2", "8", "12", "18"}),
			valVec: NA(5),
		},
		{
			name:   "normal grouped cumminer",
			vec:    Float([]float64{10, 2, 100, 1, 0, 4, 4}).GroupByIndices([][]int{{1, 3, 6}, {2, 4, 7}, {5}}),
			valVec: VectorVector([]Vector{Float([]float64{10, 10, 4}), Float([]float64{2, 1, 1}), Float([]float64{0})}),
		},
		{
			name:   "normal grouped non-cumminer",
			vec:    String([]string{"1", "2", "8", "12", "18"}).GroupByIndices([][]int{{1, 3}, {2, 5}, {4}}),
			valVec: VectorVector([]Vector{NA(2), NA(2), NA(1)}),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			valVec := data.vec.CumMin()

			if !CompareVectorsForTest(valVec, data.valVec) {
				t.Error(fmt.Sprintf("CumMax vector (%v) does not match expected (%v)",
					valVec, data.valVec))
			}
		})
	}
}
