package vector

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStringPayload_Max(t *testing.T) {
	testData := []struct {
		name    string
		payload *stringPayload
		data    []string
		sumNA   []bool
	}{
		{
			name:    "without na",
			payload: StringPayload([]string{"aaa", "aab", "bbaa", "bbba", "bbbb"}, nil).(*stringPayload),
			data:    []string{"bbbb"},
			sumNA:   []bool{false},
		},
		{
			name:    "with na",
			payload: StringPayload([]string{"aaa", "aab", "bbaa", "bbba", "bbbb"}, []bool{false, false, true, false, false}).(*stringPayload),
			data:    []string{""},
			sumNA:   []bool{true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.payload.Max().(*stringPayload)

			if !reflect.DeepEqual(payload.data, data.data) {
				t.Error(fmt.Sprintf("Max data (%v) is not equal to expected (%v)",
					payload.data, data.data))
			}

			if !reflect.DeepEqual(payload.NA, data.sumNA) {
				t.Error(fmt.Sprintf("Max na (%v) is not equal to expected (%v)",
					payload.NA, data.sumNA))
			}
		})
	}
}

func TestStringPayload_Min(t *testing.T) {
	testData := []struct {
		name    string
		payload *stringPayload
		data    []string
		sumNA   []bool
	}{
		{
			name:    "without na",
			payload: StringPayload([]string{"aaa", "aab", "bbaa", "bbba", "bbbb"}, nil).(*stringPayload),
			data:    []string{"aaa"},
			sumNA:   []bool{false},
		},
		{
			name:    "with na",
			payload: StringPayload([]string{"aaa", "aab", "bbaa", "bbba", "bbbb"}, []bool{false, false, true, false, false}).(*stringPayload),
			data:    []string{""},
			sumNA:   []bool{true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.payload.Min().(*stringPayload)

			if !reflect.DeepEqual(payload.data, data.data) {
				t.Error(fmt.Sprintf("Max data (%v) is not equal to expected (%v)",
					payload.data, data.data))
			}

			if !reflect.DeepEqual(payload.NA, data.sumNA) {
				t.Error(fmt.Sprintf("Max na (%v) is not equal to expected (%v)",
					payload.NA, data.sumNA))
			}
		})
	}
}
