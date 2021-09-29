package vector

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIntegerPayload_Sum(t *testing.T) {
	testData := []struct {
		name    string
		payload *integerPayload
		sumData []int
		sumNA   []bool
	}{
		{
			name:    "without na",
			payload: IntegerPayload([]int{-20, 10, 4, -20, 27}, nil).(*integerPayload),
			sumData: []int{1},
			sumNA:   []bool{false},
		},
		{
			name:    "with na",
			payload: IntegerPayload([]int{-20, 10, 4, -20, 27}, []bool{false, false, true, false, false}).(*integerPayload),
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
