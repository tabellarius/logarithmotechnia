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

			if !reflect.DeepEqual(payload.na, data.sumNA) {
				t.Error(fmt.Sprintf("Max na (%v) is not equal to expected (%v)",
					payload.na, data.sumNA))
			}
		})
	}
}
