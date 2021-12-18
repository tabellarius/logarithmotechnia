package vector

import (
	"fmt"
	"logarithmotechnia/util"
	"math"
	"math/cmplx"
	"reflect"
	"strconv"
	"testing"
)

func TestInteger(t *testing.T) {
	emptyNA := []bool{false, false, false, false, false}

	testData := []struct {
		name    string
		data    []int
		na      []bool
		outData []int
		isEmpty bool
	}{
		{
			name:    "normal + false na",
			data:    []int{1, 2, 3, 4, 5},
			na:      []bool{false, false, false, false, false},
			outData: []int{1, 2, 3, 4, 5},
			isEmpty: false,
		},
		{
			name:    "normal + empty na",
			data:    []int{1, 2, 3, 4, 5},
			na:      []bool{},
			outData: []int{1, 2, 3, 4, 5},
			isEmpty: false,
		},
		{
			name:    "normal + nil na",
			data:    []int{1, 2, 3, 4, 5},
			na:      nil,
			outData: []int{1, 2, 3, 4, 5},
			isEmpty: false,
		},
		{
			name:    "normal + mixed na",
			data:    []int{1, 2, 3, 4, 5},
			na:      []bool{false, true, true, true, false},
			outData: []int{1, 0, 0, 0, 5},
			isEmpty: false,
		},
		{
			name:    "normal + incorrect sized na",
			data:    []int{1, 2, 3, 4, 5},
			na:      []bool{false, false, false, false},
			outData: []int{1, 2, 3, 4, 5},
			isEmpty: true,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			v := IntegerWithNA(data.data, data.na)

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

				payload, ok := vv.payload.(*integerPayload)
				if !ok {
					t.Error("Payload is not integerPayload")
				} else {
					if !reflect.DeepEqual(payload.data, data.outData) {
						t.Error(fmt.Sprintf("Payload data (%v) is not equal to correct data (%v)\n",
							payload.data, data.data))
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
			}
		})
	}
}

func TestIntegerPayload_Type(t *testing.T) {
	vec := IntegerWithNA([]int{}, nil)
	if vec.Type() != "integer" {
		t.Error("Type is incorrect.")
	}
}

func TestIntegerPayload_Len(t *testing.T) {
	testData := []struct {
		in        []int
		outLength int
	}{
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{1, 2, 3}, 3},
		{[]int{}, 0},
		{nil, 0},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			payload := IntegerWithNA(data.in, nil).(*vector).payload
			if payload.Len() != data.outLength {
				t.Error(fmt.Sprintf("Payloads's length (%d) is not equal to out (%d)",
					payload.Len(), data.outLength))
			}
		})
	}
}

