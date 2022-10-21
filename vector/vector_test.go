package vector

import (
	"fmt"
	"logarithmotechnia/util"
	"math"
	"math/cmplx"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestVector_Len(t *testing.T) {
	testData := []struct {
		in             []int
		expectedLength int
	}{
		{
			in:             []int{1, 2, 3, 4, 5},
			expectedLength: 5,
		},
		{
			in:             []int{1, 2, 3},
			expectedLength: 3,
		},
		{
			in:             []int{1},
			expectedLength: 1,
		},
		{
			in:             []int{},
			expectedLength: 0,
		},
		{
			in:             nil,
			expectedLength: 0,
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := IntegerWithNA(data.in, nil)
			if vec.Len() != data.expectedLength {
				t.Error(fmt.Sprintf("Length (%d) is not equal to expected (%d)",
					vec.Len(), data.expectedLength))
			}
		})
	}
}

func TestVector_Integers(t *testing.T) {
	testData := []struct {
		vec       Vector
		outValues []int
		outNA     []bool
	}{
		{
			vec:       IntegerWithNA([]int{1, 2, 3, 4, 5}, []bool{false, false, true, false, false}),
			outValues: []int{1, 2, 0, 4, 5},
			outNA:     []bool{false, false, true, false, false},
		},
		{
			vec:       NA(0),
			outValues: []int{},
			outNA:     []bool{},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			integers, na := data.vec.Integers()
			if !reflect.DeepEqual(integers, data.outValues) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", integers, data.outValues))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", na, data.outNA))
			}
		})
	}
}

func TestVector_Floats(t *testing.T) {
	testData := []struct {
		vec       Vector
		outValues []float64
		outNA     []bool
	}{
		{
			vec:       IntegerWithNA([]int{1, 2, 3, 4, 5}, []bool{false, false, true, false, false}),
			outValues: []float64{1, 2, math.NaN(), 4, 5},
			outNA:     []bool{false, false, true, false, false},
		},
		{
			vec:       NA(0),
			outValues: []float64{},
			outNA:     []bool{},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			floats, na := data.vec.Floats()
			err := false
			for i, val := range floats {
				if math.IsNaN(val) {
					if !math.IsNaN(data.outValues[i]) {
						err = true
					}
				} else if val != data.outValues[i] {
					err = true
				}
			}
			if err {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", floats, data.outValues))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", na, data.outNA))
			}
		})
	}
}

func TestVector_Complexes(t *testing.T) {
	testData := []struct {
		vec       Vector
		outValues []complex128
		outNA     []bool
	}{
		{
			vec:       IntegerWithNA([]int{1, 2, 3, 4, 5}, []bool{false, false, true, false, false}),
			outValues: []complex128{1 + 0i, 2 + 0i, cmplx.NaN(), 4 + 0i, 5 + 0i},
			outNA:     []bool{false, false, true, false, false},
		},
		{
			vec:       NA(0),
			outValues: []complex128{},
			outNA:     []bool{},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			complexes, na := data.vec.Complexes()
			correct := true
			for i := 0; i < len(complexes); i++ {
				if cmplx.IsNaN(data.outValues[i]) {
					if !cmplx.IsNaN(complexes[i]) {
						correct = false
					}
				} else if complexes[i] != data.outValues[i] {
					correct = false
				}
			}
			if !correct {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", complexes, data.outValues))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", na, data.outNA))
			}
		})
	}
}

func TestVector_Booleans(t *testing.T) {
	testData := []struct {
		vec       Vector
		outValues []bool
		outNA     []bool
	}{
		{
			vec:       IntegerWithNA([]int{0, 1, 2, 3, 4, 5}, []bool{false, false, false, true, false, false}),
			outValues: []bool{false, true, true, false, true, true},
			outNA:     []bool{false, false, false, true, false, false},
		},
		{
			vec:       NA(0),
			outValues: []bool{},
			outNA:     []bool{},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			booleans, na := data.vec.Booleans()
			if !reflect.DeepEqual(booleans, data.outValues) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", booleans, data.outValues))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", na, data.outNA))
			}
		})
	}
}

func TestVector_Strings(t *testing.T) {
	testData := []struct {
		vec       Vector
		outValues []string
		outNA     []bool
	}{
		{
			vec:       IntegerWithNA([]int{1, 2, 3, 4, 5}, []bool{false, false, true, false, false}),
			outValues: []string{"1", "2", "", "4", "5"},
			outNA:     []bool{false, false, true, false, false},
		},
		{
			vec:       NA(0),
			outValues: []string{},
			outNA:     []bool{},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			strings, na := data.vec.Strings()
			if !reflect.DeepEqual(strings, data.outValues) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", strings, data.outValues))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", na, data.outNA))
			}
		})
	}
}

func TestVector_Times(t *testing.T) {
	testData := []struct {
		vec       Vector
		outValues []time.Time
		outNA     []bool
	}{
		{
			vec:       IntegerWithNA([]int{1, 2, 3, 4, 5}, []bool{false, false, true, false, false}),
			outValues: []time.Time{{}, {}, {}, {}, {}},
			outNA:     []bool{true, true, true, true, true},
		},
		{
			vec:       NA(0),
			outValues: []time.Time{},
			outNA:     []bool{},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			times, na := data.vec.Times()
			if !reflect.DeepEqual(times, data.outValues) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", times, data.outValues))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", na, data.outNA))
			}
		})
	}
}

func TestVector_AsBoolean(t *testing.T) {
	testData := []struct {
		name      string
		vec       Vector
		outValues []bool
		outNA     []bool
		isNA      bool
	}{
		{
			name:      "booleanable",
			vec:       IntegerWithNA([]int{1, 2, 0, 5, 5}, []bool{false, false, false, false, true}),
			outValues: []bool{true, true, false, true, false},
			outNA:     []bool{false, false, false, false, true},
			isNA:      false,
		},
		{
			name: "non-booleanable",
			vec: TimeWithNA(toTimeData([]string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00",
				"1800-06-10T11:00:00Z"}), nil),
			isNA: true,
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := data.vec.AsBoolean()
			if data.isNA {
				if _, ok := vec.Payload().(*naPayload); ok {
					if vec.Len() != data.vec.Len() {
						t.Error(fmt.Sprintf("NA vector length (%v) is not equal to expected (%v)", vec.Len(), data.vec.Len()))
					}
				} else {
					t.Error("Vector is not NA")
				}
			} else {
				payload := vec.Payload().(*booleanPayload)
				if !reflect.DeepEqual(payload.data, data.outValues) {
					t.Error(fmt.Sprintf("Payload data (%v) is not equal to expected (%v)", payload.data, data.outValues))
				}
				if !reflect.DeepEqual(payload.na, data.outNA) {
					t.Error(fmt.Sprintf("Payload NA (%v) is not equal to expected (%v)", payload.na, data.outNA))
				}
			}
		})
	}
}

