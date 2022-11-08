package vector

import (
	"fmt"
	"logarithmotechnia/util"
	"math/cmplx"
	"reflect"
	"testing"
)

func TestComplexPayload_Sum(t *testing.T) {
	testData := []struct {
		name    string
		payload *complexPayload
		sumData []complex128
		sumNA   []bool
	}{
		{
			name:    "without na",
			payload: ComplexPayload([]complex128{-20 + 10i, 10 - 5i, 4 + 2i, -20 + 20i, 27 - 26i}, nil).(*complexPayload),
			sumData: []complex128{1 + 1i},
			sumNA:   []bool{false},
		},
		{
			name: "with na",
			payload: ComplexPayload([]complex128{-20 + 10i, 10 - 5i, 4 + 2i, -20 + 20i, 27 - 26i},
				[]bool{false, false, true, false, false}).(*complexPayload),
			sumData: []complex128{cmplx.NaN()},
			sumNA:   []bool{true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			sumPayload := data.payload.Sum().(*complexPayload)

			if !util.EqualComplexArrays(sumPayload.data, data.sumData) {
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
