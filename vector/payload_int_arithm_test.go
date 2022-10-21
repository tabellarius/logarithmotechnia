package vector

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIntegerPayload_Add(t *testing.T) {
	testData := []struct {
		name      string
		operator  Payload
		operandum Payload
		outData   []int
		outNA     []bool
	}{
		{
			name:      "same type",
			operator:  IntegerPayload([]int{1, 2, 3, 4, 0}, []bool{false, false, false, false, true}),
			operandum: IntegerPayload([]int{11, 12, 13, 0, 15}, []bool{false, false, false, true, false}),
			outData:   []int{12, 14, 16, 0, 0},
			outNA:     []bool{false, false, false, true, true},
		},
		{
			name:      "different size",
			operator:  IntegerPayload([]int{1, 2, 3, 4, 5}, []bool{false, false, false, false, true}),
			operandum: IntegerPayload([]int{11, 12}, []bool{false, true}),
			outData:   []int{12, 0, 14, 0, 0},
			outNA:     []bool{false, true, false, true, true},
		},
		{
			name:      "different type",
			operator:  IntegerPayload([]int{1, 2, 3, 4, 0}, []bool{false, false, false, false, true}),
			operandum: FloatPayload([]float64{11, 12, 13, 0, 15}, []bool{false, false, false, true, false}),
			outData:   []int{12, 14, 16, 0, 0},
			outNA:     []bool{false, false, false, true, true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.operator.(Adder).Add(data.operandum).(*integerPayload)

			if !reflect.DeepEqual(payload.data, data.outData) {
				t.Error(fmt.Sprintf("Data (%v) do not match expected (%v)",
					payload.data, data.outData))
			}

			if !reflect.DeepEqual(payload.na, data.outNA) {
				t.Error(fmt.Sprintf("NA (%v) do not match expected (%v)",
					payload.na, data.outNA))
			}
		})
	}
}

func TestIntegerPayload_Sub(t *testing.T) {
	testData := []struct {
		name      string
		operator  Payload
		operandum Payload
		outData   []int
		outNA     []bool
	}{
		{
			name:      "same type",
			operator:  IntegerPayload([]int{21, 32, 43, 4, 0}, []bool{false, false, false, false, true}),
			operandum: IntegerPayload([]int{11, 12, 13, 0, 15}, []bool{false, false, false, true, false}),
			outData:   []int{10, 20, 30, 0, 0},
			outNA:     []bool{false, false, false, true, true},
		},
		{
			name:      "different size",
			operator:  IntegerPayload([]int{11, 22, 33, 44, 55}, []bool{false, false, false, false, true}),
			operandum: IntegerPayload([]int{11, 12}, []bool{false, true}),
			outData:   []int{0, 0, 22, 0, 0},
			outNA:     []bool{false, true, false, true, true},
		},
		{
			name:      "different type",
			operator:  IntegerPayload([]int{11, 22, 33, 44, 0}, []bool{false, false, false, false, true}),
			operandum: FloatPayload([]float64{11, 12, 13, 0, 15}, []bool{false, false, false, true, false}),
			outData:   []int{0, 10, 20, 0, 0},
			outNA:     []bool{false, false, false, true, true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.operator.(Subber).Sub(data.operandum).(*integerPayload)

			if !reflect.DeepEqual(payload.data, data.outData) {
				t.Error(fmt.Sprintf("Data (%v) do not match expected (%v)",
					payload.data, data.outData))
			}

			if !reflect.DeepEqual(payload.na, data.outNA) {
				t.Error(fmt.Sprintf("NA (%v) do not match expected (%v)",
					payload.na, data.outNA))
			}
		})
	}
}

func TestIntegerPayload_Mul(t *testing.T) {
	testData := []struct {
		name      string
		operator  Payload
		operandum Payload
		outData   []int
		outNA     []bool
	}{
		{
			name:      "same type",
			operator:  IntegerPayload([]int{1, 3, 5, 7, 0}, []bool{false, false, false, false, true}),
			operandum: IntegerPayload([]int{1, 2, 3, 0, 15}, []bool{false, false, false, true, false}),
			outData:   []int{1, 6, 15, 0, 0},
			outNA:     []bool{false, false, false, true, true},
		},
		{
			name:      "different size",
			operator:  IntegerPayload([]int{1, 2, 3, 4, 0}, []bool{false, false, false, false, true}),
			operandum: IntegerPayload([]int{2, 0}, []bool{false, true}),
			outData:   []int{2, 0, 6, 0, 0},
			outNA:     []bool{false, true, false, true, true},
		},
		{
			name:      "different type",
			operator:  IntegerPayload([]int{1, 2, 3, 4, 0}, []bool{false, false, false, false, true}),
			operandum: FloatPayload([]float64{2, 3, 4, 0, 5}, []bool{false, false, false, true, false}),
			outData:   []int{2, 6, 12, 0, 0},
			outNA:     []bool{false, false, false, true, true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.operator.(Multiplier).Mul(data.operandum).(*integerPayload)

			if !reflect.DeepEqual(payload.data, data.outData) {
				t.Error(fmt.Sprintf("Data (%v) do not match expected (%v)",
					payload.data, data.outData))
			}

			if !reflect.DeepEqual(payload.na, data.outNA) {
				t.Error(fmt.Sprintf("NA (%v) do not match expected (%v)",
					payload.na, data.outNA))
			}
		})
	}
}

func TestIntegerPayload_Div(t *testing.T) {
	testData := []struct {
		name      string
		operator  Payload
		operandum Payload
		outData   []int
		outNA     []bool
	}{
		{
			name:      "same type",
			operator:  IntegerPayload([]int{2, 6, 5, 7, 0}, []bool{false, false, false, false, true}),
			operandum: IntegerPayload([]int{1, 2, 0, 0, 15}, []bool{false, false, false, true, false}),
			outData:   []int{2, 3, 0, 0, 0},
			outNA:     []bool{false, false, true, true, true},
		},
		{
			name:      "different size",
			operator:  IntegerPayload([]int{2, 4, 6, 8, 0}, []bool{false, false, false, false, true}),
			operandum: IntegerPayload([]int{2, 0}, []bool{false, true}),
			outData:   []int{1, 0, 3, 0, 0},
			outNA:     []bool{false, true, false, true, true},
		},
		{
			name:      "different type",
			operator:  IntegerPayload([]int{2, 6, 5, 7, 0}, []bool{false, false, false, false, true}),
			operandum: FloatPayload([]float64{1, 2, 0, 0, 15}, []bool{false, false, false, true, false}),
			outData:   []int{2, 3, 0, 0, 0},
			outNA:     []bool{false, false, true, true, true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.operator.(Divider).Div(data.operandum).(*integerPayload)

			if !reflect.DeepEqual(payload.data, data.outData) {
				t.Error(fmt.Sprintf("Data (%v) do not match expected (%v)",
					payload.data, data.outData))
			}

			if !reflect.DeepEqual(payload.na, data.outNA) {
				t.Error(fmt.Sprintf("NA (%v) do not match expected (%v)",
					payload.na, data.outNA))
			}
		})
	}
}