func TestVector_AsInteger(t *testing.T) {
	testData := []struct {
		name      string
		vec       Vector
		outValues []int
		outNA     []bool
		isNA      bool
	}{
		{
			name:      "intable",
			vec:       StringWithNA([]string{"1", "2", "0", "5", "5"}, []bool{false, false, false, false, true}),
			outValues: []int{1, 2, 0, 5, 0},
			outNA:     []bool{false, false, false, false, true},
			isNA:      false,
		},
		{
			name: "non-intable",
			vec: TimeWithNA(toTimeData([]string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00",
				"1800-06-10T11:00:00Z"}), nil),
			isNA: true,
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := data.vec.AsInteger()
			if data.isNA {
				if _, ok := vec.Payload().(*naPayload); ok {
					if vec.Len() != data.vec.Len() {
						t.Error(fmt.Sprintf("NA vector length (%v) is not equal to expected (%v)", vec.Len(), data.vec.Len()))
					}
				} else {
					t.Error("Vector is not NA")
				}
			} else {
				payload := vec.Payload().(*integerPayload)
				if !reflect.DeepEqual(payload.data, data.outValues) {
					t.Error(fmt.Sprintf("Payload data (%v) is not equal to expected (%v)", payload.data, data.outValues))
				}
				if !reflect.DeepEqual(payload.na, data.outNA) {
					t.Error(fmt.Sprintf("Payload NA (%v) is not equal to expected (%v)", payload.na, data.outNA))
				}
			}
		})
	}
}

func TestVector_AsFloat(t *testing.T) {
	testData := []struct {
		name      string
		vec       Vector
		outValues []float64
		outNA     []bool
		isNA      bool
	}{
		{
			name:      "floatable",
			vec:       StringWithNA([]string{"1.1", "2", "0", "5.5", "5"}, []bool{false, false, false, false, true}),
			outValues: []float64{1.1, 2, 0, 5.5, math.NaN()},
			outNA:     []bool{false, false, false, false, true},
			isNA:      false,
		},
		{
			name: "non-floatable",
			vec: TimeWithNA(toTimeData([]string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00",
				"1800-06-10T11:00:00Z"}), nil),
			isNA: true,
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := data.vec.AsFloat()
			if data.isNA {
				if _, ok := vec.Payload().(*naPayload); ok {
					if vec.Len() != data.vec.Len() {
						t.Error(fmt.Sprintf("NA vector length (%v) is not equal to expected (%v)", vec.Len(), data.vec.Len()))
					}
				} else {
					t.Error("Vector is not NA")
				}
			} else {
				payload := vec.Payload().(*floatPayload)
				if !util.EqualFloatArrays(payload.data, data.outValues) {
					t.Error(fmt.Sprintf("Payload data (%v) is not equal to expected (%v)", payload.data, data.outValues))
				}
				if !reflect.DeepEqual(payload.na, data.outNA) {
					t.Error(fmt.Sprintf("Payload NA (%v) is not equal to expected (%v)", payload.na, data.outNA))
				}
			}
		})
	}
}

func TestVector_AsComplex(t *testing.T) {
	testData := []struct {
		name      string
		vec       Vector
		outValues []complex128
		outNA     []bool
		isNA      bool
	}{
		{
			name:      "complexable",
			vec:       StringWithNA([]string{"1.1+1.1i", "2", "0", "5.5-2.5i", "5"}, []bool{false, false, false, false, true}),
			outValues: []complex128{1.1 + 1.1i, 2, 0, 5.5 - 2.5i, cmplx.NaN()},
			outNA:     []bool{false, false, false, false, true},
			isNA:      false,
		},
		{
			name: "non-complexable",
			vec: TimeWithNA(toTimeData([]string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00",
				"1800-06-10T11:00:00Z"}), nil),
			isNA: true,
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := data.vec.AsComplex()
			if data.isNA {
				if _, ok := vec.Payload().(*naPayload); ok {
					if vec.Len() != data.vec.Len() {
						t.Error(fmt.Sprintf("NA vector length (%v) is not equal to expected (%v)", vec.Len(), data.vec.Len()))
					}
				} else {
					t.Error("Vector is not NA")
				}
			} else {
				payload := vec.Payload().(*complexPayload)
				if !util.EqualComplexArrays(payload.data, data.outValues) {
					t.Error(fmt.Sprintf("Payload data (%v) is not equal to expected (%v)", payload.data, data.outValues))
				}
				if !reflect.DeepEqual(payload.na, data.outNA) {
					t.Error(fmt.Sprintf("Payload NA (%v) is not equal to expected (%v)", payload.na, data.outNA))
				}
			}
		})
	}
}

func TestVector_AsString(t *testing.T) {
	testData := []struct {
		name      string
		vec       Vector
		outValues []string
		outNA     []bool
		isNA      bool
	}{
		{
			name:      "stringable",
			vec:       IntegerWithNA([]int{1, 2, 0, 5, 5}, []bool{false, false, false, false, true}),
			outValues: []string{"1", "2", "0", "5", ""},
			outNA:     []bool{false, false, false, false, true},
			isNA:      false,
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := data.vec.AsString()
			if data.isNA {
				if _, ok := vec.Payload().(*naPayload); ok {
					if vec.Len() != data.vec.Len() {
						t.Error(fmt.Sprintf("NA vector length (%v) is not equal to expected (%v)", vec.Len(), data.vec.Len()))
					}
				} else {
					t.Error("Vector is not NA")
				}
			} else {
				payload := vec.Payload().(*stringPayload)
				if !reflect.DeepEqual(payload.data, data.outValues) {
					t.Error(fmt.Sprintf("Payload data (%v) is not equal to expected (%v)", payload.data, data.outValues))
				}
				if !reflect.DeepEqual(payload.na, data.outNA) {
					t.Error(fmt.Sprintf("Payload NA (%v) is not equal to expected (%v)", payload.na, data.outNA))
				}
			}
		})
	}
}

func TestVector_AsInterface(t *testing.T) {
	testData := []struct {
		name      string
		vec       Vector
		outValues []any
		outNA     []bool
		isNA      bool
	}{
		{
			name:      "interfaceable",
			vec:       IntegerWithNA([]int{1, 2, 0, 5, 5}, []bool{false, false, false, false, true}),
			outValues: []any{1, 2, 0, 5, nil},
			outNA:     []bool{false, false, false, false, true},
			isNA:      false,
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := data.vec.AsInterface()
			if data.isNA {
				if _, ok := vec.Payload().(*naPayload); ok {
					if vec.Len() != data.vec.Len() {
						t.Error(fmt.Sprintf("NA vector length (%v) is not equal to expected (%v)", vec.Len(), data.vec.Len()))
					}
				} else {
					t.Error("Vector is not NA")
				}
			} else {
				payload := vec.Payload().(*anyPayload)
				if !reflect.DeepEqual(payload.data, data.outValues) {
					t.Error(fmt.Sprintf("Payload data (%v) is not equal to expected (%v)", payload.data, data.outValues))
				}
				if !reflect.DeepEqual(payload.na, data.outNA) {
					t.Error(fmt.Sprintf("Payload NA (%v) is not equal to expected (%v)", payload.na, data.outNA))
				}
			}
		})
	}
}

