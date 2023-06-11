package vector

import (
	"fmt"
	"logarithmotechnia/internal/util"
	"math"
	"reflect"
	"testing"
)

func TestFloatPayload_Add(t *testing.T) {
	testData := []struct {
		name      string
		operator  Payload
		operandum Payload
		outData   []float64
		outNA     []bool
	}{
		{
			name:      "same type",
			operator:  FloatPayload([]float64{1, 2, 3, 4, 0}, []bool{false, false, false, false, true}),
			operandum: FloatPayload([]float64{11, 12, 13, 0, 15}, []bool{false, false, false, true, false}),
			outData:   []float64{12, 14, 16, math.NaN(), math.NaN()},
			outNA:     []bool{false, false, false, true, true},
		},
		{
			name:      "different size",
			operator:  FloatPayload([]float64{1, 2, 3, 4, 5}, []bool{false, false, false, false, true}),
			operandum: FloatPayload([]float64{11, 12}, []bool{false, true}),
			outData:   []float64{12, math.NaN(), 14, math.NaN(), math.NaN()},
			outNA:     []bool{false, true, false, true, true},
		},
		{
			name:      "different type",
			operator:  FloatPayload([]float64{1, 2, 3, 4, 0}, []bool{false, false, false, false, true}),
			operandum: IntegerPayload([]int{11, 12, 13, 0, 15}, []bool{false, false, false, true, false}),
			outData:   []float64{12, 14, 16, math.NaN(), math.NaN()},
			outNA:     []bool{false, false, false, true, true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.operator.(Adder).Add(data.operandum).(*floatPayload)

			if !util.EqualFloatArrays(payload.data, data.outData) {
				t.Error(fmt.Sprintf("Data (%v) do not match expected (%v)",
					payload.data, data.outData))
			}

			if !reflect.DeepEqual(payload.NA, data.outNA) {
				t.Error(fmt.Sprintf("NA (%v) do not match expected (%v)",
					payload.NA, data.outNA))
			}
		})
	}
}

func TestFloatPayload_Sub(t *testing.T) {
	testData := []struct {
		name      string
		operator  Payload
		operandum Payload
		outData   []float64
		outNA     []bool
	}{
		{
			name:      "same type",
			operator:  FloatPayload([]float64{21, 32, 43, 4, 0}, []bool{false, false, false, false, true}),
			operandum: FloatPayload([]float64{11, 12, 13, 0, 15}, []bool{false, false, false, true, false}),
			outData:   []float64{10, 20, 30, math.NaN(), math.NaN()},
			outNA:     []bool{false, false, false, true, true},
		},
		{
			name:      "different size",
			operator:  FloatPayload([]float64{11, 22, 33, 44, 55}, []bool{false, false, false, false, true}),
			operandum: FloatPayload([]float64{11, 12}, []bool{false, true}),
			outData:   []float64{0, math.NaN(), 22, math.NaN(), math.NaN()},
			outNA:     []bool{false, true, false, true, true},
		},
		{
			name:      "different type",
			operator:  FloatPayload([]float64{11, 22, 33, 44, 0}, []bool{false, false, false, false, true}),
			operandum: IntegerPayload([]int{11, 12, 13, 0, 15}, []bool{false, false, false, true, false}),
			outData:   []float64{0, 10, 20, math.NaN(), math.NaN()},
			outNA:     []bool{false, false, false, true, true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.operator.(Subber).Sub(data.operandum).(*floatPayload)

			if !util.EqualFloatArrays(payload.data, data.outData) {
				t.Error(fmt.Sprintf("Data (%v) do not match expected (%v)",
					payload.data, data.outData))
			}

			if !reflect.DeepEqual(payload.NA, data.outNA) {
				t.Error(fmt.Sprintf("NA (%v) do not match expected (%v)",
					payload.NA, data.outNA))
			}
		})
	}
}

func TestFloatPayload_Mul(t *testing.T) {
	testData := []struct {
		name      string
		operator  Payload
		operandum Payload
		outData   []float64
		outNA     []bool
	}{
		{
			name:      "same type",
			operator:  FloatPayload([]float64{1, 3, 5, 7, 0}, []bool{false, false, false, false, true}),
			operandum: FloatPayload([]float64{1, 2, 3, 0, 15}, []bool{false, false, false, true, false}),
			outData:   []float64{1, 6, 15, math.NaN(), math.NaN()},
			outNA:     []bool{false, false, false, true, true},
		},
		{
			name:      "different size",
			operator:  FloatPayload([]float64{1, 2, 3, 4, 0}, []bool{false, false, false, false, true}),
			operandum: FloatPayload([]float64{2, 0}, []bool{false, true}),
			outData:   []float64{2, math.NaN(), 6, math.NaN(), math.NaN()},
			outNA:     []bool{false, true, false, true, true},
		},
		{
			name:      "different type",
			operator:  FloatPayload([]float64{1, 2, 3, 4, 0}, []bool{false, false, false, false, true}),
			operandum: IntegerPayload([]int{2, 3, 4, 0, 5}, []bool{false, false, false, true, false}),
			outData:   []float64{2, 6, 12, math.NaN(), math.NaN()},
			outNA:     []bool{false, false, false, true, true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.operator.(Multiplier).Mul(data.operandum).(*floatPayload)

			if !util.EqualFloatArrays(payload.data, data.outData) {
				t.Error(fmt.Sprintf("Data (%v) do not match expected (%v)",
					payload.data, data.outData))
			}

			if !reflect.DeepEqual(payload.NA, data.outNA) {
				t.Error(fmt.Sprintf("NA (%v) do not match expected (%v)",
					payload.NA, data.outNA))
			}
		})
	}
}

func TestFloatPayload_Div(t *testing.T) {
	testData := []struct {
		name      string
		operator  Payload
		operandum Payload
		outData   []float64
		outNA     []bool
	}{
		{
			name:      "same type",
			operator:  FloatPayload([]float64{2, 6, 5, 7, 0}, []bool{false, false, false, false, true}),
			operandum: FloatPayload([]float64{1, 2, 0, 0, 15}, []bool{false, false, false, true, false}),
			outData:   []float64{2, 3, math.NaN(), math.NaN(), math.NaN()},
			outNA:     []bool{false, false, true, true, true},
		},
		{
			name:      "different size",
			operator:  FloatPayload([]float64{2, 4, 6, 8, 0}, []bool{false, false, false, false, true}),
			operandum: FloatPayload([]float64{2, 0}, []bool{false, true}),
			outData:   []float64{1, math.NaN(), 3, math.NaN(), math.NaN()},
			outNA:     []bool{false, true, false, true, true},
		},
		{
			name:      "different type",
			operator:  FloatPayload([]float64{2, 6, 5, 7, 0}, []bool{false, false, false, false, true}),
			operandum: IntegerPayload([]int{1, 2, 0, 0, 15}, []bool{false, false, false, true, false}),
			outData:   []float64{2, 3, math.NaN(), math.NaN(), math.NaN()},
			outNA:     []bool{false, false, true, true, true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.operator.(Divider).Div(data.operandum).(*floatPayload)

			if !util.EqualFloatArrays(payload.data, data.outData) {
				t.Error(fmt.Sprintf("Data (%v) do not match expected (%v)",
					payload.data, data.outData))
			}

			if !reflect.DeepEqual(payload.NA, data.outNA) {
				t.Error(fmt.Sprintf("NA (%v) do not match expected (%v)",
					payload.NA, data.outNA))
			}
		})
	}
}
