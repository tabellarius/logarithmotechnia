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
		data    []interface{}
		na      []bool
		outData []interface{}
	}{
		{
			name:    "normal + false na",
			data:    []interface{}{1, 2, 3, 4, 5},
			na:      []bool{false, false, false, false, false},
			outData: []interface{}{1, 2, 3, 4, 5},
		},
		{
			name:    "normal + empty na",
			data:    []interface{}{1, 2, 3, 4, 5},
			na:      []bool{},
			outData: []interface{}{1, 2, 3, 4, 5},
		},
		{
			name:    "normal + nil na",
			data:    []interface{}{1, 2, 3, 4, 5},
			na:      nil,
			outData: []interface{}{1, 2, 3, 4, 5},
		},
		{
			name:    "normal + mixed na",
			data:    []interface{}{1, 2, 3, 4, 5},
			na:      []bool{false, true, true, true, false},
			outData: []interface{}{1, nil, nil, nil, 5},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			v := Interface(data.data, data.na).(*vector)

			length := len(data.data)
			if v.length != length {
				t.Error(fmt.Sprintf("Vector length (%d) is not equal to data length (%d)\n", v.length, length))
			}

			payload, ok := v.payload.(*interfacePayload)
			if !ok {
				t.Error("Payload is not interfacePayload")
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
	vec := Interface([]interface{}{1, 2, 3, 4, 5}, []bool{false, false, true, false})
	naPayload, ok := vec.(*vector).payload.(*naPayload)
	if !ok || naPayload.Len() > 0 {
		t.Error("Vector's payload is not empty")
	}
}

func TestInterfacePayload_Type(t *testing.T) {
	vec := Interface([]interface{}{}, nil)
	if vec.Type() != "interface" {
		t.Error("Type is incorrect.")
	}
}

func TestInterfacePayload_Len(t *testing.T) {
	testData := []struct {
		in        []interface{}
		outLength int
	}{
		{[]interface{}{1, 2, 3, 4, 5}, 5},
		{[]interface{}{1, 2, 3}, 3},
		{[]interface{}{}, 0},
		{nil, 0},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			payload := Interface(data.in, nil).(*vector).payload
			if payload.Len() != data.outLength {
				t.Error(fmt.Sprintf("Payloads's length (%d) is not equal to out (%d)",
					payload.Len(), data.outLength))
			}
		})
	}
}