func TestIntegerPayload_Booleans(t *testing.T) {
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
			vec := IntegerWithNA(data.in, data.inNA)
			payload := vec.(*vector).payload.(*integerPayload)

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

func TestIntegerPayload_Integers(t *testing.T) {
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
			vec := IntegerWithNA(data.in, data.inNA)
			payload := vec.(*vector).payload.(*integerPayload)

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

func TestIntegerPayload_Interfaces(t *testing.T) {
	testData := []struct {
		in    []int
		inNA  []bool
		out   []interface{}
		outNA []bool
	}{
		{
			in:    []int{1, 3, 0, 100, 0},
			inNA:  []bool{false, false, false, false, false},
			out:   []interface{}{1, 3, 0, 100, 0},
			outNA: []bool{false, false, false, false, false},
		},
		{
			in:    []int{10, 0, 12, 14, 1110},
			inNA:  []bool{false, false, false, true, true},
			out:   []interface{}{10, 0, 12, nil, nil},
			outNA: []bool{false, false, false, true, true},
		},
		{
			in:    []int{1, 3, 0, 100, 0, -11, -10},
			inNA:  []bool{false, false, false, false, false, false, true},
			out:   []interface{}{1, 3, 0, 100, 0, -11, nil},
			outNA: []bool{false, false, false, false, false, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := IntegerWithNA(data.in, data.inNA)
			payload := vec.(*vector).payload.(*integerPayload)

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

func TestIntegerPayload_Floats(t *testing.T) {
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
			vec := IntegerWithNA(data.in, data.inNA)
			payload := vec.(*vector).payload.(*integerPayload)

			floats, na := payload.Floats()
			if !util.EqualFloatArrays(floats, data.out) {
				t.Error(fmt.Sprintf("Floats (%v) are not equal to data.out (%v)\n", floats, data.out))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("IsNA (%v) are not equal to data.outNA (%v)\n", na, data.outNA))
			}
		})
	}
}

func TestIntegerPayload_Complexes(t *testing.T) {
	testData := []struct {
		in    []int
		inNA  []bool
		out   []complex128
		outNA []bool
	}{
		{
			in:    []int{1, 3, 0, 100, 0},
			inNA:  []bool{false, false, false, false, false},
			out:   []complex128{1 + 0i, 3 + 0i, 0 + 0i, 100 + 0i, 0 + 0i},
			outNA: []bool{false, false, false, false, false},
		},
		{
			in:    []int{10, 0, 12, 14, 1110},
			inNA:  []bool{false, false, false, true, true},
			out:   []complex128{10 + 0i, 0 + 0i, 12 + 0i, cmplx.NaN(), cmplx.NaN()},
			outNA: []bool{false, false, false, true, true},
		},
		{
			in:    []int{1, 3, 0, 100, 0, -11, -10},
			inNA:  []bool{false, false, false, false, false, false, true},
			out:   []complex128{1 + 0i, 3 + 0i, 0 + 0i, 100 + 0i, 0 + 0i, -11 + 0i, cmplx.NaN()},
			outNA: []bool{false, false, false, false, false, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := IntegerWithNA(data.in, data.inNA)
			payload := vec.(*vector).payload.(*integerPayload)

			complexes, na := payload.Complexes()
			if !util.EqualComplexArrays(complexes, data.out) {
				t.Error(fmt.Sprintf("Complexes (%v) are not equal to data.out (%v)\n", complexes, data.out))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("IsNA (%v) are not equal to data.outNA (%v)\n", na, data.outNA))
			}
		})
	}
}

func TestIntegerPayload_Strings(t *testing.T) {
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
			vec := IntegerWithNA(data.in, data.inNA)
			payload := vec.(*vector).payload.(*integerPayload)

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

func TestIntegerPayload_ByIndices(t *testing.T) {
	vec := IntegerWithNA([]int{1, 2, 3, 4, 5}, []bool{false, false, false, false, true})
	testData := []struct {
		name    string
		indices []int
		out     []int
		outNA   []bool
	}{
		{
			name:    "all",
			indices: []int{1, 2, 3, 4, 5},
			out:     []int{1, 2, 3, 4, 0},
			outNA:   []bool{false, false, false, false, true},
		},
		{
			name:    "all reverse",
			indices: []int{5, 4, 3, 2, 1},
			out:     []int{0, 4, 3, 2, 1},
			outNA:   []bool{true, false, false, false, false},
		},
		{
			name:    "some",
			indices: []int{5, 1, 3},
			out:     []int{0, 1, 3},
			outNA:   []bool{true, false, false},
		},
		{
			name:    "with zero",
			indices: []int{5, 1, 0, 3},
			out:     []int{0, 1, 0, 3},
			outNA:   []bool{true, false, true, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := vec.ByIndices(data.indices).(*vector).payload.(*integerPayload)
			if !reflect.DeepEqual(payload.data, data.out) {
				t.Error(fmt.Sprintf("payload.data (%v) is not equal to data.out (%v)", payload.data, data.out))
			}
			if !reflect.DeepEqual(payload.na, data.outNA) {
				t.Error(fmt.Sprintf("payload.data (%v) is not equal to data.out (%v)", payload.data, data.out))
			}
		})
	}
}

func TestIntegerPayload_SupportsWhicher(t *testing.T) {
	testData := []struct {
		name        string
		filter      interface{}
		isSupported bool
	}{
		{
			name:        "func(int, int, bool) bool",
			filter:      func(int, int, bool) bool { return true },
			isSupported: true,
		},
		{
			name:        "func(int, bool) bool",
			filter:      func(int, bool) bool { return true },
			isSupported: true,
		},
		{
			name:        "func(int, float64, bool) bool",
			filter:      func(int, float64, bool) bool { return true },
			isSupported: false,
		},
	}

	payload := IntegerWithNA([]int{1}, nil).(*vector).payload.(Whichable)
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if payload.SupportsWhicher(data.filter) != data.isSupported {
				t.Error("Selector's support is incorrect.")
			}
		})
	}
}