func TestVector_Transform(t *testing.T) {
	na := []bool{false, false, true}
	vec := TimeWithNA(toTimeData([]string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00",
		"1800-06-10T11:00:00Z"}), na)
	newVec := vec.Transform(func(values []any, na []bool) Payload {
		integers := make([]int, len(values))
		for i, val := range values {
			if na[i] {
				integers[i] = 0
			} else {
				integers[i] = int(val.(time.Time).Unix())
			}
		}

		return IntegerPayload(integers, na)
	})

	payload := newVec.Payload().(*integerPayload)
	expectedData := []int{1136189045, 1609493400, 0}
	if !reflect.DeepEqual(payload.data, expectedData) {
		t.Error(fmt.Sprintf("Payload data (%v) is not equal to expected (%v)", payload.data, expectedData))
	}
	if !reflect.DeepEqual(payload.na, na) {
		t.Error(fmt.Sprintf("Payload data (%v) is not equal to expected (%v)", payload.na, na))
	}
}

func TestVector_AsTime(t *testing.T) {
	testData := []struct {
		name      string
		vec       Vector
		outValues []time.Time
		outNA     []bool
		isNA      bool
	}{
		{
			name: "timeable",
			vec: TimeWithNA(toTimeData([]string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00",
				"1800-06-10T11:00:00Z"}), nil),
			outValues: toTimeData([]string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00",
				"1800-06-10T11:00:00Z"}),
			outNA: []bool{false, false, false},
			isNA:  false,
		},
		{
			name: "non-timeable",
			vec:  IntegerWithNA([]int{1, 2, 0, 5, 5}, []bool{false, false, false, false, true}),
			isNA: true,
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := data.vec.AsTime()
			if data.isNA {
				if _, ok := vec.Payload().(*naPayload); ok {
					if vec.Len() != data.vec.Len() {
						t.Error(fmt.Sprintf("NA vector length (%v) is not equal to expected (%v)", vec.Len(), data.vec.Len()))
					}
				} else {
					t.Error("Vector is not NA")
				}
			} else {
				payload := vec.Payload().(*timePayload)
				if !reflect.DeepEqual(payload.data, data.outValues) {
					t.Error(fmt.Sprintf("Payload data (%v) is not equal to expected (%v)", payload.data, data.outValues))
				}
				if !reflect.DeepEqual(payload.na, data.outNA) {
					t.Error(fmt.Sprintf("Payload NA (%v) is not equal to expected (%v)", payload.na, data.outNA))
				}
			}
		})
	}
}

func TestVector_ByIndices(t *testing.T) {
	vec := IntegerWithNA([]int{1, 2, 3, 4, 5}, []bool{false, false, true, false, false})
	testData := []struct {
		name    string
		indices []int
		out     []int
		outNA   []bool
		length  int
	}{
		{
			name:    "normal",
			indices: []int{-1, 0, 4, 3, 2, 6},
			out:     []int{0, 4, 0, 2},
			outNA:   []bool{true, false, true, false},
			length:  4,
		},
		{
			name:    "empty",
			indices: []int{},
			out:     []int{},
			outNA:   []bool{},
			length:  0,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			newVec := vec.ByIndices(data.indices)
			integers, na := newVec.Integers()
			if !reflect.DeepEqual(integers, data.out) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", integers, data.out))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("Result NA (%v) is not equal to expected NA (%v)", na, data.outNA))
			}
			if newVec.Len() != data.length {
				t.Error(fmt.Sprintf("Result length (%d) is not equal to expected length (%d)",
					newVec.Len(), data.length))
			}
		})
	}
}

func TestVector_FromTo(t *testing.T) {
	vec := IntegerWithNA(
		[]int{1, 2, 3, 4, 5},
		[]bool{false, false, true, false, false},
	)

	testData := []struct {
		name   string
		from   int
		to     int
		out    []int
		outNA  []bool
		length int
	}{
		{
			name:   "FromTo straight",
			from:   2,
			to:     4,
			out:    []int{2, 0, 4},
			outNA:  []bool{false, true, false},
			length: 3,
		},
		{
			name:   "FromTo reverse",
			from:   4,
			to:     2,
			out:    []int{4, 0, 2},
			outNA:  []bool{false, true, false},
			length: 3,
		},
		{
			name:   "FromTo with remove straight",
			from:   -2,
			to:     -4,
			out:    []int{1, 5},
			outNA:  []bool{false, false},
			length: 2,
		},
		{
			name:   "FromTo with remove reverse",
			from:   -4,
			to:     -2,
			out:    []int{1, 5},
			outNA:  []bool{false, false},
			length: 2,
		},
		{
			name:   "FromTo straight with incorrect boundaries",
			from:   0,
			to:     7,
			out:    []int{1, 2, 0, 4, 5},
			outNA:  []bool{false, false, true, false, false},
			length: 5,
		},
		{
			name:   "FromTo full remove with incorrect boundaries",
			from:   0,
			to:     -7,
			out:    []int{},
			outNA:  []bool{},
			length: 0,
		},
		{
			name:   "FromTo different signs",
			from:   -2,
			to:     4,
			out:    []int{},
			outNA:  []bool{},
			length: 0,
		},
		{
			name:   "FromTo different signs reverse",
			from:   2,
			to:     -4,
			out:    []int{},
			outNA:  []bool{},
			length: 0,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			newVec := vec.FromTo(data.from, data.to)
			integers, na := newVec.Integers()
			if !reflect.DeepEqual(integers, data.out) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", integers, data.out))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("Result NA (%v) is not equal to expected NA (%v)", na, data.outNA))
			}
			if newVec.Len() != data.length {
				t.Error(fmt.Sprintf("Result length (%d) is not equal to expected length (%d)",
					newVec.Len(), data.length))
			}
		})
	}
}

func TestVector_Filter(t *testing.T) {
	vec := IntegerWithNA(
		[]int{1, 2, 3, 4, 5},
		[]bool{false, false, true, false, false},
	)
	testData := []struct {
		name    string
		whicher any
		out     []int
		outNA   []bool
		length  int
	}{
		{
			name:    "indices",
			whicher: []int{5, 3, 1},
			out:     []int{5, 0, 1},
			outNA:   []bool{false, true, false},
			length:  3,
		},
		{
			name:    "idx",
			whicher: []int{1},
			out:     []int{1},
			outNA:   []bool{false},
			length:  1,
		},
		{
			name:    "booleanPayload",
			whicher: []bool{true, false, true, false, true},
			out:     []int{1, 0, 5},
			outNA:   []bool{false, true, false},
			length:  3,
		},
		{
			name:    "byIntFunc",
			whicher: func(_ int, val int, na bool) bool { return !na && val >= 3 },
			out:     []int{4, 5},
			outNA:   []bool{false, false},
			length:  2,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			newVec := vec.Filter(data.whicher)
			integers, na := newVec.Integers()
			if !reflect.DeepEqual(integers, data.out) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", integers, data.out))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("Result NA (%v) is not equal to expected NA (%v)", na, data.outNA))
			}
			if newVec.Len() != data.length {
				t.Error(fmt.Sprintf("Result length (%d) is not equal to expected length (%d)",
					newVec.Len(), data.length))
			}
		})
	}
}

