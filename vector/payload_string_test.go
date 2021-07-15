package vector

import (
	"fmt"
	"math"
	"math/cmplx"
	"reflect"
	"strconv"
	"testing"
)

func TestString(t *testing.T) {
	emptyNA := []bool{false, false, false, false, false}

	testData := []struct {
		name          string
		data          []string
		na            []bool
		outData       []string
		names         map[string]int
		expectedNames map[string]int
		isEmpty       bool
	}{
		{
			name:    "normal + false na",
			data:    []string{"one", "two", "three", "four", "five"},
			na:      []bool{false, false, false, false, false},
			outData: []string{"one", "two", "three", "four", "five"},
			names:   nil,
			isEmpty: false,
		},
		{
			name:    "normal + empty na",
			data:    []string{"one", "two", "three", "four", "five"},
			na:      []bool{},
			outData: []string{"one", "two", "three", "four", "five"},
			names:   nil,
			isEmpty: false,
		},
		{
			name:    "normal + nil na",
			data:    []string{"one", "two", "three", "four", "five"},
			na:      nil,
			outData: []string{"one", "two", "three", "four", "five"},
			names:   nil,
			isEmpty: false,
		},
		{
			name:    "normal + mixed na",
			data:    []string{"one", "two", "three", "four", "five"},
			na:      []bool{false, true, true, true, false},
			outData: []string{"one", "", "", "", "five"},
			names:   nil,
			isEmpty: false,
		},
		{
			name:    "normal + incorrect sized na",
			data:    []string{"one", "two", "three", "four", "five"},
			na:      []bool{false, false, false, false},
			names:   nil,
			isEmpty: true,
		},
		{
			name:          "normal + names",
			data:          []string{"one", "two", "three", "four", "five"},
			na:            []bool{false, false, false, false, false},
			outData:       []string{"one", "two", "three", "four", "five"},
			names:         map[string]int{"one": 1, "three": 3, "five": 5},
			expectedNames: map[string]int{"one": 1, "three": 3, "five": 5},
			isEmpty:       false,
		},
		{
			name:          "normal + incorrect names",
			data:          []string{"one", "two", "three", "four", "five"},
			na:            []bool{false, false, false, false, false},
			outData:       []string{"one", "two", "three", "four", "five"},
			names:         map[string]int{"zero": 0, "one": 1, "three": 3, "five": 5, "seven": 7},
			expectedNames: map[string]int{"one": 1, "three": 3, "five": 5},
			isEmpty:       false,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			var v Vector
			if data.names == nil {
				v = String(data.data, data.na)
			} else {
				config := Config{NamesMap: data.names}
				v = String(data.data, data.na, config).(*vector)
			}

			vv := v.(*vector)

			if data.isEmpty {
				naPayload, ok := vv.payload.(*naPayload)
				if !ok || naPayload.Len() > 0 {
					t.Error("Vector's payload is not empty")
				}
			} else {
				length := len(data.data)
				if vv.length != length {
					t.Error(fmt.Sprintf("Vector length (%d) is not equal to data length (%d)\n", vv.length, length))
				}

				payload, ok := vv.payload.(*stringPayload)
				if !ok {
					t.Error("Payload is not integerPayload")
				} else {
					if !reflect.DeepEqual(payload.data, data.outData) {
						t.Error(fmt.Sprintf("Payload data (%v) is not equal to correct data (%v)\n",
							payload.data, data.outData))
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
							payload.na, data.na))
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

func TestStringPayload_Type(t *testing.T) {
	vec := String([]string{}, nil)
	if vec.Type() != "string" {
		t.Error("Type is incorrect.")
	}
}

func TestStringPayload_Len(t *testing.T) {
	testData := []struct {
		in        []string
		outLength int
	}{
		{[]string{"one", "two", "three", "four", "five"}, 5},
		{[]string{"one", "two", "three"}, 3},
		{[]string{}, 0},
		{nil, 0},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			payload := String(data.in, nil).(*vector).payload
			if payload.Len() != data.outLength {
				t.Error(fmt.Sprintf("Payloads's length (%d) is not equal to out (%d)",
					payload.Len(), data.outLength))
			}
		})
	}
}

func TestStringPayload_Booleans(t *testing.T) {
	testData := []struct {
		in    []string
		inNA  []bool
		out   []bool
		outNA []bool
	}{
		{
			in:    []string{"1", "3", "", "100", ""},
			inNA:  []bool{false, false, false, false, false},
			out:   []bool{true, true, false, true, false},
			outNA: []bool{false, false, false, false, false},
		},
		{
			in:    []string{"10", "", "12", "14", "1110"},
			inNA:  []bool{false, false, false, true, true},
			out:   []bool{true, false, true, false, false},
			outNA: []bool{false, false, false, true, true},
		},
		{
			in:    []string{"1", "3", "", "100", "", "-11", "-10"},
			inNA:  []bool{false, false, false, false, false, false, true},
			out:   []bool{true, true, false, true, false, true, false},
			outNA: []bool{false, false, false, false, false, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := String(data.in, data.inNA)
			payload := vec.(*vector).payload.(*stringPayload)

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

func TestStringPayload_Integers(t *testing.T) {
	testData := []struct {
		in    []string
		inNA  []bool
		out   []int
		outNA []bool
	}{
		{
			in:    []string{"1", "3", "", "100", ""},
			inNA:  []bool{false, false, false, false, false},
			out:   []int{1, 3, 0, 100, 0},
			outNA: []bool{false, false, false, false, false},
		},
		{
			in:    []string{"10", "", "12", "14", "1110"},
			inNA:  []bool{false, false, false, true, true},
			out:   []int{10, 0, 12, 0, 0},
			outNA: []bool{false, false, false, true, true},
		},
		{
			in:    []string{"1", "3", "", "100", "", "-11", "-10"},
			inNA:  []bool{false, false, false, false, false, false, true},
			out:   []int{1, 3, 0, 100, 0, -11, 0},
			outNA: []bool{false, false, false, false, false, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := String(data.in, data.inNA)
			payload := vec.(*vector).payload.(*stringPayload)

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

func TestStringPayload_Floats(t *testing.T) {
	testData := []struct {
		in    []string
		inNA  []bool
		out   []float64
		outNA []bool
	}{
		{
			in:    []string{"1", "3", "", "100", ""},
			inNA:  []bool{false, false, false, false, false},
			out:   []float64{1, 3, 0, 100, 0},
			outNA: []bool{false, false, false, false, false},
		},
		{
			in:    []string{"10", "", "12", "14", "1110"},
			inNA:  []bool{false, false, false, true, true},
			out:   []float64{10, 0, 12, math.NaN(), math.NaN()},
			outNA: []bool{false, false, false, true, true},
		},
		{
			in:    []string{"1", "3", "", "100", "", "-11", "-10"},
			inNA:  []bool{false, false, false, false, false, false, true},
			out:   []float64{1, 3, 0, 100, 0, -11, math.NaN()},
			outNA: []bool{false, false, false, false, false, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := String(data.in, data.inNA)
			payload := vec.(*vector).payload.(*stringPayload)

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

func TestStringPayload_Complexes(t *testing.T) {
	testData := []struct {
		in    []string
		inNA  []bool
		out   []complex128
		outNA []bool
	}{
		{
			in:    []string{"1+1i", "3-3i", "0", "100 + 50i", "0+0i"},
			inNA:  []bool{false, false, false, false, false},
			out:   []complex128{1 + 1i, 3 - 3i, 0 + 0i, cmplx.NaN(), 0 + 0i},
			outNA: []bool{false, false, false, false, false},
		},
		{
			in:    []string{"10+10i", "0", "12+6i", "14+7i", "1110+0i"},
			inNA:  []bool{false, false, false, true, true},
			out:   []complex128{10 + 10i, 0 + 0i, 12 + 6i, cmplx.NaN(), cmplx.NaN()},
			outNA: []bool{false, false, false, true, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := String(data.in, data.inNA)
			payload := vec.(*vector).payload.(*stringPayload)

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

func TestStringPayload_Strings(t *testing.T) {
	testData := []struct {
		in    []string
		inNA  []bool
		out   []string
		outNA []bool
	}{
		{
			in:    []string{"1", "3", "0", "100", ""},
			inNA:  []bool{false, false, false, false, false},
			out:   []string{"1", "3", "0", "100", ""},
			outNA: []bool{false, false, false, false, false},
		},
		{
			in:    []string{"10", "", "12", "14", "1110"},
			inNA:  []bool{false, false, false, true, true},
			out:   []string{"10", "", "12", "", ""},
			outNA: []bool{false, false, false, true, true},
		},
		{
			in:    []string{"1", "3", "0", "100", "", "-11", "-10"},
			inNA:  []bool{false, false, false, false, false, false, true},
			out:   []string{"1", "3", "0", "100", "", "-11", ""},
			outNA: []bool{false, false, false, false, false, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := String(data.in, data.inNA)
			payload := vec.(*vector).payload.(*stringPayload)

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

func TestStringPayload_Interfaces(t *testing.T) {
	testData := []struct {
		in    []string
		inNA  []bool
		out   []interface{}
		outNA []bool
	}{
		{
			in:    []string{"1", "3", "0", "100", ""},
			inNA:  []bool{false, false, false, false, false},
			out:   []interface{}{"1", "3", "0", "100", ""},
			outNA: []bool{false, false, false, false, false},
		},
		{
			in:    []string{"10", "", "12", "14", "1110"},
			inNA:  []bool{false, false, false, true, true},
			out:   []interface{}{"10", "", "12", nil, nil},
			outNA: []bool{false, false, false, true, true},
		},
		{
			in:    []string{"1", "3", "0", "100", "", "-11", "-10"},
			inNA:  []bool{false, false, false, false, false, false, true},
			out:   []interface{}{"1", "3", "0", "100", "", "-11", nil},
			outNA: []bool{false, false, false, false, false, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := String(data.in, data.inNA)
			payload := vec.(*vector).payload.(*stringPayload)

			interfaces, na := payload.Interfaces()
			if !reflect.DeepEqual(interfaces, data.out) {
				t.Error(fmt.Sprintf("Interfaces (%v) are not equal to data.out (%v)\n", interfaces, data.out))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("IsNA (%v) are not equal to data.outNA (%v)\n", na, data.outNA))
			}
		})
	}
}

func TestStringPayload_ByIndices(t *testing.T) {
	vec := String([]string{"1", "2", "3", "4", "5"}, []bool{false, false, false, false, true})
	testData := []struct {
		name    string
		indices []int
		out     []string
		outNA   []bool
	}{
		{
			name:    "all",
			indices: []int{1, 2, 3, 4, 5},
			out:     []string{"1", "2", "3", "4", ""},
			outNA:   []bool{false, false, false, false, true},
		},
		{
			name:    "all reverse",
			indices: []int{5, 4, 3, 2, 1},
			out:     []string{"", "4", "3", "2", "1"},
			outNA:   []bool{true, false, false, false, false},
		},
		{
			name:    "some",
			indices: []int{5, 1, 3},
			out:     []string{"", "1", "3"},
			outNA:   []bool{true, false, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := vec.ByIndices(data.indices).(*vector).payload.(*stringPayload)
			if !reflect.DeepEqual(payload.data, data.out) {
				t.Error(fmt.Sprintf("payload.data (%v) is not equal to data.out (%v)", payload.data, data.out))
			}
			if !reflect.DeepEqual(payload.na, data.outNA) {
				t.Error(fmt.Sprintf("payload.data (%v) is not equal to data.out (%v)", payload.data, data.out))
			}
		})
	}
}

func TestStringPayload_SupportsWhicher(t *testing.T) {
	testData := []struct {
		name        string
		filter      interface{}
		isSupported bool
	}{
		{
			name:        "func(int, string, bool) bool",
			filter:      func(int, string, bool) bool { return true },
			isSupported: true,
		},
		{
			name:        "func(int, float64, bool) bool",
			filter:      func(int, float64, bool) bool { return true },
			isSupported: false,
		},
	}

	payload := String([]string{"one"}, nil).(*vector).payload.(Whichable)
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if payload.SupportsWhicher(data.filter) != data.isSupported {
				t.Error("Selector's support is incorrect.")
			}
		})
	}
}

func TestStringPayload_Whicher(t *testing.T) {
	testData := []struct {
		name string
		fn   interface{}
		out  []bool
	}{
		{
			name: "Odd",
			fn:   func(idx int, _ string, _ bool) bool { return idx%2 == 1 },
			out:  []bool{true, false, true, false, true, false, true, false, true, false},
		},
		{
			name: "Even",
			fn:   func(idx int, _ string, _ bool) bool { return idx%2 == 0 },
			out:  []bool{false, true, false, true, false, true, false, true, false, true},
		},
		{
			name: "func(_ int, val string, _ bool) bool {return val == 2}",
			fn:   func(_ int, val string, _ bool) bool { return val == "2" },
			out:  []bool{false, true, false, false, false, true, false, false, false, false},
		},
		{
			name: "func() bool {return true}",
			fn:   func() bool { return true },
			out:  []bool{false, false, false, false, false, false, false, false, false, false},
		},
	}

	payload := String([]string{"1", "2", "39", "4", "56", "2", "45", "90", "4", "3"}, nil).(*vector).payload.(Whichable)

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			result := payload.Which(data.fn)
			if !reflect.DeepEqual(result, data.out) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to out (%v)", result, data.out))
			}
		})
	}
}

func TestStringPayload_SupportsApplier(t *testing.T) {
	testData := []struct {
		name        string
		applier     interface{}
		isSupported bool
	}{
		{
			name:        "func(int, string, bool) (string, bool)",
			applier:     func(int, string, bool) (string, bool) { return "", true },
			isSupported: true,
		},
		{
			name:        "func(int, string, bool) bool",
			applier:     func(int, string, bool) bool { return true },
			isSupported: false,
		},
	}

	payload := String([]string{}, nil).(*vector).payload.(Appliable)
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if payload.SupportsApplier(data.applier) != data.isSupported {
				t.Error("Applier's support is incorrect.")
			}
		})
	}
}

func TestStringPayload_Apply(t *testing.T) {
	testData := []struct {
		name        string
		applier     interface{}
		dataIn      []string
		naIn        []bool
		dataOut     []string
		naOut       []bool
		isNAPayload bool
	}{
		{
			name: "regular",
			applier: func(_ int, val string, na bool) (string, bool) {
				return fmt.Sprintf("%s.%s", val, val), na
			},
			dataIn:      []string{"1", "9", "3", "5", "7"},
			naIn:        []bool{false, true, false, true, false},
			dataOut:     []string{"1.1", "", "3.3", "", "7.7"},
			naOut:       []bool{false, true, false, true, false},
			isNAPayload: false,
		},
		{
			name: "manipulate na",
			applier: func(idx int, val string, na bool) (string, bool) {
				if idx == 5 {
					return "1", true
				}
				return val, na
			},
			dataIn:      []string{"1", "2", "3", "4", "5"},
			naIn:        []bool{false, false, true, false, false},
			dataOut:     []string{"1", "2", "", "4", ""},
			naOut:       []bool{false, false, true, false, true},
			isNAPayload: false,
		},
		{
			name:        "incorrect applier",
			applier:     func(int, string, bool) bool { return true },
			dataIn:      []string{"1", "9", "3", "5", "7"},
			naIn:        []bool{false, true, false, true, false},
			dataOut:     []string{"", "", "", "", ""},
			naOut:       []bool{true, true, true, true, true},
			isNAPayload: true,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := String(data.dataIn, data.naIn).(*vector).payload.(Appliable).Apply(data.applier)

			if !data.isNAPayload {
				payloadOut := payload.(*stringPayload)
				if !reflect.DeepEqual(data.dataOut, payloadOut.data) {
					t.Error(fmt.Sprintf("Output data (%v) does not match expected (%v)",
						payloadOut.data, data.dataOut))
				}
				if !reflect.DeepEqual(data.naOut, payloadOut.na) {
					t.Error(fmt.Sprintf("Output NA (%v) does not match expected (%v)",
						payloadOut.na, data.naOut))
				}
			} else {
				_, ok := payload.(*naPayload)
				if !ok {
					t.Error("Payload is not NA")
				}
			}
		})
	}
}