func TestIntegerPayload_Whicher(t *testing.T) {
	testData := []struct {
		name string
		fn   interface{}
		out  []bool
	}{
		{
			name: "Odd",
			fn:   func(idx int, _ int, _ bool) bool { return idx%2 == 1 },
			out:  []bool{true, false, true, false, true, false, true, false, true, false},
		},
		{
			name: "Even",
			fn:   func(idx int, _ int, _ bool) bool { return idx%2 == 0 },
			out:  []bool{false, true, false, true, false, true, false, true, false, true},
		},
		{
			name: "Nth(3)",
			fn:   func(idx int, _ int, _ bool) bool { return idx%3 == 0 },
			out:  []bool{false, false, true, false, false, true, false, false, true, false},
		},
		{
			name: "Nth(4)",
			fn:   func(idx int, _ int, _ bool) bool { return idx%4 == 0 },
			out:  []bool{false, false, false, true, false, false, false, true, false, false},
		},
		{
			name: "Nth(5)",
			fn:   func(idx int, _ int, _ bool) bool { return idx%5 == 0 },
			out:  []bool{false, false, false, false, true, false, false, false, false, true},
		},
		{
			name: "Nth(10)",
			fn:   func(idx int, _ int, _ bool) bool { return idx%10 == 0 },
			out:  []bool{false, false, false, false, false, false, false, false, false, true},
		},
		{
			name: "Greater compact",
			fn:   func(val int, _ bool) bool { return val > 10 },
			out:  []bool{false, false, true, false, true, false, true, true, false, false},
		},
		{
			name: "func(_ int, val int, _ bool) bool {return val == 2}",
			fn:   func(_ int, val int, _ bool) bool { return val == 2 },
			out:  []bool{false, true, false, false, false, true, false, false, false, false},
		},
		{
			name: "func() bool {return true}",
			fn:   func() bool { return true },
			out:  []bool{false, false, false, false, false, false, false, false, false, false},
		},
	}

	payload := IntegerWithNA([]int{1, 2, 39, 4, 56, 2, 45, 90, 4, 3}, nil).(*vector).payload.(Whichable)

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			result := payload.Which(data.fn)
			if !reflect.DeepEqual(result, data.out) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to out (%v)", result, data.out))
			}
		})
	}
}

func TestIntegerPayload_SupportsApplier(t *testing.T) {
	testData := []struct {
		name        string
		applier     interface{}
		isSupported bool
	}{
		{
			name:        "func(int, int, bool) (int, bool)",
			applier:     func(int, int, bool) (int, bool) { return 0, true },
			isSupported: true,
		},
		{
			name:        "func(int, bool) (int, bool)",
			applier:     func(int, bool) (int, bool) { return 0, true },
			isSupported: true,
		},
		{
			name:        "func(int, float64, bool) bool",
			applier:     func(int, int, bool) bool { return true },
			isSupported: false,
		},
	}

	payload := IntegerWithNA([]int{1}, nil).(*vector).payload.(Appliable)
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if payload.SupportsApplier(data.applier) != data.isSupported {
				t.Error("Applier's support is incorrect.")
			}
		})
	}
}

