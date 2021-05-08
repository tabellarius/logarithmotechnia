package vector

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"testing"
)

func TestInteger(t *testing.T) {
	emptyNA := []bool{false, false, false, false, false}

	testData := []struct {
		name          string
		data          []int
		na            []bool
		names         map[string]int
		expectedNames map[string]int
		isEmpty       bool
	}{
		{
			name:    "normal + na",
			data:    []int{1, 2, 3, 4, 5},
			na:      []bool{false, false, false, false, false},
			names:   nil,
			isEmpty: false,
		},
		{
			name:    "normal + empty na",
			data:    []int{1, 2, 3, 4, 5},
			na:      []bool{},
			names:   nil,
			isEmpty: false,
		},
		{
			name:    "normal + nil na",
			data:    []int{1, 2, 3, 4, 5},
			na:      nil,
			names:   nil,
			isEmpty: false,
		},
		{
			name:    "normal + na",
			data:    []int{1, 2, 3, 4, 5},
			na:      []bool{false, true, true, true, false},
			names:   nil,
			isEmpty: false,
		},
		{
			name:    "normal + incorrect sized na",
			data:    []int{1, 2, 3, 4, 5},
			na:      []bool{false, false, false, false},
			names:   nil,
			isEmpty: true,
		},
		{
			name:          "normal + names",
			data:          []int{1, 2, 3, 4, 5},
			na:            []bool{false, false, false, false, false},
			names:         map[string]int{"one": 1, "three": 3, "five": 5},
			expectedNames: map[string]int{"one": 1, "three": 3, "five": 5},
			isEmpty:       false,
		},
		{
			name:          "normal + incorrect names",
			data:          []int{1, 2, 3, 4, 5},
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
				v = Integer(data.data, data.na)
			} else {
				config := Config{NamesMap: data.names}
				v = Integer(data.data, data.na, config).(*vector)
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
					t.Error(fmt.Printf("Vector length (%d) is not equal to data length (%d)\n", vv.length, length))
				}

				payload, ok := vv.payload.(*integer)
				if !ok {
					t.Error("Payload is not integer")
				} else {
					if !reflect.DeepEqual(payload.data[1:], data.data) {
						t.Error(fmt.Printf("Payload data (%v) is not equal to correct data (%v)\n",
							payload.data[1:], data.data))
					}

					if vv.length != vv.DefNameable.length || vv.length != payload.length ||
						vv.length != payload.DefNA.length {
						t.Error(fmt.Printf("Lengths are different: (vv.length - %d, "+
							"vv.DefNameable.length - %d, payload.length - %d, "+
							"payload.DefNA.length - %d",
							vv.length, vv.DefNameable.length, payload.length, payload.DefNA.length))
					}
				}

				if len(data.na) > 0 && len(data.na) == length {
					if !reflect.DeepEqual(payload.na[1:], data.na) {
						t.Error(fmt.Printf("Payload na (%v) is not equal to correct na (%v)\n",
							payload.na[1:], data.na))
					}
				} else if len(data.na) == 0 {
					if !reflect.DeepEqual(payload.na[1:], emptyNA) {
						t.Error(fmt.Printf("len(data.na) == 0 : incorrect payload.na (%v)", payload.na))
					}
				} else {
					t.Error("error")
				}

				if data.names != nil {
					if !reflect.DeepEqual(vv.names, data.expectedNames) {
						t.Error(fmt.Printf("Vector names (%v) is not equal to expected names (%v)",
							vv.names, data.expectedNames))
					}
				}

			}
		})
	}
}

func TestInteger_Booleans(t *testing.T) {
	testData := []struct {
		in    []int
		inNA  []bool
		out   []bool
		outNA []bool
	}{
		{
			in:    []int{1, 3, 0, 100, 0},
			inNA:  []bool{false, false, false, false, false},
			out:   []bool{true, true, false, true, false},
			outNA: []bool{false, false, false, false, false},
		},
		{
			in:    []int{10, 0, 12, 14, 1110},
			inNA:  []bool{false, false, false, true, true},
			out:   []bool{true, false, true, false, false},
			outNA: []bool{false, false, false, true, true},
		},
		{
			in:    []int{1, 3, 0, 100, 0, -11, -10},
			inNA:  []bool{false, false, false, false, false, false, true},
			out:   []bool{true, true, false, true, false, true, false},
			outNA: []bool{false, false, false, false, false, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := Integer(data.in, data.inNA)
			payload := vec.(*vector).payload.(*integer)

			booleans, na := payload.Booleans()
			if !reflect.DeepEqual(booleans, data.out) {
				t.Error(fmt.Printf("Booleans (%v) are not equal to data.out (%v)\n", booleans, data.out))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Printf("NA (%v) are not equal to data.outNA (%v)\n", na, data.outNA))
			}
		})
	}
}

