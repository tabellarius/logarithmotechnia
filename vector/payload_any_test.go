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

func TestInterface(t *testing.T) {
	testInterfaceEmpty(t)

	emptyNA := []bool{false, false, false, false, false}

	testData := []struct {
		name    string
		data    []any
		na      []bool
		outData []any
	}{
		{
			name:    "normal + false na",
			data:    []any{1, 2, 3, 4, 5},
			na:      []bool{false, false, false, false, false},
			outData: []any{1, 2, 3, 4, 5},
		},
		{
			name:    "normal + empty na",
			data:    []any{1, 2, 3, 4, 5},
			na:      []bool{},
			outData: []any{1, 2, 3, 4, 5},
		},
		{
			name:    "normal + nil na",
			data:    []any{1, 2, 3, 4, 5},
			na:      nil,
			outData: []any{1, 2, 3, 4, 5},
		},
		{
			name:    "normal + mixed na",
			data:    []any{1, 2, 3, 4, 5},
			na:      []bool{false, true, true, true, false},
			outData: []any{1, nil, nil, nil, 5},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			v := AnyWithNA(data.data, data.na).(*vector)

			length := len(data.data)
			if v.length != length {
				t.Error(fmt.Sprintf("Vector length (%d) is not equal to data length (%d)\n", v.length, length))
			}

			payload, ok := v.payload.(*anyPayload)
			if !ok {
				t.Error("Payload is not anyPayload")
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
		})
	}
}

func testInterfaceEmpty(t *testing.T) {
	vec := AnyWithNA([]any{1, 2, 3, 4, 5}, []bool{false, false, true, false})
	naPayload, ok := vec.(*vector).payload.(*naPayload)
	if !ok || naPayload.Len() > 0 {
		t.Error("Vector's payload is not empty")
	}
}

func TestInterfacePayload_Type(t *testing.T) {
	vec := AnyWithNA([]any{}, nil)
	if vec.Type() != "any" {
		t.Error("Type is incorrect.")
	}
}

func TestInterfacePayload_Len(t *testing.T) {
	testData := []struct {
		in        []any
		outLength int
	}{
		{[]any{1, 2, 3, 4, 5}, 5},
		{[]any{1, 2, 3}, 3},
		{[]any{}, 0},
		{nil, 0},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			payload := AnyWithNA(data.in, nil).(*vector).payload
			if payload.Len() != data.outLength {
				t.Error(fmt.Sprintf("Payloads's length (%d) is not equal to out (%d)",
					payload.Len(), data.outLength))
			}
		})
	}
}

func TestInterfacePayload_ByIndices(t *testing.T) {
	vec := AnyWithNA([]any{1, 2, 3, 4, 5}, []bool{false, false, false, false, true})
	testData := []struct {
		name    string
		indices []int
		out     []any
		outNA   []bool
	}{
		{
			name:    "all",
			indices: []int{1, 2, 3, 4, 5},
			out:     []any{1, 2, 3, 4, nil},
			outNA:   []bool{false, false, false, false, true},
		},
		{
			name:    "all reverse",
			indices: []int{5, 4, 3, 2, 1},
			out:     []any{nil, 4, 3, 2, 1},
			outNA:   []bool{true, false, false, false, false},
		},
		{
			name:    "some",
			indices: []int{5, 1, 3},
			out:     []any{nil, 1, 3},
			outNA:   []bool{true, false, false},
		},
		{
			name:    "some",
			indices: []int{5, 1, 3, 0},
			out:     []any{nil, 1, 3, nil},
			outNA:   []bool{true, false, false, true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := vec.ByIndices(data.indices).(*vector).payload.(*anyPayload)
			if !reflect.DeepEqual(payload.data, data.out) {
				t.Error(fmt.Sprintf("payload.data (%v) is not equal to data.out (%v)", payload.data, data.out))
			}
			if !reflect.DeepEqual(payload.na, data.outNA) {
				t.Error(fmt.Sprintf("payload.data (%v) is not equal to data.out (%v)", payload.data, data.out))
			}
		})
	}
}

