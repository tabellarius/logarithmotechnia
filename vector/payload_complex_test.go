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

func TestComplex(t *testing.T) {
	emptyNA := []bool{false, false, false, false, false}

	testData := []struct {
		name    string
		data    []complex128
		na      []bool
		outData []complex128
		isEmpty bool
	}{
		{
			name:    "normal + na false",
			data:    []complex128{1.1 + 0i, 2.2 + 2.2i, 3.3 + 3.3i, 4.4 + 4.4i, 5.5 - 5.5i},
			na:      []bool{false, false, false, false, false},
			outData: []complex128{1.1 + 0i, 2.2 + 2.2i, 3.3 + 3.3i, 4.4 + 4.4i, 5.5 - 5.5i},
			isEmpty: false,
		},
		{
			name:    "normal + empty na",
			data:    []complex128{1.1 + 0i, 2.2 + 2.2i, 3.3 + 3.3i, 4.4 + 4.4i, 5.5 - 5.5i},
			na:      []bool{},
			outData: []complex128{1.1 + 0i, 2.2 + 2.2i, 3.3 + 3.3i, 4.4 + 4.4i, 5.5 - 5.5i},
			isEmpty: false,
		},
		{
			name:    "normal + nil na",
			data:    []complex128{1.1 + 0i, 2.2 + 2.2i, 3.3 + 3.3i, 4.4 + 4.4i, 5.5 - 5.5i},
			na:      nil,
			outData: []complex128{1.1 + 0i, 2.2 + 2.2i, 3.3 + 3.3i, 4.4 + 4.4i, 5.5 - 5.5i},
			isEmpty: false,
		},
		{
			name:    "normal + na mixed",
			data:    []complex128{1.1 + 0i, 2.2 + 2.2i, 3.3 + 3.3i, 4.4 + 4.4i, 5.5 - 5.5i},
			na:      []bool{false, true, true, true, false},
			outData: []complex128{1.1 + 0i, cmplx.NaN(), cmplx.NaN(), cmplx.NaN(), 5.5 - 5.5i},
			isEmpty: false,
		},
		{
			name:    "normal + incorrect sized na",
			data:    []complex128{1.1 + 0i, 2.2 + 2.2i, 3.3 + 3.3i, 4.4 + 4.4i, 5.5 - 5.5i},
			na:      []bool{false, false, false, false},
			isEmpty: true,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			v := ComplexWithNA(data.data, data.na)

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

				payload, ok := vv.payload.(*complexPayload)
				if !ok {
					t.Error("Payload is not complexPayload")
				} else {
					if !util.EqualComplexArrays(payload.data, data.outData) {
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

func TestComplexPayload_Type(t *testing.T) {
	vec := ComplexWithNA([]complex128{}, nil)
	if vec.Type() != "complex" {
		t.Error("Type is incorrect.")
	}
}

func TestComplexPayload_Len(t *testing.T) {
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

func TestComplexPayload_ByIndices(t *testing.T) {
	vec := ComplexWithNA([]complex128{1, 2, 3, 4, 5}, []bool{false, false, false, false, true})
	testData := []struct {
		name    string
		indices []int
		out     []complex128
		outNA   []bool
	}{
		{
			name:    "all",
			indices: []int{1, 2, 3, 4, 5},
			out:     []complex128{1, 2, 3, 4, cmplx.NaN()},
			outNA:   []bool{false, false, false, false, true},
		},
		{
			name:    "all reverse",
			indices: []int{5, 4, 3, 2, 1},
			out:     []complex128{cmplx.NaN(), 4, 3, 2, 1},
			outNA:   []bool{true, false, false, false, false},
		},
		{
			name:    "some",
			indices: []int{5, 1, 3},
			out:     []complex128{cmplx.NaN(), 1, 3},
			outNA:   []bool{true, false, false},
		},
		{
			name:    "with zero",
			indices: []int{5, 1, 0, 3},
			out:     []complex128{cmplx.NaN(), 1, cmplx.NaN(), 3},
			outNA:   []bool{true, false, true, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := vec.ByIndices(data.indices).(*vector).payload.(*complexPayload)
			if !util.EqualComplexArrays(payload.data, data.out) {
				t.Error(fmt.Sprintf("payload.data (%v) is not equal to data.out (%v)", payload.data, data.out))
			}
			if !reflect.DeepEqual(payload.na, data.outNA) {
				t.Error(fmt.Sprintf("payload.data (%v) is not equal to data.out (%v)", payload.data, data.out))
			}
		})
	}
}

func TestComplexPayload_SupportsWhicher(t *testing.T) {
	testData := []struct {
		name        string
		filter      interface{}
		isSupported bool
	}{
		{
			name:        "func(int, complex128, bool) bool",
			filter:      func(int, complex128, bool) bool { return true },
			isSupported: true,
		},
		{
			name:        "func(complex128, bool) bool",
			filter:      func(complex128, bool) bool { return true },
			isSupported: true,
		},
		{
			name:        "func(int, int, bool) bool",
			filter:      func(int, int, bool) bool { return true },
			isSupported: false,
		},
	}

	payload := ComplexWithNA([]complex128{1}, nil).(*vector).payload.(Whichable)
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if payload.SupportsWhicher(data.filter) != data.isSupported {
				t.Error("whicher's support is incorrect.")
			}
		})
	}
}

func TestComplexPayload_Whicher(t *testing.T) {
	testData := []struct {
		name string
		fn   interface{}
		out  []bool
	}{
		{
			name: "Odd",
			fn:   func(idx int, _ complex128, _ bool) bool { return idx%2 == 1 },
			out:  []bool{true, false, true, false, true, false, true, false, true, false},
		},
		{
			name: "Even",
			fn:   func(idx int, _ complex128, _ bool) bool { return idx%2 == 0 },
			out:  []bool{false, true, false, true, false, true, false, true, false, true},
		},
		{
			name: "Nth(3)",
			fn:   func(idx int, _ complex128, _ bool) bool { return idx%3 == 0 },
			out:  []bool{false, false, true, false, false, true, false, false, true, false},
		},
		{
			name: "compact",
			fn:   func(val complex128, _ bool) bool { return val == 1 || val == 4 },
			out:  []bool{true, false, false, true, false, false, false, false, true, false},
		},
		{
			name: "func() bool {return true}",
			fn:   func() bool { return true },
			out:  []bool{false, false, false, false, false, false, false, false, false, false},
		},
	}

	payload := ComplexWithNA([]complex128{1, 2, 39, 4, 56, 2, 45, 90, 4, 3}, nil).(*vector).payload.(Whichable)

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			result := payload.Which(data.fn)
			if !reflect.DeepEqual(result, data.out) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to out (%v)", result, data.out))
			}
		})
	}
}

