package vector

import (
	"fmt"
	"logarithmotechnia/internal/util"
	"math"
	"reflect"
	"testing"
)

func TestIntegerPayload_Sum(t *testing.T) {
	testData := []struct {
		name    string
		payload *integerPayload
		data    []int
		sumNA   []bool
	}{
		{
			name:    "without na",
			payload: IntegerPayload([]int{-20, 10, 4, -20, 27}, nil).(*integerPayload),
			data:    []int{1},
			sumNA:   []bool{false},
		},
		{
			name:    "with na",
			payload: IntegerPayload([]int{-20, 10, 4, -20, 27}, []bool{false, false, true, false, false}).(*integerPayload),
			data:    []int{0},
			sumNA:   []bool{true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.payload.Sum().(*integerPayload)

			if !reflect.DeepEqual(payload.data, data.data) {
				t.Error(fmt.Sprintf("Sum data (%v) is not equal to expected (%v)",
					payload.data, data.data))
			}

			if !reflect.DeepEqual(payload.NA, data.sumNA) {
				t.Error(fmt.Sprintf("Sum na (%v) is not equal to expected (%v)",
					payload.NA, data.sumNA))
			}
		})
	}
}

func TestIntegerPayload_Max(t *testing.T) {
	testData := []struct {
		name    string
		payload *integerPayload
		data    []int
		sumNA   []bool
	}{
		{
			name:    "without na",
			payload: IntegerPayload([]int{-20, 10, 4, -20, 27}, nil).(*integerPayload),
			data:    []int{27},
			sumNA:   []bool{false},
		},
		{
			name:    "with na",
			payload: IntegerPayload([]int{-20, 10, 4, -20, 27}, []bool{false, false, true, false, false}).(*integerPayload),
			data:    []int{0},
			sumNA:   []bool{true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.payload.Max().(*integerPayload)

			if !reflect.DeepEqual(payload.data, data.data) {
				t.Error(fmt.Sprintf("Max data (%v) is not equal to expected (%v)",
					payload.data, data.data))
			}

			if !reflect.DeepEqual(payload.NA, data.sumNA) {
				t.Error(fmt.Sprintf("Max na (%v) is not equal to expected (%v)",
					payload.NA, data.sumNA))
			}
		})
	}
}

func TestIntegerPayload_Min(t *testing.T) {
	testData := []struct {
		name    string
		payload *integerPayload
		data    []int
		sumNA   []bool
	}{
		{
			name:    "without na",
			payload: IntegerPayload([]int{-20, 10, 4, -20, 27}, nil).(*integerPayload),
			data:    []int{-20},
			sumNA:   []bool{false},
		},
		{
			name:    "with na",
			payload: IntegerPayload([]int{-20, 10, 4, -20, 27}, []bool{false, false, true, false, false}).(*integerPayload),
			data:    []int{0},
			sumNA:   []bool{true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.payload.Min().(*integerPayload)

			if !reflect.DeepEqual(payload.data, data.data) {
				t.Error(fmt.Sprintf("Min data (%v) is not equal to expected (%v)",
					payload.data, data.data))
			}

			if !reflect.DeepEqual(payload.NA, data.sumNA) {
				t.Error(fmt.Sprintf("Min na (%v) is not equal to expected (%v)",
					payload.NA, data.sumNA))
			}
		})
	}
}

func TestIntegerPayload_Mean(t *testing.T) {
	testData := []struct {
		name    string
		payload *integerPayload
		data    []float64
		na      []bool
	}{
		{
			name:    "without na",
			payload: IntegerPayload([]int{-10, 10, 4, -20, 26}, nil).(*integerPayload),
			data:    []float64{2},
			na:      []bool{false},
		},
		{
			name:    "with na",
			payload: IntegerPayload([]int{-20, 10, 4, -20, 26}, []bool{false, false, true, false, false}).(*integerPayload),
			data:    []float64{math.NaN()},
			na:      []bool{true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.payload.Mean().(*floatPayload)

			if !util.EqualFloatArrays(payload.data, data.data) {
				t.Error(fmt.Sprintf("Mean data (%v) is not equal to expected (%v)",
					payload.data, data.data))
			}

			if !reflect.DeepEqual(payload.NA, data.na) {
				t.Error(fmt.Sprintf("Mean na (%v) is not equal to expected (%v)",
					payload.NA, data.na))
			}
		})
	}
}