func TestInterfacePayload_SupportsWhicher(t *testing.T) {
	testData := []struct {
		name        string
		filter      any
		isSupported bool
	}{
		{
			name:        "func(int, any, bool) bool",
			filter:      func(int, any, bool) bool { return true },
			isSupported: true,
		},
		{
			name:        "func(any, bool) bool",
			filter:      func(any, bool) bool { return true },
			isSupported: true,
		},
		{
			name:        "func(any) bool",
			filter:      func(any) bool { return true },
			isSupported: true,
		},
		{
			name:        "func(int, float64, bool) bool",
			filter:      func(int, float64, bool) bool { return true },
			isSupported: false,
		},
	}

	payload := AnyWithNA([]any{1}, nil).(*vector).payload.(Whichable)
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if payload.SupportsWhicher(data.filter) != data.isSupported {
				t.Error("Selector's support is incorrect.")
			}
		})
	}
}

func TestInterfacePayload_Which(t *testing.T) {
	testData := []struct {
		name string
		fn   any
		out  []bool
	}{
		{
			name: "Odd",
			fn:   func(idx int, _ any, _ bool) bool { return idx%2 == 1 },
			out:  []bool{true, false, true, false, true, false, true, false, true, false},
		},
		{
			name: "Even",
			fn:   func(idx int, _ any, _ bool) bool { return idx%2 == 0 },
			out:  []bool{false, true, false, true, false, true, false, true, false, true},
		},
		{
			name: "Nth(3)",
			fn:   func(idx int, _ any, _ bool) bool { return idx%3 == 0 },
			out:  []bool{false, false, true, false, false, true, false, false, true, false},
		},
		{
			name: "Not boolean compact",
			fn:   func(val any, _ bool) bool { _, ok := val.(bool); return !ok },
			out:  []bool{false, false, true, false, false, true, false, false, true, false},
		},
		{
			name: "func() bool {return true}",
			fn:   func() bool { return true },
			out:  []bool{false, false, false, false, false, false, false, false, false, false},
		},
	}

	payload := AnyWithNA([]any{true, false, 1, false, true, 2.5, true, false, "true", false}, nil).(*vector).payload.(Whichable)

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			result := payload.Which(data.fn)
			if !reflect.DeepEqual(result, data.out) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to out (%v)", result, data.out))
			}
		})
	}
}

func TestInterfacePayload_SupportsApplier(t *testing.T) {
	testData := []struct {
		name        string
		applier     any
		isSupported bool
	}{
		{
			name:        "func(int, any, bool) (bool, bool)",
			applier:     func(int, any, bool) (any, bool) { return 1, true },
			isSupported: true,
		},
		{
			name:        "func(any, bool) (bool, bool)",
			applier:     func(any, bool) (any, bool) { return 1, true },
			isSupported: true,
		},
		{
			name:        "func(int, float64, bool) bool",
			applier:     func(int, int, bool) bool { return true },
			isSupported: false,
		},
	}

	payload := AnyWithNA([]any{}, nil).(*vector).payload.(Appliable)
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if payload.SupportsApplier(data.applier) != data.isSupported {
				t.Error("Applier's support is incorrect.")
			}
		})
	}
}

