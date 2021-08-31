package vector

import (
	"time"
)

type TimeWhicherFunc = func(int, time.Time, bool) bool
type TimeWhicherCompactFunc = func(time.Time, bool) bool

type TimePrinter struct {
	Format string
}

type timePayload struct {
	length  int
	data    []time.Time
	printer TimePrinter
	DefNAble
}

func (p *timePayload) Type() string {
	return "time"
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
		length:  len(data),
		data:    data,
		printer: p.printer,
		DefNAble: DefNAble{
			na: na,
		},
	}
}

func (p *timePayload) SupportsWhicher(whicher interface{}) bool {
	if _, ok := whicher.(TimeWhicherFunc); ok {
		return true
	}

	if _, ok := whicher.(TimeWhicherCompactFunc); ok {
		return true
	}

	return false
}

func (p *timePayload) Which(whicher interface{}) []bool {
	if byFunc, ok := whicher.(TimeWhicherFunc); ok {
		return p.selectByFunc(byFunc)
	}

	if byFunc, ok := whicher.(TimeWhicherCompactFunc); ok {
		return p.selectByCompactFunc(byFunc)
	}

	return make([]bool, p.length)
}

func (p *timePayload) selectByFunc(byFunc TimeWhicherFunc) []bool {
	booleans := make([]bool, p.length)

	for idx, val := range p.data {
		if byFunc(idx+1, val, p.na[idx]) {
			booleans[idx] = true
		}
	}

	return booleans
}

func (p *timePayload) selectByCompactFunc(byFunc TimeWhicherCompactFunc) []bool {
	booleans := make([]bool, p.length)

	for idx, val := range p.data {
		if byFunc(val, p.na[idx]) {
			booleans[idx] = true
		}
	}

	return booleans
}

func (p *timePayload) SupportsApplier(applier interface{}) bool {
	if _, ok := applier.(func(int, time.Time, bool) (time.Time, bool)); ok {
		return true
	}

	return false
}

func (p *timePayload) Apply(applier interface{}) Payload {
	var data []time.Time
	var na []bool

	if applyFunc, ok := applier.(func(int, time.Time, bool) (time.Time, bool)); ok {
		data, na = p.applyByFunc(applyFunc)
	} else {
		return NAPayload(p.length)
	}

	return &timePayload{
		length:  p.length,
		data:    data,
		printer: p.printer,
		DefNAble: DefNAble{
			na: na,
		},
	}
}

func (p *timePayload) applyByFunc(applyFunc func(int, time.Time, bool) (time.Time, bool)) ([]time.Time, []bool) {
	data := make([]time.Time, p.length)
	na := make([]bool, p.length)

	for i := 0; i < p.length; i++ {
		dataVal, naVal := applyFunc(i+1, p.data[i], p.na[i])
		if naVal {
			dataVal = time.Time{}
		}
		data[i] = dataVal
		na[i] = naVal
	}

	return data, na
}

func (p *timePayload) SupportsSummarizer(summarizer interface{}) bool {
	if _, ok := summarizer.(func(int, time.Time, time.Time, bool) (time.Time, bool)); ok {
		return true
	}

	return false
}

func (p *timePayload) Summarize(summarizer interface{}) Payload {
	fn, ok := summarizer.(func(int, time.Time, time.Time, bool) (time.Time, bool))
	if !ok {
		return NAPayload(1)
	}

	val := time.Time{}
	na := false
	for i := 0; i < p.length; i++ {
		val, na = fn(i+1, val, p.data[i], p.na[i])

		if na {
			return NAPayload(1)
		}
	}

	return TimePayload([]time.Time{val}, nil)
}

func (p *timePayload) Strings() ([]string, []bool) {
	if p.length == 0 {
		return []string{}, []bool{}
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
		return []time.Time{}, []bool{}
	}

	data := make([]time.Time, p.length)
	copy(data, p.data)

	na := make([]bool, p.Len())
	copy(na, p.na)

	return data, na
}

func (p *timePayload) Interfaces() ([]interface{}, []bool) {
	if p.length == 0 {
		return []interface{}{}, []bool{}
	}

	data := make([]interface{}, p.length)
	for i := 0; i < p.length; i++ {
		if p.na[i] {
			data[i] = nil
		} else {
			data[i] = p.data[i]
		}
	}

	na := make([]bool, p.length)
	copy(na, p.na)

	return data, na
}

func (p *timePayload) Append(vec Vector) Payload {
	length := p.length + vec.Len()

	vals, na := vec.Times()

	newVals := make([]time.Time, length)
	newNA := make([]bool, length)

	copy(newVals, p.data)
	copy(newVals[p.length:], vals)
	copy(newNA, p.na)
	copy(newNA[p.length:], na)

	return TimePayload(newVals, newNA)
}

func (p *timePayload) StrForElem(idx int) string {
	return p.data[idx-1].Format(p.printer.Format)
}

func TimePayload(data []time.Time, na []bool, options ...Config) Payload {
	config := mergeConfigs(options)

	length := len(data)

	vecNA := make([]bool, length)
	if len(na) > 0 {
		if len(na) == length {
			copy(vecNA, na)
		} else {
			emp := NAPayload(0)
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

	return &timePayload{
		length:  length,
		data:    vecData,
		printer: printer,
		DefNAble: DefNAble{
			na: vecNA,
		},
	}
}

func Time(data []time.Time, na []bool, options ...Config) Vector {
	config := mergeConfigs(options)

	return New(TimePayload(data, na), config)
}