func TestVector_IsNA(t *testing.T) {
	testData := []struct {
		name  string
		vec   Vector
		notNA []bool
	}{
		{
			name:  "with NA",
			vec:   IntegerWithNA([]int{1, 2, 3}, []bool{false, false, false}),
			notNA: []bool{false, false, false},
		},
		{
			name:  "without NA",
			vec:   IntegerWithNA([]int{1, 2, 3}, []bool{false, true, false}),
			notNA: []bool{false, true, false},
		},
		{
			name:  "empty",
			vec:   NA(0),
			notNA: []bool{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			_ = data
		})
	}
}

func TestVector_NotNA(t *testing.T) {
	testData := []struct {
		name  string
		vec   Vector
		notNA []bool
	}{
		{
			name:  "with NA",
			vec:   IntegerWithNA([]int{1, 2, 3}, []bool{false, false, false}),
			notNA: []bool{true, true, true},
		},
		{
			name:  "without NA",
			vec:   IntegerWithNA([]int{1, 2, 3}, []bool{false, true, false}),
			notNA: []bool{true, false, true},
		},
		{
			name:  "empty",
			vec:   NA(0),
			notNA: []bool{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			_ = data
		})
	}
}

func TestVector_HasNA(t *testing.T) {
	testData := []struct {
		name  string
		vec   Vector
		hasNA bool
	}{
		{
			name:  "regular with nil NA",
			vec:   IntegerWithNA([]int{1, 2, 3}, nil),
			hasNA: false,
		},
		{
			name:  "regular without NA",
			vec:   IntegerWithNA([]int{1, 2, 3}, []bool{false, false, false}),
			hasNA: false,
		},
		{
			name:  "regular with NA",
			vec:   IntegerWithNA([]int{1, 2, 3}, []bool{false, true, false}),
			hasNA: true,
		},
		{
			name:  "Empty",
			vec:   NA(0),
			hasNA: false,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if data.vec.HasNA() != data.hasNA {
				t.Error(fmt.Sprintf("data.vec.HasNA() (%t) is not equal to data.hasNA (%t)",
					data.vec.HasNA(), data.hasNA))
			}
		})
	}
}

func TestVector_WithNA(t *testing.T) {
	testData := []struct {
		name   string
		vec    Vector
		withNA []int
	}{
		{
			name:   "regular with nil NA",
			vec:    IntegerWithNA([]int{1, 2, 3}, nil),
			withNA: []int{},
		},
		{
			name:   "regular without NA",
			vec:    IntegerWithNA([]int{1, 2, 3}, []bool{false, false, false}),
			withNA: []int{},
		},
		{
			name:   "regular with NA",
			vec:    IntegerWithNA([]int{1, 2, 3}, []bool{false, true, true}),
			withNA: []int{2, 3},
		},
		{
			name:   "Empty",
			vec:    NA(0),
			withNA: []int{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if !reflect.DeepEqual(data.vec.WithNA(), data.withNA) {
				t.Error(fmt.Sprintf("data.vec.WithNA() (%v) is not equal to data.WithNA (%v)",
					data.vec.WithNA(), data.withNA))
			}
		})
	}
}

func TestVector_WithoutNA(t *testing.T) {
	testData := []struct {
		name      string
		vec       Vector
		withoutNA []int
	}{
		{
			name:      "regular with nil NA",
			vec:       IntegerWithNA([]int{1, 2, 3}, nil),
			withoutNA: []int{1, 2, 3},
		},
		{
			name:      "regular without NA",
			vec:       IntegerWithNA([]int{1, 2, 3}, []bool{false, false, false}),
			withoutNA: []int{1, 2, 3},
		},
		{
			name:      "regular with NA",
			vec:       IntegerWithNA([]int{1, 2, 3}, []bool{false, true, true}),
			withoutNA: []int{1},
		},
		{
			name:      "Empty",
			vec:       NA(0),
			withoutNA: []int{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if !reflect.DeepEqual(data.vec.WithoutNA(), data.withoutNA) {
				t.Error(fmt.Sprintf("data.vec.WithoutNA() (%v) is not equal to data.WithoutNA (%v)",
					data.vec.WithoutNA(), data.withoutNA))
			}
		})
	}
}

func TestVector_IsEmpty(t *testing.T) {
	testData := []struct {
		name    string
		vec     Vector
		isEmpty bool
	}{
		{
			name:    "zero integerPayload vector",
			vec:     IntegerWithNA([]int{}, nil),
			isEmpty: true,
		},
		{
			name:    "non-zero integerPayload vector",
			vec:     IntegerWithNA([]int{1, 2, 3}, nil),
			isEmpty: false,
		},
		{
			name:    "empty vector",
			vec:     IntegerWithNA([]int{}, nil),
			isEmpty: true,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if data.vec.IsEmpty() != data.isEmpty {
				t.Error(fmt.Sprintf("data.vec.IsEmpty() (%t) is not equal to data.isEmpty (%t)",
					data.vec.IsEmpty(), data.isEmpty))
			}
		})
	}
}

func TestVector_Clone(t *testing.T) {
	vec := IntegerWithNA([]int{1, 2, 3, 4, 5}, []bool{false, true, false, true, false}).(*vector)
	newVec := vec.Clone().(*vector)

	if vec.length != newVec.length {
		t.Error(fmt.Sprintf("vec.length (%d) is not equal to newVec.length (%d)", vec.length, newVec.length))
	}

	srcAddr := &(vec.payload.(*integerPayload).data[0])
	newAddr := &(newVec.payload.(*integerPayload).data[0])

	if srcAddr != newAddr {
		t.Error("Payload data was not cloned")
	}

	srcAddrNA := &(vec.payload.(*integerPayload).na[0])
	newAddrNA := &(newVec.payload.(*integerPayload).na[0])

	if srcAddrNA != newAddrNA {
		t.Error("Payload NA data was not cloned")
	}
}

func TestVector_SupportsWhicher(t *testing.T) {
	testData := []struct {
		name            string
		vec             Vector
		whicher         any
		supportsWhicher bool
	}{
		{
			name:            "integerPayload vector + valid whicher",
			vec:             IntegerWithNA([]int{1, 2, 3}, nil),
			whicher:         func(_ int, val int, _ bool) bool { return val == 1 || val == 3 },
			supportsWhicher: true,
		},
		{
			name:            "integerPayload vector + invalid whicher",
			vec:             IntegerWithNA([]int{1, 2, 3}, nil),
			whicher:         true,
			supportsWhicher: false,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			supportsWhicher := data.vec.SupportsWhicher(data.whicher)
			if supportsWhicher != data.supportsWhicher {
				t.Error(fmt.Sprintf("data.vec.Which() (%v) is not equal to data.selected (%v)",
					supportsWhicher, data.supportsWhicher))
			}
		})
	}
}

func TestVector_Select(t *testing.T) {
	testData := []struct {
		name     string
		vec      Vector
		whicher  any
		selected []bool
	}{
		{
			name:     "integerPayload vector + valid whicher",
			vec:      IntegerWithNA([]int{1, 2, 3}, nil),
			whicher:  func(_ int, val int, _ bool) bool { return val == 1 || val == 3 },
			selected: []bool{true, false, true},
		},
		{
			name:     "integerPayload vector + invalid whicher",
			vec:      IntegerWithNA([]int{1, 2, 3}, nil),
			whicher:  true,
			selected: []bool{false, false, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			selected := data.vec.Which(data.whicher)
			if !reflect.DeepEqual(selected, data.selected) {
				t.Error(fmt.Sprintf("data.vec.Which() (%v) is not equal to data.selected (%v)",
					selected, data.selected))
			}
		})
	}
}

