package vector

import (
	"fmt"
	"math"
	"math/cmplx"
	"reflect"
	"strconv"
	"testing"
)

func TestFactorPayload(t *testing.T) {
	testData := []struct {
		name    string
		payload Payload
		levels  []string
		data    []uint32
		isEmpty bool
	}{
		{
			name: "with NA",
			payload: FactorPayload(
				[]string{"one", "two", "one", "three", "two", "one", "three"},
				[]bool{false, false, false, false, false, true, false},
			),
			levels:  []string{"", "one", "two", "three"},
			data:    []uint32{1, 2, 1, 3, 2, 0, 3},
			isEmpty: false,
		},
		{
			name: "without NA",
			payload: FactorPayload(
				[]string{"one", "two", "one", "three", "two", "one", "three"},
				nil,
			),
			levels:  []string{"", "one", "two", "three"},
			data:    []uint32{1, 2, 1, 3, 2, 1, 3},
			isEmpty: false,
		},
		{
			name: "empty",
			payload: FactorPayload(
				[]string{},
				[]bool{},
			),
			levels:  []string{""},
			data:    []uint32{},
			isEmpty: false,
		},
		{
			name: "empty with nil NA",
			payload: FactorPayload(
				[]string{},
				nil,
			),
			levels:  []string{""},
			data:    []uint32{},
			isEmpty: false,
		},
		{
			name: "invalid NA (2)",
			payload: FactorPayload(
				[]string{"one", "two", "one", "three", "two", "one", "three"},
				[]bool{false, false},
			),
			isEmpty: true,
		},
		{
			name: "invalid NA (0)",
			payload: FactorPayload(
				[]string{"one", "two", "one", "three", "two", "one", "three"},
				[]bool{},
			),
			isEmpty: true,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if data.isEmpty {
				payload, ok := data.payload.(*naPayload)
				if ok {
					if payload.length != 0 {
						t.Error("Payload is NAPayload but not zero-length")
					}
				} else {
					t.Error("Payload is not NAPayload")
				}

				return
			}

			payload := data.payload.(*factorPayload)

			if !reflect.DeepEqual(payload.levels, data.levels) {
				t.Error(fmt.Sprintf("Payload levels (%v) are not equal to expected (%v)",
					payload.levels, data.levels))
			}

			if !reflect.DeepEqual(payload.data, data.data) {
				t.Error(fmt.Sprintf("Payload data (%v) are not equal to expected (%v)",
					payload.data, data.data))
			}
		})
	}
}

func TestFactor(t *testing.T) {
	testData := []struct {
		name     string
		inData   []string
		isFactor bool
	}{
		{
			name:     "normal",
			inData:   []string{"one", "two", "one", "three", "two", "one", "three"},
			isFactor: true,
		},
		{
			name:     "empty",
			inData:   []string{},
			isFactor: true,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			vec := Factor(data.inData)
			_, isFactor := vec.Payload().(*factorPayload)

			if isFactor != data.isFactor {
				t.Error(fmt.Sprintf("isFactor (%v) is not equal to expected (%v)",
					isFactor, data.isFactor))
			}
		})
	}
}

func TestFactorWithNA(t *testing.T) {
	testData := []struct {
		name     string
		inData   []string
		inNA     []bool
		isFactor bool
	}{
		{
			name:     "normal",
			inData:   []string{"one", "two", "one", "three", "two", "one", "three"},
			inNA:     []bool{false, true, false, true, false, true, false},
			isFactor: true,
		},
		{
			name:     "invalid",
			inData:   []string{"one", "two", "one", "three", "two", "one", "three"},
			inNA:     []bool{false, true, false},
			isFactor: false,
		},
		{
			name:     "empty",
			inData:   []string{},
			inNA:     []bool{},
			isFactor: true,
		},
		{
			name:     "empty with nil NA",
			inData:   []string{},
			inNA:     nil,
			isFactor: true,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			vec := FactorWithNA(data.inData, data.inNA)
			_, isFactor := vec.Payload().(*factorPayload)

			if isFactor != data.isFactor {
				t.Error(fmt.Sprintf("isFactor (%v) is not equal to expected (%v)",
					isFactor, data.isFactor))
			}
		})
	}
}

func TestFactorPayload_Type(t *testing.T) {
	payload := FactorPayload([]string{"one"}, []bool{false})

	if payload.Type() != "factor" {
		t.Error("Type is not 'factor'")
	}
}

