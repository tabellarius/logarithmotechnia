package vector

import (
	"fmt"
	"logarithmotechnia/internal/util"
	"math"
	"math/cmplx"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestNA(t *testing.T) {
	testData := []struct {
		name      string
		inLength  int
		outLength int
	}{
		{
			name:      "normal",
			inLength:  5,
			outLength: 5,
		},
		{
			name:      "zero",
			inLength:  0,
			outLength: 0,
		},
		{
			name:      "negative",
			inLength:  -1,
			outLength: 0,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := NA(data.inLength).(*vector).payload.(*naPayload)
			if payload.length != data.outLength {
				t.Error(fmt.Sprintf("payload.length (%d) is not equal to expected (%d)",
					payload.length, data.outLength))
			}
		})
	}
}

func TestNAPayload(t *testing.T) {
	testData := []struct {
		name      string
		inLength  int
		outLength int
	}{
		{
			name:      "normal",
			inLength:  5,
			outLength: 5,
		},
		{
			name:      "zero",
			inLength:  0,
			outLength: 0,
		},
		{
			name:      "negative",
			inLength:  -1,
			outLength: 0,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := NAPayload(data.inLength).(*naPayload)
			if payload.length != data.outLength {
				t.Error(fmt.Sprintf("payload.length (%d) is not equal to expected (%d)",
					payload.length, data.outLength))
			}
		})
	}
}

func TestNaPayload_Type(t *testing.T) {
	vec := NA(5)
	if vec.Type() != "na" {
		t.Error("Type is incorrect.")
	}
}

func TestNaPayload_Len(t *testing.T) {
	testData := []struct {
		inLength  int
		outLength int
	}{
		{5, 5},
		{3, 3},
		{0, 0},
		{-10, 0},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			payload := NA(data.inLength).(*vector).payload
			if payload.Len() != data.outLength {
				t.Error(fmt.Sprintf("Payloads's length (%d) is not equal to out (%d)",
					payload.Len(), data.outLength))
			}
		})
	}
}

func TestNaPayload_ByIndices(t *testing.T) {
	testData := []struct {
		indices   []int
		outLength int
	}{
		{
			indices:   []int{1, 2, 3, 4, 5},
			outLength: 5,
		},
		{
			indices:   []int{5, 3, 1},
			outLength: 3,
		},
		{
			indices:   []int{},
			outLength: 0,
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			payload := NA(5).ByIndices(data.indices).(*vector).payload
			if payload.Len() != data.outLength {
				t.Error(fmt.Sprintf("Payloads's length (%d) is not equal to out (%d)",
					payload.Len(), data.outLength))
			}
		})
	}
}

func TestNaPayload_IsNA(t *testing.T) {
	testData := []struct {
		length int
		out    []bool
	}{
		{
			length: 5,
			out:    []bool{true, true, true, true, true},
		},
		{
			length: 3,
			out:    []bool{true, true, true},
		},
		{
			length: 0,
			out:    []bool{},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			payload := NA(data.length).(*vector).payload.(*naPayload)
			na := payload.IsNA()
			if !reflect.DeepEqual(na, data.out) {
				t.Error(fmt.Sprintf("payload.isNA() (%v) is not equal to expected (%v)",
					na, data.out))
			}
		})
	}
}

func TestNaPayload_NotNA(t *testing.T) {
	testData := []struct {
		length int
		out    []bool
	}{
		{
			length: 5,
			out:    []bool{false, false, false, false, false},
		},
		{
			length: 3,
			out:    []bool{false, false, false},
		},
		{
			length: 0,
			out:    []bool{},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			payload := NA(data.length).(*vector).payload.(*naPayload)
			notNA := payload.NotNA()
			if !reflect.DeepEqual(notNA, data.out) {
				t.Error(fmt.Sprintf("payload.notNA() (%v) is not equal to expected (%v)",
					notNA, data.out))
			}
		})
	}
}

