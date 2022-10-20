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

func TestFloat(t *testing.T) {
	emptyNA := []bool{false, false, false, false, false}

	testData := []struct {
		name          string
		data          []float64
		na            []bool
		outData       []float64
		names         map[string]int
		expectedNames map[string]int
		isEmpty       bool
	}{
		{
			name:    "normal + na false",
			data:    []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			na:      []bool{false, false, false, false, false},
			outData: []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			names:   nil,
			isEmpty: false,
		},
		{
			name:    "normal + empty na",
			data:    []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			na:      []bool{},
			outData: []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			names:   nil,
			isEmpty: false,
		},
		{
			name:    "normal + nil na",
			data:    []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			na:      nil,
			outData: []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			names:   nil,
			isEmpty: false,
		},
		{
			name:    "normal + na mixed",
			data:    []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			na:      []bool{false, true, true, true, false},
			outData: []float64{1.1, math.NaN(), math.NaN(), math.NaN(), 5.5},
			names:   nil,
			isEmpty: false,
		},
		{
			name:    "normal + incorrect sized na",
			data:    []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			na:      []bool{false, false, false, false},
			names:   nil,
			isEmpty: true,
		},
		{
			name:          "normal + names",
			data:          []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			na:            []bool{false, false, false, false, false},
			outData:       []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			names:         map[string]int{"one": 1, "three": 3, "five": 5},
			expectedNames: map[string]int{"one": 1, "three": 3, "five": 5},
			isEmpty:       false,
		},
		{
			name:          "normal + incorrect names",
			data:          []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			na:            []bool{false, false, false, false, false},
			outData:       []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			names:         map[string]int{"zero": 0, "one": 1, "three": 3, "five": 5, "seven": 7},
			expectedNames: map[string]int{"one": 1, "three": 3, "five": 5},
			isEmpty:       false,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			var v Vector
			v = FloatWithNA(data.data, data.na)

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

				payload, ok := vv.payload.(*floatPayload)
				if !ok {
					t.Error("Payload is not floatPayload")
				} else {
					if !util.EqualFloatArrays(payload.data, data.outData) {
						t.Error(fmt.Sprintf("Payload data (%v) is not equal to correct data (%v)\n",
							payload.data, data.outData))
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

func TestFloatPayload_Type(t *testing.T) {
	vec := FloatWithNA([]float64{}, nil)
	if vec.Type() != "float" {
		t.Error("Type is incorrect.")
	}
}

func TestFloatPayload_Len(t *testing.T) {
	testData := []struct {
		in        []float64
		outLength int
	}{
		{[]float64{1, 2, 3, 4, 5}, 5},
		{[]float64{1, 2, 3}, 3},
		{[]float64{}, 0},
		{nil, 0},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			payload := FloatWithNA(data.in, nil).(*vector).payload
			if payload.Len() != data.outLength {
				t.Error(fmt.Sprintf("Payloads's length (%d) is not equal to out (%d)",
					payload.Len(), data.outLength))
			}
		})
	}
}

func TestFloatPayload_Booleans(t *testing.T) {
	testData := []struct {
		in    []float64
		inNA  []bool
		out   []bool
		outNA []bool
	}{
		{
			in:    []float64{1, 3, 0, 100, 0},
			inNA:  []bool{false, false, false, false, false},
			out:   []bool{true, true, false, true, false},
			outNA: []bool{false, false, false, false, false},
		},
		{
			in:    []float64{10, 0, 12, 14, 1110},
			inNA:  []bool{false, false, false, true, true},
			out:   []bool{true, false, true, false, false},
			outNA: []bool{false, false, false, true, true},
		},
		{
			in:    []float64{1, 3, 0, 100, 0, -11, -10},
			inNA:  []bool{false, false, false, false, false, false, true},
			out:   []bool{true, true, false, true, false, true, false},
			outNA: []bool{false, false, false, false, false, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := FloatWithNA(data.in, data.inNA)
			payload := vec.(*vector).payload.(*floatPayload)

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

func TestFloatPayload_Integers(t *testing.T) {
	testData := []struct {
		in    []float64
		inNA  []bool
		out   []int
		outNA []bool
	}{
		{
			in:    []float64{1, 3, 0, 100, 0},
			inNA:  []bool{false, false, false, false, false},
			out:   []int{1, 3, 0, 100, 0},
			outNA: []bool{false, false, false, false, false},
		},
		{
			in:    []float64{10, 0, 12, 14, 1110},
			inNA:  []bool{false, false, false, true, true},
			out:   []int{10, 0, 12, 0, 0},
			outNA: []bool{false, false, false, true, true},
		},
		{
			in:    []float64{1, 3, 0, 100, 0, -11, -10},
			inNA:  []bool{false, false, false, false, false, false, true},
			out:   []int{1, 3, 0, 100, 0, -11, 0},
			outNA: []bool{false, false, false, false, false, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := FloatWithNA(data.in, data.inNA)
			payload := vec.(*vector).payload.(*floatPayload)

			integers, na := payload.Integers()
			if !reflect.DeepEqual(integers, data.out) {
				t.Error(fmt.Sprintf("Integers (%v) are not equal to data.out (%v)\n", integers, data.out))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("NA (%v) are not equal to data.outNA (%v)\n", na, data.outNA))
			}
		})
	}
}

func TestFloatPayload_Interfaces(t *testing.T) {
	testData := []struct {
		in    []float64
		inNA  []bool
		out   []any
		outNA []bool
	}{
		{
			in:    []float64{1, 3, 0, 100, 0},
			inNA:  []bool{false, false, false, false, false},
			out:   []any{1.0, 3.0, 0.0, 100.0, 0.0},
			outNA: []bool{false, false, false, false, false},
		},
		{
			in:    []float64{10, 0, 12, 14, 1110},
			inNA:  []bool{false, false, false, true, true},
			out:   []any{10.0, 0.0, 12.0, nil, nil},
			outNA: []bool{false, false, false, true, true},
		},
		{
			in:    []float64{1, 3, 0, 100, 0, -11, -10},
			inNA:  []bool{false, false, false, false, false, false, true},
			out:   []any{1.0, 3.0, 0.0, 100.0, 0.0, -11.0, nil},
			outNA: []bool{false, false, false, false, false, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := FloatWithNA(data.in, data.inNA)
			payload := vec.(*vector).payload.(*floatPayload)

			interfaces, na := payload.Anies()
			if !reflect.DeepEqual(interfaces, data.out) {
				t.Error(fmt.Sprintf("Anies (%v) are not equal to data.out (%v)\n", interfaces, data.out))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("NA (%v) are not equal to data.outNA (%v)\n", na, data.outNA))
			}
		})
	}
}

func TestFloatPayload_Floats(t *testing.T) {
	testData := []struct {
		in    []float64
		inNA  []bool
		out   []float64
		outNA []bool
	}{
		{
			in:    []float64{1, 3, 0, 100, 0},
			inNA:  []bool{false, false, false, false, false},
			out:   []float64{1, 3, 0, 100, 0},
			outNA: []bool{false, false, false, false, false},
		},
		{
			in:    []float64{10, 0, 12, 14, 1110},
			inNA:  []bool{false, false, false, true, true},
			out:   []float64{10, 0, 12, math.NaN(), math.NaN()},
			outNA: []bool{false, false, false, true, true},
		},
		{
			in:    []float64{1, 3, 0, 100, 0, -11, -10},
			inNA:  []bool{false, false, false, false, false, false, true},
			out:   []float64{1, 3, 0, 100, 0, -11, math.NaN()},
			outNA: []bool{false, false, false, false, false, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := FloatWithNA(data.in, data.inNA)
			payload := vec.(*vector).payload.(*floatPayload)

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

func TestFloatPayload_Complexes(t *testing.T) {
	testData := []struct {
		in    []float64
		inNA  []bool
		out   []complex128
		outNA []bool
	}{
		{
			in:    []float64{1, 3, 0, 100, 0, math.NaN()},
			inNA:  []bool{false, false, false, false, false, false},
			out:   []complex128{1 + 0i, 3 + 0i, 0 + 0i, 100 + 0i, 0 + 0i, cmplx.NaN()},
			outNA: []bool{false, false, false, false, false, false},
		},
		{
			in:    []float64{10, 0, 12, 14, 1110},
			inNA:  []bool{false, false, false, true, true},
			out:   []complex128{10 + 0i, 0 + 0i, 12 + 0i, cmplx.NaN(), cmplx.NaN()},
			outNA: []bool{false, false, false, true, true},
		},
		{
			in:    []float64{1, 3, 0, 100, 0, -11, -10},
			inNA:  []bool{false, false, false, false, false, false, true},
			out:   []complex128{1 + 0i, 3 + 0i, 0 + 0i, 100 + 0i, 0 + 0i, -11 + 0i, cmplx.NaN()},
			outNA: []bool{false, false, false, false, false, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := FloatWithNA(data.in, data.inNA)
			payload := vec.(*vector).payload.(*floatPayload)

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

func TestFloatPayload_Strings(t *testing.T) {
	testData := []struct {
		in    []float64
		inNA  []bool
		out   []string
		outNA []bool
	}{
		{
			in:    []float64{1, 3, math.NaN(), 100, 0, math.Inf(+1), math.Inf(-1)},
			inNA:  []bool{false, false, false, false, false, false, false},
			out:   []string{"1.000", "3.000", "NaN", "100.000", "0.000", "+Inf", "-Inf"},
			outNA: []bool{false, false, false, false, false, false, false},
		},
		{
			in:    []float64{10, 0, 12, 14, 1110},
			inNA:  []bool{false, false, false, true, true},
			out:   []string{"10.000", "0.000", "12.000", "NA", "NA"},
			outNA: []bool{false, false, false, true, true},
		},
		{
			in:    []float64{1, 3, math.NaN(), 100, 0, -11, -10},
			inNA:  []bool{false, false, false, false, false, false, true},
			out:   []string{"1.000", "3.000", "NaN", "100.000", "0.000", "-11.000", "NA"},
			outNA: []bool{false, false, false, false, false, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := FloatWithNA(data.in, data.inNA)
			payload := vec.(*vector).payload.(*floatPayload)

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

func TestFloatPayload_ByIndices(t *testing.T) {
	vec := FloatWithNA([]float64{1, 2, 3, 4, 5}, []bool{false, false, false, false, true})
	testData := []struct {
		name    string
		indices []int
		out     []float64
		outNA   []bool
	}{
		{
			name:    "all",
			indices: []int{1, 2, 3, 4, 5},
			out:     []float64{1, 2, 3, 4, math.NaN()},
			outNA:   []bool{false, false, false, false, true},
		},
		{
			name:    "all reverse",
			indices: []int{5, 4, 3, 2, 1},
			out:     []float64{math.NaN(), 4, 3, 2, 1},
			outNA:   []bool{true, false, false, false, false},
		},
		{
			name:    "some",
			indices: []int{5, 1, 3},
			out:     []float64{math.NaN(), 1, 3},
			outNA:   []bool{true, false, false},
		},
		{
			name:    "with zero",
			indices: []int{5, 1, 0, 3},
			out:     []float64{math.NaN(), 1, math.NaN(), 3},
			outNA:   []bool{true, false, true, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := vec.ByIndices(data.indices).(*vector).payload.(*floatPayload)
			if !util.EqualFloatArrays(payload.data, data.out) {
				t.Error(fmt.Sprintf("payload.data (%v) is not equal to data.out (%v)", payload.data, data.out))
			}
			if !reflect.DeepEqual(payload.na, data.outNA) {
				t.Error(fmt.Sprintf("payload.data (%v) is not equal to data.out (%v)", payload.data, data.out))
			}
		})
	}
}

func TestFloatPayload_SupportsWhicher(t *testing.T) {
	testData := []struct {
		name        string
		filter      any
		isSupported bool
	}{
		{
			name:        "func(int, float64, bool) bool",
			filter:      func(int, float64, bool) bool { return true },
			isSupported: true,
		},
		{
			name:        "func(float64, bool) bool",
			filter:      func(float64, bool) bool { return true },
			isSupported: true,
		},
		{
			name:        "func(float64) bool",
			filter:      func(float64) bool { return true },
			isSupported: true,
		},
		{
			name:        "func(int, int, bool) bool",
			filter:      func(int, int, bool) bool { return true },
			isSupported: false,
		},
	}

	payload := FloatWithNA([]float64{1}, nil).(*vector).payload.(Whichable)
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if payload.SupportsWhicher(data.filter) != data.isSupported {
				t.Error("whicher's support is incorrect.")
			}
		})
	}
}

func TestFloatPayload_Whicher(t *testing.T) {
	testData := []struct {
		name string
		fn   any
		out  []bool
	}{
		{
			name: "Odd",
			fn:   func(idx int, _ float64, _ bool) bool { return idx%2 == 1 },
			out:  []bool{true, false, true, false, true, false, true, false, true, false},
		},
		{
			name: "Even",
			fn:   func(idx int, _ float64, _ bool) bool { return idx%2 == 0 },
			out:  []bool{false, true, false, true, false, true, false, true, false, true},
		},
		{
			name: "Nth(3)",
			fn:   func(idx int, _ float64, _ bool) bool { return idx%3 == 0 },
			out:  []bool{false, false, true, false, false, true, false, false, true, false},
		},
		{
			name: "Greater compact",
			fn:   func(val float64, _ bool) bool { return val > 10 },
			out:  []bool{false, false, true, false, true, false, true, true, false, false},
		},
		{
			name: "func() bool {return true}",
			fn:   func() bool { return true },
			out:  []bool{false, false, false, false, false, false, false, false, false, false},
		},
	}

	payload := FloatWithNA([]float64{1, 2, 39, 4, 56, 2, 45, 90, 4, 3}, nil).(*vector).payload.(Whichable)

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			result := payload.Which(data.fn)
			if !reflect.DeepEqual(result, data.out) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to out (%v)", result, data.out))
			}
		})
	}
}

func TestFloatPayload_SupportsApplier(t *testing.T) {
	testData := []struct {
		name        string
		applier     any
		isSupported bool
	}{
		{
			name:        "func(int, float64, bool) (float64, bool)",
			applier:     func(int, float64, bool) (float64, bool) { return 0, true },
			isSupported: true,
		},
		{
			name:        "func(float64, bool) (float64, bool)",
			applier:     func(float64, bool) (float64, bool) { return 0, true },
			isSupported: true,
		},
		{
			name:        "func(int, float64, bool) bool",
			applier:     func(int, float64, bool) bool { return true },
			isSupported: false,
		},
	}

	payload := FloatWithNA([]float64{1}, nil).(*vector).payload.(Appliable)
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if payload.SupportsApplier(data.applier) != data.isSupported {
				t.Error("Applier's support is incorrect.")
			}
		})
	}
}

func TestFloatPayload_Apply(t *testing.T) {
	testData := []struct {
		name        string
		applier     any
		dataIn      []float64
		naIn        []bool
		dataOut     []float64
		naOut       []bool
		isNAPayload bool
	}{
		{
			name: "regular",
			applier: func(_ int, val float64, na bool) (float64, bool) {
				return val * 2, na
			},
			dataIn:      []float64{1, 9, 3, 5, 7},
			naIn:        []bool{false, true, false, true, false},
			dataOut:     []float64{2, math.NaN(), 6, math.NaN(), 14},
			naOut:       []bool{false, true, false, true, false},
			isNAPayload: false,
		},
		{
			name: "regular compact",
			applier: func(val float64, na bool) (float64, bool) {
				return val * 2, na
			},
			dataIn:      []float64{1, 9, 3, 5, 7},
			naIn:        []bool{false, true, false, true, false},
			dataOut:     []float64{2, math.NaN(), 6, math.NaN(), 14},
			naOut:       []bool{false, true, false, true, false},
			isNAPayload: false,
		},
		{
			name: "regular brief",
			applier: func(val float64) float64 {
				return val * 2
			},
			dataIn:      []float64{1, 9, 3, 5, 7},
			naIn:        []bool{false, true, false, true, false},
			dataOut:     []float64{2, math.NaN(), 6, math.NaN(), 14},
			naOut:       []bool{false, true, false, true, false},
			isNAPayload: false,
		},
		{
			name: "manipulate na",
			applier: func(idx int, val float64, na bool) (float64, bool) {
				if idx == 5 {
					return 0, true
				}
				return val, na
			},
			dataIn:      []float64{1, 2, 3, 4, 5},
			naIn:        []bool{false, false, true, false, false},
			dataOut:     []float64{1, 2, math.NaN(), 4, math.NaN()},
			naOut:       []bool{false, false, true, false, true},
			isNAPayload: false,
		},
		{
			name:        "incorrect applier",
			applier:     func(int, float64, bool) bool { return true },
			dataIn:      []float64{1, 9, 3, 5, 7},
			naIn:        []bool{false, true, false, true, false},
			dataOut:     []float64{math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN()},
			naOut:       []bool{true, true, true, true, true},
			isNAPayload: true,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := FloatWithNA(data.dataIn, data.naIn).(*vector).payload.(Appliable).Apply(data.applier)

			if !data.isNAPayload {
				payloadOut := payload.(*floatPayload)
				if !util.EqualFloatArrays(data.dataOut, payloadOut.data) {
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

func TestFloatPayload_SupportsSummarizer(t *testing.T) {
	testData := []struct {
		name        string
		summarizer  any
		isSupported bool
	}{
		{
			name:        "valid",
			summarizer:  func(int, complex128, complex128, bool) (complex128, bool) { return 0 + 0i, true },
			isSupported: true,
		},
		{
			name:        "invalid",
			summarizer:  func(int, int, bool) bool { return true },
			isSupported: false,
		},
	}

	payload := ComplexWithNA([]complex128{}, nil).(*vector).payload.(Summarizable)
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if payload.SupportsSummarizer(data.summarizer) != data.isSupported {
				t.Error("Summarizer's support is incorrect.")
			}
		})
	}
}

func TestFloatPayload_Summarize(t *testing.T) {
	summarizer := func(idx int, prev float64, cur float64, na bool) (float64, bool) {
		if idx == 1 {
			return cur, false
		}

		return prev + cur, na
	}

	testData := []struct {
		name       string
		summarizer any
		dataIn     []float64
		naIn       []bool
		dataOut    []float64
		naOut      []bool
	}{
		{
			name:       "true",
			summarizer: summarizer,
			dataIn:     []float64{1, 2, 1.5, 5.5, 5},
			naIn:       []bool{false, false, false, false, false},
			dataOut:    []float64{15},
			naOut:      []bool{false},
		},
		{
			name:       "NA",
			summarizer: summarizer,
			dataIn:     []float64{1, 2, 1.5, 5.5, 5},
			naIn:       []bool{false, false, false, false, true},
			dataOut:    []float64{math.NaN()},
			naOut:      []bool{true},
		},
		{
			name:       "incorrect applier",
			summarizer: func(int, int, bool) bool { return true },
			dataIn:     []float64{1, 2, 1.5, 5.5, 5},
			naIn:       []bool{false, false, false, true, false},
			dataOut:    []float64{math.NaN()},
			naOut:      []bool{true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := FloatWithNA(data.dataIn, data.naIn).(*vector).payload.(Summarizable).Summarize(data.summarizer)

			payloadOut := payload.(*floatPayload)
			if !util.EqualFloatArrays(data.dataOut, payloadOut.data) {
				t.Error(fmt.Sprintf("Output data (%v) does not match expected (%v)",
					data.dataOut, payloadOut.data))
			}
			if !reflect.DeepEqual(data.naOut, payloadOut.na) {
				t.Error(fmt.Sprintf("Output NA (%v) does not match expected (%v)",
					data.naOut, payloadOut.na))
			}
		})
	}
}

func TestFloatPayload_Append(t *testing.T) {
	payload := FloatPayload([]float64{1.1, 2.2, 3.3}, nil)

	testData := []struct {
		name    string
		vec     Vector
		outData []float64
		outNA   []bool
	}{
		{
			name:    "float",
			vec:     FloatWithNA([]float64{4.4, 5.5}, []bool{true, false}),
			outData: []float64{1.1, 2.2, 3.3, math.NaN(), 5.5},
			outNA:   []bool{false, false, false, true, false},
		},
		{
			name:    "integer",
			vec:     IntegerWithNA([]int{4, 5}, []bool{true, false}),
			outData: []float64{1.1, 2.2, 3.3, math.NaN(), 5},
			outNA:   []bool{false, false, false, true, false},
		},
		{
			name:    "na",
			vec:     NA(2),
			outData: []float64{1.1, 2.2, 3.3, math.NaN(), math.NaN()},
			outNA:   []bool{false, false, false, true, true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			outPayload := payload.Append(data.vec.Payload()).(*floatPayload)

			if !util.EqualFloatArrays(data.outData, outPayload.data) {
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

func TestFloatPayload_Adjust(t *testing.T) {
	payload5 := FloatPayload([]float64{1, 2.5, 5.5, 0, -2}, nil).(*floatPayload)
	payload3 := FloatPayload([]float64{1, 2.5, 5.5}, []bool{false, false, true}).(*floatPayload)

	testData := []struct {
		name       string
		inPayload  *floatPayload
		size       int
		outPaylout *floatPayload
	}{
		{
			inPayload:  payload5,
			name:       "same",
			size:       5,
			outPaylout: FloatPayload([]float64{1 + 0i, 2.5, 5.5, 0, -2}, nil).(*floatPayload),
		},
		{
			inPayload:  payload5,
			name:       "lesser",
			size:       3,
			outPaylout: FloatPayload([]float64{1 + 0i, 2.5, 5.5}, nil).(*floatPayload),
		},
		{
			inPayload: payload3,
			name:      "bigger",
			size:      10,
			outPaylout: FloatPayload([]float64{1 + 0i, 2.5, math.NaN(), 1 + 0i, 2.5, math.NaN(),
				1 + 0i, 2.5, math.NaN(), 1 + 0i},
				[]bool{false, false, true, false, false, true, false, false, true, false}).(*floatPayload),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			outPayload := data.inPayload.Adjust(data.size).(*floatPayload)

			if !util.EqualFloatArrays(outPayload.data, data.outPaylout.data) {
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

func TestFloatPayload_PrecisionOption(t *testing.T) {
	testData := []struct {
		name              string
		payload           *floatPayload
		expectedPrecision int
	}{
		{
			name:              "precision 4",
			payload:           FloatPayload(nil, nil, OptionPrecision(4)).(*floatPayload),
			expectedPrecision: 4,
		},
		{
			name:              "precision 5",
			payload:           FloatPayload(nil, nil, OptionPrecision(5)).(*floatPayload),
			expectedPrecision: 5,
		},
		{
			name:              "default precision",
			payload:           FloatPayload(nil, nil).(*floatPayload),
			expectedPrecision: 3,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if data.payload.printer.Precision != data.expectedPrecision {
				t.Error(fmt.Sprintf("Precision (%v) does not match expected (%v)",
					data.payload.printer.Precision, data.expectedPrecision))
			}
		})
	}
}

func TestFloatPayload_Find(t *testing.T) {
	payload := FloatPayload([]float64{1, 2, 1.0, 4.1, 0}, nil).(*floatPayload)

	testData := []struct {
		name   string
		needle any
		pos    int
	}{
		{"existent", 4.1, 4},
		{"int", 2, 2},
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

func TestFloatPayload_FindAll(t *testing.T) {
	payload := FloatPayload([]float64{1, 2, 1.0, 4.1, 0}, nil).(*floatPayload)

	testData := []struct {
		name   string
		needle any
		pos    []int
	}{
		{"existent", 1, []int{1, 3}},
		{"int", 1, []int{1, 3}},
		{"non-existent", -10.5, []int{}},
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

func TestFloatPayload_Eq(t *testing.T) {
	payload := FloatPayload([]float64{1.5, 0, 1.5, 1.5, 1}, []bool{false, false, true, false, false}).(*floatPayload)

	testData := []struct {
		eq  any
		cmp []bool
	}{
		{1.5, []bool{true, false, false, true, false}},
		{1.5 + 0i, []bool{true, false, false, true, false}},
		{complex64(1.5 + 0i), []bool{true, false, false, true, false}},
		{complex128(1.5 + 0i), []bool{true, false, false, true, false}},
		{"1.5", []bool{false, false, false, false, false}},
		{float32(1), []bool{false, false, false, false, true}},
		{int64(1), []bool{false, false, false, false, true}},
		{int32(1), []bool{false, false, false, false, true}},
		{uint64(1), []bool{false, false, false, false, true}},
		{uint32(1), []bool{false, false, false, false, true}},
		{true, []bool{false, false, false, false, false}},
		{1.5 + 1i, []bool{false, false, false, false, false}},
		{complex64(1.5 + 1i), []bool{false, false, false, false, false}},
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

func TestFloatPayload_Neq(t *testing.T) {
	payload := FloatPayload([]float64{1.5, 0, 1.5, 1.5, 1}, []bool{false, false, true, false, false}).(*floatPayload)

	testData := []struct {
		eq  any
		cmp []bool
	}{
		{1.5, []bool{false, true, true, false, true}},
		{1.5 + 0i, []bool{false, true, true, false, true}},
		{complex64(1.5 + 0i), []bool{false, true, true, false, true}},

		{float32(1), []bool{true, true, true, true, false}},
		{int64(1), []bool{true, true, true, true, false}},
		{int32(1), []bool{true, true, true, true, false}},
		{uint64(1), []bool{true, true, true, true, false}},
		{uint32(1), []bool{true, true, true, true, false}},

		{true, []bool{true, true, true, true, true}},
		{1.5 + 1i, []bool{true, true, true, true, true}},
		{complex64(1.5 + 1i), []bool{true, true, true, true, true}},
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

func TestFloatPayload_Gt(t *testing.T) {
	payload := FloatPayload([]float64{1.5, 0, 1.5, 1.5, 1}, []bool{false, false, true, false, false}).(*floatPayload)

	testData := []struct {
		val any
		cmp []bool
	}{
		{1.0, []bool{true, false, false, true, false}},
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

func TestFloatPayload_Lt(t *testing.T) {
	payload := FloatPayload([]float64{1.5, 0, 1.5, 1.5, 1}, []bool{false, false, true, false, false}).(*floatPayload)

	testData := []struct {
		val any
		cmp []bool
	}{
		{1.0, []bool{false, true, false, false, false}},
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

func TestFloatPayload_Gte(t *testing.T) {
	payload := FloatPayload([]float64{1.5, 0, 1.5, 1.5, 1}, []bool{false, false, true, false, false}).(*floatPayload)

	testData := []struct {
		val any
		cmp []bool
	}{
		{1.0, []bool{true, false, false, true, true}},
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

func TestFloatPayload_Lte(t *testing.T) {
	payload := FloatPayload([]float64{1.5, 0, 1.5, 1.5, 1}, []bool{false, false, true, false, false}).(*floatPayload)

	testData := []struct {
		val any
		cmp []bool
	}{
		{1.0, []bool{false, true, false, false, true}},
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

func TestFloatPayload_Groups(t *testing.T) {
	testData := []struct {
		name    string
		payload Payload
		groups  [][]int
		values  []any
	}{
		{
			name:    "normal",
			payload: FloatPayload([]float64{-20, 10, 4, -20, 7, -20, 10, -20, 4, 10}, nil),
			groups:  [][]int{{1, 4, 6, 8}, {2, 7, 10}, {3, 9}, {5}},
			values:  []any{-20.0, 10.0, 4.0, 7.0},
		},
		{
			name: "with NA",
			payload: FloatPayload([]float64{-20, 10, 4, -20, 10, -20, 10, -20, 4, 7},
				[]bool{false, false, false, false, false, false, true, true, false, false}),
			groups: [][]int{{1, 4, 6}, {2, 5}, {3, 9}, {10}, {7, 8}},
			values: []any{-20.0, 10.0, 4.0, 7.0, nil},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			groups, values := data.payload.(*floatPayload).Groups()

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

func TestFloatPayload_IsUnique(t *testing.T) {
	testData := []struct {
		name     string
		payload  Payload
		booleans []bool
	}{
		{
			name: "without NA",
			payload: FloatPayload([]float64{1, 2, 1, 3, 2, 3, 2, math.NaN(), math.NaN(),
				math.Inf(1), math.Inf(-1), math.Inf(1), math.Inf(-1)}, nil),
			booleans: []bool{true, true, false, true, false, false, false, true, false, true, true, false, false},
		},
		{
			name:     "with NA",
			payload:  FloatPayload([]float64{1, 2, 1, 3, 2, 3, 2}, []bool{false, true, true, false, false, false, false}),
			booleans: []bool{true, true, false, true, true, false, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			booleans := data.payload.(*floatPayload).IsUnique()

			if !reflect.DeepEqual(booleans, data.booleans) {
				t.Error(fmt.Sprintf("Result of IsUnique() (%v) do not match expected (%v)",
					booleans, data.booleans))
			}
		})
	}
}

func TestFloatPayload_Coalesce(t *testing.T) {
	testData := []struct {
		name         string
		coalescer    Payload
		coalescendum Payload
		outData      []float64
		outNA        []bool
	}{
		{
			name:         "empty",
			coalescer:    FloatPayload(nil, nil),
			coalescendum: FloatPayload([]float64{}, nil),
			outData:      []float64{},
			outNA:        []bool{},
		},
		{
			name:         "same type",
			coalescer:    FloatPayload([]float64{1, 0, 0, 0, 5}, []bool{false, true, true, true, false}),
			coalescendum: FloatPayload([]float64{11, 12, 0, 14, 15}, []bool{false, false, true, false, false}),
			outData:      []float64{1, 12, math.NaN(), 14, 5},
			outNA:        []bool{false, false, true, false, false},
		},
		{
			name:         "same type + different size",
			coalescer:    FloatPayload([]float64{1, 0, 0, 0, 5}, []bool{false, true, true, true, false}),
			coalescendum: FloatPayload([]float64{0, 11}, []bool{true, false}),
			outData:      []float64{1, 11, math.NaN(), 11, 5},
			outNA:        []bool{false, false, true, false, false},
		},
		{
			name:         "different type",
			coalescer:    FloatPayload([]float64{1, 0, 0, 0, 5}, []bool{false, true, true, true, false}),
			coalescendum: IntegerPayload([]int{0, 10, 0, 112, 0}, []bool{false, false, true, false, false}),
			outData:      []float64{1, 10, math.NaN(), 112, 5},
			outNA:        []bool{false, false, true, false, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.coalescer.(Coalescer).Coalesce(data.coalescendum).(*floatPayload)

			if !util.EqualFloatArrays(payload.data, data.outData) {
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

func TestFloatPayload_Pick(t *testing.T) {
	payload := FloatPayload([]float64{1, 2, 3, 4, 5}, []bool{false, false, true, true, false})

	testData := []struct {
		name string
		idx  int
		val  any
	}{
		{
			name: "normal 2",
			idx:  2,
			val:  any(2.0),
		},
		{
			name: "normal 5",
			idx:  5,
			val:  any(5.0),
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

func TestFloatPayload_Data(t *testing.T) {
	testData := []struct {
		name    string
		payload Payload
		outData []any
	}{
		{
			name:    "empty",
			payload: FloatPayload([]float64{}, []bool{}),
			outData: []any{},
		},
		{
			name:    "non-empty",
			payload: FloatPayload([]float64{1, 2, 3, 4, 5}, []bool{false, false, true, true, false}),
			outData: []any{1.0, 2.0, nil, nil, 5.0},
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

func TestFloatPayload_ApplyTo(t *testing.T) {
	srcPayload := FloatPayload([]float64{1, 2, 3, 4, 5}, []bool{false, true, false, true, false})

	testData := []struct {
		name        string
		indices     []int
		applier     any
		dataOut     []float64
		naOut       []bool
		isNAPayload bool
	}{
		{
			name:    "regular",
			indices: []int{1, 2, 5},
			applier: func(idx int, val float64, na bool) (float64, bool) {
				if idx == 5 {
					val = val * 2
				}
				if na {
					val = 0
				}
				return val, false
			},
			dataOut:     []float64{1, 0, 3, math.NaN(), 10},
			naOut:       []bool{false, false, false, true, false},
			isNAPayload: false,
		},
		{
			name:    "regular compact",
			indices: []int{1, 2, 5},
			applier: func(val float64, na bool) (float64, bool) {
				return val * 3, false
			},
			dataOut:     []float64{3, math.NaN(), 3, math.NaN(), 15},
			naOut:       []bool{false, false, false, true, false},
			isNAPayload: false,
		},
		{
			name:    "regular brief",
			indices: []int{1, 2, 5},
			applier: func(val float64) float64 {
				return val * 3
			},
			dataOut:     []float64{3, math.NaN(), 3, math.NaN(), 15},
			naOut:       []bool{false, true, false, true, false},
			isNAPayload: false,
		},
		{
			name:        "incorrect applier",
			indices:     []int{1, 2, 5},
			applier:     func(int, int, bool) bool { return true },
			isNAPayload: true,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := srcPayload.(Appliable).ApplyTo(data.indices, data.applier)

			if !data.isNAPayload {
				payloadOut := payload.(*floatPayload)
				if !util.EqualFloatArrays(data.dataOut, payloadOut.data) {
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