func TestComplexPayload_SupportsApplier(t *testing.T) {
	testData := []struct {
		name        string
		applier     interface{}
		isSupported bool
	}{
		{
			name:        "func(int, complex128, bool) (complex128, bool)",
			applier:     func(int, complex128, bool) (complex128, bool) { return 0, true },
			isSupported: true,
		},
		{
			name:        "func(complex128, bool) (complex128, bool)",
			applier:     func(complex128, bool) (complex128, bool) { return 0, true },
			isSupported: true,
		},
		{
			name:        "func(int, complex128, bool) bool",
			applier:     func(int, complex128, bool) bool { return true },
			isSupported: false,
		},
	}

	payload := ComplexWithNA([]complex128{}, nil).(*vector).payload.(Appliable)
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if payload.SupportsApplier(data.applier) != data.isSupported {
				t.Error("Applier's support is incorrect.")
			}
		})
	}
}

func TestComplexPayload_Apply(t *testing.T) {
	testData := []struct {
		name        string
		applier     interface{}
		dataIn      []complex128
		naIn        []bool
		dataOut     []complex128
		naOut       []bool
		isNAPayload bool
	}{
		{
			name: "regular",
			applier: func(_ int, val complex128, na bool) (complex128, bool) {
				return val * 2, na
			},
			dataIn:      []complex128{1, 9, 3, 5, 7},
			naIn:        []bool{false, true, false, true, false},
			dataOut:     []complex128{2, cmplx.NaN(), 6, cmplx.NaN(), 14},
			naOut:       []bool{false, true, false, true, false},
			isNAPayload: false,
		},
		{
			name: "regular compact",
			applier: func(val complex128, na bool) (complex128, bool) {
				return val * 2, na
			},
			dataIn:      []complex128{1, 9, 3, 5, 7},
			naIn:        []bool{false, true, false, true, false},
			dataOut:     []complex128{2, cmplx.NaN(), 6, cmplx.NaN(), 14},
			naOut:       []bool{false, true, false, true, false},
			isNAPayload: false,
		},
		{
			name: "manipulate na",
			applier: func(idx int, val complex128, na bool) (complex128, bool) {
				if idx == 5 {
					return cmplx.NaN(), true
				}
				return val, na
			},
			dataIn:      []complex128{1, 2, 3, 4, 5},
			naIn:        []bool{false, false, true, false, false},
			dataOut:     []complex128{1, 2, cmplx.NaN(), 4, cmplx.NaN()},
			naOut:       []bool{false, false, true, false, true},
			isNAPayload: false,
		},
		{
			name:        "incorrect applier",
			applier:     func(int, complex128, bool) bool { return true },
			dataIn:      []complex128{1, 9, 3, 5, 7},
			naIn:        []bool{false, true, false, true, false},
			dataOut:     []complex128{cmplx.NaN(), cmplx.NaN(), cmplx.NaN(), cmplx.NaN(), cmplx.NaN()},
			naOut:       []bool{true, true, true, true, true},
			isNAPayload: true,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := ComplexWithNA(data.dataIn, data.naIn).(*vector).payload.(Appliable).Apply(data.applier)

			if !data.isNAPayload {
				payloadOut := payload.(*complexPayload)
				if !util.EqualComplexArrays(data.dataOut, payloadOut.data) {
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

func TestComplexPayload_Booleans(t *testing.T) {
	testData := []struct {
		in    []complex128
		inNA  []bool
		out   []bool
		outNA []bool
	}{
		{
			in:    []complex128{1, 3, 0, 100, 0},
			inNA:  []bool{false, false, false, false, false},
			out:   []bool{true, true, false, true, false},
			outNA: []bool{false, false, false, false, false},
		},
		{
			in:    []complex128{10, 0, 12, 14, 1110},
			inNA:  []bool{false, false, false, true, true},
			out:   []bool{true, false, true, false, false},
			outNA: []bool{false, false, false, true, true},
		},
		{
			in:    []complex128{1, 3, 0, 100, 0, -11, -10},
			inNA:  []bool{false, false, false, false, false, false, true},
			out:   []bool{true, true, false, true, false, true, false},
			outNA: []bool{false, false, false, false, false, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := ComplexWithNA(data.in, data.inNA)
			payload := vec.(*vector).payload.(*complexPayload)

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

func TestComplexPayload_Integers(t *testing.T) {
	testData := []struct {
		in    []complex128
		inNA  []bool
		out   []int
		outNA []bool
	}{
		{
			in:    []complex128{1, 3, 0, 100, 0},
			inNA:  []bool{false, false, false, false, false},
			out:   []int{1, 3, 0, 100, 0},
			outNA: []bool{false, false, false, false, false},
		},
		{
			in:    []complex128{10, 0, 12, 14, 1110},
			inNA:  []bool{false, false, false, true, true},
			out:   []int{10, 0, 12, 0, 0},
			outNA: []bool{false, false, false, true, true},
		},
		{
			in:    []complex128{1, 3, 0, 100, 0, -11, -10},
			inNA:  []bool{false, false, false, false, false, false, true},
			out:   []int{1, 3, 0, 100, 0, -11, 0},
			outNA: []bool{false, false, false, false, false, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := ComplexWithNA(data.in, data.inNA)
			payload := vec.(*vector).payload.(*complexPayload)

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

func TestComplexPayload_Interfaces(t *testing.T) {
	testData := []struct {
		in    []complex128
		inNA  []bool
		out   []interface{}
		outNA []bool
	}{
		{
			in:    []complex128{1 + 0i, 3 + 0i, 0 + 0i, 100 + 0i, 0 + 0i},
			inNA:  []bool{false, false, false, false, false},
			out:   []interface{}{1 + 0i, 3 + 0i, 0 + 0i, 100 + 0i, 0 + 0i},
			outNA: []bool{false, false, false, false, false},
		},
		{
			in:    []complex128{10 + 0i, 0 + 0i, 12 + 0i, 14 + 0i, 1110 + 0i},
			inNA:  []bool{false, false, false, true, true},
			out:   []interface{}{10 + 0i, 0 + 0i, 12 + 0i, nil, nil},
			outNA: []bool{false, false, false, true, true},
		},
		{
			in:    []complex128{1 + 0i, 3 + 0i, 0 + 0i, 100 + 0i, 0 + 0i, -11 + 0i, -10 + 0i},
			inNA:  []bool{false, false, false, false, false, false, true},
			out:   []interface{}{1 + 0i, 3 + 0i, 0 + 0i, 100 + 0i, 0 + 0i, -11 + 0i, nil},
			outNA: []bool{false, false, false, false, false, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := ComplexWithNA(data.in, data.inNA)
			payload := vec.(*vector).payload.(*complexPayload)

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

func TestComplexPayload_Floats(t *testing.T) {
	testData := []struct {
		in    []complex128
		inNA  []bool
		out   []float64
		outNA []bool
	}{
		{
			in:    []complex128{1, 3, 0, 100, 0},
			inNA:  []bool{false, false, false, false, false},
			out:   []float64{1, 3, 0, 100, 0},
			outNA: []bool{false, false, false, false, false},
		},
		{
			in:    []complex128{10, 0, 12, 14, 1110},
			inNA:  []bool{false, false, false, true, true},
			out:   []float64{10, 0, 12, math.NaN(), math.NaN()},
			outNA: []bool{false, false, false, true, true},
		},
		{
			in:    []complex128{1, 3, 0, 100, 0, -11, -10},
			inNA:  []bool{false, false, false, false, false, false, true},
			out:   []float64{1, 3, 0, 100, 0, -11, math.NaN()},
			outNA: []bool{false, false, false, false, false, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := ComplexWithNA(data.in, data.inNA)
			payload := vec.(*vector).payload.(*complexPayload)

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

func TestComplexPayload_Complexes(t *testing.T) {
	testData := []struct {
		in    []complex128
		inNA  []bool
		out   []complex128
		outNA []bool
	}{
		{
			in:    []complex128{1, 3, 0, 100, 0, cmplx.NaN()},
			inNA:  []bool{false, false, false, false, false, false},
			out:   []complex128{1 + 0i, 3 + 0i, 0 + 0i, 100 + 0i, 0 + 0i, cmplx.NaN()},
			outNA: []bool{false, false, false, false, false, false},
		},
		{
			in:    []complex128{10, 0, 12, 14, 1110},
			inNA:  []bool{false, false, false, true, true},
			out:   []complex128{10 + 0i, 0 + 0i, 12 + 0i, cmplx.NaN(), cmplx.NaN()},
			outNA: []bool{false, false, false, true, true},
		},
		{
			in:    []complex128{1, 3, 0, 100, 0, -11, -10},
			inNA:  []bool{false, false, false, false, false, false, true},
			out:   []complex128{1 + 0i, 3 + 0i, 0 + 0i, 100 + 0i, 0 + 0i, -11 + 0i, cmplx.NaN()},
			outNA: []bool{false, false, false, false, false, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := ComplexWithNA(data.in, data.inNA)
			payload := vec.(*vector).payload.(*complexPayload)

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

func TestComplexPayload_Strings(t *testing.T) {
	testData := []struct {
		in    []complex128
		inNA  []bool
		out   []string
		outNA []bool
	}{
		{
			in:    []complex128{1, 3, cmplx.NaN(), 100, 0, cmplx.Inf()},
			inNA:  []bool{false, false, false, false, false, false},
			out:   []string{"(1.000+0.000i)", "(3.000+0.000i)", "NaN", "(100.000+0.000i)", "(0.000+0.000i)", "Inf"},
			outNA: []bool{false, false, false, false, false, false},
		},
		{
			in:    []complex128{10, 0, 12, 14, 1110},
			inNA:  []bool{false, false, false, true, true},
			out:   []string{"(10.000+0.000i)", "(0.000+0.000i)", "(12.000+0.000i)", "NA", "NA"},
			outNA: []bool{false, false, false, true, true},
		},
		{
			in:   []complex128{1, 3, cmplx.NaN(), 100, 0, -11, -10},
			inNA: []bool{false, false, false, false, false, false, true},
			out: []string{"(1.000+0.000i)", "(3.000+0.000i)", "NaN", "(100.000+0.000i)", "(0.000+0.000i)",
				"(-11.000+0.000i)", "NA"},
			outNA: []bool{false, false, false, false, false, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := ComplexWithNA(data.in, data.inNA)
			payload := vec.(*vector).payload.(*complexPayload)

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

func TestComplexPayload_SupportsSummarizer(t *testing.T) {
	testData := []struct {
		name        string
		summarizer  interface{}
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

func TestComplexPayload_Summarize(t *testing.T) {
	summarizer := func(idx int, prev complex128, cur complex128, na bool) (complex128, bool) {
		if idx == 1 {
			return cur, false
		}

		return prev + cur, na
	}

	testData := []struct {
		name        string
		summarizer  interface{}
		dataIn      []complex128
		naIn        []bool
		dataOut     []complex128
		naOut       []bool
		isNAPayload bool
	}{
		{
			name:        "true",
			summarizer:  summarizer,
			dataIn:      []complex128{1 + 1i, 2 + 3.5i, 1.5 + 1.5i, 5.5 + 5i, 5 + 11i},
			naIn:        []bool{false, false, false, false, false},
			dataOut:     []complex128{15 + 22i},
			naOut:       []bool{false},
			isNAPayload: false,
		},
		{
			name:        "NA",
			summarizer:  summarizer,
			dataIn:      []complex128{1 + 1i, 2 + 3.5i, 1.5 + 1.5i, 5.5 + 5i, 5 + 11i},
			naIn:        []bool{false, false, false, false, true},
			isNAPayload: true,
		},
		{
			name:        "incorrect applier",
			summarizer:  func(int, int, bool) bool { return true },
			dataIn:      []complex128{1 + 1i, 2 + 3.5i, 1.5 + 1.5i, 5.5 + 5i, 5 + 11i},
			naIn:        []bool{false, true, false, true, false},
			isNAPayload: true,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := ComplexWithNA(data.dataIn, data.naIn).(*vector).payload.(Summarizable).Summarize(data.summarizer)

			if !data.isNAPayload {
				payloadOut := payload.(*complexPayload)
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

func TestComplexPayload_Append(t *testing.T) {
	payload := ComplexPayload([]complex128{1 + 1i, 2 + 2i, 3 + 3i}, nil)

	testData := []struct {
		name    string
		vec     Vector
		outData []complex128
		outNA   []bool
	}{
		{
			name:    "complex",
			vec:     ComplexWithNA([]complex128{4 + 4i, 5 + 5i}, []bool{true, false}),
			outData: []complex128{1 + 1i, 2 + 2i, 3 + 3i, cmplx.NaN(), 5 + 5i},
			outNA:   []bool{false, false, false, true, false},
		},
		{
			name:    "integer",
			vec:     IntegerWithNA([]int{4, 5}, []bool{true, false}),
			outData: []complex128{1 + 1i, 2 + 2i, 3 + 3i, cmplx.NaN(), 5 + 0i},
			outNA:   []bool{false, false, false, true, false},
		},
		{
			name:    "na",
			vec:     NA(2),
			outData: []complex128{1 + 1i, 2 + 2i, 3 + 3i, cmplx.NaN(), cmplx.NaN()},
			outNA:   []bool{false, false, false, true, true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			outPayload := payload.Append(data.vec.Payload()).(*complexPayload)

			if !util.EqualComplexArrays(data.outData, outPayload.data) {
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

func TestComplexPayload_Adjust(t *testing.T) {
	payload5 := ComplexPayload([]complex128{1 + 0i, 2.5 + 1i, 5.5i, 0, -2 - 1.5i}, nil).(*complexPayload)
	payload3 := ComplexPayload([]complex128{1 + 0i, 2.5 + 1i, 5.5i}, []bool{false, false, true}).(*complexPayload)

	testData := []struct {
		name       string
		inPayload  *complexPayload
		size       int
		outPaylout *complexPayload
	}{
		{
			inPayload:  payload5,
			name:       "same",
			size:       5,
			outPaylout: ComplexPayload([]complex128{1 + 0i, 2.5 + 1i, 5.5i, 0, -2 - 1.5i}, nil).(*complexPayload),
		},
		{
			inPayload:  payload5,
			name:       "lesser",
			size:       3,
			outPaylout: ComplexPayload([]complex128{1 + 0i, 2.5 + 1i, 5.5i}, nil).(*complexPayload),
		},
		{
			inPayload: payload3,
			name:      "bigger",
			size:      10,
			outPaylout: ComplexPayload([]complex128{1 + 0i, 2.5 + 1i, cmplx.NaN(), 1 + 0i, 2.5 + 1i, cmplx.NaN(),
				1 + 0i, 2.5 + 1i, cmplx.NaN(), 1 + 0i},
				[]bool{false, false, true, false, false, true, false, false, true, false}).(*complexPayload),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			outPayload := data.inPayload.Adjust(data.size).(*complexPayload)

			if !util.EqualComplexArrays(outPayload.data, data.outPaylout.data) {
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

func TestComplexPayload_PrecisionOption(t *testing.T) {
	testData := []struct {
		name              string
		payload           *complexPayload
		expectedPrecision int
	}{
		{
			name:              "precision 4",
			payload:           ComplexPayload(nil, nil, OptionPrecision(4)).(*complexPayload),
			expectedPrecision: 4,
		},
		{
			name:              "precision 5",
			payload:           ComplexPayload(nil, nil, OptionPrecision(5)).(*complexPayload),
			expectedPrecision: 5,
		},
		{
			name:              "default precision",
			payload:           ComplexPayload(nil, nil).(*complexPayload),
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

func TestComlexPayload_Find(t *testing.T) {
	payload := ComplexPayload([]complex128{1 + 0i, 2, 1 + 0i, 3 + 2i, 0 + 0i}, nil).(*complexPayload)

	testData := []struct {
		name   string
		needle interface{}
		pos    int
	}{
		{"existent", 3 + 2i, 4},
		{"float64", 2.0, 2},
		{"int", 2, 2},
		{"non-existent", 5i, 0},
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

func TestComplexPayload_FindAll(t *testing.T) {
	payload := ComplexPayload([]complex128{1 + 0i, 2, 1 + 0i, 3 + 2i, 0 + 0i}, nil).(*complexPayload)

	testData := []struct {
		name   string
		needle interface{}
		pos    []int
	}{
		{"existent", 1 + 0i, []int{1, 3}},
		{"float64", 1.0, []int{1, 3}},
		{"int", 1, []int{1, 3}},
		{"non-existent", 5i, []int{}},
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

func TestComplexPayload_Eq(t *testing.T) {
	payload := ComplexPayload([]complex128{1 + 1.5i, 0, 1 + 1.5i, 1 + 1.5i, 1},
		[]bool{false, false, true, false, false}).(*complexPayload)

	testData := []struct {
		eq  interface{}
		cmp []bool
	}{
		{1 + 1.5i, []bool{true, false, false, true, false}},
		{0 + 0i, []bool{false, true, false, false, false}},
		{1, []bool{false, false, false, false, true}},
		{complex64(1 + 1.5i), []bool{true, false, false, true, false}},
		{float64(1), []bool{false, false, false, false, true}},
		{float32(1), []bool{false, false, false, false, true}},
		{int64(1), []bool{false, false, false, false, true}},
		{int32(1), []bool{false, false, false, false, true}},
		{uint64(1), []bool{false, false, false, false, true}},
		{uint32(1), []bool{false, false, false, false, true}},
		{true, []bool{false, false, false, false, false}},
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

func TestComplexPayload_Neq(t *testing.T) {
	payload := ComplexPayload([]complex128{1 + 1.5i, 0, 1 + 1.5i, 1 + 1.5i, 1},
		[]bool{false, false, true, false, false}).(*complexPayload)

	testData := []struct {
		val interface{}
		cmp []bool
	}{
		{1 + 1.5i, []bool{false, true, true, false, true}},
		{0 + 0i, []bool{true, false, true, true, true}},
		{1, []bool{true, true, true, true, false}},
		{complex64(1 + 1.5i), []bool{false, true, true, false, true}},
		{float64(1), []bool{true, true, true, true, false}},
		{float32(1), []bool{true, true, true, true, false}},
		{int64(1), []bool{true, true, true, true, false}},
		{int32(1), []bool{true, true, true, true, false}},
		{uint64(1), []bool{true, true, true, true, false}},
		{uint32(1), []bool{true, true, true, true, false}},
		{true, []bool{true, true, true, true, true}},
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

func TestComplexPayload_Lt(t *testing.T) {
	payload := ComplexPayload([]complex128{1 + 1.5i, 0, 1 + 1.5i, 1 + 1.5i, 1},
		[]bool{false, false, true, false, false}).(*complexPayload)

	testData := []struct {
		val interface{}
		cmp []bool
	}{
		{1 + 1.5i, []bool{false, false, false, false, false}},
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

func TestComplexPayload_Gt(t *testing.T) {
	payload := ComplexPayload([]complex128{1 + 1.5i, 0, 1 + 1.5i, 1 + 1.5i, 1},
		[]bool{false, false, true, false, false}).(*complexPayload)

	testData := []struct {
		val interface{}
		cmp []bool
	}{
		{1 + 1.5i, []bool{false, false, false, false, false}},
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

func TestComplexPayload_Lte(t *testing.T) {
	payload := ComplexPayload([]complex128{1 + 1.5i, 0, 1 + 1.5i, 1 + 1.5i, 1},
		[]bool{false, false, true, false, false}).(*complexPayload)

	testData := []struct {
		val interface{}
		cmp []bool
	}{
		{1 + 1.5i, []bool{false, false, false, false, false}},
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

func TestComplexPayload_Gte(t *testing.T) {
	payload := ComplexPayload([]complex128{1 + 1.5i, 0, 1 + 1.5i, 1 + 1.5i, 1},
		[]bool{false, false, true, false, false}).(*complexPayload)

	testData := []struct {
		val interface{}
		cmp []bool
	}{
		{1 + 1.5i, []bool{false, false, false, false, false}},
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

func TestComplexPayload_Groups(t *testing.T) {
	testData := []struct {
		name    string
		payload Payload
		groups  [][]int
		values  []interface{}
	}{
		{
			name:    "normal",
			payload: ComplexPayload([]complex128{-20, 10, 4, -20, 7, -20, 10, -20, 4, 10}, nil),
			groups:  [][]int{{1, 4, 6, 8}, {2, 7, 10}, {3, 9}, {5}},
			values:  []interface{}{-20 + 0i, 10 + 0i, 4 + 0i, 7 + 0i},
		},
		{
			name: "with NA",
			payload: ComplexPayload([]complex128{-20, 10, 4, -20, 10, -20, 10, -20, 4, 7},
				[]bool{false, false, false, false, false, false, true, true, false, false}),
			groups: [][]int{{1, 4, 6}, {2, 5}, {3, 9}, {10}, {7, 8}},
			values: []interface{}{-20 + 0i, 10 + 0i, 4 + 0i, 7 + 0i, nil},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			groups, values := data.payload.(*complexPayload).Groups()

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

func TestComplexPayload_IsUnique(t *testing.T) {
	testData := []struct {
		name     string
		payload  Payload
		booleans []bool
	}{
		{
			name: "without NA",
			payload: ComplexPayload([]complex128{1, 2, 1, 3, 2, 3, 2, cmplx.NaN(), cmplx.NaN(),
				cmplx.Inf(), cmplx.Inf()}, nil),
			booleans: []bool{true, true, false, true, false, false, false, true, false, true, false},
		},
		{
			name:     "with NA",
			payload:  ComplexPayload([]complex128{1, 2, 1, 3, 2, 3, 2}, []bool{false, true, true, false, false, false, false}),
			booleans: []bool{true, true, false, true, true, false, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			booleans := data.payload.(*complexPayload).IsUnique()

			if !reflect.DeepEqual(booleans, data.booleans) {
				t.Error(fmt.Sprintf("Result of IsUnique() (%v) do not match expected (%v)",
					booleans, data.booleans))
			}
		})
	}
}

func TestComplexPayload_Coalesce(t *testing.T) {
	testData := []struct {
		name         string
		coalescer    Payload
		coalescendum Payload
		outData      []complex128
		outNA        []bool
	}{
		{
			name:         "empty",
			coalescer:    ComplexPayload(nil, nil),
			coalescendum: ComplexPayload([]complex128{}, nil),
			outData:      []complex128{},
			outNA:        []bool{},
		},
		{
			name:         "same type",
			coalescer:    ComplexPayload([]complex128{1, 0, 0, 0, 5}, []bool{false, true, true, true, false}),
			coalescendum: ComplexPayload([]complex128{11, 12, 0, 14, 15}, []bool{false, false, true, false, false}),
			outData:      []complex128{1, 12, cmplx.NaN(), 14, 5},
			outNA:        []bool{false, false, true, false, false},
		},
		{
			name:         "same type + different size",
			coalescer:    ComplexPayload([]complex128{1, 0, 0, 0, 5}, []bool{false, true, true, true, false}),
			coalescendum: ComplexPayload([]complex128{0, 11}, []bool{true, false}),
			outData:      []complex128{1, 11, cmplx.NaN(), 11, 5},
			outNA:        []bool{false, false, true, false, false},
		},
		{
			name:         "different type",
			coalescer:    ComplexPayload([]complex128{1, 0, 0, 0, 5}, []bool{false, true, true, true, false}),
			coalescendum: IntegerPayload([]int{0, 10, 0, 112, 0}, []bool{false, false, true, false, false}),
			outData:      []complex128{1, 10, cmplx.NaN(), 112, 5},
			outNA:        []bool{false, false, true, false, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.coalescer.(Coalescer).Coalesce(data.coalescendum).(*complexPayload)

			if !util.EqualComplexArrays(payload.data, data.outData) {
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
