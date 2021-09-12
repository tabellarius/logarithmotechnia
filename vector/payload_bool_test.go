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
		name    string
		data    []bool
		na      []bool
		outData []bool
		isEmpty bool
	}{
		{
			name:    "normal + na false",
			data:    []bool{true, false, true, false, true},
			na:      []bool{false, false, false, false, false},
			outData: []bool{true, false, true, false, true},
			isEmpty: false,
		},
		{
			name:    "normal + empty na",
			data:    []bool{true, false, true, false, true},
			na:      []bool{},
			outData: []bool{true, false, true, false, true},
			isEmpty: false,
		},
		{
			name:    "normal + nil na",
			data:    []bool{true, false, true, false, true},
			na:      nil,
			outData: []bool{true, false, true, false, true},
			isEmpty: false,
		},
		{
			name:    "normal + na mixed",
			data:    []bool{true, false, true, false, true},
			na:      []bool{false, true, true, true, false},
			outData: []bool{true, false, false, false, true},
			isEmpty: false,
		},
		{
			name:    "normal + incorrect sized na",
			data:    []bool{true, false, true, false, true},
			na:      []bool{false, false, false, false},
			isEmpty: true,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			v := Boolean(data.data, data.na)
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

				payload, ok := vv.payload.(*booleanPayload)
				if !ok {
					t.Error("Payload is not booleanPayload")
				} else {
					if !reflect.DeepEqual(payload.data, data.outData) {
						t.Error(fmt.Sprintf("Payload data (%v) is not equal to correct data (%v)\n",
							payload.data, data.outData))
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
			}
		})
	}
}

func TestBooleanPayload_Type(t *testing.T) {
	vec := Boolean([]bool{}, nil)
	if vec.Type() != "boolean" {
		t.Error("Type is incorrect.")
	}
}

func TestBooleanPayload_Len(t *testing.T) {
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

func TestBooleanPayload_Booleans(t *testing.T) {
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
			payload := vec.(*vector).payload.(*booleanPayload)

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

func TestBooleanPayload_Integers(t *testing.T) {
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
			payload := vec.(*vector).payload.(*booleanPayload)

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

func TestBooleanPayload_Floats(t *testing.T) {
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
			payload := vec.(*vector).payload.(*booleanPayload)

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

func TestBooleanPayload_Complexes(t *testing.T) {
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
			payload := vec.(*vector).payload.(*booleanPayload)

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

func TestBooleanPayload_Strings(t *testing.T) {
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
			payload := vec.(*vector).payload.(*booleanPayload)

			strings, na := payload.Strings()
			if !reflect.DeepEqual(strings, data.out) {
				t.Error(fmt.Sprintf("Strings (%v) are not equal to data.out (%v)\n", strings, data.out))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("NA (%v) are not equal to data.outNA (%v)\n", na, data.outNA))
			}
		})
	}
}

func TestBooleanPayload_Interfaces(t *testing.T) {
	testData := []struct {
		in    []bool
		inNA  []bool
		out   []interface{}
		outNA []bool
	}{
		{
			in:    []bool{true, true, false, true, false},
			inNA:  []bool{false, false, false, false, false},
			out:   []interface{}{true, true, false, true, false},
			outNA: []bool{false, false, false, false, false},
		},
		{
			in:    []bool{true, false, true, true, true},
			inNA:  []bool{false, false, false, true, true},
			out:   []interface{}{true, false, true, nil, nil},
			outNA: []bool{false, false, false, true, true},
		},
		{
			in:    []bool{true, true, false, true, false, true, true},
			inNA:  []bool{false, false, false, false, false, false, true},
			out:   []interface{}{true, true, false, true, false, true, nil},
			outNA: []bool{false, false, false, false, false, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := Boolean(data.in, data.inNA)
			payload := vec.(*vector).payload.(*booleanPayload)

			interfaces, na := payload.Interfaces()
			if !reflect.DeepEqual(interfaces, data.out) {
				t.Error(fmt.Sprintf("Interfaces (%v) are not equal to data.out (%v)\n", interfaces, data.out))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("NA (%v) are not equal to data.outNA (%v)\n", na, data.outNA))
			}
		})
	}
}

