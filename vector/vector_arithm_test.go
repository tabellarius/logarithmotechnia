package vector

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestVector_Add(t *testing.T) {
	vec := IntegerWithNA([]int{1, 2, 3, 4, 5, 0, 7, 8, 9, 10}, []bool{false, false, false, false, false, true,
		false, false, false, false})
	vectors := []Vector{
		IntegerWithNA([]int{1, 4, 9, 16, 25, 36, 0, 64, 81, 100}, []bool{false, false, false, false, false, false,
			true, false, false, false}),
		Float([]float64{2, 4, 6, 8, 10, 12, 14, math.NaN(), math.Inf(1), math.Inf(-1)}),
		Integer([]int{1}),
	}
	outData := []int{5, 11, 19, 29, 41, 0, 0, 0, 0, 0}
	outNA := []bool{false, false, false, false, false, true, true, true, true, true}

	outVec := vec.Add(vectors...)
	payload := outVec.Payload().(*integerPayload)

	if !reflect.DeepEqual(payload.data, outData) {
		t.Error(fmt.Sprintf("payload.data (%v) do not match outData (%v)",
			payload.data, outData))
	}

	if !reflect.DeepEqual(payload.NA, outNA) {
		t.Error(fmt.Sprintf("payload.NA (%v) do not match outNA (%v)",
			payload.NA, outNA))
	}
}

func TestVector_Sub(t *testing.T) {
	vec := IntegerWithNA([]int{1, 2, 3, 4, 5, 0, 7, 8, 9, 10}, []bool{false, false, false, false, false, true,
		false, false, false, false})
	vectors := []Vector{
		IntegerWithNA([]int{1, 4, 9, 16, 25, 36, 0, 64, 81, 100}, []bool{false, false, false, false, false, false,
			true, false, false, false}),
		Float([]float64{2, 4, 6, 8, 10, 12, 14, math.NaN(), math.Inf(1), math.Inf(-1)}),
		Integer([]int{1}),
	}
	outData := []int{-3, -7, -13, -21, -31, 0, 0, 0, 0, 0}
	outNA := []bool{false, false, false, false, false, true, true, true, true, true}

	outVec := vec.Sub(vectors...)
	payload := outVec.Payload().(*integerPayload)

	if !reflect.DeepEqual(payload.data, outData) {
		t.Error(fmt.Sprintf("payload.data (%v) do not match outData (%v)",
			payload.data, outData))
	}

	if !reflect.DeepEqual(payload.NA, outNA) {
		t.Error(fmt.Sprintf("payload.NA (%v) do not match outNA (%v)",
			payload.NA, outNA))
	}
}

func TestVector_Mul(t *testing.T) {
	vec := IntegerWithNA([]int{1, 2, 3, 4, 5, 0, 7, 8, 9, 10}, []bool{false, false, false, false, false, true,
		false, false, false, false})
	vectors := []Vector{
		IntegerWithNA([]int{1, 2, 3, 4, 5, 6, 0, 64, 81, 100}, []bool{false, false, false, false, false, false,
			true, false, false, false}),
		Float([]float64{2, 3, 4, 5, 6, 7, 8, math.NaN(), math.Inf(1), math.Inf(-1)}),
		Integer([]int{2}),
	}
	outData := []int{4, 24, 72, 160, 300, 0, 0, 0, 0, 0}
	outNA := []bool{false, false, false, false, false, true, true, true, true, true}

	outVec := vec.Mul(vectors...)
	payload := outVec.Payload().(*integerPayload)

	if !reflect.DeepEqual(payload.data, outData) {
		t.Error(fmt.Sprintf("payload.data (%v) do not match outData (%v)",
			payload.data, outData))
	}

	if !reflect.DeepEqual(payload.NA, outNA) {
		t.Error(fmt.Sprintf("payload.NA (%v) do not match outNA (%v)",
			payload.NA, outNA))
	}
}

func TestVector_Div(t *testing.T) {
	vec := IntegerWithNA([]int{8, 36, 96, 200, 5, 0, 7, 8, 9, 10}, []bool{false, false, false, false, false, true,
		false, false, false, false})
	vectors := []Vector{
		IntegerWithNA([]int{1, 2, 3, 4, 5, 6, 0, 64, 81, 100}, []bool{false, false, false, false, false, false,
			true, false, false, false}),
		Float([]float64{2, 3, 4, 5, 0, 7, 8, math.NaN(), math.Inf(1), math.Inf(-1)}),
		Integer([]int{2}),
	}
	outData := []int{2, 3, 4, 5, 0, 0, 0, 0, 0, 0}
	outNA := []bool{false, false, false, false, true, true, true, true, true, true}

	outVec := vec.Div(vectors...)
	payload := outVec.Payload().(*integerPayload)

	if !reflect.DeepEqual(payload.data, outData) {
		t.Error(fmt.Sprintf("payload.data (%v) do not match outData (%v)",
			payload.data, outData))
	}

	if !reflect.DeepEqual(payload.NA, outNA) {
		t.Error(fmt.Sprintf("payload.NA (%v) do not match outNA (%v)",
			payload.NA, outNA))
	}
}