func TestFactorPayload_Len(t *testing.T) {
	testData := []struct {
		name   string
		inData []string
		inNA   []bool
		length int
	}{
		{
			name:   "normal (5)",
			inData: []string{"one", "two", "one", "three", "two", "one", "three"},
			inNA:   []bool{false, true, false, true, false, true, false},
			length: 7,
		},
		{
			name:   "normal (3)",
			inData: []string{"one", "two", "one"},
			inNA:   []bool{false, true, false},
			length: 3,
		},
		{
			name:   "invalid",
			inData: []string{"one", "two", "one", "three", "two", "one", "three"},
			inNA:   []bool{false, true, false},
			length: 0,
		},
		{
			name:   "empty",
			inData: []string{},
			inNA:   []bool{},
			length: 0,
		},
		{
			name:   "empty with nil NA",
			inData: []string{},
			inNA:   nil,
			length: 0,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := FactorPayload(data.inData, data.inNA)
			length := payload.Len()

			if length != data.length {
				t.Error(fmt.Sprintf("length (%v) is not equal to expected (%v)",
					length, data.length))
			}
		})
	}
}

func TestFactorPayload_ByIndices(t *testing.T) {
	srcPayload := FactorPayload([]string{"one", "two", "two", "three", "one"}, []bool{false, false, false, false, true})

	testData := []struct {
		name    string
		indices []int
		levels  []string
		out     []uint32
		outNA   []bool
	}{
		{
			name:    "all",
			indices: []int{1, 2, 3, 4, 5},
			levels:  []string{"", "one", "two", "three"},
			out:     []uint32{1, 2, 2, 3, 0},
		},
		{
			name:    "all reverse",
			indices: []int{5, 4, 3, 2, 1},
			levels:  []string{"", "one", "two", "three"},
			out:     []uint32{0, 3, 2, 2, 1},
		},
		{
			name:    "some",
			indices: []int{5, 1, 3},
			levels:  []string{"", "one", "two", "three"},
			out:     []uint32{0, 1, 2},
		},
		{
			name:    "with zero",
			indices: []int{5, 1, 0, 3},
			levels:  []string{"", "one", "two", "three"},
			out:     []uint32{0, 1, 0, 2},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := srcPayload.ByIndices(data.indices).(*factorPayload)

			if !reflect.DeepEqual(payload.data, data.out) {
				t.Error(fmt.Sprintf("payload.data (%v) is not equal to data.out (%v)", payload.data, data.out))
			}

			if !reflect.DeepEqual(payload.levels, data.levels) {
				t.Error(fmt.Sprintf("payload.levels (%v) is not equal to data.levels (%v)", payload.data, data.out))
			}
		})
	}
}

func TestFactorPayload_Adjust(t *testing.T) {
	payload5 := FactorPayload([]string{"one", "two", "two", "three", "one"}, nil)
	payload3 := FactorPayload([]string{"one", "two", "two"}, []bool{false, false, true})

	testData := []struct {
		name      string
		inPayload Payload
		size      int
		outLevels []string
		outData   []uint32
	}{
		{
			inPayload: payload5,
			name:      "same",
			size:      5,
			outLevels: []string{"", "one", "two", "three"},
			outData:   []uint32{1, 2, 2, 3, 1},
		},
		{
			inPayload: payload5,
			name:      "lesser",
			size:      3,
			outLevels: []string{"", "one", "two", "three"},
			outData:   []uint32{1, 2, 2},
		},
		{
			inPayload: payload3,
			name:      "bigger",
			size:      10,
			outLevels: []string{"", "one", "two"},
			outData:   []uint32{1, 2, 0, 1, 2, 0, 1, 2, 0, 1},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			outPayload := data.inPayload.Adjust(data.size).(*factorPayload)

			if !reflect.DeepEqual(outPayload.data, data.outData) {
				t.Error(fmt.Sprintf("Output data (%v) does not match expected (%v)",
					outPayload.data, data.outData))
			}
			if !reflect.DeepEqual(outPayload.levels, data.outLevels) {
				t.Error(fmt.Sprintf("Output levels (%v) do not match expected (%v)",
					outPayload.levels, data.outLevels))
			}
		})
	}
}

