package vector

import (
	"fmt"
	"testing"
)

func TestVector_Sum(t *testing.T) {
	testData := []struct {
		name   string
		vec    Vector
		sumVec Vector
	}{
		{
			name:   "normal summer",
			vec:    Integer([]int{10, 2, 8, 12, 18}),
			sumVec: Integer([]int{50}),
		},
		{
			name:   "normal non-summer",
			vec:    String([]string{"one", "two", "8", "12", "18"}),
			sumVec: NA(1),
		},
		{
			name:   "normal grouped summer",
			vec:    Integer([]int{10, 2, 10, 12, 0, 10, 2}).GroupByIndices([][]int{{1, 3, 6}, {2, 4, 7}, {5}}),
			sumVec: Integer([]int{30, 16, 0}),
		},
		{
			name:   "normal grouped non-summer",
			vec:    String([]string{"one", "two", "8", "12", "18"}).GroupByIndices([][]int{{1, 3}, {2, 5}, {4}}),
			sumVec: NA(3),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			sumVec := data.vec.Sum()

			if !CompareVectorsForTest(sumVec, data.sumVec) {
				t.Error(fmt.Sprintf("Sum vector (%v) does not match expected (%v)",
					sumVec, data.sumVec))
			}
		})
	}
}