func TestNaPayload_HasNA(t *testing.T) {
	testData := []struct {
		length int
		hasNA  bool
	}{
		{5, true}, {0, false}, {-1, false},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			payload := NA(data.length).(*vector).payload.(*naPayload)
			hasNA := payload.HasNA()
			if !reflect.DeepEqual(hasNA, data.hasNA) {
				t.Error(fmt.Sprintf("payload.hasNA() (%v) is not equal to expected (%v)",
					hasNA, data.hasNA))
			}
		})
	}
}

func TestNaPayload_WithNA(t *testing.T) {
	testData := []struct {
		length int
		out    []int
	}{
		{
			length: 5,
			out:    []int{1, 2, 3, 4, 5},
		},
		{
			length: 3,
			out:    []int{1, 2, 3},
		},
		{
			length: 0,
			out:    []int{},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			payload := NA(data.length).(*vector).payload.(*naPayload)
			withNA := payload.WithNA()
			if !reflect.DeepEqual(withNA, data.out) {
				t.Error(fmt.Sprintf("payload.withNA() (%v) is not equal to expected (%v)",
					withNA, data.out))
			}
		})
	}
}

func TestNaPayload_WithoutNA(t *testing.T) {
	testData := []struct {
		length int
		out    []int
	}{
		{
			length: 5,
			out:    []int{},
		},
		{
			length: 3,
			out:    []int{},
		},
		{
			length: 0,
			out:    []int{},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			payload := NA(data.length).(*vector).payload.(*naPayload)
			withoutNA := payload.WithoutNA()
			if !reflect.DeepEqual(withoutNA, data.out) {
				t.Error(fmt.Sprintf("payload.withoutNA() (%v) is not equal to expected (%v)",
					withoutNA, data.out))
			}
		})
	}
}

func TestNaPayload_Integers(t *testing.T) {
	testData := []struct {
		length  int
		outData []int
		outNA   []bool
	}{
		{
			length:  5,
			outData: []int{0, 0, 0, 0, 0},
			outNA:   []bool{true, true, true, true, true},
		},
		{
			length:  3,
			outData: []int{0, 0, 0},
			outNA:   []bool{true, true, true},
		},
		{
			length:  0,
			outData: []int{},
			outNA:   []bool{},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			payload := NA(data.length).(*vector).payload.(*naPayload)
			integers, na := payload.Integers()
			if !reflect.DeepEqual(integers, data.outData) {
				t.Error(fmt.Sprintf("Integers (%v) are not equal to expected (%v)",
					integers, data.outData))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("NA (%v) are not equal to expected (%v)",
					na, data.outNA))
			}
		})
	}
}

func TestNaPayload_Floats(t *testing.T) {
	testData := []struct {
		length  int
		outData []float64
		outNA   []bool
	}{
		{
			length:  5,
			outData: []float64{math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN()},
			outNA:   []bool{true, true, true, true, true},
		},
		{
			length:  3,
			outData: []float64{math.NaN(), math.NaN(), math.NaN()},
			outNA:   []bool{true, true, true},
		},
		{
			length:  0,
			outData: []float64{},
			outNA:   []bool{},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			payload := NA(data.length).(*vector).payload.(*naPayload)
			floats, na := payload.Floats()
			if !util.EqualFloatArrays(floats, data.outData) {
				t.Error(fmt.Sprintf("Floats (%v) are not equal to expected (%v)",
					floats, data.outData))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("NA (%v) are not equal to expected (%v)",
					na, data.outNA))
			}
		})
	}
}

func TestNaPayload_Complexes(t *testing.T) {
	testData := []struct {
		length  int
		outData []complex128
		outNA   []bool
	}{
		{
			length:  5,
			outData: []complex128{cmplx.NaN(), cmplx.NaN(), cmplx.NaN(), cmplx.NaN(), cmplx.NaN()},
			outNA:   []bool{true, true, true, true, true},
		},
		{
			length:  3,
			outData: []complex128{cmplx.NaN(), cmplx.NaN(), cmplx.NaN()},
			outNA:   []bool{true, true, true},
		},
		{
			length:  0,
			outData: []complex128{},
			outNA:   []bool{},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			payload := NA(data.length).(*vector).payload.(*naPayload)
			complexes, na := payload.Complexes()
			if !util.EqualComplexArrays(complexes, data.outData) {
				t.Error(fmt.Sprintf("Complexes (%v) are not equal to expected (%v)",
					complexes, data.outData))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("NA (%v) are not equal to expected (%v)",
					na, data.outNA))
			}
		})
	}
}