func TestFactorPayload_Append(t *testing.T) {
	srcPayload := FactorPayload([]string{"one", "two", "two"}, []bool{false, false, true})

	var testData = []struct {
		name      string
		inPayload Payload
		levels    []string
		data      []uint32
	}{
		{
			name:      "append similar factor",
			inPayload: FactorPayload([]string{"one", "two", "one"}, nil),
			levels:    []string{"", "one", "two"},
			data:      []uint32{1, 2, 0, 1, 2, 1},
		},
		{
			name:      "append non-similar factor",
			inPayload: FactorPayload([]string{"one", "three", "one"}, nil),
			levels:    []string{"", "one", "two", "three"},
			data:      []uint32{1, 2, 0, 1, 3, 1},
		},
	}
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			outPayload := srcPayload.Append(data.inPayload)

			payload := outPayload.(*factorPayload)

			if !reflect.DeepEqual(payload.data, data.data) {
				t.Error(fmt.Sprintf("Factor data (%v) does not match expected (%v)",
					payload.data, data.data))
			}

			if !reflect.DeepEqual(payload.levels, data.levels) {
				t.Error(fmt.Sprintf("Factor levels (%v) do not match expected (%v)",
					payload.levels, data.levels))
			}
		})
	}
}

func TestFactorPayload_SupportsWhicher(t *testing.T) {
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
			name:        "func(string, bool) bool",
			filter:      func(string, bool) bool { return true },
			isSupported: true,
		},
		{
			name:        "func(int, float64, bool) bool",
			filter:      func(int, float64, bool) bool { return true },
			isSupported: false,
		},
	}

	payload := FactorPayload([]string{"one"}, nil).(Whichable)

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if payload.SupportsWhicher(data.filter) != data.isSupported {
				t.Error("Selector's support is incorrect.")
			}
		})
	}
}

func TestFactorPayload_Whicher(t *testing.T) {
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
			name: "Comparer compact",
			fn:   func(val string, _ bool) bool { return val == "39" || val == "90" },
			out:  []bool{false, false, true, false, false, false, false, true, false, false},
		},
		{
			name: "func() bool {return true}",
			fn:   func() bool { return true },
			out:  []bool{false, false, false, false, false, false, false, false, false, false},
		},
	}

	payload := FactorPayload([]string{"1", "2", "39", "4", "56", "2", "45", "90", "4", "3"}, nil).(Whichable)

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			result := payload.Which(data.fn)
			if !reflect.DeepEqual(result, data.out) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to out (%v)", result, data.out))
			}
		})
	}
}

func TestFactorPayload_SupportsApplier(t *testing.T) {
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
			name:        "func(string, bool) (string, bool)",
			applier:     func(string, bool) (string, bool) { return "", true },
			isSupported: true,
		},
		{
			name:        "func(int, string, bool) bool",
			applier:     func(int, string, bool) bool { return true },
			isSupported: false,
		},
	}

	payload := FactorPayload([]string{}, nil).(Appliable)
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if payload.SupportsApplier(data.applier) != data.isSupported {
				t.Error("Applier's support is incorrect.")
			}
		})
	}
}

func TestFactorPayload_Apply(t *testing.T) {
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
			name: "regular compact",
			applier: func(val string, na bool) (string, bool) {
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
			payload := FactorPayload(data.dataIn, data.naIn).(Appliable).Apply(data.applier)

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

func TestFactorPayload_SupportsSummarizer(t *testing.T) {
	testData := []struct {
		name        string
		summarizer  interface{}
		isSupported bool
	}{
		{
			name:        "valid",
			summarizer:  func(int, string, string, bool) (string, bool) { return "", false },
			isSupported: true,
		},
		{
			name:        "invalid",
			summarizer:  func(int, int, bool) bool { return true },
			isSupported: false,
		},
	}

	payload := StringWithNA([]string{}, nil).(*vector).payload.(Summarizable)
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if payload.SupportsSummarizer(data.summarizer) != data.isSupported {
				t.Error("Summarizer's support is incorrect.")
			}
		})
	}
}