func TestVector_SupportsApplier(t *testing.T) {
	testData := []struct {
		name            string
		vec             Vector
		whicher         any
		supportsApplier bool
	}{
		{
			name:            "integerPayload vector + valid applier",
			vec:             IntegerWithNA([]int{1, 2, 3}, nil),
			whicher:         func(_ int, val int, na bool) (int, bool) { return 10 * val, na },
			supportsApplier: true,
		},
		{
			name:            "integerPayload vector + invalid applier",
			vec:             IntegerWithNA([]int{1, 2, 3}, nil),
			whicher:         true,
			supportsApplier: false,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			supportsApplier := data.vec.SupportsApplier(data.whicher)
			if supportsApplier != data.supportsApplier {
				t.Error(fmt.Sprintf("Applier's support (%v) is not equal to expected (%v)",
					supportsApplier, data.supportsApplier))
			}
		})
	}
}

func TestVector_Apply(t *testing.T) {
	testData := []struct {
		name    string
		vec     Vector
		applier any
		dataOut []int
		NAOut   []bool
	}{
		{
			name: "integerPayload vector + valid applier",
			vec:  IntegerWithNA([]int{1, 2, 3, 4, 5}, nil),
			applier: func(idx int, val int, na bool) (int, bool) {
				if idx == 5 {
					return val, true
				}
				return 10 * val, na
			},
			dataOut: []int{10, 20, 30, 40, 0},
			NAOut:   []bool{false, false, false, false, true},
		},
		{
			name:    "integerPayload vector + invalid applier",
			vec:     IntegerWithNA([]int{1, 2, 3, 4, 5}, nil),
			applier: true,
			dataOut: []int{0, 0, 0, 0, 0},
			NAOut:   []bool{true, true, true, true, true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			newVec := data.vec.Apply(data.applier)
			integers, na := newVec.Integers()
			if !reflect.DeepEqual(integers, data.dataOut) {
				t.Error(fmt.Sprintf("Integers (%v) is not equal to expected (%v)",
					integers, data.dataOut))
			}
			if !reflect.DeepEqual(na, data.NAOut) {
				t.Error(fmt.Sprintf("NA (%v) is not equal to expected (%v)",
					na, data.NAOut))
			}
		})
	}
}

func TestVector_ApplyTo(t *testing.T) {
	testData := []struct {
		name    string
		vec     Vector
		whicher any
		applier any
		dataOut []int
		NAOut   []bool
	}{
		{
			name:    "[]int whicher",
			whicher: []int{1, 2, 5},
			vec:     IntegerWithNA([]int{1, 2, 3, 4, 5}, nil),
			applier: func(idx int, val int, na bool) (int, bool) {
				if idx == 5 {
					return val, true
				}
				return 10 * val, na
			},
			dataOut: []int{10, 20, 3, 4, 0},
			NAOut:   []bool{false, false, false, false, true},
		},
		{
			name:    "[]bool whicher 5",
			whicher: []bool{true, false, false, false, true},
			vec:     IntegerWithNA([]int{1, 2, 3, 4, 5}, nil),
			applier: func(idx int, val int, na bool) (int, bool) {
				if idx == 5 {
					return val, true
				}
				return 10 * val, na
			},
			dataOut: []int{10, 2, 3, 4, 0},
			NAOut:   []bool{false, false, false, false, true},
		},
		{
			name:    "[]bool whicher 7",
			whicher: []bool{true, false, false, false, true, true, true},
			vec:     IntegerWithNA([]int{1, 2, 3, 4, 5}, nil),
			applier: func(idx int, val int, na bool) (int, bool) {
				if idx == 5 {
					return val, true
				}
				return 10 * val, na
			},
			dataOut: []int{10, 2, 3, 4, 0},
			NAOut:   []bool{false, false, false, false, true},
		},
		{
			name:    "[]bool whicher 2",
			whicher: []bool{true, false},
			vec:     IntegerWithNA([]int{1, 2, 3, 4, 5}, nil),
			applier: func(idx int, val int, na bool) (int, bool) {
				if idx == 5 {
					return val, true
				}
				return 10 * val, na
			},
			dataOut: []int{10, 2, 30, 4, 0},
			NAOut:   []bool{false, false, false, false, true},
		},
		{
			name: "[]bool whicher func",
			whicher: func(val int, na bool) bool {
				if val%2 == 0 {
					return false
				}
				return true
			},
			vec: IntegerWithNA([]int{1, 2, 3, 4, 5}, nil),
			applier: func(idx int, val int, na bool) (int, bool) {
				if idx == 5 {
					return val, true
				}
				return 10 * val, na
			},
			dataOut: []int{10, 2, 30, 4, 0},
			NAOut:   []bool{false, false, false, false, true},
		},
		{
			name:    "invalid applier",
			vec:     IntegerWithNA([]int{1, 2, 3, 4, 5}, nil),
			applier: true,
			dataOut: []int{0, 0, 0, 0, 0},
			NAOut:   []bool{true, true, true, true, true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			newVec := data.vec.ApplyTo(data.whicher, data.applier)
			integers, na := newVec.Integers()
			if !reflect.DeepEqual(integers, data.dataOut) {
				t.Error(fmt.Sprintf("Integers (%v) is not equal to expected (%v)",
					integers, data.dataOut))
			}
			if !reflect.DeepEqual(na, data.NAOut) {
				t.Error(fmt.Sprintf("NA (%v) is not equal to expected (%v)",
					na, data.NAOut))
			}
		})
	}
}

func TestVector_Append(t *testing.T) {
	vec := IntegerWithNA([]int{1, 2, 3}, nil)

	testData := []struct {
		name    string
		vec     Vector
		outData []int
		outNA   []bool
	}{
		{
			name:    "boolean",
			vec:     BooleanWithNA([]bool{true, true}, []bool{true, false}),
			outData: []int{1, 2, 3, 0, 1},
			outNA:   []bool{false, false, false, true, false},
		},
		{
			name:    "integer",
			vec:     IntegerWithNA([]int{4, 5}, []bool{true, false}),
			outData: []int{1, 2, 3, 0, 5},
			outNA:   []bool{false, false, false, true, false},
		},
		{
			name:    "empty",
			vec:     IntegerWithNA([]int{}, []bool{}),
			outData: []int{1, 2, 3},
			outNA:   []bool{false, false, false},
		},
		{
			name:    "na",
			vec:     NA(2),
			outData: []int{1, 2, 3, 0, 0},
			outNA:   []bool{false, false, false, true, true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			outVector := vec.Append(data.vec).(*vector)
			outPayload := outVector.payload.(*integerPayload)

			if outVector.Len() != vec.Len()+data.vec.Len() {
				t.Error(fmt.Sprintf("Output length (%d) does not match expected (%d)",
					outVector.Len(), vec.Len()+data.vec.Len()))
			}
			if !reflect.DeepEqual(data.outData, outPayload.data) {
				t.Error(fmt.Sprintf("Output data (%v) does not match expected (%v)",
					outPayload.data, data.outData))
			}
			if !reflect.DeepEqual(data.outNA, outPayload.na) {
				t.Error(fmt.Sprintf("Output NA (%v) does not match expected (%v)",
					outPayload.na, data.outNA))
			}
		})
	}
}

