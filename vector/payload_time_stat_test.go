package vector

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestTimePayload_Max(t *testing.T) {
	testData := []struct {
		name    string
		payload *timePayload
		data    []time.Time
		sumNA   []bool
	}{
		{
			name: "without na",
			payload: TimePayload(toTimeData([]string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00", "1800-06-10T11:00:00Z"}),
				nil).(*timePayload),
			data:  toTimeData([]string{"2021-01-01T12:30:00+03:00"}),
			sumNA: []bool{false},
		},
		{
			name: "with na",
			payload: TimePayload(toTimeData([]string{"2006-01-02T15:04:05+07:00", "2021-01-01T12:30:00+03:00", "1800-06-10T11:00:00Z"}),
				[]bool{false, false, true}).(*timePayload),
			data:  toTimeData([]string{"0001-01-01T00:00:00Z"}),
			sumNA: []bool{true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			payload := data.payload.Max().(*timePayload)

			if !reflect.DeepEqual(payload.data, data.data) {
				t.Error(fmt.Sprintf("Max data (%v) is not equal to expected (%v)",
					payload.data, data.data))
			}

			if !reflect.DeepEqual(payload.na, data.sumNA) {
				t.Error(fmt.Sprintf("Max na (%v) is not equal to expected (%v)",
					payload.na, data.sumNA))
			}
		})
	}
}