func TestIntegerPayload_Apply(t *testing.T) {
	testData := []struct {
		name        string
		applier     interface{}
		dataIn      []int
		naIn        []bool
		dataOut     []int
		naOut       []bool
		isNAPayload bool
	}{
		{
			name: "regular",
			applier: func(_ int, val int, na bool) (int, bool) {
				return val * 2, na
			},
			dataIn:      []int{1, 9, 3, 5, 7},
			naIn:        []bool{false, true, false, true, false},
			dataOut:     []int{2, 0, 6, 0, 14},
			naOut:       []bool{false, true, false, true, false},
			isNAPayload: false,
		},
		{
			name: "regular compact",
			applier: func(val int, na bool) (int, bool) {
				return val * 2, na
			},
			dataIn:      []int{1, 9, 3, 5, 7},
			naIn:        []bool{false, true, false, true, false},
			dataOut:     []int{2, 0, 6, 0, 14},
			naOut:       []bool{false, true, false, true, false},
			isNAPayload: false,
		},
		{
			name: "manipulate na",
			applier: func(idx int, val int, na bool) (int, bool) {
				newNA := na
				if idx == 5 {
					newNA = true
				}
				return val, newNA
			},
			dataIn:      []int{1, 2, 3, 4, 5},
			naIn:        []bool{false, false, true, false, false},
			dataOut:     []int{1, 2, 0, 4, 0},
			naOut:       []bool{false, false, true, false, true},
			isNAPayload: false,
		},
		{
			name:        "incorrect applier",
			applier:     func(int, int, bool) bool { return true },
			dataIn:      []int{1, 9, 3, 5, 7},
			naIn:        []bool{false, true, false, true, false},
			dataOut:     []int{0, 0, 0, 0, 0},
			naOut:       []bool{true, true, true, true, true},
			isNAPayload: true,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := IntegerWithNA(data.dataIn, data.naIn).(*vector).payload.(Appliable).Apply(data.applier)

			if !data.isNAPayload {
				payloadOut := payload.(*integerPayload)
				if !reflect.DeepEqual(data.dataOut, payloadOut.data) {
					t.Error(fmt.Sprintf("Output data (%v) does not match expected (%v)",
						data.dataOut, payloadOut.data))
				}
				if !reflect.DeepEqual(data.naOut, payloadOut.na) {
					t.Error(fmt.Sprintf("Output NA (%v) does not match expected (%v)",
						data.naOut, payloadOut.na))
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

func TestIntegerPayload_SupportsSummarizer(t *testing.T) {
	testData := []struct {
		name        string
		summarizer  interface{}
		isSupported bool
	}{
		{
			name:        "valid",
			summarizer:  func(int, int, int, bool) (int, bool) { return 0, true },
			isSupported: true,
		},
		{
			name:        "invalid",
			summarizer:  func(int, int, bool) bool { return true },
			isSupported: false,
		},
	}

	payload := IntegerWithNA([]int{}, nil).(*vector).payload.(Summarizable)
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if payload.SupportsSummarizer(data.summarizer) != data.isSupported {
				t.Error("Summarizer's support is incorrect.")
			}
		})
	}
}

func TestIntegerPayload_Summarize(t *testing.T) {
	summarizer := func(idx int, prev int, cur int, na bool) (int, bool) {
		if idx == 1 {
			return cur, false
		}

		return prev + cur, na
	}

	testData := []struct {
		name        string
		summarizer  interface{}
		dataIn      []int
		naIn        []bool
		dataOut     []int
		naOut       []bool
		isNAPayload bool
	}{
		{
			name:        "true",
			summarizer:  summarizer,
			dataIn:      []int{1, 2, 1, 6, 5},
			naIn:        []bool{false, false, false, false, false},
			dataOut:     []int{15},
			naOut:       []bool{false},
			isNAPayload: false,
		},
		{
			name:        "NA",
			summarizer:  summarizer,
			dataIn:      []int{1, 2, 1, 6, 5},
			naIn:        []bool{false, false, false, false, true},
			isNAPayload: true,
		},
		{
			name:        "incorrect applier",
			summarizer:  func(int, int, bool) bool { return true },
			dataIn:      []int{1, 2, 1, 6, 5},
			naIn:        []bool{false, true, false, true, false},
			isNAPayload: true,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := IntegerWithNA(data.dataIn, data.naIn).(*vector).payload.(Summarizable).Summarize(data.summarizer)

			if !data.isNAPayload {
				payloadOut := payload.(*integerPayload)
				if !reflect.DeepEqual(data.dataOut, payloadOut.data) {
					t.Error(fmt.Sprintf("Output data (%v) does not match expected (%v)",
						data.dataOut, payloadOut.data))
				}
				if !reflect.DeepEqual(data.naOut, payloadOut.na) {
					t.Error(fmt.Sprintf("Output NA (%v) does not match expected (%v)",
						data.naOut, payloadOut.na))
				}
			} else {
				naPayload, ok := payload.(*naPayload)
				if ok {
					if naPayload.length != 1 {
						t.Error("Incorrect length of NA payload (not 1)")
					}
				} else {
					t.Error("Payload is not NA")
				}
			}
		})
	}
}

func TestIntegerPayload_Append(t *testing.T) {
	payload := IntegerPayload([]int{1, 2, 3}, nil)

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
			name:    "na",
			vec:     NA(2),
			outData: []int{1, 2, 3, 0, 0},
			outNA:   []bool{false, false, false, true, true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			outPayload := payload.Append(data.vec.Payload()).(*integerPayload)

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

func TestIntegerPayload_Adjust(t *testing.T) {
	payload5 := IntegerPayload([]int{1, 2, 3, 4, 5}, nil).(*integerPayload)
	payload3 := IntegerPayload([]int{1, 2, 3}, []bool{false, false, true}).(*integerPayload)

	testData := []struct {
		name       string
		inPayload  *integerPayload
		size       int
		outPaylout *integerPayload
	}{
		{
			inPayload:  payload5,
			name:       "same",
			size:       5,
			outPaylout: IntegerPayload([]int{1, 2, 3, 4, 5}, nil).(*integerPayload),
		},
		{
			inPayload:  payload5,
			name:       "lesser",
			size:       3,
			outPaylout: IntegerPayload([]int{1, 2, 3}, nil).(*integerPayload),
		},
		{
			inPayload: payload3,
			name:      "bigger",
			size:      10,
			outPaylout: IntegerPayload([]int{1, 2, 0, 1, 2, 0, 1, 2, 0, 1},
				[]bool{false, false, true, false, false, true, false, false, true, false}).(*integerPayload),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			outPayload := data.inPayload.Adjust(data.size).(*integerPayload)

			if !reflect.DeepEqual(outPayload.data, data.outPaylout.data) {
				t.Error(fmt.Sprintf("Output data (%v) does not match expected (%v)",
					outPayload.data, data.outPaylout.data))
			}
			if !reflect.DeepEqual(outPayload.na, data.outPaylout.na) {
				t.Error(fmt.Sprintf("Output NA (%v) does not match expected (%v)",
					outPayload.na, data.outPaylout.na))
			}
		})
	}
}

func TestIntegerPayload_Find(t *testing.T) {
	payload := IntegerPayload([]int{1, 2, 1, 4, 0}, nil).(*integerPayload)

	testData := []struct {
		name   string
		needle interface{}
		pos    int
	}{
		{"existent", 4, 4},
		{"float64", 2.0, 2},
		{"float64", 2.1, 0},
		{"non-existent", -10, 0},
		{"incorrect type", "true", 0},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			pos := payload.Find(data.needle)

			if pos != data.pos {
				t.Error(fmt.Sprintf("Position (%v) does not match expected (%v)",
					pos, data.pos))
			}
		})
	}
}