func TestVector_Adjust(t *testing.T) {
	vec := IntegerWithNA([]int{1, 2, 3, 4, 5}, []bool{false, false, true, false, false})

	testData := []struct {
		name    string
		size    int
		outData []int
		outNA   []bool
	}{
		{
			name:    "same",
			size:    5,
			outData: []int{1, 2, 0, 4, 5},
			outNA:   []bool{false, false, true, false, false},
		},
		{
			name:    "less",
			size:    3,
			outData: []int{1, 2, 0},
			outNA:   []bool{false, false, true},
		},
		{
			name:    "more",
			size:    12,
			outData: []int{1, 2, 0, 4, 5, 1, 2, 0, 4, 5, 1, 2},
			outNA:   []bool{false, false, true, false, false, false, false, true, false, false, false, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			outPayload := vec.Adjust(data.size).Payload().(*integerPayload)

			if !reflect.DeepEqual(data.outData, outPayload.data) {
				t.Error(fmt.Sprintf("Output data (%v) does not match expected (%v)",
					outPayload.data, data.outData))
			}
			if !reflect.DeepEqual(data.outNA, outPayload.na) {
				t.Error(fmt.Sprintf("Output NA (%v) does not match expected (%v)",
					outPayload.na, data.outNA))
			}
		})
	}
}

func TestVector_Find(t *testing.T) {
	vec := IntegerWithNA([]int{1, 2, 1, 4, 0}, nil)

	testData := []struct {
		name   string
		needle any
		pos    int
	}{
		{"existent", 4, 4},
		{"float64", 2.0, 2},
		{"float64 with floating part", 2.1, 0},
		{"non-existent", -10, 0},
		{"incorrect type", "true", 0},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			pos := vec.Find(data.needle)

			if pos != data.pos {
				t.Error(fmt.Sprintf("Position (%v) does not match expected (%v)",
					pos, data.pos))
			}
		})
	}
}

func TestVector_FindAll(t *testing.T) {
	vec := IntegerWithNA([]int{1, 2, 1, 4, 0}, nil)

	testData := []struct {
		name   string
		needle any
		pos    []int
	}{
		{"existent", 1, []int{1, 3}},
		{"float", 1.0, []int{1, 3}},
		{"float with floating part", 1.2, []int{}},
		{"non-existent", -10, []int{}},
		{"incorrect type", false, []int{}},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			pos := vec.FindAll(data.needle)

			if !reflect.DeepEqual(pos, data.pos) {
				t.Error(fmt.Sprintf("Positions (%v) does not match expected (%v)",
					pos, data.pos))
			}
		})
	}
}

func TestVector_Has(t *testing.T) {
	vec := IntegerWithNA([]int{1, 2, 1, 4, 0}, nil)

	testData := []struct {
		name   string
		needle any
		has    bool
	}{
		{"existent", 4, true},
		{"float64", 2.0, true},
		{"float64 with floating part", 2.1, false},
		{"non-existent", -10, false},
		{"incorrect type", "true", false},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			has := vec.Has(data.needle)

			if has != data.has {
				t.Error(fmt.Sprintf("Result (%v) does not match expected (%v)",
					has, data.has))
			}
		})
	}
}

func TestVector_Eq(t *testing.T) {
	testData := []struct {
		vec Vector
		val any
		cmp []bool
	}{
		{
			vec: IntegerWithNA([]int{2, 0, 2, 2, 1}, []bool{false, false, true, false, false}),
			val: 2,
			cmp: []bool{true, false, false, true, false},
		},
		{
			vec: NA(5),
			val: 2,
			cmp: []bool{false, false, false, false, false},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			cmp := data.vec.Eq(data.val)

			if !reflect.DeepEqual(cmp, data.cmp) {
				t.Error(fmt.Sprintf("Comparator results (%v) do not match expected (%v)",
					cmp, data.cmp))
			}
		})
	}
}

func TestVector_Neq(t *testing.T) {
	testData := []struct {
		vec Vector
		val any
		cmp []bool
	}{
		{
			vec: IntegerWithNA([]int{2, 0, 1, 2, 1}, []bool{false, false, true, false, false}),
			val: 2,
			cmp: []bool{false, true, true, false, true},
		},
		{
			vec: NA(5),
			val: 2,
			cmp: []bool{true, true, true, true, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			cmp := data.vec.Neq(data.val)

			if !reflect.DeepEqual(cmp, data.cmp) {
				t.Error(fmt.Sprintf("Comparator results (%v) do not match expected (%v)",
					cmp, data.cmp))
			}
		})
	}
}

func TestVector_Gt(t *testing.T) {
	testData := []struct {
		vec Vector
		val any
		cmp []bool
	}{
		{
			vec: IntegerWithNA([]int{2, 0, 1, 2, 1}, []bool{false, false, true, false, false}),
			val: 1,
			cmp: []bool{true, false, false, true, false},
		},
		{
			vec: NA(5),
			val: 2,
			cmp: []bool{false, false, false, false, false},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			cmp := data.vec.Gt(data.val)

			if !reflect.DeepEqual(cmp, data.cmp) {
				t.Error(fmt.Sprintf("Comparator results (%v) do not match expected (%v)",
					cmp, data.cmp))
			}
		})
	}
}

func TestVector_Lt(t *testing.T) {
	testData := []struct {
		vec Vector
		val any
		cmp []bool
	}{
		{
			vec: IntegerWithNA([]int{2, 0, 1, 2, 1}, []bool{false, false, true, false, false}),
			val: 2,
			cmp: []bool{false, true, false, false, true},
		},
		{
			vec: NA(5),
			val: 2,
			cmp: []bool{false, false, false, false, false},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			cmp := data.vec.Lt(data.val)

			if !reflect.DeepEqual(cmp, data.cmp) {
				t.Error(fmt.Sprintf("Comparator results (%v) do not match expected (%v)",
					cmp, data.cmp))
			}
		})
	}
}

func TestVector_Gte(t *testing.T) {
	testData := []struct {
		vec Vector
		val any
		cmp []bool
	}{
		{
			vec: IntegerWithNA([]int{2, 0, 1, 2, 1}, []bool{false, false, true, false, false}),
			val: 1,
			cmp: []bool{true, false, false, true, true},
		},
		{
			vec: NA(5),
			val: 2,
			cmp: []bool{false, false, false, false, false},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			cmp := data.vec.Gte(data.val)

			if !reflect.DeepEqual(cmp, data.cmp) {
				t.Error(fmt.Sprintf("Comparator results (%v) do not match expected (%v)",
					cmp, data.cmp))
			}
		})
	}
}