func TestNaPayload_Strings(t *testing.T) {
	testData := []struct {
		length  int
		outData []string
		outNA   []bool
	}{
		{
			length:  5,
			outData: []string{"", "", "", "", ""},
			outNA:   []bool{true, true, true, true, true},
		},
		{
			length:  3,
			outData: []string{"", "", ""},
			outNA:   []bool{true, true, true},
		},
		{
			length:  0,
			outData: []string{},
			outNA:   []bool{},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			payload := NA(data.length).(*vector).payload.(*naPayload)
			strings, na := payload.Strings()
			if !reflect.DeepEqual(strings, data.outData) {
				t.Error(fmt.Sprintf("Strings (%v) are not equal to expected (%v)",
					strings, data.outData))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("NA (%v) are not equal to expected (%v)",
					na, data.outNA))
			}
		})
	}
}

func TestNaPayload_Booleans(t *testing.T) {
	testData := []struct {
		length  int
		outData []bool
		outNA   []bool
	}{
		{
			length:  5,
			outData: []bool{false, false, false, false, false},
			outNA:   []bool{true, true, true, true, true},
		},
		{
			length:  3,
			outData: []bool{false, false, false},
			outNA:   []bool{true, true, true},
		},
		{
			length:  0,
			outData: []bool{},
			outNA:   []bool{},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			payload := NA(data.length).(*vector).payload.(*naPayload)
			booleans, na := payload.Booleans()
			if !reflect.DeepEqual(booleans, data.outData) {
				t.Error(fmt.Sprintf("Booleans (%v) are not equal to expected (%v)",
					booleans, data.outData))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("NA (%v) are not equal to expected (%v)",
					na, data.outNA))
			}
		})
	}
}

func TestNaPayload_Times(t *testing.T) {
	testData := []struct {
		length  int
		outData []time.Time
		outNA   []bool
	}{
		{
			length:  5,
			outData: []time.Time{{}, {}, {}, {}, {}},
			outNA:   []bool{true, true, true, true, true},
		},
		{
			length:  3,
			outData: []time.Time{{}, {}, {}},
			outNA:   []bool{true, true, true},
		},
		{
			length:  0,
			outData: []time.Time{},
			outNA:   []bool{},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			payload := NA(data.length).(*vector).payload.(*naPayload)
			times, na := payload.Times()
			if !reflect.DeepEqual(times, data.outData) {
				t.Error(fmt.Sprintf("Times (%v) are not equal to expected (%v)",
					times, data.outData))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("NA (%v) are not equal to expected (%v)",
					na, data.outNA))
			}
		})
	}
}

func TestNaPayload_Interfaces(t *testing.T) {
	testData := []struct {
		length  int
		outData []any
		outNA   []bool
	}{
		{
			length:  5,
			outData: []any{nil, nil, nil, nil, nil},
			outNA:   []bool{true, true, true, true, true},
		},
		{
			length:  3,
			outData: []any{nil, nil, nil},
			outNA:   []bool{true, true, true},
		},
		{
			length:  0,
			outData: []any{},
			outNA:   []bool{},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			payload := NA(data.length).(*vector).payload.(*naPayload)

			interfaces, na := payload.Anies()
			if !reflect.DeepEqual(interfaces, data.outData) {
				t.Error(fmt.Sprintf("Anies (%v) are not equal to expected (%v)",
					interfaces, data.outData))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("NA (%v) are not equal to expected (%v)",
					na, data.outNA))
			}
		})
	}
}

func TestNaPayload_StrForElem(t *testing.T) {
	payload := NA(5).(*vector).payload.(*naPayload)

	for i := 1; i <= 5; i++ {
		if payload.StrForElem(i) != "NA" {
			t.Error(fmt.Sprintf("StrForElem is not equal to 'NA' for %d-th element", i))
		}
	}

}

