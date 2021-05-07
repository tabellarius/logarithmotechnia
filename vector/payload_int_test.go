package vector

import (
	"fmt"
	"reflect"
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