func TestFactorPayload_Summarize(t *testing.T) {
	summarizer := func(idx int, prev string, cur string, na bool) (string, bool) {
		return prev + cur, na
	}

	testData := []struct {
		name        string
		summarizer  interface{}
		dataIn      []string
		naIn        []bool
		dataOut     []string
		naOut       []bool
		isNAPayload bool
	}{
		{
			name:        "true",
			summarizer:  summarizer,
			dataIn:      []string{"1", "2", "1", "6", "5"},
			naIn:        []bool{false, false, false, false, false},
			dataOut:     []string{"12165"},
			naOut:       []bool{false},
			isNAPayload: false,
		},
		{
			name:        "NA",
			summarizer:  summarizer,
			dataIn:      []string{"1", "2", "1", "6", "5"},
			naIn:        []bool{false, false, false, false, true},
			isNAPayload: true,
		},
		{
			name:        "incorrect applier",
			summarizer:  func(int, int, bool) bool { return true },
			dataIn:      []string{"1", "2", "1", "6", "5"},
			naIn:        []bool{false, true, false, true, false},
			isNAPayload: true,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := StringWithNA(data.dataIn, data.naIn).(*vector).payload.(Summarizable).Summarize(data.summarizer)

			if !data.isNAPayload {
				payloadOut := payload.(*stringPayload)
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

func TestFactorPayload_Booleans(t *testing.T) {
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
			payload := FactorPayload(data.in, data.inNA).(*factorPayload)

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

func TestFactorPayload_Integers(t *testing.T) {
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
			outNA: []bool{false, false, true, false, true},
		},
		{
			in:    []string{"10", "", "12", "14", "1110"},
			inNA:  []bool{false, false, false, true, true},
			out:   []int{10, 0, 12, 0, 0},
			outNA: []bool{false, true, false, true, true},
		},
		{
			in:    []string{"1", "3", "", "100", "", "-11", "-10"},
			inNA:  []bool{false, false, false, false, false, false, true},
			out:   []int{1, 3, 0, 100, 0, -11, 0},
			outNA: []bool{false, false, true, false, true, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			payload := FactorPayload(data.in, data.inNA).(*factorPayload)

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

func TestFactorPayload_Floats(t *testing.T) {
	testData := []struct {
		in    []string
		inNA  []bool
		out   []float64
		outNA []bool
	}{
		{
			in:    []string{"1", "3", "", "100", ""},
			inNA:  []bool{false, false, false, false, false},
			out:   []float64{1, 3, math.NaN(), 100, math.NaN()},
			outNA: []bool{false, false, true, false, true},
		},
		{
			in:    []string{"10", "", "12", "14", "1110"},
			inNA:  []bool{false, false, false, true, true},
			out:   []float64{10, math.NaN(), 12, math.NaN(), math.NaN()},
			outNA: []bool{false, true, false, true, true},
		},
		{
			in:    []string{"1", "3", "", "100", "", "-11", "-10"},
			inNA:  []bool{false, false, false, false, false, false, true},
			out:   []float64{1, 3, math.NaN(), 100, math.NaN(), -11, math.NaN()},
			outNA: []bool{false, false, true, false, true, false, true},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			payload := FactorPayload(data.in, data.inNA).(*factorPayload)

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

func TestFactorPayload_Complexes(t *testing.T) {
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
			outNA: []bool{false, false, false, true, false},
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
			payload := FactorPayload(data.in, data.inNA).(*factorPayload)

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

func TestFactorPayload_Strings(t *testing.T) {
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
			payload := FactorPayload(data.in, data.inNA).(*factorPayload)

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

func TestFactorPayload_Interfaces(t *testing.T) {
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
			payload := FactorPayload(data.in, data.inNA).(*factorPayload)

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

func TestFactorPayload_IsNA(t *testing.T) {
	testData := []struct {
		payload Payload
		outNA   []bool
	}{
		{
			payload: FactorPayload([]string{"1", "1", "1", "1", "1", "1"}, []bool{true, true, false, false, false, true}),
			outNA:   []bool{true, true, false, false, false, true},
		},
		{
			payload: FactorPayload([]string{"1", "1", "1"}, []bool{true, true, false}),
			outNA:   []bool{true, true, false},
		},
		{
			payload: FactorPayload([]string{}, []bool{}),
			outNA:   []bool{},
		},
		{
			payload: FactorPayload([]string{}, nil),
			outNA:   []bool{},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			na := data.payload.(NAble).IsNA()

			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("Value IsNA(%v) is not equal to out(%v)", na, data.outNA))
			}
		})
	}
}

func TestFactorPayload_NotNA(t *testing.T) {
	testData := []struct {
		payload Payload
		outNA   []bool
	}{
		{
			payload: FactorPayload([]string{"1", "1", "1", "1", "1", "1"}, []bool{true, true, false, false, false, true}),
			outNA:   []bool{false, false, true, true, true, false},
		},
		{
			payload: FactorPayload([]string{"1", "1", "1"}, []bool{true, true, false}),
			outNA:   []bool{false, false, true},
		},
		{
			payload: FactorPayload([]string{}, []bool{}),
			outNA:   []bool{},
		},
		{
			payload: FactorPayload([]string{}, nil),
			outNA:   []bool{},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			na := data.payload.(NAble).NotNA()

			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("Value NotNA(%v) is not equal to out(%v)", na, data.outNA))
			}
		})
	}
}

func TestFactorPayload_HasNA(t *testing.T) {
	testData := []struct {
		name    string
		payload Payload
		hasNA   bool
	}{
		{
			name:    "no NA (1)",
			payload: FactorPayload([]string{"1", "1", "1", "1", "1"}, []bool{false, false, false, false, false}),
			hasNA:   false,
		},
		{
			name:    "no NA (2)",
			payload: FactorPayload([]string{"1", "1", "1"}, []bool{true, true, false}),
			hasNA:   true,
		},
		{
			name:    "with NA",
			payload: FactorPayload([]string{"1", "1", "1", "1", "1"}, []bool{true, false, true, false, true}),
			hasNA:   true,
		},
		{
			name:    "empty",
			payload: FactorPayload([]string{}, []bool{}),
			hasNA:   false,
		},
		{
			name:    "empty (nil)",
			payload: FactorPayload([]string{}, nil),
			hasNA:   false,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			hasNA := data.payload.(NAble).HasNA()

			if !reflect.DeepEqual(hasNA, data.hasNA) {
				t.Error(fmt.Sprintf("Value NotNA(%v) is not equal to out(%v)", hasNA, data.hasNA))
			}
		})
	}
}

func TestFactorPayload_WithNA(t *testing.T) {
	testData := []struct {
		name    string
		payload Payload
		withNA  []int
	}{
		{
			name:    "without na",
			payload: FactorPayload([]string{"1", "1", "1", "1", "1"}, []bool{false, false, false, false, false}),
			withNA:  []int{},
		},
		{
			name:    "with na#1",
			payload: FactorPayload([]string{"1", "1", "1", "1", "1"}, []bool{true, false, true, false, true}),
			withNA:  []int{1, 3, 5},
		},
		{
			name:    "with na#2",
			payload: FactorPayload([]string{"1", "1", "1", "1", "1"}, []bool{false, false, false, true, true}),
			withNA:  []int{4, 5},
		},
		{
			name:    "with na#3",
			payload: FactorPayload([]string{"1", "1", "1", "1", "1"}, []bool{false, false, false, false, true}),
			withNA:  []int{5},
		},
		{
			name:    "all na",
			payload: FactorPayload([]string{"1", "1", "1", "1", "1"}, []bool{true, true, true, true, true}),
			withNA:  []int{1, 2, 3, 4, 5},
		},
		{
			name:    "empty",
			payload: FactorPayload([]string{}, []bool{}),
			withNA:  []int{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			withNA := data.payload.(NAble).WithNA()

			if !reflect.DeepEqual(withNA, data.withNA) {
				t.Error(fmt.Sprintf("Value NotNA(%v) is not equal to out(%v)", withNA, data.withNA))
			}
		})
	}
}

func TestFactorPayload_WithoutNA(t *testing.T) {
	testData := []struct {
		name      string
		payload   Payload
		withoutNA []int
	}{
		{
			name:      "without na",
			payload:   FactorPayload([]string{"1", "1", "1", "1", "1"}, []bool{false, false, false, false, false}),
			withoutNA: []int{1, 2, 3, 4, 5},
		},
		{
			name:      "with na#1",
			payload:   FactorPayload([]string{"1", "1", "1", "1", "1"}, []bool{true, false, true, false, true}),
			withoutNA: []int{2, 4},
		},
		{
			name:      "with na#2",
			payload:   FactorPayload([]string{"1", "1", "1", "1", "1"}, []bool{false, false, false, true, true}),
			withoutNA: []int{1, 2, 3},
		},
		{
			name:      "with na#3",
			payload:   FactorPayload([]string{"1", "1", "1", "1", "1"}, []bool{false, false, false, false, true}),
			withoutNA: []int{1, 2, 3, 4},
		},
		{
			name:      "all na",
			payload:   FactorPayload([]string{"1", "1", "1", "1", "1"}, []bool{true, true, true, true, true}),
			withoutNA: []int{},
		},
		{
			name:      "empty",
			payload:   FactorPayload([]string{}, []bool{}),
			withoutNA: []int{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			withoutNA := data.payload.(NAble).WithoutNA()

			if !reflect.DeepEqual(withoutNA, data.withoutNA) {
				t.Error(fmt.Sprintf("Value NotNA(%v) is not equal to out(%v)", withoutNA, data.withoutNA))
			}
		})
	}
}

func TestFactorPayload_Find(t *testing.T) {
	payload := FactorPayload([]string{"1", "2", "1", "4", "0"}, nil).(*factorPayload)

	testData := []struct {
		name   string
		needle interface{}
		pos    int
	}{
		{"existent", "4", 4},
		{"non-existent", "non", 0},
		{"incorrect type", true, 0},
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

func TestFactorPayload_FindAll(t *testing.T) {
	payload := FactorPayload([]string{"1", "2", "1", "4", "0"}, nil).(*factorPayload)

	testData := []struct {
		name   string
		needle interface{}
		pos    []int
	}{
		{"existent", "1", []int{1, 3}},
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

func TestFactorPayload_Eq(t *testing.T) {
	payload := StringPayload([]string{"2", "zero", "2", "2", "1"},
		[]bool{false, false, true, false, false}).(*stringPayload)

	testData := []struct {
		eq  interface{}
		cmp []bool
	}{
		{"2", []bool{true, false, false, true, false}},
		{"zero", []bool{false, true, false, false, false}},
		{2, []bool{false, false, false, false, false}},
		{int64(1), []bool{false, false, false, false, false}},
		{int32(1), []bool{false, false, false, false, false}},
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

func TestFactorPayload_Neq(t *testing.T) {
	payload := FactorPayload([]string{"2", "zero", "2", "2", "1"},
		[]bool{false, false, true, false, false}).(*factorPayload)

	testData := []struct {
		eq  interface{}
		cmp []bool
	}{
		{"2", []bool{false, true, true, false, true}},
		{2, []bool{true, true, true, true, true}},
		{int64(1), []bool{true, true, true, true, true}},
		{int32(1), []bool{true, true, true, true, true}},
		{true, []bool{true, true, true, true, true}},
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

func TestFactorPayload_Gt(t *testing.T) {
	payload := FactorPayload([]string{"alpha", "zero", "zeroth", "zero", "gamma"},
		[]bool{false, false, false, true, false}).(*factorPayload)

	testData := []struct {
		val interface{}
		cmp []bool
	}{
		{"zero", []bool{false, false, true, false, false}},
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

func TestFactorPayload_Lt(t *testing.T) {
	payload := FactorPayload([]string{"alpha", "zero", "zeroth", "zeroth", "gamma"},
		[]bool{false, false, false, true, false}).(*factorPayload)

	testData := []struct {
		val interface{}
		cmp []bool
	}{
		{"zero", []bool{true, false, false, false, true}},
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

func TestFactorPayload_Gte(t *testing.T) {
	payload := FactorPayload([]string{"alpha", "zero", "zeroth", "zero", "gamma"},
		[]bool{false, false, false, true, false}).(*factorPayload)

	testData := []struct {
		val interface{}
		cmp []bool
	}{
		{"zero", []bool{false, true, true, false, false}},
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

func TestFactorPayload_Lte(t *testing.T) {
	payload := FactorPayload([]string{"alpha", "zero", "zeroth", "zero", "gamma"},
		[]bool{false, false, false, true, false}).(*factorPayload)

	testData := []struct {
		val interface{}
		cmp []bool
	}{
		{"zero", []bool{true, true, false, false, true}},
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

func TestFactorPayload_Groups(t *testing.T) {
	testData := []struct {
		name    string
		payload Payload
		groups  [][]int
		values  []interface{}
	}{
		{
			name:    "normal",
			payload: FactorPayload([]string{"aa", "bb", "bb", "aa", "cc", "dd", "aa"}, nil),
			groups:  [][]int{{1, 4, 7}, {2, 3}, {5}, {6}},
			values:  []interface{}{"aa", "bb", "cc", "dd"},
		},
		{
			name: "with NA",
			payload: FactorPayload([]string{"aa", "bb", "bb", "aa", "cc", "dd", "aa"},
				[]bool{false, false, false, false, true, false, false}),
			groups: [][]int{{1, 4, 7}, {2, 3}, {6}, {5}},
			values: []interface{}{"aa", "bb", "dd", nil},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			groups, values := data.payload.(*factorPayload).Groups()

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

func TestFactorPayload_IsUnique(t *testing.T) {
	testData := []struct {
		name     string
		payload  Payload
		booleans []bool
	}{
		{
			name:     "without NA",
			payload:  FactorPayload([]string{"1", "2", "1", "3", "2", "3", "2"}, nil),
			booleans: []bool{true, true, false, true, false, false, false},
		},
		{
			name: "with NA",
			payload: FactorPayload([]string{"1", "2", "1", "3", "2", "3", "2"},
				[]bool{false, true, true, false, false, false, false}),
			booleans: []bool{true, true, false, true, true, false, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			booleans := data.payload.(*factorPayload).IsUnique()

			if !reflect.DeepEqual(booleans, data.booleans) {
				t.Error(fmt.Sprintf("Result of IsUnique() (%v) do not match expected (%v)",
					booleans, data.booleans))
			}
		})
	}
}

func TestFactorPayload_Coalesce(t *testing.T) {
	testData := []struct {
		name         string
		coalescer    Payload
		coalescendum Payload
		outData      []uint32
		outLevels    []string
	}{
		{
			name:         "empty",
			coalescer:    FactorPayload(nil, nil),
			coalescendum: FactorPayload([]string{}, nil),
			outData:      []uint32{},
			outLevels:    []string{""},
		},
		{
			name:         "same type",
			coalescer:    FactorPayload([]string{"1", "", "", "", "5"}, []bool{false, true, true, true, false}),
			coalescendum: FactorPayload([]string{"11", "12", "", "14", "15"}, []bool{false, false, true, false, false}),
			outData:      []uint32{1, 2, 0, 3, 4},
			outLevels:    []string{"", "1", "12", "14", "5"},
		},
		{
			name:         "same type + different size",
			coalescer:    FactorPayload([]string{"1", "", "", "", "5"}, []bool{false, true, true, true, false}),
			coalescendum: FactorPayload([]string{"", "11"}, []bool{true, false}),
			outData:      []uint32{1, 2, 0, 2, 3},
			outLevels:    []string{"", "1", "11", "5"},
		},
		{
			name:         "different type",
			coalescer:    FactorPayload([]string{"1", "0", "0", "0", "5"}, []bool{false, true, true, true, false}),
			coalescendum: IntegerPayload([]int{0, 10, 0, 112, 0}, []bool{false, false, true, false, false}),
			outData:      []uint32{1, 2, 0, 3, 4},
			outLevels:    []string{"", "1", "10", "112", "5"},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.coalescer.(Coalescer).Coalesce(data.coalescendum).(*factorPayload)

			if !reflect.DeepEqual(payload.data, data.outData) {
				t.Error(fmt.Sprintf("Data (%v) do not match expected (%v)",
					payload.data, data.outData))
			}

			if !reflect.DeepEqual(payload.levels, data.outLevels) {
				t.Error(fmt.Sprintf("Data (%v) do not match expected (%v)",
					payload.data, data.outData))
			}
		})
	}
}

func TestFactorPayload_SortedIndices(t *testing.T) {
	testData := []struct {
		name          string
		payload       Payload
		sortedIndices []int
	}{
		{
			name:          "with NA",
			payload:       FactorPayload([]string{"delta", "beta", "alpha", "zeroth", "zero"}, []bool{false, false, true, true, false}),
			sortedIndices: []int{2, 1, 5, 3, 4},
		},
		{
			name:          "without NA",
			payload:       FactorPayload([]string{"delta", "beta", "alpha", "zeroth", "zero"}, nil),
			sortedIndices: []int{3, 2, 1, 5, 4},
		},
		{
			name:          "empty",
			payload:       FactorPayload([]string{}, nil),
			sortedIndices: []int{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			sortedIndices := data.payload.(*factorPayload).SortedIndices()

			if !reflect.DeepEqual(sortedIndices, data.sortedIndices) {
				t.Error(fmt.Sprintf("Comparator results (%v) do not match expected (%v)",
					sortedIndices, data.sortedIndices))
			}
		})
	}
}

func TestFactorPayload_SortedIndicesWithRanks(t *testing.T) {
	testData := []struct {
		name          string
		payload       Payload
		sortedIndices []int
		ranks         []int
	}{
		{
			name:          "with NA",
			payload:       FactorPayload([]string{"delta", "beta", "alpha", "zeroth", "zero"}, []bool{false, false, true, true, false}),
			sortedIndices: []int{2, 1, 5, 3, 4},
			ranks:         []int{1, 2, 3, 4, 4},
		},
		{
			name:          "without NA",
			payload:       FactorPayload([]string{"delta", "beta", "alpha", "zeroth", "zero"}, nil),
			sortedIndices: []int{3, 2, 1, 5, 4},
			ranks:         []int{1, 2, 3, 4, 5},
		},
		{
			name:          "empty",
			payload:       FactorPayload([]string{}, nil),
			sortedIndices: []int{},
			ranks:         []int{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			sortedIndices, ranks := data.payload.(*factorPayload).SortedIndicesWithRanks()

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

func TestFactorPayload_Levels(t *testing.T) {
	testData := []struct {
		name    string
		payload Payload
		levels  []string
	}{
		{
			name:    "normal",
			payload: FactorPayload([]string{"one", "two", "two", "one", "three"}, nil),
			levels:  []string{"one", "two", "three"},
		},
		{
			name:    "empty",
			payload: FactorPayload([]string{}, nil),
			levels:  []string{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			levels := data.payload.(*factorPayload).Levels()

			if !reflect.DeepEqual(levels, data.levels) {
				t.Error(fmt.Sprintf("Data (%v) do not match expected (%v)",
					levels, data.levels))
			}
		})
	}
}

func TestFactorPayload_Level(t *testing.T) {
	payload := FactorPayload([]string{"one", "two", "two", "one", "three"}, nil).(*factorPayload)

	testData := []struct {
		level    string
		position int
	}{
		{
			level:    "one",
			position: 1,
		},
		{
			level:    "two",
			position: 2,
		},
		{
			level:    "three",
			position: 3,
		},
		{
			level:    "four",
			position: 0,
		},
	}

	for _, data := range testData {
		t.Run(data.level, func(t *testing.T) {
			position := payload.Level(data.level)

			if !reflect.DeepEqual(position, data.position) {
				t.Error(fmt.Sprintf("Position (%v) do not match expected (%v)",
					position, data.position))
			}
		})
	}
}

func TestFactorPayload_HasLevel(t *testing.T) {
	payload := FactorPayload([]string{"one", "two", "two", "one", "three"}, nil).(*factorPayload)

	testData := []struct {
		level    string
		hasLevel bool
	}{
		{
			level:    "one",
			hasLevel: true,
		},
		{
			level:    "two",
			hasLevel: true,
		},
		{
			level:    "three",
			hasLevel: true,
		},
		{
			level:    "four",
			hasLevel: false,
		},
	}

	for _, data := range testData {
		t.Run(data.level, func(t *testing.T) {
			hasLevel := payload.HasLevel(data.level)

			if hasLevel != data.hasLevel {
				t.Error(fmt.Sprintf("hasLevel (%v) do not match expected (%v)",
					hasLevel, data.hasLevel))
			}
		})
	}
}

func TestFactorPayload_IsSameLevels(t *testing.T) {
	testData := []struct {
		name         string
		factor1      Payload
		factor2      Payload
		isSameLevels bool
	}{
		{
			name:         "same payloads",
			factor1:      FactorPayload([]string{"one", "two", "two", "one", "three"}, nil),
			factor2:      FactorPayload([]string{"one", "two", "two", "one", "three"}, nil),
			isSameLevels: true,
		},
		{
			name:         "different payloads, same levels",
			factor1:      FactorPayload([]string{"one", "two", "two", "one", "three"}, nil),
			factor2:      FactorPayload([]string{"one", "two", "two", "one", "two", "one", "three", "three"}, nil),
			isSameLevels: true,
		},
		{
			name:         "different number of levels",
			factor1:      FactorPayload([]string{"one", "two", "two", "one", "three"}, nil),
			factor2:      FactorPayload([]string{"one", "two", "two", "one", "two", "one"}, nil),
			isSameLevels: false,
		},
		{
			name:         "different order of levels",
			factor1:      FactorPayload([]string{"one", "two", "two", "one", "three"}, nil),
			factor2:      FactorPayload([]string{"one", "three", "two", "one", "two"}, nil),
			isSameLevels: false,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			isSameLevels := data.factor1.(*factorPayload).IsSameLevels(data.factor2.(*factorPayload))

			if isSameLevels != data.isSameLevels {
				t.Error(fmt.Sprintf("isSameLevels (%v) do not match expected (%v)",
					isSameLevels, data.isSameLevels))
			}
		})
	}
}