func TestNaPayload_Append(t *testing.T) {
	payload := NAPayload(3)

	testData := []struct {
		name   string
		vec    Vector
		outLen int
	}{
		{
			name:   "na",
			vec:    NA(2),
			outLen: 5,
		},
		{
			name:   "integer",
			vec:    IntegerWithNA([]int{1, 1}, []bool{true, false}),
			outLen: 5,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			outPayload := payload.Append(data.vec.Payload()).(*naPayload)

			if outPayload.length != data.outLen {
				t.Error(fmt.Sprintf("Out length (%d) is not equal to expected (%d)",
					outPayload.length, data.outLen))
			}
		})
	}
}

func TestNaPayload_Adjust(t *testing.T) {
	payload := NAPayload(3).(*naPayload)
	newPayload := payload.Adjust(5).(*naPayload)

	if newPayload.length != 5 {
		t.Error(fmt.Sprintf("New payload's length is wrong (%v instead of %v)", newPayload.length, 5))
	}
}

func TestNaPayload_IsUnique(t *testing.T) {
	testData := []struct {
		name     string
		payload  Payload
		booleans []bool
	}{
		{
			name:     "zero",
			payload:  NAPayload(0),
			booleans: []bool{},
		},
		{
			name:     "normal",
			payload:  NAPayload(5),
			booleans: []bool{true, false, false, false, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			booleans := data.payload.(*naPayload).IsUnique()

			if !reflect.DeepEqual(booleans, data.booleans) {
				t.Error(fmt.Sprintf("Result of IsUnique() (%v) do not match expected (%v)",
					booleans, data.booleans))
			}
		})
	}
}

func TestNaPayload_Coalesce(t *testing.T) {
	testData := []struct {
		name         string
		coalescer    Payload
		coalescendum Payload
		outData      []int
		outNA        []bool
	}{
		{
			name:         "empty",
			coalescer:    NAPayload(0),
			coalescendum: IntegerPayload([]int{}, nil),
			outData:      []int{},
			outNA:        []bool{},
		},
		{
			name:         "non-empty",
			coalescer:    NAPayload(5),
			coalescendum: IntegerPayload([]int{11, 12, 0, 14, 15}, []bool{false, false, true, false, false}),
			outData:      []int{11, 12, 0, 14, 15},
			outNA:        []bool{false, false, true, false, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.coalescer.(Coalescer).Coalesce(data.coalescendum).(*integerPayload)

			if !reflect.DeepEqual(payload.data, data.outData) {
				t.Error(fmt.Sprintf("Data (%v) do not match expected (%v)",
					payload.data, data.outData))
			}

			if !reflect.DeepEqual(payload.NA, data.outNA) {
				t.Error(fmt.Sprintf("NA (%v) do not match expected (%v)",
					payload.NA, data.outNA))
			}
		})
	}
}

func TestNaPayload_Pick(t *testing.T) {
	payload := NAPayload(5)

	testData := []struct {
		name string
		idx  int
		val  any
	}{
		{
			name: "normal 2",
			idx:  2,
			val:  nil,
		},
		{
			name: "normal 5",
			idx:  5,
			val:  nil,
		},
		{
			name: "na",
			idx:  3,
			val:  nil,
		},
		{
			name: "out of bounds -1",
			idx:  -1,
			val:  nil,
		},
		{
			name: "out of bounds 6",
			idx:  6,
			val:  nil,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			val := payload.Pick(data.idx)

			if val != data.val {
				t.Error(fmt.Sprintf("Result of Pick() (%v) do not match expected (%v)",
					val, data.val))
			}
		})
	}
}

func TestNaPayload_Data(t *testing.T) {
	testData := []struct {
		name    string
		payload Payload
		outData []any
	}{
		{
			name:    "empty",
			payload: NAPayload(0),
			outData: []any{},
		},
		{
			name:    "non-empty",
			payload: NAPayload(5),
			outData: []any{nil, nil, nil, nil, nil},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payloadData := data.payload.Data()

			if !reflect.DeepEqual(payloadData, data.outData) {
				t.Error(fmt.Sprintf("Result of Data() (%v) do not match expected (%v)",
					payloadData, data.outData))
			}
		})
	}
}