func TestInterfacePayload_ByIndices(t *testing.T) {
	vec := Interface([]interface{}{1, 2, 3, 4, 5}, []bool{false, false, false, false, true})
	testData := []struct {
		name    string
		indices []int
		out     []interface{}
		outNA   []bool
	}{
		{
			name:    "all",
			indices: []int{1, 2, 3, 4, 5},
			out:     []interface{}{1, 2, 3, 4, nil},
			outNA:   []bool{false, false, false, false, true},
		},
		{
			name:    "all reverse",
			indices: []int{5, 4, 3, 2, 1},
			out:     []interface{}{nil, 4, 3, 2, 1},
			outNA:   []bool{true, false, false, false, false},
		},
		{
			name:    "some",
			indices: []int{5, 1, 3},
			out:     []interface{}{nil, 1, 3},
			outNA:   []bool{true, false, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := vec.ByIndices(data.indices).(*vector).payload.(*interfacePayload)
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
		filter      interface{}
		isSupported bool
	}{
		{
			name:        "func(int, interface{}, bool) bool",
			filter:      func(int, interface{}, bool) bool { return true },
			isSupported: true,
		},
		{
			name:        "func(interface{}, bool) bool",
			filter:      func(interface{}, bool) bool { return true },
			isSupported: true,
		},
		{
			name:        "func(int, float64, bool) bool",
			filter:      func(int, float64, bool) bool { return true },
			isSupported: false,
		},
	}

	payload := Interface([]interface{}{1}, nil).(*vector).payload.(Whichable)
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
		fn   interface{}
		out  []bool
	}{
		{
			name: "Odd",
			fn:   func(idx int, _ interface{}, _ bool) bool { return idx%2 == 1 },
			out:  []bool{true, false, true, false, true, false, true, false, true, false},
		},
		{
			name: "Even",
			fn:   func(idx int, _ interface{}, _ bool) bool { return idx%2 == 0 },
			out:  []bool{false, true, false, true, false, true, false, true, false, true},
		},
		{
			name: "Nth(3)",
			fn:   func(idx int, _ interface{}, _ bool) bool { return idx%3 == 0 },
			out:  []bool{false, false, true, false, false, true, false, false, true, false},
		},
		{
			name: "Not boolean compact",
			fn:   func(val interface{}, _ bool) bool { _, ok := val.(bool); return !ok },
			out:  []bool{false, false, true, false, false, true, false, false, true, false},
		},
		{
			name: "func() bool {return true}",
			fn:   func() bool { return true },
			out:  []bool{false, false, false, false, false, false, false, false, false, false},
		},
	}

	payload := Interface([]interface{}{true, false, 1, false, true, 2.5, true, false, "true", false}, nil).(*vector).payload.(Whichable)

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
		applier     interface{}
		isSupported bool
	}{
		{
			name:        "func(int, interface{}, bool) (bool, bool)",
			applier:     func(int, interface{}, bool) (interface{}, bool) { return 1, true },
			isSupported: true,
		},
		{
			name:        "func(interface{}, bool) (bool, bool)",
			applier:     func(interface{}, bool) (interface{}, bool) { return 1, true },
			isSupported: true,
		},
		{
			name:        "func(int, float64, bool) bool",
			applier:     func(int, int, bool) bool { return true },
			isSupported: false,
		},
	}

	payload := Interface([]interface{}{}, nil).(*vector).payload.(Appliable)
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
		applier     interface{}
		dataIn      []interface{}
		naIn        []bool
		dataOut     []interface{}
		naOut       []bool
		isNAPayload bool
	}{
		{
			name: "regular",
			applier: func(idx int, val interface{}, na bool) (interface{}, bool) {
				if idx == 5 {
					return 5, na
				}
				return val, na
			},
			dataIn:      []interface{}{true, true, true, false, false},
			naIn:        []bool{false, true, false, true, false},
			dataOut:     []interface{}{true, nil, true, nil, 5},
			naOut:       []bool{false, true, false, true, false},
			isNAPayload: false,
		},
		{
			name: "regular compact",
			applier: func(val interface{}, na bool) (interface{}, bool) {
				if val == false {
					return 0, na
				}

				return val, na
			},
			dataIn:      []interface{}{true, true, true, false, false},
			naIn:        []bool{false, true, false, true, false},
			dataOut:     []interface{}{true, nil, true, nil, 0},
			naOut:       []bool{false, true, false, true, false},
			isNAPayload: false,
		},
		{
			name: "manipulate na",
			applier: func(idx int, val interface{}, na bool) (interface{}, bool) {
				newNA := na
				if idx == 5 {
					newNA = true
				}
				return val, newNA
			},
			dataIn:      []interface{}{true, true, false, false, true},
			naIn:        []bool{false, true, false, true, false},
			dataOut:     []interface{}{true, nil, false, nil, nil},
			naOut:       []bool{false, true, false, true, true},
			isNAPayload: false,
		},
		{
			name:        "incorrect applier",
			applier:     func(int, int, bool) bool { return true },
			dataIn:      []interface{}{true, true, false, false, true},
			naIn:        []bool{false, true, false, true, false},
			dataOut:     []interface{}{nil, nil, nil, nil, nil},
			naOut:       []bool{true, true, true, true, true},
			isNAPayload: true,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := Interface(data.dataIn, data.naIn).(*vector).payload.(Appliable).Apply(data.applier)

			if !data.isNAPayload {
				payloadOut := payload.(*interfacePayload)
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
	convertor := func(idx int, val interface{}, na bool) (int, bool) {
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
		dataIn    []interface{}
		naIn      []bool
		convertor func(idx int, val interface{}, na bool) (int, bool)
		dataOut   []int
		naOut     []bool
	}{
		{
			name:      "regular",
			dataIn:    []interface{}{1, 2.5, "three", 4 + 3i, 5, 0},
			naIn:      []bool{false, false, false, false, true, false},
			convertor: convertor,
			dataOut:   []int{1, 2, 0, 0, 0, 0},
			naOut:     []bool{false, false, true, true, true, false},
		},
		{
			name:      "without convertor",
			dataIn:    []interface{}{1, 2.5, "three", 4 + 3i, 5, 0},
			naIn:      []bool{false, false, false, false, true, false},
			convertor: nil,
			dataOut:   []int{0, 0, 0, 0, 0, 0},
			naOut:     []bool{true, true, true, true, true, true},
		},
		{
			name:      "empty",
			dataIn:    []interface{}{},
			naIn:      []bool{},
			convertor: convertor,
			dataOut:   []int{},
			naOut:     []bool{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			vec := Interface(data.dataIn, data.naIn,
				OptionInterfaceConvertors(&InterfaceConvertors{Intabler: data.convertor}))
			payload := vec.(*vector).payload.(*interfacePayload)

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
	convertor := func(idx int, val interface{}, na bool) (float64, bool) {
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
		dataIn    []interface{}
		naIn      []bool
		convertor func(idx int, val interface{}, na bool) (float64, bool)
		dataOut   []float64
		naOut     []bool
	}{
		{
			name:      "regular",
			dataIn:    []interface{}{1, 2.5, "three", 4 + 3i, 5, 0},
			naIn:      []bool{false, false, false, false, true, false},
			convertor: convertor,
			dataOut:   []float64{1, 2.5, math.NaN(), math.NaN(), math.NaN(), 0},
			naOut:     []bool{false, false, true, true, true, false},
		},
		{
			name:      "without convertor",
			dataIn:    []interface{}{1, 2.5, "three", 4 + 3i, 5, 0},
			naIn:      []bool{false, false, false, false, true, false},
			convertor: nil,
			dataOut:   []float64{math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN()},
			naOut:     []bool{true, true, true, true, true, true},
		},
		{
			name:      "empty",
			dataIn:    []interface{}{},
			naIn:      []bool{},
			convertor: convertor,
			dataOut:   []float64{},
			naOut:     []bool{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			vec := Interface(data.dataIn, data.naIn,
				OptionInterfaceConvertors(&InterfaceConvertors{Floatabler: data.convertor}))
			payload := vec.(*vector).payload.(*interfacePayload)

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
	convertor := func(idx int, val interface{}, na bool) (complex128, bool) {
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
		dataIn    []interface{}
		naIn      []bool
		convertor func(idx int, val interface{}, na bool) (complex128, bool)
		dataOut   []complex128
		naOut     []bool
	}{
		{
			name:      "regular",
			dataIn:    []interface{}{1, 2.5, "three", 4 + 3i, 5, 0},
			naIn:      []bool{false, false, false, false, true, false},
			convertor: convertor,
			dataOut:   []complex128{1, 2.5, cmplx.NaN(), 4 + 3i, cmplx.NaN(), 0},
			naOut:     []bool{false, false, true, false, true, false},
		},
		{
			name:      "without convertor",
			dataIn:    []interface{}{1, 2.5, "three", 4 + 3i, 5, 0},
			naIn:      []bool{false, false, false, false, true, false},
			convertor: nil,
			dataOut:   []complex128{cmplx.NaN(), cmplx.NaN(), cmplx.NaN(), cmplx.NaN(), cmplx.NaN(), cmplx.NaN()},
			naOut:     []bool{true, true, true, true, true, true},
		},
		{
			name:      "empty",
			dataIn:    []interface{}{},
			naIn:      []bool{},
			convertor: convertor,
			dataOut:   []complex128{},
			naOut:     []bool{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			vec := Interface(data.dataIn, data.naIn,
				OptionInterfaceConvertors(&InterfaceConvertors{Complexabler: data.convertor}))
			payload := vec.(*vector).payload.(*interfacePayload)

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
	convertor := func(idx int, val interface{}, na bool) (bool, bool) {
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
		dataIn    []interface{}
		naIn      []bool
		convertor func(idx int, val interface{}, na bool) (bool, bool)
		dataOut   []bool
		naOut     []bool
	}{
		{
			name:      "regular",
			dataIn:    []interface{}{1, -2, "three", true, 5, 0},
			naIn:      []bool{false, false, false, false, true, false},
			convertor: convertor,
			dataOut:   []bool{true, false, false, true, false, false},
			naOut:     []bool{false, false, true, false, true, false},
		},
		{
			name:      "without convertor",
			dataIn:    []interface{}{1, -2, "three", 4 + 3i, 5, 0},
			naIn:      []bool{false, false, false, false, true, false},
			convertor: nil,
			dataOut:   []bool{false, false, false, false, false, false},
			naOut:     []bool{true, true, true, true, true, true},
		},
		{
			name:      "empty",
			dataIn:    []interface{}{},
			naIn:      []bool{},
			convertor: convertor,
			dataOut:   []bool{},
			naOut:     []bool{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			vec := Interface(data.dataIn, data.naIn,
				OptionInterfaceConvertors(&InterfaceConvertors{Boolabler: data.convertor}))
			payload := vec.(*vector).payload.(*interfacePayload)

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
	convertor := func(idx int, val interface{}, na bool) (string, bool) {
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
		dataIn    []interface{}
		naIn      []bool
		convertor func(idx int, val interface{}, na bool) (string, bool)
		dataOut   []string
		naOut     []bool
	}{
		{
			name:      "regular",
			dataIn:    []interface{}{1, 2.5, "three", 4 + 3i, 5, 0},
			naIn:      []bool{false, false, false, false, true, false},
			convertor: convertor,
			dataOut:   []string{"1", "", "three", "", "", "0"},
			naOut:     []bool{false, true, false, true, true, false},
		},
		{
			name:      "without convertor",
			dataIn:    []interface{}{1, 2.5, "three", 4 + 3i, 5, 0},
			naIn:      []bool{false, false, false, false, true, false},
			convertor: nil,
			dataOut:   []string{"", "", "", "", "", ""},
			naOut:     []bool{true, true, true, true, true, true},
		},
		{
			name:      "empty",
			dataIn:    []interface{}{},
			naIn:      []bool{},
			convertor: convertor,
			dataOut:   []string{},
			naOut:     []bool{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			vec := Interface(data.dataIn, data.naIn,
				OptionInterfaceConvertors(&InterfaceConvertors{Stringabler: data.convertor}))
			payload := vec.(*vector).payload.(*interfacePayload)

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
	convertor := func(idx int, val interface{}, na bool) (time.Time, bool) {
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
		dataIn    []interface{}
		naIn      []bool
		convertor func(idx int, val interface{}, na bool) (time.Time, bool)
		dataOut   []time.Time
		naOut     []bool
	}{
		{
			name:      "regular",
			dataIn:    []interface{}{1625270725, "three", 1625270725},
			naIn:      []bool{false, false, true},
			convertor: convertor,
			dataOut:   toTimeData([]string{"2021-07-03T03:05:25+03:00", "0001-01-01T00:00:00Z", "0001-01-01T00:00:00Z"}),
			naOut:     []bool{false, true, true},
		},
		{
			name:      "without convertor",
			dataIn:    []interface{}{1625270725, "three", 1625270725},
			naIn:      []bool{false, false, true},
			convertor: nil,
			dataOut:   toTimeData([]string{"0001-01-01T00:00:00Z", "0001-01-01T00:00:00Z", "0001-01-01T00:00:00Z"}),
			naOut:     []bool{true, true, true},
		},
		{
			name:      "empty",
			dataIn:    []interface{}{},
			naIn:      []bool{},
			convertor: convertor,
			dataOut:   []time.Time{},
			naOut:     []bool{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			vec := Interface(data.dataIn, data.naIn,
				OptionInterfaceConvertors(&InterfaceConvertors{Timeabler: data.convertor}))
			payload := vec.(*vector).payload.(*interfacePayload)

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
	convertor := func(idx int, val interface{}, na bool) (int, bool) {
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
		dataIn    []interface{}
		naIn      []bool
		convertor func(idx int, val interface{}, na bool) (int, bool)
		dataOut   []interface{}
		naOut     []bool
	}{
		{
			name:      "regular",
			dataIn:    []interface{}{1, 2.5, "three", 4 + 3i, 5, 0},
			naIn:      []bool{false, false, false, false, true, false},
			convertor: convertor,
			dataOut:   []interface{}{1, 2.5, "three", 4 + 3i, nil, 0},
			naOut:     []bool{false, false, false, false, true, false},
		},
		{
			name:      "empty",
			dataIn:    []interface{}{},
			naIn:      []bool{},
			convertor: convertor,
			dataOut:   []interface{}{},
			naOut:     []bool{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			vec := Interface(data.dataIn, data.naIn,
				OptionInterfaceConvertors(&InterfaceConvertors{Intabler: data.convertor}))
			payload := vec.(*vector).payload.(*interfacePayload)

			interfaces, na := payload.Interfaces()
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
		summarizer  interface{}
		isSupported bool
	}{
		{
			name:        "valid",
			summarizer:  func(int, interface{}, interface{}, bool) (interface{}, bool) { return 0, true },
			isSupported: true,
		},
		{
			name:        "invalid",
			summarizer:  func(int, int, bool) bool { return true },
			isSupported: false,
		},
	}

	payload := Interface([]interface{}{}, nil).(*vector).payload.(Summarizable)
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if payload.SupportsSummarizer(data.summarizer) != data.isSupported {
				t.Error("Summarizer's support is incorrect.")
			}
		})
	}
}

func TestInterfacePayload_Summarize(t *testing.T) {
	summarizer := func(idx int, prev interface{}, cur interface{}, na bool) (interface{}, bool) {
		if idx == 1 {
			return cur, false
		}

		return interface{}(idx), na
	}

	testData := []struct {
		name        string
		summarizer  interface{}
		dataIn      []interface{}
		naIn        []bool
		dataOut     []interface{}
		naOut       []bool
		isNAPayload bool
	}{
		{
			name:        "true",
			summarizer:  summarizer,
			dataIn:      []interface{}{1, 2, 1, 6, 5},
			naIn:        []bool{false, false, false, false, false},
			dataOut:     []interface{}{5},
			naOut:       []bool{false},
			isNAPayload: false,
		},
		{
			name:        "NA",
			summarizer:  summarizer,
			dataIn:      []interface{}{1, 2, 1, 6, 5},
			naIn:        []bool{false, false, false, false, true},
			isNAPayload: true,
		},
		{
			name:        "incorrect applier",
			summarizer:  func(int, int, bool) bool { return true },
			dataIn:      []interface{}{1, 2, 1, 6, 5},
			naIn:        []bool{false, true, false, true, false},
			isNAPayload: true,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := Interface(data.dataIn, data.naIn).(*vector).payload.(Summarizable).Summarize(data.summarizer)

			if !data.isNAPayload {
				payloadOut := payload.(*interfacePayload)
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

func TestInterfacePayload_Append(t *testing.T) {
	payload := InterfacePayload([]interface{}{1, 2, 3}, nil)

	testData := []struct {
		name    string
		vec     Vector
		outData []interface{}
		outNA   []bool
	}{
		{
			name:    "boolean",
			vec:     Boolean([]bool{true, true}, []bool{true, false}),
			outData: []interface{}{1, 2, 3, nil, true},
			outNA:   []bool{false, false, false, true, false},
		},
		{
			name:    "integer",
			vec:     Integer([]int{4, 5}, []bool{true, false}),
			outData: []interface{}{1, 2, 3, nil, 5},
			outNA:   []bool{false, false, false, true, false},
		},
		{
			name:    "na",
			vec:     NA(2),
			outData: []interface{}{1, 2, 3, nil, nil},
			outNA:   []bool{false, false, false, true, true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			outPayload := payload.Append(data.vec).(*interfacePayload)

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