func TestVector_Lte(t *testing.T) {
	testData := []struct {
		vec Vector
		val any
		cmp []bool
	}{
		{
			vec: IntegerWithNA([]int{2, 0, 2, 2, 1}, []bool{false, false, true, false, false}),
			val: 2,
			cmp: []bool{true, true, false, true, true},
		},
		{
			vec: NA(5),
			val: 2,
			cmp: []bool{false, false, false, false, false},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			cmp := data.vec.Lte(data.val)

			if !reflect.DeepEqual(cmp, data.cmp) {
				t.Error(fmt.Sprintf("Comparator results (%v) do not match expected (%v)",
					cmp, data.cmp))
			}
		})
	}
}

func TestVector_SortedIndices(t *testing.T) {
	testData := []struct {
		name          string
		vec           Vector
		sortedIndices []int
	}{
		{
			name:          "integer with NA",
			vec:           IntegerWithNA([]int{12, -8, 0, -4, 5}, []bool{false, false, true, false, false}),
			sortedIndices: []int{2, 4, 5, 1, 3},
		},
		{
			name:          "integer without NA",
			vec:           IntegerWithNA([]int{12, -8, 0, -4, 5}, nil),
			sortedIndices: []int{2, 4, 3, 5, 1},
		},
		{
			name:          "boolean with NA",
			vec:           BooleanWithNA([]bool{true, true, false, false, true}, []bool{false, false, false, true, true}),
			sortedIndices: []int{3, 1, 2, 4, 5},
		},
		{
			name:          "boolean without NA",
			vec:           BooleanWithNA([]bool{true, true, false, false, true}, nil),
			sortedIndices: []int{3, 4, 1, 2, 5},
		},
		{
			name:          "float with NA",
			vec:           FloatWithNA([]float64{12, -8, 0, -4, 5}, []bool{false, false, true, false, false}),
			sortedIndices: []int{2, 4, 5, 1, 3},
		},
		{
			name:          "float without NA",
			vec:           FloatWithNA([]float64{12, -8, 0, -4, 5}, nil),
			sortedIndices: []int{2, 4, 3, 5, 1},
		},
		{
			name:          "string with NA",
			vec:           StringWithNA([]string{"delta", "beta", "alpha", "zeroth", "zero"}, []bool{false, false, true, true, false}),
			sortedIndices: []int{2, 1, 5, 3, 4},
		},
		{
			name:          "string without NA",
			vec:           StringWithNA([]string{"delta", "beta", "alpha", "zeroth", "zero"}, nil),
			sortedIndices: []int{3, 2, 1, 5, 4},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			sortedIndices := data.vec.SortedIndices()

			if !reflect.DeepEqual(sortedIndices, data.sortedIndices) {
				t.Error(fmt.Sprintf("Comparator results (%v) do not match expected (%v)",
					sortedIndices, data.sortedIndices))
			}
		})
	}
}

func TestVector_SortedIndicesWithRanks(t *testing.T) {
	testData := []struct {
		name          string
		vec           Vector
		sortedIndices []int
		ranks         []int
	}{
		{
			name:          "integer with NA",
			vec:           IntegerWithNA([]int{12, -4, 0, -4, 5}, []bool{false, false, true, false, false}),
			sortedIndices: []int{2, 4, 5, 1, 3},
			ranks:         []int{1, 1, 2, 3, 4},
		},
		{
			name:          "integer without NA",
			vec:           IntegerWithNA([]int{12, -4, 0, -4, 5}, nil),
			sortedIndices: []int{2, 4, 3, 5, 1},
			ranks:         []int{1, 1, 2, 3, 4},
		},
		{
			name:          "boolean with NA",
			vec:           BooleanWithNA([]bool{true, true, false, false, true}, []bool{false, false, false, true, true}),
			sortedIndices: []int{3, 1, 2, 4, 5},
			ranks:         []int{1, 2, 2, 3, 3},
		},
		{
			name:          "boolean without NA",
			vec:           BooleanWithNA([]bool{true, true, false, false, true}, nil),
			sortedIndices: []int{3, 4, 1, 2, 5},
			ranks:         []int{1, 1, 2, 2, 2},
		},
		{
			name:          "float with NA",
			vec:           FloatWithNA([]float64{12, -4, 0, -4, 5}, []bool{false, false, true, false, false}),
			sortedIndices: []int{2, 4, 5, 1, 3},
			ranks:         []int{1, 1, 2, 3, 4},
		},
		{
			name:          "float without NA",
			vec:           FloatWithNA([]float64{12, -4, 0, -4, 5}, nil),
			sortedIndices: []int{2, 4, 3, 5, 1},
			ranks:         []int{1, 1, 2, 3, 4},
		},
		{
			name:          "string with NA",
			vec:           StringWithNA([]string{"delta", "beta", "alpha", "zero", "zero"}, []bool{false, false, true, true, false}),
			sortedIndices: []int{2, 1, 5, 3, 4},
			ranks:         []int{1, 2, 3, 4, 4},
		},
		{
			name:          "string without NA",
			vec:           StringWithNA([]string{"delta", "beta", "alpha", "zero", "zero"}, nil),
			sortedIndices: []int{3, 2, 1, 4, 5},
			ranks:         []int{1, 2, 3, 4, 4},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			sortedIndices, ranks := data.vec.SortedIndicesWithRanks()

			if !reflect.DeepEqual(sortedIndices, data.sortedIndices) {
				t.Error(fmt.Sprintf("Sorted indices (%v) do not match expected (%v)",
					sortedIndices, data.sortedIndices))
			}

			if !reflect.DeepEqual(ranks, data.ranks) {
				t.Error(fmt.Sprintf("Sorted ranks (%v) do not match expected (%v)",
					ranks, data.ranks))
			}
		})
	}
}

func TestVector_Groups(t *testing.T) {
	testData := []struct {
		name   string
		vec    Vector
		groups [][]int
		values []any
	}{
		{
			name:   "normal",
			vec:    Integer([]int{-20, 10, 4, -20, 7, -20, 10, -20, 4, 10}, nil),
			groups: [][]int{{1, 4, 6, 8}, {2, 7, 10}, {3, 9}, {5}},
			values: []any{-20, 10, 4, 7},
		},
		{
			name: "with NA",
			vec: IntegerWithNA([]int{-20, 10, 4, -20, 7, -20, 10, -20, 4, 10},
				[]bool{false, false, false, false, false, false, true, true, false, false}),
			groups: [][]int{{1, 4, 6}, {2, 10}, {3, 9}, {5}, {7, 8}},
			values: []any{-20, 10, 4, 7, nil},
		},
		{
			name:   "empty",
			vec:    Integer([]int{}, nil),
			groups: [][]int{},
			values: []any{},
		},
		{
			name:   "non-groupable",
			vec:    NA(10),
			groups: [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			values: []any{nil},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			groups, _ := data.vec.Groups()

			if !reflect.DeepEqual(groups, data.groups) {
				t.Error(fmt.Sprintf("Groups (%v) do not match expected (%v)",
					groups, data.groups))
			}
		})
	}
}

