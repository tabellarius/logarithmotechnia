package vector

import (
	"time"
)

type TimePrinter struct {
	Format string
}

type timePayload struct {
	length  int
	data    []time.Time
	printer TimePrinter
	DefNAble
}

func (p *timePayload) Len() int {
	return p.length
}

func (p *timePayload) ByIndices(indices []int) Payload {
	data := make([]time.Time, 0, len(indices))
	na := make([]bool, 0, len(indices))

	for _, idx := range indices {
		data = append(data, p.data[idx-1])
		na = append(na, p.na[idx-1])
	}

	return &timePayload{
		length: len(data),
		data:   data,
		DefNAble: DefNAble{
			na: na,
		},
	}
}

func (p *timePayload) SupportsSelector(selector interface{}) bool {
	if _, ok := selector.(func(int, time.Time, bool) bool); ok {
		return true
	}

	return false
}

func (p *timePayload) Select(selector interface{}) []bool {
	if byFunc, ok := selector.(func(int, time.Time, bool) bool); ok {
		return p.selectByFunc(byFunc)
	}

	return make([]bool, p.length)
}

func (p *timePayload) selectByFunc(byFunc func(int, time.Time, bool) bool) []bool {
	booleans := make([]bool, p.length)

	for idx, val := range p.data {
		if byFunc(idx+1, val, p.na[idx]) {
			booleans[idx] = true
		}
	}

	return booleans
}

func (p *timePayload) Strings() ([]string, []bool) {
	if p.length == 0 {
		return []string{}, nil
	}

	data := make([]string, p.length)

	for i := 0; i < p.length; i++ {
		if p.na[i] {
			data[i] = ""
		} else {
			data[i] = p.StrForElem(i + 1)
		}
	}

	na := make([]bool, p.Len())
	copy(na, p.na)

	return data, na
}

func (p *timePayload) Times() ([]time.Time, []bool) {
	if p.length == 0 {
		return []time.Time{}, nil
	}

	data := make([]time.Time, p.length)
	copy(data, p.data)

	na := make([]bool, p.Len())
	copy(na, p.na)

	return data, na
}

func (p *timePayload) StrForElem(idx int) string {
	return p.data[idx-1].Format(p.printer.Format)
}

func (p *timePayload) NAPayload() Payload {
	data := make([]string, p.length)
	na := make([]bool, p.length)
	for i := 0; i < p.length; i++ {
		data[i] = ""
		na[i] = true
	}

	return &stringPayload{
		length: p.length,
		data:   data,
		DefNAble: DefNAble{
			na: na,
		},
	}
}

func Time(data []time.Time, na []bool, options ...Config) Vector {
	config := mergeConfigs(options)

	length := len(data)

	vecNA := make([]bool, length)
	if len(na) > 0 {
		if len(na) == length {
			copy(vecNA, na)
		} else {
			emp := Empty()
			emp.Report().AddError("Float(): data length is not equal to na's length")
			return emp
		}
	}

	vecData := make([]time.Time, length)
	for i := 0; i < length; i++ {
		if vecNA[i] {
			vecData[i] = time.Time{}
		} else {
			vecData[i] = data[i]
		}
	}

	printer := TimePrinter{Format: time.RFC3339}
	if config.TimePrinter != nil {
		printer = *config.TimePrinter
	}

	payload := &timePayload{
		length:  length,
		data:    vecData,
		printer: printer,
		DefNAble: DefNAble{
			na: vecNA,
		},
	}

	return New(payload, config)
}
