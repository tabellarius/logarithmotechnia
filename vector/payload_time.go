package vector

import (
	"time"
)

type TimeWhicherFunc = func(int, time.Time, bool) bool
type TimeWhicherCompactFunc = func(time.Time, bool) bool
type TimeToTimeApplierFunc = func(int, time.Time, bool) (time.Time, bool)
type TimeToTimeApplierCompactFunc = func(time.Time, bool) (time.Time, bool)
type TimeSummarizerFunc = func(int, time.Time, time.Time, bool) (time.Time, bool)

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
	if _, ok := applier.(TimeToTimeApplierFunc); ok {
		return true
	}

	if _, ok := applier.(TimeToTimeApplierCompactFunc); ok {
		return true
	}

	return false
}

func (p *timePayload) Apply(applier interface{}) Payload {
	if applyFunc, ok := applier.(TimeToTimeApplierFunc); ok {
		return p.applyToTimeByFunc(applyFunc)
	}

	if applyFunc, ok := applier.(TimeToTimeApplierCompactFunc); ok {
		return p.applyToTimeByCompactFunc(applyFunc)
	}

	return NAPayload(p.length)
}

func (p *timePayload) applyToTimeByFunc(applyFunc TimeToTimeApplierFunc) Payload {
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

	return TimePayload(data, na)
}

func (p *timePayload) applyToTimeByCompactFunc(applyFunc TimeToTimeApplierCompactFunc) Payload {
	data := make([]time.Time, p.length)
	na := make([]bool, p.length)

	for i := 0; i < p.length; i++ {
		dataVal, naVal := applyFunc(p.data[i], p.na[i])
		if naVal {
			dataVal = time.Time{}
		}
		data[i] = dataVal
		na[i] = naVal
	}

	return TimePayload(data, na)
}

func (p *timePayload) SupportsSummarizer(summarizer interface{}) bool {
	if _, ok := summarizer.(TimeSummarizerFunc); ok {
		return true
	}

	return false
}

func (p *timePayload) Summarize(summarizer interface{}) Payload {
	fn, ok := summarizer.(TimeSummarizerFunc)
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

func (p *timePayload) Append(payload Payload) Payload {
	length := p.length + payload.Len()

	var vals []time.Time
	var na []bool

	if timeable, ok := payload.(Timeable); ok {
		vals, na = timeable.Times()
	} else {
		vals, na = NAPayload(payload.Len()).(Timeable).Times()
	}

	newVals := make([]time.Time, length)
	newNA := make([]bool, length)

	copy(newVals, p.data)
	copy(newVals[p.length:], vals)
	copy(newNA, p.na)
	copy(newNA[p.length:], na)

	return TimePayload(newVals, newNA)
}

func (p *timePayload) Adjust(size int) Payload {
	if size < p.length {
		return p.adjustToLesserSize(size)
	}

	if size > p.length {
		return p.adjustToBiggerSize(size)
	}

	return p
}

func (p *timePayload) adjustToLesserSize(size int) Payload {
	data := make([]time.Time, size)
	na := make([]bool, size)

	copy(data, p.data)
	copy(na, p.na)

	return TimePayload(data, na)
}

func (p *timePayload) adjustToBiggerSize(size int) Payload {
	cycles := size / p.length
	if size%p.length > 0 {
		cycles++
	}

	data := make([]time.Time, cycles*p.length)
	na := make([]bool, cycles*p.length)

	for i := 0; i < cycles; i++ {
		copy(data[i*p.length:], p.data)
		copy(na[i*p.length:], p.na)
	}

	data = data[:size]
	na = na[:size]

	return TimePayload(data, na)
}

func (p *timePayload) StrForElem(idx int) string {
	return p.data[idx-1].Format(p.printer.Format)
}

/* Finder interface */

func (p *timePayload) Find(needle interface{}) int {
	val, ok := needle.(time.Time)
	if !ok {
		return 0
	}

	for i, datum := range p.data {
		if val.Equal(datum) {
			return i + 1
		}
	}

	return 0
}

func (p *timePayload) FindAll(needle interface{}) []int {
	val, ok := needle.(time.Time)
	if !ok {
		return []int{}
	}

	found := []int{}
	for i, datum := range p.data {
		if val.Equal(datum) {
			found = append(found, i+1)
		}
	}

	return found
}

/* Comparable interface */

func (p *timePayload) Eq(val interface{}) []bool {
	cmp := make([]bool, p.length)

	v, ok := val.(time.Time)
	if !ok {
		return cmp
	}

	for i, datum := range p.data {
		if p.na[i] {
			cmp[i] = false
		} else {
			cmp[i] = datum.Equal(v)
		}
	}

	return cmp
}

func (p *timePayload) Neq(val interface{}) []bool {
	cmp := make([]bool, p.length)

	v, ok := val.(time.Time)
	if !ok {
		for i := range p.data {
			cmp[i] = true
		}

		return cmp
	}

	for i, datum := range p.data {
		if p.na[i] {
			cmp[i] = true
		} else {
			cmp[i] = !datum.Equal(v)
		}
	}

	return cmp
}

func (p *timePayload) Gt(val interface{}) []bool {
	cmp := make([]bool, p.length)

	v, ok := val.(time.Time)
	if !ok {
		return cmp
	}

	for i, datum := range p.data {
		if p.na[i] {
			cmp[i] = false
		} else {
			if p.na[i] {
				cmp[i] = false
			} else {
				cmp[i] = datum.After(v)
			}
		}
	}

	return cmp
}

func (p *timePayload) Lt(val interface{}) []bool {
	cmp := make([]bool, p.length)

	v, ok := val.(time.Time)
	if !ok {
		return cmp
	}

	for i, datum := range p.data {
		if p.na[i] {
			cmp[i] = false
		} else {
			if p.na[i] {
				cmp[i] = false
			} else {
				cmp[i] = datum.Before(v)
			}
		}
	}

	return cmp
}

func (p *timePayload) Gte(val interface{}) []bool {
	cmp := make([]bool, p.length)

	v, ok := val.(time.Time)
	if !ok {
		return cmp
	}

	for i, datum := range p.data {
		if p.na[i] {
			cmp[i] = false
		} else {
			cmp[i] = datum.After(v) || datum.Equal(v)
		}
	}

	return cmp
}

func (p *timePayload) Lte(val interface{}) []bool {
	cmp := make([]bool, p.length)

	v, ok := val.(time.Time)
	if !ok {
		return cmp
	}

	for i, datum := range p.data {
		if p.na[i] {
			cmp[i] = false
		} else {
			cmp[i] = datum.Before(v) || datum.Equal(v)
		}
	}

	return cmp
}

func TimePayload(data []time.Time, na []bool) Payload {
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

	return &timePayload{
		length:  length,
		data:    vecData,
		printer: printer,
		DefNAble: DefNAble{
			na: vecNA,
		},
	}
}

func Time(data []time.Time, na []bool) Vector {
	return New(TimePayload(data, na))
}