func TestIntegerPayload_FindAll(t *testing.T) {
	payload := IntegerPayload([]int{1, 2, 1, 4, 0}, nil).(*integerPayload)

	testData := []struct {
		name   string
		needle interface{}
		pos    []int
	}{
		{"existent", 1, []int{1, 3}},
		{"float", 1.0, []int{1, 3}},
		{"float", 1.2, []int{}},
		{"non-existent", -10, []int{}},
		{"incorrect type", false, []int{}},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			pos := payload.FindAll(data.needle)

			if !reflect.DeepEqual(pos, data.pos) {
				t.Error(fmt.Sprintf("Positions (%v) does not match expected (%v)",
					pos, data.pos))
			}
		})
	}
}

func TestIntegerPayload_Eq(t *testing.T) {
	payload := IntegerPayload([]int{2, 0, 2, 2, 1}, []bool{false, false, true, false, false}).(*integerPayload)

	testData := []struct {
		eq  interface{}
		cmp []bool
	}{
		{2, []bool{true, false, false, true, false}},
		{2.0, []bool{true, false, false, true, false}},
		{2 + 0i, []bool{true, false, false, true, false}},
		{complex64(2 + 0i), []bool{true, false, false, true, false}},
		{"2", []bool{false, false, false, false, false}},

		{float32(1), []bool{false, false, false, false, true}},
		{int64(1), []bool{false, false, false, false, true}},
		{int32(1), []bool{false, false, false, false, true}},
		{uint64(1), []bool{false, false, false, false, true}},
		{uint32(1), []bool{false, false, false, false, true}},

		{true, []bool{false, false, false, false, false}},
		{2 + 1i, []bool{false, false, false, false, false}},
		{2.5 + 0i, []bool{false, false, false, false, false}},
		{2.5, []bool{false, false, false, false, false}},
		{complex64(2 + 1i), []bool{false, false, false, false, false}},
		{complex64(2.5 + 0i), []bool{false, false, false, false, false}},
		{"three", []bool{false, false, false, false, false}},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			cmp := payload.Eq(data.eq)

			if !reflect.DeepEqual(cmp, data.cmp) {
				t.Error(fmt.Sprintf("Comparator results (%v) do not match expected (%v)",
					cmp, data.cmp))
			}
		})
	}
}

