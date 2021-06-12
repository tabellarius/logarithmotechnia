package vector

import (
	"fmt"
	"math"
	"math/cmplx"
	"reflect"
	"strconv"
	"testing"
)

func TestBoolean(t *testing.T) {
	emptyNA := []bool{false, false, false, false, false}

	testData := []struct {
		name          string
		data          []bool
		na            []bool
		names         map[string]int
		expectedNames map[string]int
		isEmpty       bool
	}{
		{
			name:    "normal + na",
			data:    []bool{true, false, true, false, true},
			na:      []bool{false, false, false, false, false},
			names:   nil,
			isEmpty: false,
		},
		{
			name:    "normal + empty na",
			data:    []bool{true, false, true, false, true},
			na:      []bool{},
			names:   nil,
			isEmpty: false,
		},
		{
			name:    "normal + nil na",
			data:    []bool{true, false, true, false, true},
			na:      nil,
			names:   nil,
			isEmpty: false,
		},
		{
			name:    "normal + na",
			data:    []bool{true, false, true, false, true},
			na:      []bool{false, true, true, true, false},
			names:   nil,
			isEmpty: false,
		},
		{
			name:    "normal + incorrect sized na",
			data:    []bool{true, false, true, false, true},
			na:      []bool{false, false, false, false},
			names:   nil,
			isEmpty: true,
		},
		{
			name:          "normal + names",
			data:          []bool{true, false, true, false, true},
			na:            []bool{false, false, false, false, false},
			names:         map[string]int{"one": 1, "three": 3, "five": 5},
			expectedNames: map[string]int{"one": 1, "three": 3, "five": 5},
			isEmpty:       false,
		},
		{
			name:          "normal + incorrect names",
			data:          []bool{true, false, true, false, true},
			na:            []bool{false, false, false, false, false},
			names:         map[string]int{"zero": 0, "one": 1, "three": 3, "five": 5, "seven": 7},
			expectedNames: map[string]int{"one": 1, "three": 3, "five": 5},
			isEmpty:       false,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			var v Vector
			if data.names == nil {
				v = Boolean(data.data, data.na)
			} else {
				config := Config{NamesMap: data.names}
				v = Boolean(data.data, data.na, config).(*vector)
			}

			vv := v.(*vector)

			if data.isEmpty {
				_, ok := vv.payload.(*emptyPayload)
				if !ok {
					t.Error("Vector's payload is not empty")
				}
			} else {
				length := len(data.data)
				if vv.length != length {
					t.Error(fmt.Sprintf("Vector length (%d) is not equal to data length (%d)\n", vv.length, length))
				}

				payload, ok := vv.payload.(*boolean)
				if !ok {
					t.Error("Payload is not boolean")
				} else {
					if !reflect.DeepEqual(payload.data, data.data) {
						t.Error(fmt.Sprintf("Payload data (%v) is not equal to correct data (%v)\n",
							payload.data[1:], data.data))
					}

					if vv.length != vv.DefNameable.length || vv.length != payload.length {
						t.Error(fmt.Sprintf("Lengths are different: (vv.length - %d, "+
							"vv.DefNameable.length - %d, payload.length - %d, ",
							vv.length, vv.DefNameable.length, payload.length))
					}
				}

				if len(data.na) > 0 && len(data.na) == length {
					if !reflect.DeepEqual(payload.na, data.na) {
						t.Error(fmt.Sprintf("Payload na (%v) is not equal to correct na (%v)\n",
							payload.na[1:], data.na))
					}
				} else if len(data.na) == 0 {
					if !reflect.DeepEqual(payload.na, emptyNA) {
						t.Error(fmt.Sprintf("len(data.na) == 0 : incorrect payload.na (%v)", payload.na))
					}
				} else {
					t.Error("error")
				}

				if data.names != nil {
					if !reflect.DeepEqual(vv.names, data.expectedNames) {
						t.Error(fmt.Sprintf("Vector names (%v) is not equal to out names (%v)",
							vv.names, data.expectedNames))
					}
				}

			}
		})
	}
}

func TestBoolean_Len(t *testing.T) {
	testData := []struct {
		in        []bool
		outLength int
	}{
		{[]bool{true, false, true, false, true}, 5},
		{[]bool{true, false, true}, 3},
		{[]bool{}, 0},
		{nil, 0},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			payload := Boolean(data.in, nil).(*vector).payload
			if payload.Len() != data.outLength {
				t.Error(fmt.Sprintf("Payloads's length (%d) is not equal to out (%d)",
					payload.Len(), data.outLength))
			}
		})
	}
}

