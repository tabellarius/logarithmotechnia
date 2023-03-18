package vector

import (
	"fmt"
	"logarithmotechnia/util"
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

			if !reflect.DeepEqual(payload.na, data.sumNA) {
				t.Error(fmt.Sprintf("Sum na (%v) is not equal to expected (%v)",
					payload.na, data.sumNA))
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

			if !reflect.DeepEqual(payload.na, data.sumNA) {
				t.Error(fmt.Sprintf("Max na (%v) is not equal to expected (%v)",
					payload.na, data.sumNA))
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

			if !reflect.DeepEqual(payload.na, data.sumNA) {
				t.Error(fmt.Sprintf("Min na (%v) is not equal to expected (%v)",
					payload.na, data.sumNA))
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

			if !reflect.DeepEqual(payload.na, data.na) {
				t.Error(fmt.Sprintf("Mean na (%v) is not equal to expected (%v)",
					payload.na, data.na))
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

			if !reflect.DeepEqual(payload.na, data.na) {
				t.Error(fmt.Sprintf("Mediann na (%v) is not equal to expected (%v)",
					payload.na, data.na))
			}
		})
	}
}