func TestIntegerPayload_Median(t *testing.T) {
	testData := []struct {
		name    string
		payload *integerPayload
		data    []int
		na      []bool
	}{
		{
			name:    "without na odd",
			payload: IntegerPayload([]int{10, 26, 4, -20, 26}, nil).(*integerPayload),
			data:    []int{10},
			na:      []bool{false},
		},
		{
			name:    "without na even",
			payload: IntegerPayload([]int{10, 26, 4, -20, 26, 16}, nil).(*integerPayload),
			data:    []int{13},
			na:      []bool{false},
		},
		{
			name:    "with na",
			payload: IntegerPayload([]int{10, 26, 4, -20, 26}, []bool{false, false, true, false, false}).(*integerPayload),
			data:    []int{0},
			na:      []bool{true},
		},
		{
			name:    "one element",
			payload: IntegerPayload([]int{10}, nil).(*integerPayload),
			data:    []int{10},
			na:      []bool{false},
		},
		{
			name:    "zero elements",
			payload: IntegerPayload([]int{}, nil).(*integerPayload),
			data:    []int{0},
			na:      []bool{true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.payload.Median().(*integerPayload)

			if !reflect.DeepEqual(payload.data, data.data) {
				t.Error(fmt.Sprintf("Median data (%v) is not equal to expected (%v)",
					payload.data, data.data))
			}

			if !reflect.DeepEqual(payload.NA, data.na) {
				t.Error(fmt.Sprintf("Mediann na (%v) is not equal to expected (%v)",
					payload.NA, data.na))
			}
		})
	}
}

func TestIntegerPayload_Prod(t *testing.T) {
	testData := []struct {
		name    string
		payload *integerPayload
		data    []int
		na      []bool
	}{
		{
			name:    "without na",
			payload: IntegerPayload([]int{10, 26, 4, -20, 26}, nil).(*integerPayload),
			data:    []int{-540800},
			na:      []bool{false},
		},
		{
			name:    "with na",
			payload: IntegerPayload([]int{10, 26, 4, -20, 26}, []bool{false, false, true, false, false}).(*integerPayload),
			data:    []int{0},
			na:      []bool{true},
		},
		{
			name:    "one element",
			payload: IntegerPayload([]int{10}, nil).(*integerPayload),
			data:    []int{10},
			na:      []bool{false},
		},
		{
			name:    "zero elements",
			payload: IntegerPayload([]int{}, nil).(*integerPayload),
			data:    []int{0},
			na:      []bool{false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.payload.Prod().(*integerPayload)

			if !reflect.DeepEqual(payload.data, data.data) {
				t.Error(fmt.Sprintf("Prod data (%v) is not equal to expected (%v)",
					payload.data, data.data))
			}

			if !reflect.DeepEqual(payload.NA, data.na) {
				t.Error(fmt.Sprintf("Prod na (%v) is not equal to expected (%v)",
					payload.NA, data.na))
			}
		})
	}
}

func TestIntegerPayload_CumSum(t *testing.T) {
	testData := []struct {
		name    string
		payload *integerPayload
		data    []int
		na      []bool
	}{
		{
			name:    "without na",
			payload: IntegerPayload([]int{10, 26, 4, -20, 26}, nil).(*integerPayload),
			data:    []int{10, 36, 40, 20, 46},
			na:      []bool{false, false, false, false, false},
		},
		{
			name:    "with na",
			payload: IntegerPayload([]int{10, 26, 4, -20, 26}, []bool{false, false, true, false, false}).(*integerPayload),
			data:    []int{10, 36, 0, 0, 0},
			na:      []bool{false, false, true, true, true},
		},
		{
			name:    "one element",
			payload: IntegerPayload([]int{10}, nil).(*integerPayload),
			data:    []int{10},
			na:      []bool{false},
		},
		{
			name:    "zero elements",
			payload: IntegerPayload([]int{}, nil).(*integerPayload),
			data:    []int{},
			na:      []bool{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.payload.CumSum().(*integerPayload)

			if !reflect.DeepEqual(payload.data, data.data) {
				t.Error(fmt.Sprintf("CumSum data (%v) is not equal to expected (%v)",
					payload.data, data.data))
			}

			if !reflect.DeepEqual(payload.NA, data.na) {
				t.Error(fmt.Sprintf("CumSum na (%v) is not equal to expected (%v)",
					payload.NA, data.na))
			}
		})
	}
}