func TestBooleanPayload_ByIndices(t *testing.T) {
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
			out:     []bool{true, false, true, false, false},
			outNA:   []bool{false, false, false, false, true},
		},
		{
			name:    "all reverse",
			indices: []int{5, 4, 3, 2, 1},
			out:     []bool{false, false, true, false, true},
			outNA:   []bool{true, false, false, false, false},
		},
		{
			name:    "some",
			indices: []int{5, 1, 3},
			out:     []bool{false, true, true},
			outNA:   []bool{true, false, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := vec.ByIndices(data.indices).(*vector).payload.(*booleanPayload)
			if !reflect.DeepEqual(payload.data, data.out) {
				t.Error(fmt.Sprintf("payload.data (%v) is not equal to data.out (%v)", payload.data, data.out))
			}
			if !reflect.DeepEqual(payload.na, data.outNA) {
				t.Error(fmt.Sprintf("payload.data (%v) is not equal to data.out (%v)", payload.data, data.out))
			}
		})
	}
}

func TestBooleanPayload_SupportsWhicher(t *testing.T) {
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
			name:        "func(bool, bool) bool",
			filter:      func(bool, bool) bool { return true },
			isSupported: true,
		},
		{
			name:        "func(int, float64, bool) bool",
			filter:      func(int, float64, bool) bool { return true },
			isSupported: false,
		},
	}

	payload := Boolean([]bool{true}, nil).(*vector).payload.(Whichable)
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if payload.SupportsWhicher(data.filter) != data.isSupported {
				t.Error("Selector's support is incorrect.")
			}
		})
	}
}

func TestBooleanPayload_Whicher(t *testing.T) {
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
			name: "Inverse compact",
			fn:   func(val bool, _ bool) bool { return !val },
			out:  []bool{false, true, false, true, false, true, false, true, false, true},
		},
		{
			name: "func() bool {return true}",
			fn:   func() bool { return true },
			out:  []bool{false, false, false, false, false, false, false, false, false, false},
		},
	}

	payload := Boolean([]bool{true, false, true, false, true, false, true, false, true, false}, nil).(*vector).payload.(Whichable)

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			result := payload.Which(data.fn)
			if !reflect.DeepEqual(result, data.out) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to out (%v)", result, data.out))
			}
		})
	}
}

func TestBooleanPayload_SupportsApplier(t *testing.T) {
	testData := []struct {
		name        string
		applier     interface{}
		isSupported bool
	}{
		{
			name:        "func(int, bool, bool) (bool, bool)",
			applier:     func(int, bool, bool) (bool, bool) { return true, true },
			isSupported: true,
		},
		{
			name:        "func(bool, bool) (bool, bool)",
			applier:     func(bool, bool) (bool, bool) { return true, true },
			isSupported: true,
		},
		{
			name:        "func(int, float64, bool) bool",
			applier:     func(int, int, bool) bool { return true },
			isSupported: false,
		},
	}

	payload := Boolean([]bool{}, nil).(*vector).payload.(Appliable)
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if payload.SupportsApplier(data.applier) != data.isSupported {
				t.Error("Applier's support is incorrect.")
			}
		})
	}
}

