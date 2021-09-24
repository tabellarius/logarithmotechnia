package vector

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBoolPayload_Sum(t *testing.T) {
	testData := []struct {
		name    string
		payload *booleanPayload
		sumData []int
		sumNA   []bool
	}{
		{
			name:    "without na",
			payload: BooleanPayload([]bool{true, false, true, false, true}, nil).(*booleanPayload),
			sumData: []int{3},
			sumNA:   []bool{false},
		},
		{
			name: "with na",
			payload: BooleanPayload([]bool{true, false, true, false, true},
				[]bool{false, false, true, false, false}).(*booleanPayload),
			sumData: []int{0},
			sumNA:   []bool{true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			sumPayload := data.payload.Sum().(*integerPayload)

			if !reflect.DeepEqual(sumPayload.data, data.sumData) {
				t.Error(fmt.Sprintf("Sum data (%v) is not equal to expected (%v)",
					sumPayload.data, data.sumData))
			}

			if !reflect.DeepEqual(sumPayload.na, data.sumNA) {
				t.Error(fmt.Sprintf("Sum data (%v) is not equal to expected (%v)",
					sumPayload.na, data.sumNA))
			}
		})
	}
}