func TestInteger_Integers(t *testing.T) {
	testData := []struct {
		in    []int
		inNA  []bool
		out   []int
		outNA []bool
	}{
		{
			in:    []int{1, 3, 0, 100, 0},
			inNA:  []bool{false, false, false, false, false},
			out:   []int{1, 3, 0, 100, 0},
			outNA: []bool{false, false, false, false, false},
		},
		{
			in:    []int{10, 0, 12, 14, 1110},
			inNA:  []bool{false, false, false, true, true},
			out:   []int{10, 0, 12, 0, 0},
			outNA: []bool{false, false, false, true, true},
		},
		{
			in:    []int{1, 3, 0, 100, 0, -11, -10},
			inNA:  []bool{false, false, false, false, false, false, true},
			out:   []int{1, 3, 0, 100, 0, -11, 0},
			outNA: []bool{false, false, false, false, false, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := Integer(data.in, data.inNA)
			payload := vec.(*vector).payload.(*integer)

			integers, na := payload.Integers()
			if !reflect.DeepEqual(integers, data.out) {
				t.Error(fmt.Printf("Integers (%v) are not equal to data.out (%v)\n", integers, data.out))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Printf("NA (%v) are not equal to data.outNA (%v)\n", na, data.outNA))
			}
		})
	}
}

func TestInteger_Floats(t *testing.T) {
	testData := []struct {
		in    []int
		inNA  []bool
		out   []float64
		outNA []bool
	}{
		{
			in:    []int{1, 3, 0, 100, 0},
			inNA:  []bool{false, false, false, false, false},
			out:   []float64{1, 3, 0, 100, 0},
			outNA: []bool{false, false, false, false, false},
		},
		{
			in:    []int{10, 0, 12, 14, 1110},
			inNA:  []bool{false, false, false, true, true},
			out:   []float64{10, 0, 12, math.NaN(), math.NaN()},
			outNA: []bool{false, false, false, true, true},
		},
		{
			in:    []int{1, 3, 0, 100, 0, -11, -10},
			inNA:  []bool{false, false, false, false, false, false, true},
			out:   []float64{1, 3, 0, 100, 0, -11, math.NaN()},
			outNA: []bool{false, false, false, false, false, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := Integer(data.in, data.inNA)
			payload := vec.(*vector).payload.(*integer)

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
				t.Error(fmt.Printf("Floats (%v) are not equal to data.out (%v)\n", floats, data.out))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Printf("NA (%v) are not equal to data.outNA (%v)\n", na, data.outNA))
			}
		})
	}
}

func TestInteger_Strings(t *testing.T) {
	testData := []struct {
		in    []int
		inNA  []bool
		out   []string
		outNA []bool
	}{
		{
			in:    []int{1, 3, 0, 100, 0},
			inNA:  []bool{false, false, false, false, false},
			out:   []string{"1", "3", "0", "100", "0"},
			outNA: []bool{false, false, false, false, false},
		},
		{
			in:    []int{10, 0, 12, 14, 1110},
			inNA:  []bool{false, false, false, true, true},
			out:   []string{"10", "0", "12", "", ""},
			outNA: []bool{false, false, false, true, true},
		},
		{
			in:    []int{1, 3, 0, 100, 0, -11, -10},
			inNA:  []bool{false, false, false, false, false, false, true},
			out:   []string{"1", "3", "0", "100", "0", "-11", ""},
			outNA: []bool{false, false, false, false, false, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := Integer(data.in, data.inNA)
			payload := vec.(*vector).payload.(*integer)

			strings, na := payload.Strings()
			if !reflect.DeepEqual(strings, data.out) {
				t.Error(fmt.Printf("Integers (%v) are not equal to data.out (%v)\n", strings, data.out))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Printf("NA (%v) are not equal to data.outNA (%v)\n", na, data.outNA))
			}
		})
	}
}

func TestInteger_ByIndices(t *testing.T) {
	testData := []struct {
		name string
	}{
		{},
	}

	_ = testData
}
