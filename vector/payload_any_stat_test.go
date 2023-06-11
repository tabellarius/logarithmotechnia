package vector

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAnyPayload_Max(t *testing.T) {
	callbacks := AnyCallbacks{
		Lt: func(one, two any) bool { return one.(int) < two.(int) },
		Eq: func(one, two any) bool { return one.(int) == two.(int) },
	}

	testData := []struct {
		name    string
		payload *anyPayload
		data    []any
		sumNA   []bool
	}{
		{
			name:    "without na",
			payload: AnyPayload([]any{-20, 10, 4, -20, 27}, nil, OptionAnyCallbacks(callbacks)).(*anyPayload),
			data:    []any{27},
			sumNA:   []bool{false},
		},
		{
			name:    "without functions",
			payload: AnyPayload([]any{-20, 10, 4, -20, 27}, nil).(*anyPayload),
			data:    []any{nil},
			sumNA:   []bool{true},
		},
		{
			name:    "with na",
			payload: AnyPayload([]any{-20, 10, 4, -20, 27}, []bool{false, false, true, false, false}).(*anyPayload),
			data:    []any{nil},
			sumNA:   []bool{true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.payload.Max().(*anyPayload)

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

func TestAnyPayload_Min(t *testing.T) {
	callbacks := AnyCallbacks{
		Lt: func(one, two any) bool { return one.(int) < two.(int) },
		Eq: func(one, two any) bool { return one.(int) == two.(int) },
	}

	testData := []struct {
		name    string
		payload *anyPayload
		data    []any
		sumNA   []bool
	}{
		{
			name:    "without na",
			payload: AnyPayload([]any{-20, 10, 4, -20, 27}, nil, OptionAnyCallbacks(callbacks)).(*anyPayload),
			data:    []any{-20},
			sumNA:   []bool{false},
		},
		{
			name:    "without functions",
			payload: AnyPayload([]any{-20, 10, 4, -20, 27}, nil).(*anyPayload),
			data:    []any{nil},
			sumNA:   []bool{true},
		},
		{
			name:    "with na",
			payload: AnyPayload([]any{-20, 10, 4, -20, 27}, []bool{false, false, true, false, false}).(*anyPayload),
			data:    []any{nil},
			sumNA:   []bool{true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.payload.Min().(*anyPayload)

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

func TestAnyPayload_CumMax(t *testing.T) {
	callbacks := AnyCallbacks{
		Lt: func(one, two any) bool { return one.(int) < two.(int) },
		Eq: func(one, two any) bool { return one.(int) == two.(int) },
	}

	testData := []struct {
		name    string
		payload *anyPayload
		data    []any
		sumNA   []bool
	}{
		{
			name:    "without na",
			payload: AnyPayload([]any{-20, 10, 4, -20, 27}, nil, OptionAnyCallbacks(callbacks)).(*anyPayload),
			data:    []any{-20, 10, 10, 10, 27},
			sumNA:   []bool{false, false, false, false, false},
		},
		{
			name:    "without functions",
			payload: AnyPayload([]any{-20, 10, 4, -20, 27}, nil).(*anyPayload),
			data:    []any{nil, nil, nil, nil, nil},
			sumNA:   []bool{true, true, true, true, true},
		},
		{
			name: "with na",
			payload: AnyPayload(
				[]any{-20, 10, 4, -20, 27},
				[]bool{false, false, true, false, false},
				OptionAnyCallbacks(callbacks),
			).(*anyPayload),
			data:  []any{-20, 10, nil, nil, nil},
			sumNA: []bool{false, false, true, true, true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.payload.CumMax().(*anyPayload)

			if !reflect.DeepEqual(payload.data, data.data) {
				t.Error(fmt.Sprintf("CumMax data (%v) is not equal to expected (%v)",
					payload.data, data.data))
			}

			if !reflect.DeepEqual(payload.NA, data.sumNA) {
				t.Error(fmt.Sprintf("CumMax na (%v) is not equal to expected (%v)",
					payload.NA, data.sumNA))
			}
		})
	}
}

func TestAnyPayload_CumMin(t *testing.T) {
	callbacks := AnyCallbacks{
		Lt: func(one, two any) bool { return one.(int) < two.(int) },
		Eq: func(one, two any) bool { return one.(int) == two.(int) },
	}

	testData := []struct {
		name    string
		payload *anyPayload
		data    []any
		sumNA   []bool
	}{
		{
			name:    "without na",
			payload: AnyPayload([]any{-10, 10, 4, -20, 27}, nil, OptionAnyCallbacks(callbacks)).(*anyPayload),
			data:    []any{-10, -10, -10, -20, -20},
			sumNA:   []bool{false, false, false, false, false},
		},
		{
			name:    "without functions",
			payload: AnyPayload([]any{-20, 10, 4, -20, 27}, nil).(*anyPayload),
			data:    []any{nil, nil, nil, nil, nil},
			sumNA:   []bool{true, true, true, true, true},
		},
		{
			name: "with na",
			payload: AnyPayload(
				[]any{-20, 10, 4, -20, 27},
				[]bool{false, false, true, false, false},
				OptionAnyCallbacks(callbacks),
			).(*anyPayload),
			data:  []any{-20, -20, nil, nil, nil},
			sumNA: []bool{false, false, true, true, true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.payload.CumMin().(*anyPayload)

			if !reflect.DeepEqual(payload.data, data.data) {
				t.Error(fmt.Sprintf("CumMin data (%v) is not equal to expected (%v)",
					payload.data, data.data))
			}

			if !reflect.DeepEqual(payload.NA, data.sumNA) {
				t.Error(fmt.Sprintf("CumMin na (%v) is not equal to expected (%v)",
					payload.NA, data.sumNA))
			}
		})
	}
}
