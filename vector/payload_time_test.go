package vector

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	emptyNA := []bool{false, false, false}

	testData := []struct {
		name          string
		data          []string
		na            []bool
		outData       []string
		names         map[string]int
		expectedNames map[string]int
		isEmpty       bool
	}{
		{
			name:    "normal + false na",
			data:    []string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00", "1800-06-10T11:00:00Z"},
			na:      []bool{false, false, false},
			outData: []string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00", "1800-06-10T11:00:00Z"},
			names:   nil,
			isEmpty: false,
		},
		{
			name:    "normal + empty na",
			data:    []string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00", "1800-06-10T11:00:00Z"},
			na:      []bool{},
			outData: []string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00", "1800-06-10T11:00:00Z"},
			names:   nil,
			isEmpty: false,
		},
		{
			name:    "normal + nil na",
			data:    []string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00", "1800-06-10T11:00:00Z"},
			na:      nil,
			outData: []string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00", "1800-06-10T11:00:00Z"},
			names:   nil,
			isEmpty: false,
		},
		{
			name:    "normal + mixed na",
			data:    []string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00", "1800-06-10T11:00:00Z"},
			na:      []bool{false, false, true},
			outData: []string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00", "0001-01-01T00:00:00Z"},
			names:   nil,
			isEmpty: false,
		},
		{
			name:    "normal + incorrect sized na",
			data:    []string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00", "1800-06-10T11:00:00Z"},
			na:      []bool{false, false, false, false},
			names:   nil,
			isEmpty: true,
		},
		{
			name:          "normal + names",
			data:          []string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00", "1800-06-10T11:00:00Z"},
			na:            []bool{false, false, false},
			outData:       []string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00", "1800-06-10T11:00:00Z"},
			names:         map[string]int{"one": 1, "three": 3},
			expectedNames: map[string]int{"one": 1, "three": 3},
			isEmpty:       false,
		},
		{
			name:          "normal + incorrect names",
			data:          []string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00", "1800-06-10T11:00:00Z"},
			na:            []bool{false, false, false},
			outData:       []string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00", "1800-06-10T11:00:00Z"},
			names:         map[string]int{"zero": 0, "one": 1, "three": 3, "five": 5},
			expectedNames: map[string]int{"one": 1, "three": 3},
			isEmpty:       false,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			var v Vector

			timeData := toTimeData(data.data)
			outTimeData := toTimeData(data.outData)

			if data.names == nil {
				v = Time(timeData, data.na)
			} else {
				config := Config{NamesMap: data.names}
				v = Time(timeData, data.na, config).(*vector)
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
					t.Error(fmt.Sprintf("Vector length (%d) is not equal to data length (%d)\n", vv.length, length))
				}

				payload, ok := vv.payload.(*timePayload)
				if !ok {
					t.Error("Payload is not floatPayload")
				} else {
					if !reflect.DeepEqual(payload.data, outTimeData) {
						t.Error(fmt.Sprintf("Payload data (%v) is not equal to correct data (%v)\n",
							payload.data[1:], timeData))
					}

					if vv.length != vv.DefNameable.length || vv.length != payload.length {
						t.Error(fmt.Sprintf("Lengths are different: (vv.length - %d, "+
							"vv.DefNameable.length - %d, payload.length - %d, ",
							vv.length, vv.DefNameable.length, payload.length))
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

				if data.names != nil {
					if !reflect.DeepEqual(vv.names, data.expectedNames) {
						t.Error(fmt.Sprintf("Vector names (%v) is not equal to out names (%v)",
							vv.names, data.expectedNames))
					}
				}

			}
		})
	}
}

func TestTimePayload_Len(t *testing.T) {
	testData := []struct {
		in        []string
		outLength int
	}{
		{
			in:        []string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00", "1800-06-10T11:00:00Z"},
			outLength: 3,
		},
		{
			in:        []string{"2006-01-02T15:04:05+07:00"},
			outLength: 1,
		},
		{
			in:        []string{},
			outLength: 0,
		},
		{
			in:        nil,
			outLength: 0,
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			payload := Time(toTimeData(data.in), nil).(*vector).payload
			if payload.Len() != data.outLength {
				t.Error(fmt.Sprintf("Payloads's length (%d) is not equal to out (%d)",
					payload.Len(), data.outLength))
			}
		})
	}
}