func TestIntegerPayload_Neq(t *testing.T) {
	payload := IntegerPayload([]int{2, 0, 2, 2, 1}, []bool{false, false, true, false, false}).(*integerPayload)

	testData := []struct {
		eq  interface{}
		cmp []bool
	}{
		{2, []bool{false, true, true, false, true}},
		{2.0, []bool{false, true, true, false, true}},
		{2 + 0i, []bool{false, true, true, false, true}},
		{complex64(2 + 0i), []bool{false, true, true, false, true}},
		{"2", []bool{true, true, true, true, true}},

		{float32(1), []bool{true, true, true, true, false}},
		{int64(1), []bool{true, true, true, true, false}},
		{int32(1), []bool{true, true, true, true, false}},
		{uint64(1), []bool{true, true, true, true, false}},
		{uint32(1), []bool{true, true, true, true, false}},

		{true, []bool{true, true, true, true, true}},
		{2 + 1i, []bool{true, true, true, true, true}},
		{2.5 + 0i, []bool{true, true, true, true, true}},
		{2.5, []bool{true, true, true, true, true}},
		{complex64(2 + 1i), []bool{true, true, true, true, true}},
		{complex64(2.5 + 0i), []bool{true, true, true, true, true}},
		{"three", []bool{true, true, true, true, true}},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			cmp := payload.Neq(data.eq)

			if !reflect.DeepEqual(cmp, data.cmp) {
				t.Error(fmt.Sprintf("Comparator results (%v) do not match expected (%v)",
					cmp, data.cmp))
			}
		})
	}
}

func TestIntegerPayload_Gt(t *testing.T) {
	payload := IntegerPayload([]int{2, 0, 2, 2, 1}, []bool{false, false, true, false, false}).(*integerPayload)

	testData := []struct {
		val interface{}
		cmp []bool
	}{
		{1, []bool{true, false, false, true, false}},
		{true, []bool{false, false, false, false, false}},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			cmp := payload.Gt(data.val)

			if !reflect.DeepEqual(cmp, data.cmp) {
				t.Error(fmt.Sprintf("Comparator results (%v) do not match expected (%v)",
					cmp, data.cmp))
			}
		})
	}
}

func TestIntegerPayload_Lt(t *testing.T) {
	payload := IntegerPayload([]int{2, 0, 2, 2, 1}, []bool{false, false, true, false, false}).(*integerPayload)

	testData := []struct {
		val interface{}
		cmp []bool
	}{
		{1, []bool{false, true, false, false, false}},
		{true, []bool{false, false, false, false, false}},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			cmp := payload.Lt(data.val)

			if !reflect.DeepEqual(cmp, data.cmp) {
				t.Error(fmt.Sprintf("Comparator results (%v) do not match expected (%v)",
					cmp, data.cmp))
			}
		})
	}
}

func TestIntegerPayload_Gte(t *testing.T) {
	payload := IntegerPayload([]int{2, 0, 2, 2, 1}, []bool{false, false, true, false, false}).(*integerPayload)

	testData := []struct {
		val interface{}
		cmp []bool
	}{
		{1, []bool{true, false, false, true, true}},
		{true, []bool{false, false, false, false, false}},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			cmp := payload.Gte(data.val)

			if !reflect.DeepEqual(cmp, data.cmp) {
				t.Error(fmt.Sprintf("Comparator results (%v) do not match expected (%v)",
					cmp, data.cmp))
			}
		})
	}
}

