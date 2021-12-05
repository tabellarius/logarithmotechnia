package vector

import (
	"fmt"
	"reflect"
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