func TestTime_Strings(t *testing.T) {
	testData := []struct {
		in    []string
		inNA  []bool
		out   []string
		outNA []bool
	}{
		{
			in:    []string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00", "1800-06-10T11:00:00Z"},
			inNA:  []bool{false, false, false},
			out:   []string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00", "1800-06-10T11:00:00Z"},
			outNA: []bool{false, false, false},
		},
		{
			in:    []string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00", "1800-06-10T11:00:00Z"},
			inNA:  []bool{false, true, false},
			out:   []string{"2006-01-02T15:04:05+07:00", "", "1800-06-10T11:00:00Z"},
			outNA: []bool{false, true, false},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			//			fmt.Println(toTimeData(data.in))
			vec := Time(toTimeData(data.in), data.inNA)
			payload := vec.(*vector).payload.(*timePayload)

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

func TestTimePayload_Times(t *testing.T) {
	testData := []struct {
		in    []string
		inNA  []bool
		out   []time.Time
		outNA []bool
	}{
		{
			in:    []string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00", "1800-06-10T11:00:00Z"},
			inNA:  []bool{false, false, false},
			out:   toTimeData([]string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00", "1800-06-10T11:00:00Z"}),
			outNA: []bool{false, false, false},
		},
		{
			in:    []string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00", "1800-06-10T11:00:00Z"},
			inNA:  []bool{false, true, false},
			out:   toTimeData([]string{"2006-01-02T15:04:05+07:00", "0001-01-01T00:00:00Z", "1800-06-10T11:00:00Z"}),
			outNA: []bool{false, true, false},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			timeData := toTimeData(data.in)
			vec := Time(timeData, data.inNA)
			payload := vec.(*vector).payload.(*timePayload)

			times, na := payload.Times()
			if !reflect.DeepEqual(times, data.out) {
				t.Error(fmt.Sprintf("Times (%v) are not equal to timeData (%v)\n", times, data.out))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("IsNA (%v) are not equal to data.outNA (%v)\n", na, data.outNA))
			}
		})
	}
}

func TestTimePayload_ByIndices(t *testing.T) {
	vec := Time(toTimeData([]string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00", "1800-06-10T11:00:00Z"}),
		[]bool{false, false, true},
	)
	testData := []struct {
		name    string
		indices []int
		out     []time.Time
		outNA   []bool
	}{
		{
			name:    "all",
			indices: []int{1, 2, 3},
			out:     toTimeData([]string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00", "0001-01-01T00:00:00Z"}),
			outNA:   []bool{false, false, true},
		},
		{
			name:    "all reverse",
			indices: []int{3, 2, 1},
			out:     toTimeData([]string{"0001-01-01T00:00:00Z", "2021-01-01T12:30:00+03:00", "2006-01-02T15:04:05+07:00"}),
			outNA:   []bool{true, false, false},
		},
		{
			name:    "some",
			indices: []int{3, 1},
			out:     toTimeData([]string{"0001-01-01T00:00:00Z", "2006-01-02T15:04:05+07:00"}),
			outNA:   []bool{true, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := vec.ByIndices(data.indices).(*vector).payload.(*timePayload)
			if !reflect.DeepEqual(payload.data, data.out) {
				t.Error(fmt.Sprintf("payload.data (%v) is not equal to data.out (%v)", payload.data, data.out))
			}
			if !reflect.DeepEqual(payload.na, data.outNA) {
				t.Error(fmt.Sprintf("payload.na (%v) is not equal to data.out (%v)", payload.na, data.out))
			}
		})
	}
}

func TestTimePayload_SupportsSelector(t *testing.T) {
	testData := []struct {
		name        string
		filter      interface{}
		isSupported bool
	}{
		{
			name:        "func(int, time.Time, bool) bool",
			filter:      func(int, time.Time, bool) bool { return true },
			isSupported: true,
		},
		{
			name:        "func(int, int, bool) bool",
			filter:      func(int, int, bool) bool { return true },
			isSupported: false,
		},
	}

	payload := Time([]time.Time{}, nil).(*vector).payload.(Selectable)
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if payload.SupportsSelector(data.filter) != data.isSupported {
				t.Error("Selector's support is incorrect.")
			}
		})
	}
}

func TestTimePayload_Select(t *testing.T) {
	testData := []struct {
		name   string
		filter interface{}
		out    []bool
	}{
		{
			name:   "func(int, time.Time, bool) bool",
			filter: func(idx int, val time.Time, na bool) bool { return idx == 1 || na == true },
			out:    []bool{true, false, true},
		},
		{
			name:   "func() bool",
			filter: func() bool { return true },
			out:    []bool{false, false, false},
		},
	}

	payload := Time(toTimeData([]string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00", "1800-06-10T11:00:00Z"}),
		[]bool{false, false, true}).(*vector).payload.(Selectable)

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			filtered := payload.Select(data.filter)
			if !reflect.DeepEqual(payload.Select(data.filter), data.out) {
				t.Error(fmt.Sprintf("payload.Select() (%v) is not equal to out value (%v)",
					filtered, data.out))
			}
		})
	}
}

func toTimeData(times []string) []time.Time {
	timeData := make([]time.Time, len(times))

	for i := 0; i < len(times); i++ {
		t, err := time.Parse(time.RFC3339, times[i])
		if err != nil {
			fmt.Println(err)
		}
		timeData[i] = t
	}

	return timeData
}