func TestBooleanPayload_Apply(t *testing.T) {
	testData := []struct {
		name        string
		applier     interface{}
		dataIn      []bool
		naIn        []bool
		dataOut     []bool
		naOut       []bool
		isNAPayload bool
	}{
		{
			name: "regular",
			applier: func(idx int, val bool, na bool) (bool, bool) {
				if idx == 5 {
					return true, na
				}
				return val, na
			},
			dataIn:      []bool{true, true, true, false, false},
			naIn:        []bool{false, true, false, true, false},
			dataOut:     []bool{true, false, true, false, true},
			naOut:       []bool{false, true, false, true, false},
			isNAPayload: false,
		},
		{
			name: "manipulate na",
			applier: func(idx int, val bool, na bool) (bool, bool) {
				newNA := na
				if idx == 5 {
					newNA = true
				}
				return val, newNA
			},
			dataIn:      []bool{true, true, false, false, true},
			naIn:        []bool{false, true, false, true, false},
			dataOut:     []bool{true, false, false, false, false},
			naOut:       []bool{false, true, false, true, true},
			isNAPayload: false,
		},
		{
			name:        "regular compact",
			applier:     func(val bool, na bool) (bool, bool) { return !val, na },
			dataIn:      []bool{true, true, true, false, false},
			naIn:        []bool{false, true, false, true, false},
			dataOut:     []bool{false, false, false, false, true},
			naOut:       []bool{false, true, false, true, false},
			isNAPayload: false,
		},
		{
			name:        "incorrect applier",
			applier:     func(int, int, bool) bool { return true },
			dataIn:      []bool{true, true, false, false, true},
			naIn:        []bool{false, true, false, true, false},
			dataOut:     []bool{false, false, false, false, false},
			naOut:       []bool{true, true, true, true, true},
			isNAPayload: true,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := Boolean(data.dataIn, data.naIn).(*vector).payload.(Appliable).Apply(data.applier)

			if !data.isNAPayload {
				payloadOut := payload.(*booleanPayload)
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

func TestBooleanPayload_SupportsSummarizer(t *testing.T) {
	testData := []struct {
		name        string
		summarizer  interface{}
		isSupported bool
	}{
		{
			name:        "valid",
			summarizer:  func(int, bool, bool, bool) (bool, bool) { return true, true },
			isSupported: true,
		},
		{
			name:        "invalid",
			summarizer:  func(int, int, bool) bool { return true },
			isSupported: false,
		},
	}

	payload := Boolean([]bool{}, nil).(*vector).payload.(Summarizable)
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if payload.SupportsSummarizer(data.summarizer) != data.isSupported {
				t.Error("Summarizer's support is incorrect.")
			}
		})
	}
}

func TestBooleanPayload_Summarize(t *testing.T) {
	summarizer := func(idx int, prev bool, cur bool, na bool) (bool, bool) {
		if idx == 1 {
			return cur, false
		}

		return prev && cur, na
	}

	testData := []struct {
		name        string
		summarizer  interface{}
		dataIn      []bool
		naIn        []bool
		dataOut     []bool
		naOut       []bool
		isNAPayload bool
	}{
		{
			name:        "true",
			summarizer:  summarizer,
			dataIn:      []bool{true, true, true, true, true},
			naIn:        []bool{false, false, false, false, false},
			dataOut:     []bool{true},
			naOut:       []bool{false},
			isNAPayload: false,
		},
		{
			name:        "false",
			summarizer:  summarizer,
			dataIn:      []bool{true, true, false, true, true},
			naIn:        []bool{false, false, false, false, false},
			dataOut:     []bool{false},
			naOut:       []bool{false},
			isNAPayload: false,
		},
		{
			name:        "NA",
			summarizer:  summarizer,
			dataIn:      []bool{true, true, true, true, true},
			naIn:        []bool{false, false, false, false, true},
			isNAPayload: true,
		},
		{
			name:        "incorrect applier",
			summarizer:  func(int, int, bool) bool { return true },
			dataIn:      []bool{true, true, false, false, true},
			naIn:        []bool{false, true, false, true, false},
			isNAPayload: true,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := Boolean(data.dataIn, data.naIn).(*vector).payload.(Summarizable).Summarize(data.summarizer)

			if !data.isNAPayload {
				payloadOut := payload.(*booleanPayload)
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

func TestBooleanPayload_Append(t *testing.T) {
	payload := BooleanPayload([]bool{true, false, true}, nil)

	testData := []struct {
		name    string
		vec     Vector
		outData []bool
		outNA   []bool
	}{
		{
			name:    "boolean",
			vec:     Boolean([]bool{true, true}, []bool{true, false}),
			outData: []bool{true, false, true, false, true},
			outNA:   []bool{false, false, false, true, false},
		},
		{
			name:    "integer",
			vec:     Integer([]int{1, 1}, []bool{true, false}),
			outData: []bool{true, false, true, false, true},
			outNA:   []bool{false, false, false, true, false},
		},
		{
			name:    "na",
			vec:     NA(2),
			outData: []bool{true, false, true, false, false},
			outNA:   []bool{false, false, false, true, true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			outPayload := payload.Append(data.vec.Payload()).(*booleanPayload)

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

func TestBooleanPayload_Adjust(t *testing.T) {
	payload5 := BooleanPayload([]bool{true, false, true, false, true}, nil).(*booleanPayload)
	payload3 := BooleanPayload([]bool{true, false, true}, []bool{false, false, true}).(*booleanPayload)

	testData := []struct {
		name       string
		inPayload  *booleanPayload
		size       int
		outPaylout *booleanPayload
	}{
		{
			inPayload:  payload5,
			name:       "same",
			size:       5,
			outPaylout: BooleanPayload([]bool{true, false, true, false, true}, nil).(*booleanPayload),
		},
		{
			inPayload:  payload5,
			name:       "lesser",
			size:       3,
			outPaylout: BooleanPayload([]bool{true, false, true}, nil).(*booleanPayload),
		},
		{
			inPayload: payload3,
			name:      "bigger",
			size:      10,
			outPaylout: BooleanPayload([]bool{true, false, false, true, false, false, true, false, false, true},
				[]bool{false, false, true, false, false, true, false, false, true, false}).(*booleanPayload),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			outPayload := data.inPayload.Adjust(data.size).(*booleanPayload)

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

func TestBooleanPayload_Find(t *testing.T) {
	payload := BooleanPayload([]bool{true, true, true, true, true}, nil).(*booleanPayload)

	testData := []struct {
		name   string
		needle interface{}
		pos    int
	}{
		{"existent", true, 1},
		{"non-existent", false, 0},
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

func TestBooleanPayload_FindAll(t *testing.T) {
	payload := BooleanPayload([]bool{true, false, true, false, true}, nil).(*booleanPayload)

	testData := []struct {
		name   string
		needle interface{}
		pos    []int
	}{
		{"existent", true, []int{1, 3, 5}},
		{"incorrect type", "true", []int{}},
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

	payload = BooleanPayload([]bool{true, true, true}, nil).(*booleanPayload)
	pos := payload.FindAll(false)

	if !reflect.DeepEqual(pos, []int{}) {
		t.Error(fmt.Sprintf("Positions (%v) does not match expected (%v)",
			pos, []int{}))
	}
}

func TestBooleanPayload_Eq(t *testing.T) {
	payload := BooleanPayload([]bool{true, false, true, false, true}, nil).(*booleanPayload)

	testData := []struct {
		eq  interface{}
		cmp []bool
	}{
		{true, []bool{true, false, true, false, true}},
		{false, []bool{false, true, false, true, false}},
		{1, []bool{false, false, false, false, false}},
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

func TestBooleanPayload_Neq(t *testing.T) {
	payload := BooleanPayload([]bool{true, false, true, false, true}, nil).(*booleanPayload)

	testData := []struct {
		val interface{}
		cmp []bool
	}{
		{true, []bool{false, true, false, true, false}},
		{false, []bool{true, false, true, false, true}},
		{1, []bool{true, true, true, true, true}},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			cmp := payload.Neq(data.val)

			if !reflect.DeepEqual(cmp, data.cmp) {
				t.Error(fmt.Sprintf("Comparator results (%v) do not match expected (%v)",
					cmp, data.cmp))
			}
		})
	}
}

func TestBooleanPayload_Lt(t *testing.T) {
	payload := BooleanPayload([]bool{true, false, true, false, true}, nil).(*booleanPayload)

	testData := []struct {
		val interface{}
		cmp []bool
	}{
		{true, []bool{false, false, false, false, false}},
		{false, []bool{false, false, false, false, false}},
		{1, []bool{false, false, false, false, false}},
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

func TestBooleanPayload_Gt(t *testing.T) {
	payload := BooleanPayload([]bool{true, false, true, false, true}, nil).(*booleanPayload)

	testData := []struct {
		val interface{}
		cmp []bool
	}{
		{true, []bool{false, false, false, false, false}},
		{false, []bool{false, false, false, false, false}},
		{1, []bool{false, false, false, false, false}},
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

func TestBooleanPayload_Lte(t *testing.T) {
	payload := BooleanPayload([]bool{true, false, true, false, true}, nil).(*booleanPayload)

	testData := []struct {
		val interface{}
		cmp []bool
	}{
		{true, []bool{false, false, false, false, false}},
		{false, []bool{false, false, false, false, false}},
		{1, []bool{false, false, false, false, false}},
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

func TestBooleanPayload_Gte(t *testing.T) {
	payload := BooleanPayload([]bool{true, false, true, false, true}, nil).(*booleanPayload)

	testData := []struct {
		val interface{}
		cmp []bool
	}{
		{true, []bool{false, false, false, false, false}},
		{false, []bool{false, false, false, false, false}},
		{1, []bool{false, false, false, false, false}},
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
