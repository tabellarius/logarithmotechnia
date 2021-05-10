package vector

import (
	"fmt"
	"strconv"
	"testing"
)

func TestOdd(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	out := []bool{true, false, true, false, true, false, true, false, true, false}
	fOdd := Odd()

	for i := 0; i < len(in); i++ {
		result := fOdd(in[i], in[i], false)
		if out[i] != result {
			t.Error(fmt.Sprintf("Result (%t) for in[%d] is not equal to %t", result, in[i], out[i]))
			break
		}
	}
}

func TestEven(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	out := []bool{false, true, false, true, false, true, false, true, false, true}
	fEven := Even()

	for i := 0; i < len(in); i++ {
		result := fEven(in[i], in[i], false)
		if out[i] != result {
			t.Error(fmt.Sprintf("Result (%t) for in[%d] is not equal to %t", result, in[i], out[i]))
			break
		}
	}
}

func TestNth(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	testData := []struct {
		nth int
		out []bool
	}{
		{
			nth: 3,
			out: []bool{false, false, true, false, false, true, false, false, true, false},
		},
		{
			nth: 4,
			out: []bool{false, false, false, true, false, false, false, true, false, false},
		},
		{
			nth: 5,
			out: []bool{false, false, false, false, true, false, false, false, false, true},
		},
		{
			nth: 10,
			out: []bool{false, false, false, false, false, false, false, false, false, true},
		},
	}

	for _, data := range testData {
		t.Run(strconv.Itoa(data.nth), func(t *testing.T) {
			fNth := Nth(data.nth)

			for i := 0; i < len(in); i++ {
				result := fNth(in[i], in[i], false)
				if data.out[i] != result {
					t.Error(fmt.Sprintf("Result (%t) for in[%d] is not equal to %t", result, in[i], data.out[i]))
					break
				}
			}
		})
	}
}