func TestBoolean_Booleans(t *testing.T) {
	testData := []struct {
		in    []bool
		inNA  []bool
		out   []bool
		outNA []bool
	}{
		{
			in:    []bool{true, true, false, true, false},
			inNA:  []bool{false, false, false, false, false},
			out:   []bool{true, true, false, true, false},
			outNA: []bool{false, false, false, false, false},
		},
		{
			in:    []bool{true, false, true, true, true},
			inNA:  []bool{false, false, false, true, true},
			out:   []bool{true, false, true, false, false},
			outNA: []bool{false, false, false, true, true},
		},
		{
			in:    []bool{true, true, false, true, false, true, true},
			inNA:  []bool{false, false, false, false, false, false, true},
			out:   []bool{true, true, false, true, false, true, false},
			outNA: []bool{false, false, false, false, false, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := Boolean(data.in, data.inNA)
			payload := vec.(*vector).payload.(*boolean)

			booleans, na := payload.Booleans()
			if !reflect.DeepEqual(booleans, data.out) {
				t.Error(fmt.Sprintf("Booleans (%v) are not equal to data.out (%v)\n", booleans, data.out))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("IsNA (%v) are not equal to data.outNA (%v)\n", na, data.outNA))
			}
		})
	}
}

func TestBoolean_Integers(t *testing.T) {
	testData := []struct {
		in    []bool
		inNA  []bool
		out   []int
		outNA []bool
	}{
		{
			in:    []bool{true, true, false, true, false},
			inNA:  []bool{false, false, false, false, false},
			out:   []int{1, 1, 0, 1, 0},
			outNA: []bool{false, false, false, false, false},
		},
		{
			in:    []bool{true, false, true, true, true},
			inNA:  []bool{false, false, false, true, true},
			out:   []int{1, 0, 1, 0, 0},
			outNA: []bool{false, false, false, true, true},
		},
		{
			in:    []bool{true, true, false, true, false, true, true},
			inNA:  []bool{false, false, false, false, false, false, true},
			out:   []int{1, 1, 0, 1, 0, 1, 0},
			outNA: []bool{false, false, false, false, false, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := Boolean(data.in, data.inNA)
			payload := vec.(*vector).payload.(*boolean)

			integers, na := payload.Integers()
			if !reflect.DeepEqual(integers, data.out) {
				t.Error(fmt.Sprintf("Integers (%v) are not equal to data.out (%v)\n", integers, data.out))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("IsNA (%v) are not equal to data.outNA (%v)\n", na, data.outNA))
			}
		})
	}
}

func TestBoolean_Floats(t *testing.T) {
	testData := []struct {
		in    []bool
		inNA  []bool
		out   []float64
		outNA []bool
	}{
		{
			in:    []bool{true, true, false, true, false},
			inNA:  []bool{false, false, false, false, false},
			out:   []float64{1, 1, 0, 1, 0},
			outNA: []bool{false, false, false, false, false},
		},
		{
			in:    []bool{true, false, true, true, true},
			inNA:  []bool{false, false, false, true, true},
			out:   []float64{1, 0, 1, math.NaN(), math.NaN()},
			outNA: []bool{false, false, false, true, true},
		},
		{
			in:    []bool{true, true, false, true, false, true, true},
			inNA:  []bool{false, false, false, false, false, false, true},
			out:   []float64{1, 1, 0, 1, 0, 1, math.NaN()},
			outNA: []bool{false, false, false, false, false, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := Boolean(data.in, data.inNA)
			payload := vec.(*vector).payload.(*boolean)

			floats, na := payload.Floats()
			correct := true
			for i := 0; i < len(floats); i++ {
				if math.IsNaN(data.out[i]) {
					if !math.IsNaN(floats[i]) {
						correct = false
					}
				} else if floats[i] != data.out[i] {
					correct = false
				}
			}
			if !correct {
				t.Error(fmt.Sprintf("Floats (%v) are not equal to data.out (%v)\n", floats, data.out))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("IsNA (%v) are not equal to data.outNA (%v)\n", na, data.outNA))
			}
		})
	}
}

func TestBoolean_Complexes(t *testing.T) {
	testData := []struct {
		in    []bool
		inNA  []bool
		out   []complex128
		outNA []bool
	}{
		{
			in:    []bool{true, true, false, true, false},
			inNA:  []bool{false, false, false, false, false},
			out:   []complex128{1 + 0i, 1 + 0i, 0 + 0i, 1 + 0i, 0 + 0i},
			outNA: []bool{false, false, false, false, false},
		},
		{
			in:    []bool{true, false, true, true, true},
			inNA:  []bool{false, false, false, true, true},
			out:   []complex128{1 + 0i, 0 + 0i, 1 + 0i, cmplx.NaN(), cmplx.NaN()},
			outNA: []bool{false, false, false, true, true},
		},
		{
			in:    []bool{true, true, false, true, false, true, true},
			inNA:  []bool{false, false, false, false, false, false, true},
			out:   []complex128{1 + 0i, 1 + 0i, 0 + 0i, 1 + 0i, 0 + 0i, 1 + 0i, cmplx.NaN()},
			outNA: []bool{false, false, false, false, false, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := Boolean(data.in, data.inNA)
			payload := vec.(*vector).payload.(*boolean)

			complexes, na := payload.Complexes()
			correct := true
			for i := 0; i < len(complexes); i++ {
				if cmplx.IsNaN(data.out[i]) {
					if !cmplx.IsNaN(complexes[i]) {
						correct = false
					}
				} else if complexes[i] != data.out[i] {
					correct = false
				}
			}
			if !correct {
				t.Error(fmt.Sprintf("Complexes (%v) are not equal to data.out (%v)\n", complexes, data.out))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("IsNA (%v) are not equal to data.outNA (%v)\n", na, data.outNA))
			}
		})
	}
}