func TestVector_GroupByIndices(t *testing.T) {
	testData := []struct {
		name      string
		vec       Vector
		indices   [][]int
		groups    []Vector
		isGrouped bool
	}{
		{
			name:      "normal",
			vec:       Integer([]int{-20, 10, 4, -20, 7, -20, 10, -20, 4, 10}, nil),
			indices:   [][]int{{1, 4, 7, 10}, {2, 3, 5, 9}, {6, 8}},
			isGrouped: true,
		},
		{
			name:      "with non-existent indices",
			vec:       Integer([]int{-20, 10, 4, -20, 7, -20, 10, -20, 4, 10}, nil),
			indices:   [][]int{{1, 4, 7, 10}, {2, 3, 5, 9, 12}, {6, 8, 11}},
			isGrouped: true,
		},
		{
			name:      "empty group index",
			vec:       Integer([]int{-20, 10, 4, -20, 7, -20, 10, -20, 4, 10}, nil),
			indices:   [][]int{},
			isGrouped: false,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			groupedVec := data.vec.GroupByIndices(data.indices).(*vector)

			if data.isGrouped {
				if !reflect.DeepEqual([][]int(groupedVec.groupIndex), data.indices) {
					t.Error(fmt.Sprintf("Group index (%v) do not match expected (%v)",
						groupedVec.groupIndex, data.indices))
				}
			} else {
				if groupedVec.groupIndex != nil {
					t.Error("Group index is not nil")
				}
			}
		})
	}
}

func TestVector_IsUnique(t *testing.T) {
	testData := []struct {
		name     string
		vector   Vector
		booleans []bool
	}{
		{
			name:     "without NA",
			vector:   Integer([]int{1, 0, 1, 3, 2, 3, 2, 0}),
			booleans: []bool{true, true, false, true, true, false, false, false},
		},
		{
			name: "with NA",
			vector: IntegerWithNA([]int{1, 0, 1, 3, 2, 3, 2, 0},
				[]bool{false, true, true, false, false, false, false, false}),
			booleans: []bool{true, true, false, true, true, false, false, true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			booleans := data.vector.IsUnique()

			if !reflect.DeepEqual(booleans, data.booleans) {
				t.Error(fmt.Sprintf("Result of IsUnique() (%v) do not match expected (%v)",
					booleans, data.booleans))
			}
		})
	}
}

func TestVector_Unique(t *testing.T) {
	testData := []struct {
		name      string
		vector    Vector
		outVector Vector
	}{
		{
			name:      "without NA",
			vector:    Integer([]int{1, 0, 1, 3, 2, 3, 2, 0}),
			outVector: Integer([]int{1, 0, 3, 2}),
		},
		{
			name: "with NA",
			vector: IntegerWithNA([]int{1, 0, 1, 3, 2, 3, 2, 0},
				[]bool{false, true, true, false, false, false, false, false}),
			outVector: IntegerWithNA([]int{1, 0, 3, 2, 0}, []bool{false, true, false, false, false}),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			v := data.vector.Unique()

			if !CompareVectorsForTest(v, data.outVector) {
				t.Error(fmt.Sprintf("Result of Unique() (%v) do not match expected (%v)",
					v, data.outVector))
			}
		})
	}
}

func TestVector_Coalesce(t *testing.T) {
	testData := []struct {
		name         string
		coalescer    Vector
		coalescendum []Vector
		coalescens   Vector
	}{
		{
			name:         "empty",
			coalescer:    Integer(nil, nil),
			coalescendum: []Vector{Integer([]int{}, nil)},
			coalescens:   Integer([]int{}),
		},
		{
			name:      "normal",
			coalescer: IntegerWithNA([]int{1, 0, 0, 0, 0, 0, 7}, []bool{false, true, true, true, true, true, false}),
			coalescendum: []Vector{
				IntegerWithNA([]int{10, 20, 0, 0, 0, 60, 70}, []bool{false, false, true, true, true, false, false}),
				FloatWithNA([]float64{100, 200, 300, 0, 500, 600, 700},
					[]bool{false, false, false, true, false, false, false}),
			},
			coalescens: IntegerWithNA([]int{1, 20, 300, 0, 500, 60, 7},
				[]bool{false, false, false, true, false, false, false}),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			vec := data.coalescer.Coalesce(data.coalescendum...)

			if !CompareVectorsForTest(vec, data.coalescens) {
				t.Error(fmt.Sprintf("Data (%v) do not match expected (%v)",
					vec, data.coalescens))
			}
		})
	}
}

func TestVector_Even(t *testing.T) {
	testData := []struct {
		name        string
		vec         Vector
		outBooleans []bool
	}{
		{
			name:        "five",
			vec:         Integer([]int{1, 2, 3, 4, 5}),
			outBooleans: []bool{false, true, false, true, false},
		},
		{
			name:        "eight",
			vec:         Float([]float64{1, 2, 3, 4, 5, 6, 7, 8}),
			outBooleans: []bool{false, true, false, true, false, true, false, true},
		},
		{
			name:        "empty",
			vec:         Any([]any{}),
			outBooleans: []bool{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			outBooleans := data.vec.Even()
			if !reflect.DeepEqual(outBooleans, data.outBooleans) {
				t.Error(fmt.Sprintf("vec.Even() (%v) does not match data.outBooleans (%v)", outBooleans, data.outBooleans))
			}
		})
	}
}

func TestVector_Odd(t *testing.T) {
	testData := []struct {
		name        string
		vec         Vector
		outBooleans []bool
	}{
		{
			name:        "five",
			vec:         Integer([]int{1, 2, 3, 4, 5}),
			outBooleans: []bool{true, false, true, false, true},
		},
		{
			name:        "eight",
			vec:         Float([]float64{1, 2, 3, 4, 5, 6, 7, 8}),
			outBooleans: []bool{true, false, true, false, true, false, true, false},
		},
		{
			name:        "empty",
			vec:         Time([]time.Time{}),
			outBooleans: []bool{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			outBooleans := data.vec.Odd()
			if !reflect.DeepEqual(outBooleans, data.outBooleans) {
				t.Error(fmt.Sprintf("vec.Even() (%v) does not match data.outBooleans (%v)", outBooleans, data.outBooleans))
			}
		})
	}
}

func TestVector_Nth(t *testing.T) {
	testData := []struct {
		name        string
		nth         int
		vec         Vector
		outBooleans []bool
	}{
		{
			name:        "every 3",
			nth:         3,
			vec:         Integer([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			outBooleans: []bool{false, false, true, false, false, true, false, false, true, false},
		},
		{
			name:        "every 5",
			nth:         5,
			vec:         Integer([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			outBooleans: []bool{false, false, false, false, true, false, false, false, false, true},
		},
		{
			name:        "every 1",
			nth:         1,
			vec:         Integer([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			outBooleans: []bool{true, true, true, true, true, true, true, true, true, true},
		},
		{
			name:        "empty",
			nth:         3,
			vec:         Time([]time.Time{}),
			outBooleans: []bool{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			outBooleans := data.vec.Nth(data.nth)
			if !reflect.DeepEqual(outBooleans, data.outBooleans) {
				t.Error(fmt.Sprintf("vec.Even() (%v) does not match data.outBooleans (%v)", outBooleans, data.outBooleans))
			}
		})
	}
}
