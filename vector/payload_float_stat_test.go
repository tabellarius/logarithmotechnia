package vector

import (
	"fmt"
	"logarithmotechnia/util"
	"math"
	"reflect"
	"testing"
)

func TestFloatPayload_Sum(t *testing.T) {
	testData := []struct {
		name    string
		payload *floatPayload
		data    []float64
		sumNA   []bool
	}{
		{
			name:    "without na",
			payload: FloatPayload([]float64{-20, 10, 4, -20, 27}, nil).(*floatPayload),
			data:    []float64{1},
			sumNA:   []bool{false},
		},
		{
			name:    "with na",
			payload: FloatPayload([]float64{-20, 10, 4, -20, 27}, []bool{false, false, true, false, false}).(*floatPayload),
			data:    []float64{math.NaN()},
			sumNA:   []bool{true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.payload.Sum().(*floatPayload)

			if !util.EqualFloatArrays(payload.data, data.data) {
				t.Error(fmt.Sprintf("Sum (%v) is not equal to expected (%v)",
					payload.data, data.data))
			}

			if !reflect.DeepEqual(payload.na, data.sumNA) {
				t.Error(fmt.Sprintf("Sum (%v) is not equal to expected (%v)",
					payload.na, data.sumNA))
			}
		})
	}
}

func TestFloatPayload_Max(t *testing.T) {
	testData := []struct {
		name    string
		payload *floatPayload
		data    []float64
		sumNA   []bool
	}{
		{
			name:    "without na",
			payload: FloatPayload([]float64{-20, 10, 4, -20, 27}, nil).(*floatPayload),
			data:    []float64{27},
			sumNA:   []bool{false},
		},
		{
			name:    "with na",
			payload: FloatPayload([]float64{-20, 10, 4, -20, 27}, []bool{false, false, true, false, false}).(*floatPayload),
			data:    []float64{math.NaN()},
			sumNA:   []bool{true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.payload.Max().(*floatPayload)

			if !util.EqualFloatArrays(payload.data, data.data) {
				t.Error(fmt.Sprintf("Max (%v) is not equal to expected (%v)",
					payload.data, data.data))
			}

			if !reflect.DeepEqual(payload.na, data.sumNA) {
				t.Error(fmt.Sprintf("Max (%v) is not equal to expected (%v)",
					payload.na, data.sumNA))
			}
		})
	}
}

func TestFloatPayload_Min(t *testing.T) {
	testData := []struct {
		name    string
		payload *floatPayload
		data    []float64
		sumNA   []bool
	}{
		{
			name:    "without na",
			payload: FloatPayload([]float64{-20, 10, 4, -20, 27}, nil).(*floatPayload),
			data:    []float64{-20},
			sumNA:   []bool{false},
		},
		{
			name:    "with na",
			payload: FloatPayload([]float64{-20, 10, 4, -20, 27}, []bool{false, false, true, false, false}).(*floatPayload),
			data:    []float64{math.NaN()},
			sumNA:   []bool{true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.payload.Min().(*floatPayload)

			if !util.EqualFloatArrays(payload.data, data.data) {
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