func TestBoolean_Strings(t *testing.T) {
	testData := []struct {
		in    []bool
		inNA  []bool
		out   []string
		outNA []bool
	}{
		{
			in:    []bool{true, true, false, true, false},
			inNA:  []bool{false, false, false, false, false},
			out:   []string{"true", "true", "false", "true", "false"},
			outNA: []bool{false, false, false, false, false},
		},
		{
			in:    []bool{true, false, true, true, true},
			inNA:  []bool{false, false, false, true, true},
			out:   []string{"true", "false", "true", "NA", "NA"},
			outNA: []bool{false, false, false, true, true},
		},
		{
			in:    []bool{true, true, false, true, false, true, true},
			inNA:  []bool{false, false, false, false, false, false, true},
			out:   []string{"true", "true", "false", "true", "false", "true", "NA"},
			outNA: []bool{false, false, false, false, false, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := Boolean(data.in, data.inNA)
			payload := vec.(*vector).payload.(*boolean)

			strings, na := payload.Strings()
			if !reflect.DeepEqual(strings, data.out) {
				t.Error(fmt.Sprintf("Strings (%v) are not equal to data.out (%v)\n", strings, data.out))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("IsNA (%v) are not equal to data.outNA (%v)\n", na, data.outNA))
			}
		})
	}
}

func TestBoolean_ByIndices(t *testing.T) {
	vec := Boolean([]bool{true, false, true, false, true}, []bool{false, false, false, false, true})
	testData := []struct {
		name    string
		indices []int
		out     []bool
		outNA   []bool
	}{
		{
			name:    "all",
			indices: []int{1, 2, 3, 4, 5},
			out:     []bool{true, false, true, false, true},
			outNA:   []bool{false, false, false, false, true},
		},
		{
			name:    "all reverse",
			indices: []int{5, 4, 3, 2, 1},
			out:     []bool{true, false, true, false, true},
			outNA:   []bool{true, false, false, false, false},
		},
		{
			name:    "some",
			indices: []int{5, 1, 3},
			out:     []bool{true, true, true},
			outNA:   []bool{true, false, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := vec.ByIndices(data.indices).(*vector).payload.(*boolean)
			if !reflect.DeepEqual(payload.data, data.out) {
				t.Error(fmt.Sprintf("payload.data (%v) is not equal to data.out (%v)", payload.data, data.out))
			}
			if !reflect.DeepEqual(payload.na, data.outNA) {
				t.Error(fmt.Sprintf("payload.data (%v) is not equal to data.out (%v)", payload.data, data.out))
			}
		})
	}
}

func TestBoolean_SupportsSelector(t *testing.T) {
	testData := []struct {
		name        string
		filter      interface{}
		isSupported bool
	}{
		{
			name:        "func(int, bool, bool) bool",
			filter:      func(int, bool, bool) bool { return true },
			isSupported: true,
		},
		{
			name:        "func(int, float64, bool) bool",
			filter:      func(int, float64, bool) bool { return true },
			isSupported: false,
		},
	}

	payload := Boolean([]bool{true}, nil).(*vector).payload
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if payload.SupportsSelector(data.filter) != data.isSupported {
				t.Error("Selector's support is incorrect.")
			}
		})
	}
}

func TestBoolean_Select(t *testing.T) {
	testData := []struct {
		name string
		fn   interface{}
		out  []bool
	}{
		{
			name: "Odd",
			fn:   func(idx int, _ bool, _ bool) bool { return idx%2 == 1 },
			out:  []bool{true, false, true, false, true, false, true, false, true, false},
		},
		{
			name: "Even",
			fn:   func(idx int, _ bool, _ bool) bool { return idx%2 == 0 },
			out:  []bool{false, true, false, true, false, true, false, true, false, true},
		},
		{
			name: "Nth(3)",
			fn:   func(idx int, _ bool, _ bool) bool { return idx%3 == 0 },
			out:  []bool{false, false, true, false, false, true, false, false, true, false},
		},
		{
			name: "func() bool {return true}",
			fn:   func() bool { return true },
			out:  []bool{false, false, false, false, false, false, false, false, false, false},
		},
	}

	payload := Boolean([]bool{true, false, true, false, true, false, true, false, true, false}, nil).(*vector).payload

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			result := payload.Select(data.fn)
			if !reflect.DeepEqual(result, data.out) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to out (%v)", result, data.out))
			}
		})
	}
}