func TestIntegerPayload_Lte(t *testing.T) {
	payload := IntegerPayload([]int{2, 0, 2, 2, 1}, []bool{false, false, true, false, false}).(*integerPayload)

	testData := []struct {
		val interface{}
		cmp []bool
	}{
		{1, []bool{false, true, false, false, true}},
		{true, []bool{false, false, false, false, false}},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			cmp := payload.Lte(data.val)

			if !reflect.DeepEqual(cmp, data.cmp) {
				t.Error(fmt.Sprintf("Comparator results (%v) do not match expected (%v)",
					cmp, data.cmp))
			}
		})
	}
}

func TestIntegerPayload_Groups(t *testing.T) {
	testData := []struct {
		name    string
		payload Payload
		groups  [][]int
		values  []interface{}
	}{
		{
			name:    "normal",
			payload: IntegerPayload([]int{-20, 10, 4, -20, 7, -20, 10, -20, 4, 10}, nil),
			groups:  [][]int{{1, 4, 6, 8}, {2, 7, 10}, {3, 9}, {5}},
			values:  []interface{}{-20, 10, 4, 7},
		},
		{
			name: "with NA",
			payload: IntegerPayload([]int{-20, 10, 4, -20, 10, -20, 10, -20, 4, 7},
				[]bool{false, false, false, false, false, false, true, true, false, false}),
			groups: [][]int{{1, 4, 6}, {2, 5}, {3, 9}, {10}, {7, 8}},
			values: []interface{}{-20, 10, 4, 7, nil},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			groups, values := data.payload.(*integerPayload).Groups()

			if !reflect.DeepEqual(groups, data.groups) {
				t.Error(fmt.Sprintf("Groups (%v) do not match expected (%v)",
					groups, data.groups))
			}

			if !reflect.DeepEqual(values, data.values) {
				t.Error(fmt.Sprintf("Groups (%v) do not match expected (%v)",
					values, data.values))
			}
		})
	}
}

func TestIntegerPayload_IsUnique(t *testing.T) {
	testData := []struct {
		name     string
		payload  Payload
		booleans []bool
	}{
		{
			name:     "without NA",
			payload:  IntegerPayload([]int{1, 2, 1, 3, 2, 3, 2}, nil),
			booleans: []bool{true, true, false, true, false, false, false},
		},
		{
			name:     "with NA",
			payload:  IntegerPayload([]int{1, 2, 1, 3, 2, 3, 2}, []bool{false, true, true, false, false, false, false}),
			booleans: []bool{true, true, false, true, true, false, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			booleans := data.payload.(*integerPayload).IsUnique()

			if !reflect.DeepEqual(booleans, data.booleans) {
				t.Error(fmt.Sprintf("Result of IsUnique() (%v) do not match expected (%v)",
					booleans, data.booleans))
			}
		})
	}
}

func TestIntegerPayload_Coalesce(t *testing.T) {
	testData := []struct {
		name         string
		coalescer    Payload
		coalescendum Payload
		outData      []int
		outNA        []bool
	}{
		{
			name:         "empty",
			coalescer:    IntegerPayload(nil, nil),
			coalescendum: IntegerPayload([]int{}, nil),
			outData:      []int{},
			outNA:        []bool{},
		},
		{
			name:         "same type",
			coalescer:    IntegerPayload([]int{1, 0, 0, 0, 5}, []bool{false, true, true, true, false}),
			coalescendum: IntegerPayload([]int{11, 12, 0, 14, 15}, []bool{false, false, true, false, false}),
			outData:      []int{1, 12, 0, 14, 5},
			outNA:        []bool{false, false, true, false, false},
		},
		{
			name:         "same type + different size",
			coalescer:    IntegerPayload([]int{1, 0, 0, 0, 5}, []bool{false, true, true, true, false}),
			coalescendum: IntegerPayload([]int{0, 11}, []bool{true, false}),
			outData:      []int{1, 11, 0, 11, 5},
			outNA:        []bool{false, false, true, false, false},
		},
		{
			name:         "different type",
			coalescer:    IntegerPayload([]int{1, 0, 0, 0, 5}, []bool{false, true, true, true, false}),
			coalescendum: FloatPayload([]float64{0, 10, 0, 112, 0}, []bool{false, false, true, false, false}),
			outData:      []int{1, 10, 0, 112, 5},
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

			if !reflect.DeepEqual(payload.na, data.outNA) {
				t.Error(fmt.Sprintf("NA (%v) do not match expected (%v)",
					payload.na, data.outNA))
			}
		})
	}
}
