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
		sumData []float64
		sumNA   []bool
	}{
		{
			name:    "without na",
			payload: FloatPayload([]float64{-20, 10, 4, -20, 27}, nil).(*floatPayload),
			sumData: []float64{1},
			sumNA:   []bool{false},
		},
		{
			name:    "with na",
			payload: FloatPayload([]float64{-20, 10, 4, -20, 27}, []bool{false, false, true, false, false}).(*floatPayload),
			sumData: []float64{math.NaN()},
			sumNA:   []bool{true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			sumPayload := data.payload.Sum().(*floatPayload)

			if !util.EqualFloatArrays(sumPayload.data, data.sumData) {
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
