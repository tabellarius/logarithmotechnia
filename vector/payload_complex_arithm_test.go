package vector

import (
	"fmt"
	"logarithmotechnia/util"
	"math/cmplx"
	"reflect"
	"testing"
)

func TestComplexPayload_Add(t *testing.T) {
	testData := []struct {
		name      string
		operator  Payload
		operandum Payload
		outData   []complex128
		outNA     []bool
	}{
		{
			name:      "same type",
			operator:  ComplexPayload([]complex128{1, 2, 3, 4, 0}, []bool{false, false, false, false, true}),
			operandum: ComplexPayload([]complex128{11, 12, 13, 0, 15}, []bool{false, false, false, true, false}),
			outData:   []complex128{12, 14, 16, cmplx.NaN(), cmplx.NaN()},
			outNA:     []bool{false, false, false, true, true},
		},
		{
			name:      "different size",
			operator:  ComplexPayload([]complex128{1, 2, 3, 4, 5}, []bool{false, false, false, false, true}),
			operandum: ComplexPayload([]complex128{11, 12}, []bool{false, true}),
			outData:   []complex128{12, cmplx.NaN(), 14, cmplx.NaN(), cmplx.NaN()},
			outNA:     []bool{false, true, false, true, true},
		},
		{
			name:      "different type",
			operator:  ComplexPayload([]complex128{1, 2, 3, 4, 0}, []bool{false, false, false, false, true}),
			operandum: IntegerPayload([]int{11, 12, 13, 0, 15}, []bool{false, false, false, true, false}),
			outData:   []complex128{12, 14, 16, cmplx.NaN(), cmplx.NaN()},
			outNA:     []bool{false, false, false, true, true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.operator.(Adder).Add(data.operandum).(*complexPayload)

			if !util.EqualComplexArrays(payload.data, data.outData) {
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

func TestComplexPayload_Sub(t *testing.T) {
	testData := []struct {
		name      string
		operator  Payload
		operandum Payload
		outData   []complex128
		outNA     []bool
	}{
		{
			name:      "same type",
			operator:  ComplexPayload([]complex128{21, 32, 43, 4, 0}, []bool{false, false, false, false, true}),
			operandum: ComplexPayload([]complex128{11, 12, 13, 0, 15}, []bool{false, false, false, true, false}),
			outData:   []complex128{10, 20, 30, cmplx.NaN(), cmplx.NaN()},
			outNA:     []bool{false, false, false, true, true},
		},
		{
			name:      "different size",
			operator:  ComplexPayload([]complex128{11, 22, 33, 44, 55}, []bool{false, false, false, false, true}),
			operandum: ComplexPayload([]complex128{11, 12}, []bool{false, true}),
			outData:   []complex128{0, cmplx.NaN(), 22, cmplx.NaN(), cmplx.NaN()},
			outNA:     []bool{false, true, false, true, true},
		},
		{
			name:      "different type",
			operator:  ComplexPayload([]complex128{11, 22, 33, 44, 0}, []bool{false, false, false, false, true}),
			operandum: IntegerPayload([]int{11, 12, 13, 0, 15}, []bool{false, false, false, true, false}),
			outData:   []complex128{0, 10, 20, cmplx.NaN(), cmplx.NaN()},
			outNA:     []bool{false, false, false, true, true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.operator.(Subber).Sub(data.operandum).(*complexPayload)

			if !util.EqualComplexArrays(payload.data, data.outData) {
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

func TestComplexPayload_Mul(t *testing.T) {
	testData := []struct {
		name      string
		operator  Payload
		operandum Payload
		outData   []complex128
		outNA     []bool
	}{
		{
			name:      "same type",
			operator:  ComplexPayload([]complex128{1, 3, 5, 7, 0}, []bool{false, false, false, false, true}),
			operandum: ComplexPayload([]complex128{1, 2, 3, 0, 15}, []bool{false, false, false, true, false}),
			outData:   []complex128{1, 6, 15, cmplx.NaN(), cmplx.NaN()},
			outNA:     []bool{false, false, false, true, true},
		},
		{
			name:      "different size",
			operator:  ComplexPayload([]complex128{1, 2, 3, 4, 0}, []bool{false, false, false, false, true}),
			operandum: ComplexPayload([]complex128{2, 0}, []bool{false, true}),
			outData:   []complex128{2, cmplx.NaN(), 6, cmplx.NaN(), cmplx.NaN()},
			outNA:     []bool{false, true, false, true, true},
		},
		{
			name:      "different type",
			operator:  ComplexPayload([]complex128{1, 2, 3, 4, 0}, []bool{false, false, false, false, true}),
			operandum: IntegerPayload([]int{2, 3, 4, 0, 5}, []bool{false, false, false, true, false}),
			outData:   []complex128{2, 6, 12, cmplx.NaN(), cmplx.NaN()},
			outNA:     []bool{false, false, false, true, true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.operator.(Multiplier).Mul(data.operandum).(*complexPayload)

			if !util.EqualComplexArrays(payload.data, data.outData) {
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

func TestComplexPayload_Div(t *testing.T) {
	testData := []struct {
		name      string
		operator  Payload
		operandum Payload
		outData   []complex128
		outNA     []bool
	}{
		{
			name:      "same type",
			operator:  ComplexPayload([]complex128{2, 6, 5, 7, 0}, []bool{false, false, false, false, true}),
			operandum: ComplexPayload([]complex128{1, 2, 0, 0, 15}, []bool{false, false, false, true, false}),
			outData:   []complex128{2, 3, cmplx.NaN(), cmplx.NaN(), cmplx.NaN()},
			outNA:     []bool{false, false, true, true, true},
		},
		{
			name:      "different size",
			operator:  ComplexPayload([]complex128{2, 4, 6, 8, 0}, []bool{false, false, false, false, true}),
			operandum: ComplexPayload([]complex128{2, 0}, []bool{false, true}),
			outData:   []complex128{1, cmplx.NaN(), 3, cmplx.NaN(), cmplx.NaN()},
			outNA:     []bool{false, true, false, true, true},
		},
		{
			name:      "different type",
			operator:  ComplexPayload([]complex128{2, 6, 5, 7, 0}, []bool{false, false, false, false, true}),
			operandum: IntegerPayload([]int{1, 2, 0, 0, 15}, []bool{false, false, false, true, false}),
			outData:   []complex128{2, 3, cmplx.NaN(), cmplx.NaN(), cmplx.NaN()},
			outNA:     []bool{false, false, true, true, true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.operator.(Divider).Div(data.operandum).(*complexPayload)

			if !util.EqualComplexArrays(payload.data, data.outData) {
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