func TestInterfacePayload_Apply(t *testing.T) {
	testData := []struct {
		name        string
		applier     any
		dataIn      []any
		naIn        []bool
		dataOut     []any
		naOut       []bool
		isNAPayload bool
	}{
		{
			name: "regular",
			applier: func(idx int, val any, na bool) (any, bool) {
				if idx == 5 {
					return 5, na
				}
				return val, na
			},
			dataIn:      []any{true, true, true, false, false},
			naIn:        []bool{false, true, false, true, false},
			dataOut:     []any{true, nil, true, nil, 5},
			naOut:       []bool{false, true, false, true, false},
			isNAPayload: false,
		},
		{
			name: "regular compact",
			applier: func(val any, na bool) (any, bool) {
				if val == false {
					return 0, na
				}

				return val, na
			},
			dataIn:      []any{true, true, true, false, false},
			naIn:        []bool{false, true, false, true, false},
			dataOut:     []any{true, nil, true, nil, 0},
			naOut:       []bool{false, true, false, true, false},
			isNAPayload: false,
		},
		{
			name: "regular brief",
			applier: func(val any) any {
				if val == false {
					return 0
				}

				return val
			},
			dataIn:      []any{true, true, true, false, false},
			naIn:        []bool{false, true, false, true, false},
			dataOut:     []any{true, nil, true, nil, 0},
			naOut:       []bool{false, true, false, true, false},
			isNAPayload: false,
		},
		{
			name: "manipulate na",
			applier: func(idx int, val any, na bool) (any, bool) {
				newNA := na
				if idx == 5 {
					newNA = true
				}
				return val, newNA
			},
			dataIn:      []any{true, true, false, false, true},
			naIn:        []bool{false, true, false, true, false},
			dataOut:     []any{true, nil, false, nil, nil},
			naOut:       []bool{false, true, false, true, true},
			isNAPayload: false,
		},
		{
			name:        "incorrect applier",
			applier:     func(int, int, bool) bool { return true },
			dataIn:      []any{true, true, false, false, true},
			naIn:        []bool{false, true, false, true, false},
			dataOut:     []any{nil, nil, nil, nil, nil},
			naOut:       []bool{true, true, true, true, true},
			isNAPayload: true,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := AnyWithNA(data.dataIn, data.naIn).(*vector).payload.(Appliable).Apply(data.applier)

			if !data.isNAPayload {
				payloadOut := payload.(*anyPayload)
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

func TestInterfacePayload_Integers(t *testing.T) {
	convertor := func(idx int, val any, na bool) (int, bool) {
		if na {
			return 0, true
		}

		switch v := val.(type) {
		case float64:
			return int(v), false
		case int:
			return v, false
		default:
			return 0, true
		}
	}

	testData := []struct {
		name      string
		dataIn    []any
		naIn      []bool
		convertor func(idx int, val any, na bool) (int, bool)
		dataOut   []int
		naOut     []bool
	}{
		{
			name:      "regular",
			dataIn:    []any{1, 2.5, "three", 4 + 3i, 5, 0},
			naIn:      []bool{false, false, false, false, true, false},
			convertor: convertor,
			dataOut:   []int{1, 2, 0, 0, 0, 0},
			naOut:     []bool{false, false, true, true, true, false},
		},
		{
			name:      "without convertor",
			dataIn:    []any{1, 2.5, "three", 4 + 3i, 5, 0},
			naIn:      []bool{false, false, false, false, true, false},
			convertor: nil,
			dataOut:   []int{0, 0, 0, 0, 0, 0},
			naOut:     []bool{true, true, true, true, true, true},
		},
		{
			name:      "empty",
			dataIn:    []any{},
			naIn:      []bool{},
			convertor: convertor,
			dataOut:   []int{},
			naOut:     []bool{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			vec := AnyWithNA(data.dataIn, data.naIn,
				OptionInterfaceConvertors(&AnyConvertors{Intabler: data.convertor}))
			payload := vec.(*vector).payload.(*anyPayload)

			integers, na := payload.Integers()
			if !reflect.DeepEqual(integers, data.dataOut) {
				t.Error(fmt.Sprintf("Result data (%v) is not equal to expected (%v)", integers, data.dataOut))
			}
			if !reflect.DeepEqual(na, data.naOut) {
				t.Error(fmt.Sprintf("Result na (%v) is not equal to expected (%v)", na, data.naOut))
			}
		})
	}
}

func TestInterfacePayload_Floats(t *testing.T) {
	convertor := func(idx int, val any, na bool) (float64, bool) {
		if na {
			return math.NaN(), true
		}

		switch v := val.(type) {
		case float64:
			return v, false
		case int:
			return float64(v), false
		default:
			return math.NaN(), true
		}
	}

	testData := []struct {
		name      string
		dataIn    []any
		naIn      []bool
		convertor func(idx int, val any, na bool) (float64, bool)
		dataOut   []float64
		naOut     []bool
	}{
		{
			name:      "regular",
			dataIn:    []any{1, 2.5, "three", 4 + 3i, 5, 0},
			naIn:      []bool{false, false, false, false, true, false},
			convertor: convertor,
			dataOut:   []float64{1, 2.5, math.NaN(), math.NaN(), math.NaN(), 0},
			naOut:     []bool{false, false, true, true, true, false},
		},
		{
			name:      "without convertor",
			dataIn:    []any{1, 2.5, "three", 4 + 3i, 5, 0},
			naIn:      []bool{false, false, false, false, true, false},
			convertor: nil,
			dataOut:   []float64{math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN()},
			naOut:     []bool{true, true, true, true, true, true},
		},
		{
			name:      "empty",
			dataIn:    []any{},
			naIn:      []bool{},
			convertor: convertor,
			dataOut:   []float64{},
			naOut:     []bool{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			vec := AnyWithNA(data.dataIn, data.naIn,
				OptionInterfaceConvertors(&AnyConvertors{Floatabler: data.convertor}))
			payload := vec.(*vector).payload.(*anyPayload)

			floats, na := payload.Floats()
			if !util.EqualFloatArrays(floats, data.dataOut) {
				t.Error(fmt.Sprintf("Result data (%v) is not equal to expected (%v)", floats, data.dataOut))
			}
			if !reflect.DeepEqual(na, data.naOut) {
				t.Error(fmt.Sprintf("Result na (%v) is not equal to expected (%v)", na, data.naOut))
			}
		})
	}
}

func TestInterfacePayload_Complexes(t *testing.T) {
	convertor := func(idx int, val any, na bool) (complex128, bool) {
		if na {
			return cmplx.NaN(), true
		}

		switch v := val.(type) {
		case complex128:
			return v, false
		case float64:
			return complex(v, 0), false
		case int:
			return complex(float64(v), 0), false
		default:
			return cmplx.NaN(), true
		}
	}

	testData := []struct {
		name      string
		dataIn    []any
		naIn      []bool
		convertor func(idx int, val any, na bool) (complex128, bool)
		dataOut   []complex128
		naOut     []bool
	}{
		{
			name:      "regular",
			dataIn:    []any{1, 2.5, "three", 4 + 3i, 5, 0},
			naIn:      []bool{false, false, false, false, true, false},
			convertor: convertor,
			dataOut:   []complex128{1, 2.5, cmplx.NaN(), 4 + 3i, cmplx.NaN(), 0},
			naOut:     []bool{false, false, true, false, true, false},
		},
		{
			name:      "without convertor",
			dataIn:    []any{1, 2.5, "three", 4 + 3i, 5, 0},
			naIn:      []bool{false, false, false, false, true, false},
			convertor: nil,
			dataOut:   []complex128{cmplx.NaN(), cmplx.NaN(), cmplx.NaN(), cmplx.NaN(), cmplx.NaN(), cmplx.NaN()},
			naOut:     []bool{true, true, true, true, true, true},
		},
		{
			name:      "empty",
			dataIn:    []any{},
			naIn:      []bool{},
			convertor: convertor,
			dataOut:   []complex128{},
			naOut:     []bool{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			vec := AnyWithNA(data.dataIn, data.naIn,
				OptionInterfaceConvertors(&AnyConvertors{Complexabler: data.convertor}))
			payload := vec.(*vector).payload.(*anyPayload)

			complexes, na := payload.Complexes()
			if !util.EqualComplexArrays(complexes, data.dataOut) {
				t.Error(fmt.Sprintf("Result data (%v) is not equal to expected (%v)", complexes, data.dataOut))
			}
			if !reflect.DeepEqual(na, data.naOut) {
				t.Error(fmt.Sprintf("Result na (%v) is not equal to expected (%v)", na, data.naOut))
			}
		})
	}
}

func TestInterfacePayload_Booleans(t *testing.T) {
	convertor := func(idx int, val any, na bool) (bool, bool) {
		if na {
			return false, true
		}

		switch v := val.(type) {
		case bool:
			return v, false
		case int:
			return v > 0, false
		default:
			return false, true
		}
	}

	testData := []struct {
		name      string
		dataIn    []any
		naIn      []bool
		convertor func(idx int, val any, na bool) (bool, bool)
		dataOut   []bool
		naOut     []bool
	}{
		{
			name:      "regular",
			dataIn:    []any{1, -2, "three", true, 5, 0},
			naIn:      []bool{false, false, false, false, true, false},
			convertor: convertor,
			dataOut:   []bool{true, false, false, true, false, false},
			naOut:     []bool{false, false, true, false, true, false},
		},
		{
			name:      "without convertor",
			dataIn:    []any{1, -2, "three", 4 + 3i, 5, 0},
			naIn:      []bool{false, false, false, false, true, false},
			convertor: nil,
			dataOut:   []bool{false, false, false, false, false, false},
			naOut:     []bool{true, true, true, true, true, true},
		},
		{
			name:      "empty",
			dataIn:    []any{},
			naIn:      []bool{},
			convertor: convertor,
			dataOut:   []bool{},
			naOut:     []bool{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			vec := AnyWithNA(data.dataIn, data.naIn,
				OptionInterfaceConvertors(&AnyConvertors{Boolabler: data.convertor}))
			payload := vec.(*vector).payload.(*anyPayload)

			bools, na := payload.Booleans()
			if !reflect.DeepEqual(bools, data.dataOut) {
				t.Error(fmt.Sprintf("Result data (%v) is not equal to expected (%v)", bools, data.dataOut))
			}
			if !reflect.DeepEqual(na, data.naOut) {
				t.Error(fmt.Sprintf("Result na (%v) is not equal to expected (%v)", na, data.naOut))
			}
		})
	}
}

func TestInterfacePayload_Strings(t *testing.T) {
	convertor := func(idx int, val any, na bool) (string, bool) {
		if na {
			return "", true
		}

		switch v := val.(type) {
		case string:
			return v, false
		case int:
			return strconv.Itoa(v), false
		default:
			return "", true
		}
	}

	testData := []struct {
		name      string
		dataIn    []any
		naIn      []bool
		convertor func(idx int, val any, na bool) (string, bool)
		dataOut   []string
		naOut     []bool
	}{
		{
			name:      "regular",
			dataIn:    []any{1, 2.5, "three", 4 + 3i, 5, 0},
			naIn:      []bool{false, false, false, false, true, false},
			convertor: convertor,
			dataOut:   []string{"1", "", "three", "", "", "0"},
			naOut:     []bool{false, true, false, true, true, false},
		},
		{
			name:      "without convertor",
			dataIn:    []any{1, 2.5, "three", 4 + 3i, 5, 0},
			naIn:      []bool{false, false, false, false, true, false},
			convertor: nil,
			dataOut:   []string{"", "", "", "", "", ""},
			naOut:     []bool{true, true, true, true, true, true},
		},
		{
			name:      "empty",
			dataIn:    []any{},
			naIn:      []bool{},
			convertor: convertor,
			dataOut:   []string{},
			naOut:     []bool{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			vec := AnyWithNA(data.dataIn, data.naIn,
				OptionInterfaceConvertors(&AnyConvertors{Stringabler: data.convertor}))
			payload := vec.(*vector).payload.(*anyPayload)

			strings, na := payload.Strings()
			if !reflect.DeepEqual(strings, data.dataOut) {
				t.Error(fmt.Sprintf("Result data (%v) is not equal to expected (%v)", strings, data.dataOut))
			}
			if !reflect.DeepEqual(na, data.naOut) {
				t.Error(fmt.Sprintf("Result na (%v) is not equal to expected (%v)", na, data.naOut))
			}
		})
	}
}

func TestInterfacePayload_Times(t *testing.T) {
	convertor := func(idx int, val any, na bool) (time.Time, bool) {
		if na {
			return time.Time{}, true
		}

		switch v := val.(type) {
		case int:
			return time.Unix(int64(v), 0), false
		default:
			return time.Time{}, true
		}
	}

	testData := []struct {
		name      string
		dataIn    []any
		naIn      []bool
		convertor func(idx int, val any, na bool) (time.Time, bool)
		dataOut   []time.Time
		naOut     []bool
	}{
		{
			name:      "regular",
			dataIn:    []any{1625270725, "three", 1625270725},
			naIn:      []bool{false, false, true},
			convertor: convertor,
			dataOut:   toTimeData([]string{"2021-07-03T03:05:25+03:00", "0001-01-01T00:00:00Z", "0001-01-01T00:00:00Z"}),
			naOut:     []bool{false, true, true},
		},
		{
			name:      "without convertor",
			dataIn:    []any{1625270725, "three", 1625270725},
			naIn:      []bool{false, false, true},
			convertor: nil,
			dataOut:   toTimeData([]string{"0001-01-01T00:00:00Z", "0001-01-01T00:00:00Z", "0001-01-01T00:00:00Z"}),
			naOut:     []bool{true, true, true},
		},
		{
			name:      "empty",
			dataIn:    []any{},
			naIn:      []bool{},
			convertor: convertor,
			dataOut:   []time.Time{},
			naOut:     []bool{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			vec := AnyWithNA(data.dataIn, data.naIn,
				OptionInterfaceConvertors(&AnyConvertors{Timeabler: data.convertor}))
			payload := vec.(*vector).payload.(*anyPayload)

			times, na := payload.Times()
			if !reflect.DeepEqual(times, data.dataOut) {
				t.Error(fmt.Sprintf("Result data (%v) is not equal to expected (%v)", times, data.dataOut))
			}
			if !reflect.DeepEqual(na, data.naOut) {
				t.Error(fmt.Sprintf("Result na (%v) is not equal to expected (%v)", na, data.naOut))
			}
		})
	}
}

func TestInterfacePayload_Interfaces(t *testing.T) {
	convertor := func(idx int, val any, na bool) (int, bool) {
		if na {
			return 0, true
		}

		switch v := val.(type) {
		case float64:
			return int(v), false
		case int:
			return v, false
		default:
			return 0, true
		}
	}

	testData := []struct {
		name      string
		dataIn    []any
		naIn      []bool
		convertor func(idx int, val any, na bool) (int, bool)
		dataOut   []any
		naOut     []bool
	}{
		{
			name:      "regular",
			dataIn:    []any{1, 2.5, "three", 4 + 3i, 5, 0},
			naIn:      []bool{false, false, false, false, true, false},
			convertor: convertor,
			dataOut:   []any{1, 2.5, "three", 4 + 3i, nil, 0},
			naOut:     []bool{false, false, false, false, true, false},
		},
		{
			name:      "empty",
			dataIn:    []any{},
			naIn:      []bool{},
			convertor: convertor,
			dataOut:   []any{},
			naOut:     []bool{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			vec := AnyWithNA(data.dataIn, data.naIn,
				OptionInterfaceConvertors(&AnyConvertors{Intabler: data.convertor}))
			payload := vec.(*vector).payload.(*anyPayload)

			interfaces, na := payload.Anies()
			if !reflect.DeepEqual(interfaces, data.dataOut) {
				t.Error(fmt.Sprintf("Result data (%v) is not equal to expected (%v)", interfaces, data.dataOut))
			}
			if !reflect.DeepEqual(na, data.naOut) {
				t.Error(fmt.Sprintf("Result na (%v) is not equal to expected (%v)", na, data.naOut))
			}
		})
	}
}

func TestInterfacePayload_SupportsSummarizer(t *testing.T) {
	testData := []struct {
		name        string
		summarizer  any
		isSupported bool
	}{
		{
			name:        "valid",
			summarizer:  func(int, any, any, bool) (any, bool) { return 0, true },
			isSupported: true,
		},
		{
			name:        "invalid",
			summarizer:  func(int, int, bool) bool { return true },
			isSupported: false,
		},
	}

	payload := AnyWithNA([]any{}, nil).(*vector).payload.(Summarizable)
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if payload.SupportsSummarizer(data.summarizer) != data.isSupported {
				t.Error("Summarizer's support is incorrect.")
			}
		})
	}
}

func TestInterfacePayload_Summarize(t *testing.T) {
	summarizer := func(idx int, prev any, cur any, na bool) (any, bool) {
		if idx == 1 {
			return cur, false
		}

		return any(idx), na
	}

	testData := []struct {
		name       string
		summarizer any
		dataIn     []any
		naIn       []bool
		dataOut    []any
		naOut      []bool
	}{
		{
			name:       "true",
			summarizer: summarizer,
			dataIn:     []any{1, 2, 1, 6, 5},
			naIn:       []bool{false, false, false, false, false},
			dataOut:    []any{5},
			naOut:      []bool{false},
		},
		{
			name:       "NA",
			summarizer: summarizer,
			dataIn:     []any{1, 2, 1, 6, 5},
			naIn:       []bool{false, false, false, false, true},
			dataOut:    []any{nil},
			naOut:      []bool{true},
		},
		{
			name:       "incorrect applier",
			summarizer: func(int, int, bool) bool { return true },
			dataIn:     []any{1, 2, 1, 6, 5},
			naIn:       []bool{false, true, false, true, false},
			dataOut:    []any{nil},
			naOut:      []bool{true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := AnyWithNA(data.dataIn, data.naIn).(*vector).payload.(Summarizable).Summarize(data.summarizer)

			payloadOut := payload.(*anyPayload)
			if !reflect.DeepEqual(data.dataOut, payloadOut.data) {
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

func TestInterfacePayload_Append(t *testing.T) {
	payload := InterfacePayload([]any{1, 2, 3}, nil)

	testData := []struct {
		name    string
		vec     Vector
		outData []any
		outNA   []bool
	}{
		{
			name:    "boolean",
			vec:     BooleanWithNA([]bool{true, true}, []bool{true, false}),
			outData: []any{1, 2, 3, nil, true},
			outNA:   []bool{false, false, false, true, false},
		},
		{
			name:    "integer",
			vec:     IntegerWithNA([]int{4, 5}, []bool{true, false}),
			outData: []any{1, 2, 3, nil, 5},
			outNA:   []bool{false, false, false, true, false},
		},
		{
			name:    "na",
			vec:     NA(2),
			outData: []any{1, 2, 3, nil, nil},
			outNA:   []bool{false, false, false, true, true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			outPayload := payload.Append(data.vec.Payload()).(*anyPayload)

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

func TestInterfacePayload_Adjust(t *testing.T) {
	payload5 := InterfacePayload([]any{1, 2, 3, 4, 5}, nil).(*anyPayload)
	payload3 := InterfacePayload([]any{1, 2, 3}, []bool{false, false, true}).(*anyPayload)

	testData := []struct {
		name       string
		inPayload  *anyPayload
		size       int
		outPaylout *anyPayload
	}{
		{
			inPayload:  payload5,
			name:       "same",
			size:       5,
			outPaylout: InterfacePayload([]any{1, 2, 3, 4, 5}, nil).(*anyPayload),
		},
		{
			inPayload:  payload5,
			name:       "lesser",
			size:       3,
			outPaylout: InterfacePayload([]any{1, 2, 3}, nil).(*anyPayload),
		},
		{
			inPayload: payload3,
			name:      "bigger",
			size:      10,
			outPaylout: InterfacePayload([]any{1, 2, 0, 1, 2, 0, 1, 2, 0, 1},
				[]bool{false, false, true, false, false, true, false, false, true, false}).(*anyPayload),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			outPayload := data.inPayload.Adjust(data.size).(*anyPayload)

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

func TestInterfacePayload_Pick(t *testing.T) {
	payload := InterfacePayload([]any{"a", "b", "a", "b", "c"}, []bool{false, false, true, true, false})

	testData := []struct {
		name string
		idx  int
		val  any
	}{
		{
			name: "normal 2",
			idx:  2,
			val:  any("b"),
		},
		{
			name: "normal 5",
			idx:  5,
			val:  any("c"),
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

func TestInterfacePayload_Data(t *testing.T) {
	testData := []struct {
		name    string
		payload Payload
		outData []any
	}{
		{
			name:    "empty",
			payload: InterfacePayload([]any{}, []bool{}),
			outData: []any{},
		},
		{
			name:    "non-empty",
			payload: InterfacePayload([]any{"a", "b", "a", "b", "c"}, []bool{false, false, true, true, false}),
			outData: []any{"a", "b", nil, nil, "c"},
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

func TestInterfacePayload_ApplyTo(t *testing.T) {
	srcPayload := InterfacePayload([]any{1, 2, 3, 4, 5}, []bool{false, true, false, true, false})

	testData := []struct {
		name        string
		indices     []int
		applier     any
		dataOut     []any
		naOut       []bool
		isNAPayload bool
	}{
		{
			name:    "regular",
			indices: []int{1, 2, 5},
			applier: func(idx int, val any, na bool) (any, bool) {
				iVal := 0
				if val != nil {
					iVal = val.(int)
				}
				if idx == 5 {
					iVal = iVal * 2
				}
				if na {
					iVal = 0
				}
				return iVal, false
			},
			dataOut:     []any{1, 0, 3, nil, 10},
			naOut:       []bool{false, false, false, true, false},
			isNAPayload: false,
		},
		{
			name:    "regular compact",
			indices: []int{1, 2, 5},
			applier: func(val any) any {
				if val == nil {
					return 0
				}
				return val.(int) * 3
			},
			dataOut:     []any{3, nil, 3, nil, 15},
			naOut:       []bool{false, true, false, true, false},
			isNAPayload: false,
		},
		{
			name:    "regular compact",
			indices: []int{1, 2, 5},
			applier: func(val any, na bool) (any, bool) {
				if val == nil {
					return 0, false
				}
				return val.(int) * 3, false
			},
			dataOut:     []any{3, 0, 3, nil, 15},
			naOut:       []bool{false, false, false, true, false},
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
				payloadOut := payload.(*anyPayload)
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