func TestIntegerPayload_CumProd(t *testing.T) {
	testData := []struct {
		name    string
		payload *integerPayload
		data    []int
		na      []bool
	}{
		{
			name:    "without na",
			payload: IntegerPayload([]int{10, 26, 4, -20, 26}, nil).(*integerPayload),
			data:    []int{10, 260, 1040, -20800, -540800},
			na:      []bool{false, false, false, false, false},
		},
		{
			name:    "with na",
			payload: IntegerPayload([]int{10, 26, 4, -20, 26}, []bool{false, false, true, false, false}).(*integerPayload),
			data:    []int{10, 260, 0, 0, 0},
			na:      []bool{false, false, true, true, true},
		},
		{
			name:    "one element",
			payload: IntegerPayload([]int{10}, nil).(*integerPayload),
			data:    []int{10},
			na:      []bool{false},
		},
		{
			name:    "zero elements",
			payload: IntegerPayload([]int{}, nil).(*integerPayload),
			data:    []int{},
			na:      []bool{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.payload.CumProd().(*integerPayload)

			if !reflect.DeepEqual(payload.data, data.data) {
				t.Error(fmt.Sprintf("CumProd data (%v) is not equal to expected (%v)",
					payload.data, data.data))
			}

			if !reflect.DeepEqual(payload.NA, data.na) {
				t.Error(fmt.Sprintf("CumProd na (%v) is not equal to expected (%v)",
					payload.NA, data.na))
			}
		})
	}
}

func TestIntegerPayload_CumMax(t *testing.T) {
	testData := []struct {
		name    string
		payload *integerPayload
		data    []int
		na      []bool
	}{
		{
			name:    "without na",
			payload: IntegerPayload([]int{10, 26, 4, 35, -2}, nil).(*integerPayload),
			data:    []int{10, 26, 26, 35, 35},
			na:      []bool{false, false, false, false, false},
		},
		{
			name:    "with na",
			payload: IntegerPayload([]int{10, 26, 4, -20, 26}, []bool{false, false, true, false, false}).(*integerPayload),
			data:    []int{10, 26, 0, 0, 0},
			na:      []bool{false, false, true, true, true},
		},
		{
			name:    "one element",
			payload: IntegerPayload([]int{10}, nil).(*integerPayload),
			data:    []int{10},
			na:      []bool{false},
		},
		{
			name:    "zero elements",
			payload: IntegerPayload([]int{}, nil).(*integerPayload),
			data:    []int{},
			na:      []bool{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.payload.CumMax().(*integerPayload)

			if !reflect.DeepEqual(payload.data, data.data) {
				t.Error(fmt.Sprintf("CumMax data (%v) is not equal to expected (%v)",
					payload.data, data.data))
			}

			if !reflect.DeepEqual(payload.NA, data.na) {
				t.Error(fmt.Sprintf("CumMax na (%v) is not equal to expected (%v)",
					payload.NA, data.na))
			}
		})
	}
}

func TestIntegerPayload_CumMin(t *testing.T) {
	testData := []struct {
		name    string
		payload *integerPayload
		data    []int
		na      []bool
	}{
		{
			name:    "without na",
			payload: IntegerPayload([]int{10, 26, 4, 35, -2}, nil).(*integerPayload),
			data:    []int{10, 10, 4, 4, -2},
			na:      []bool{false, false, false, false, false},
		},
		{
			name:    "with na",
			payload: IntegerPayload([]int{10, 26, 4, -20, 26}, []bool{false, false, true, false, false}).(*integerPayload),
			data:    []int{10, 10, 0, 0, 0},
			na:      []bool{false, false, true, true, true},
		},
		{
			name:    "one element",
			payload: IntegerPayload([]int{10}, nil).(*integerPayload),
			data:    []int{10},
			na:      []bool{false},
		},
		{
			name:    "zero elements",
			payload: IntegerPayload([]int{}, nil).(*integerPayload),
			data:    []int{},
			na:      []bool{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.payload.CumMin().(*integerPayload)

			if !reflect.DeepEqual(payload.data, data.data) {
				t.Error(fmt.Sprintf("CumMin data (%v) is not equal to expected (%v)",
					payload.data, data.data))
			}

			if !reflect.DeepEqual(payload.NA, data.na) {
				t.Error(fmt.Sprintf("CumMin na (%v) is not equal to expected (%v)",
					payload.NA, data.na))
			}
		})
	}
}
