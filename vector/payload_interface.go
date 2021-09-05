package vector

import (
	"math"
	"math/cmplx"
	"time"
)

type InterfaceWhicherFunc = func(int, interface{}, bool) bool
type InterfaceWhicherCompactFunc = func(interface{}, bool) bool
type InterfaceToInterfaceApplierFunc = func(int, interface{}, bool) (interface{}, bool)
type InterfaceToInterfaceApplierCompactFunc = func(interface{}, bool) (interface{}, bool)
type InterfaceSummarizerFunc = func(int, interface{}, interface{}, bool) (interface{}, bool)
type InterfacePrinterFunc = func(interface{}) string

type InterfaceConvertors struct {
	Intabler     func(idx int, val interface{}, na bool) (int, bool)
	Floatabler   func(idx int, val interface{}, na bool) (float64, bool)
	Complexabler func(idx int, val interface{}, na bool) (complex128, bool)
	Boolabler    func(idx int, val interface{}, na bool) (bool, bool)
	Stringabler  func(idx int, val interface{}, na bool) (string, bool)
	Timeabler    func(idx int, val interface{}, na bool) (time.Time, bool)
}

type interfacePayload struct {
	length     int
	data       []interface{}
	printer    InterfacePrinterFunc
	convertors *InterfaceConvertors
	DefNAble
}

func (p *interfacePayload) Type() string {
	return "interface"
}

func (p *interfacePayload) Len() int {
	return p.length
}

func (p *interfacePayload) ByIndices(indices []int) Payload {
	data := make([]interface{}, 0, len(indices))
	na := make([]bool, 0, len(indices))

	for _, idx := range indices {
		data = append(data, p.data[idx-1])
		na = append(na, p.na[idx-1])
	}

	return &interfacePayload{
		length:  len(data),
		data:    data,
		printer: p.printer,
		DefNAble: DefNAble{
			na: na,
		},
	}
}

func (p *interfacePayload) StrForElem(idx int) string {
	if p.na[idx-1] {
		return "NA"
	}

	if p.printer != nil {
		return p.printer(p.data[idx-1])
	}

	return ""
}

func (p *interfacePayload) SupportsWhicher(whicher interface{}) bool {
	if _, ok := whicher.(InterfaceWhicherFunc); ok {
		return true
	}

	if _, ok := whicher.(InterfaceWhicherCompactFunc); ok {
		return true
	}

	return false
}

func (p *interfacePayload) Which(whicher interface{}) []bool {
	if byFunc, ok := whicher.(InterfaceWhicherFunc); ok {
		return p.selectByFunc(byFunc)
	}

	if byFunc, ok := whicher.(InterfaceWhicherCompactFunc); ok {
		return p.selectByCompactFunc(byFunc)
	}

	return make([]bool, p.length)
}

func (p *interfacePayload) selectByFunc(byFunc InterfaceWhicherFunc) []bool {
	booleans := make([]bool, p.length)

	for idx, val := range p.data {
		if byFunc(idx+1, val, p.na[idx]) {
			booleans[idx] = true
		}
	}

	return booleans
}

func (p *interfacePayload) selectByCompactFunc(byFunc InterfaceWhicherCompactFunc) []bool {
	booleans := make([]bool, p.length)

	for idx, val := range p.data {
		if byFunc(val, p.na[idx]) {
			booleans[idx] = true
		}
	}

	return booleans
}

func (p *interfacePayload) SupportsApplier(applier interface{}) bool {
	if _, ok := applier.(InterfaceToInterfaceApplierFunc); ok {
		return true
	}

	if _, ok := applier.(InterfaceToInterfaceApplierCompactFunc); ok {
		return true
	}

	return false
}

func (p *interfacePayload) Apply(applier interface{}) Payload {
	if applyFunc, ok := applier.(InterfaceToInterfaceApplierFunc); ok {
		return p.applyToInterfaceByFunc(applyFunc)
	}

	if applyFunc, ok := applier.(InterfaceToInterfaceApplierCompactFunc); ok {
		return p.applyToInterfaceByCompactFunc(applyFunc)
	}

	return NAPayload(p.length)
}

func (p *interfacePayload) applyToInterfaceByFunc(applyFunc InterfaceToInterfaceApplierFunc) Payload {
	data := make([]interface{}, p.length)
	na := make([]bool, p.length)

	for i := 0; i < p.length; i++ {
		dataVal, naVal := applyFunc(i+1, p.data[i], p.na[i])
		if naVal {
			dataVal = nil
		}
		data[i] = dataVal
		na[i] = naVal
	}

	return InterfacePayload(data, na)
}

func (p *interfacePayload) applyToInterfaceByCompactFunc(applyFunc InterfaceToInterfaceApplierCompactFunc) Payload {
	data := make([]interface{}, p.length)
	na := make([]bool, p.length)

	for i := 0; i < p.length; i++ {
		dataVal, naVal := applyFunc(p.data[i], p.na[i])
		if naVal {
			dataVal = nil
		}
		data[i] = dataVal
		na[i] = naVal
	}

	return InterfacePayload(data, na)
}

func (p *interfacePayload) SupportsSummarizer(summarizer interface{}) bool {
	if _, ok := summarizer.(InterfaceSummarizerFunc); ok {
		return true
	}

	return false
}

func (p *interfacePayload) Summarize(summarizer interface{}) Payload {
	fn, ok := summarizer.(InterfaceSummarizerFunc)
	if !ok {
		return NAPayload(1)
	}

	val := interface{}(nil)
	na := false
	for i := 0; i < p.length; i++ {
		val, na = fn(i+1, val, p.data[i], p.na[i])

		if na {
			return NAPayload(1)
		}
	}

	return InterfacePayload([]interface{}{val}, nil)
}

func (p *interfacePayload) Integers() ([]int, []bool) {
	if p.length == 0 {
		return []int{}, []bool{}
	}

	data := make([]int, p.length)
	na := make([]bool, p.Len())

	for i := 0; i < p.length; i++ {
		val, naVal := 0, true
		if p.convertors != nil && p.convertors.Intabler != nil {
			val, naVal = p.convertors.Intabler(i+1, p.data[i], p.na[i])
		}
		data[i] = val
		na[i] = naVal
	}

	return data, na
}

func (p *interfacePayload) Floats() ([]float64, []bool) {
	if p.length == 0 {
		return []float64{}, []bool{}
	}

	data := make([]float64, p.length)
	na := make([]bool, p.Len())

	for i := 0; i < p.length; i++ {
		val, naVal := math.NaN(), true
		if p.convertors != nil && p.convertors.Floatabler != nil {
			val, naVal = p.convertors.Floatabler(i+1, p.data[i], p.na[i])
		}
		data[i] = val
		na[i] = naVal
	}

	return data, na
}

func (p *interfacePayload) Complexes() ([]complex128, []bool) {
	if p.length == 0 {
		return []complex128{}, []bool{}
	}

	data := make([]complex128, p.length)
	na := make([]bool, p.Len())

	for i := 0; i < p.length; i++ {
		val, naVal := cmplx.NaN(), true
		if p.convertors != nil && p.convertors.Complexabler != nil {
			val, naVal = p.convertors.Complexabler(i+1, p.data[i], p.na[i])
		}
		data[i] = val
		na[i] = naVal
	}

	return data, na
}

func (p *interfacePayload) Booleans() ([]bool, []bool) {
	if p.length == 0 {
		return []bool{}, []bool{}
	}

	data := make([]bool, p.length)
	na := make([]bool, p.Len())

	for i := 0; i < p.length; i++ {
		val, naVal := false, true
		if p.convertors != nil && p.convertors.Boolabler != nil {
			val, naVal = p.convertors.Boolabler(i+1, p.data[i], p.na[i])
		}
		data[i] = val
		na[i] = naVal
	}

	return data, na
}

func (p *interfacePayload) Strings() ([]string, []bool) {
	if p.length == 0 {
		return []string{}, []bool{}
	}

	data := make([]string, p.length)
	na := make([]bool, p.Len())

	for i := 0; i < p.length; i++ {
		val, naVal := "", true
		if p.convertors != nil && p.convertors.Stringabler != nil {
			val, naVal = p.convertors.Stringabler(i+1, p.data[i], p.na[i])
		}
		data[i] = val
		na[i] = naVal
	}

	return data, na
}

func (p *interfacePayload) Times() ([]time.Time, []bool) {
	if p.length == 0 {
		return []time.Time{}, []bool{}
	}

	data := make([]time.Time, p.length)
	na := make([]bool, p.Len())

	for i := 0; i < p.length; i++ {
		val, naVal := time.Time{}, true
		if p.convertors != nil && p.convertors.Timeabler != nil {
			val, naVal = p.convertors.Timeabler(i+1, p.data[i], p.na[i])
		}
		data[i] = val
		na[i] = naVal
	}

	return data, na
}

func (p *interfacePayload) Interfaces() ([]interface{}, []bool) {
	if p.length == 0 {
		return []interface{}{}, []bool{}
	}

	data := make([]interface{}, p.length)
	copy(data, p.data)

	na := make([]bool, p.length)
	copy(na, p.na)

	return data, na
}

func (p *interfacePayload) Append(payload Payload) Payload {
	length := p.length + payload.Len()

	var vals []interface{}
	var na []bool

	if interfaceable, ok := payload.(Interfaceable); ok {
		vals, na = interfaceable.Interfaces()
	} else {
		vals, na = NAPayload(payload.Len()).(Interfaceable).Interfaces()
	}

	newVals := make([]interface{}, length)
	newNA := make([]bool, length)

	copy(newVals, p.data)
	copy(newVals[p.length:], vals)
	copy(newNA, p.na)
	copy(newNA[p.length:], na)

	return InterfacePayload(newVals, newNA)
}

func (p *interfacePayload) Adjust(size int) Payload {
	if size < p.length {
		return p.adjustToLesserSize(size)
	}

	if size > p.length {
		return p.adjustToBiggerSize(size)
	}

	return p
}

func (p *interfacePayload) adjustToLesserSize(size int) Payload {
	data := make([]interface{}, size)
	na := make([]bool, size)

	copy(data, p.data)
	copy(na, p.na)

	return InterfacePayload(data, na)
}

func (p *interfacePayload) adjustToBiggerSize(size int) Payload {
	cycles := size / p.length
	if size%p.length > 0 {
		cycles++
	}

	data := make([]interface{}, cycles*p.length)
	na := make([]bool, cycles*p.length)

	for i := 0; i < cycles; i++ {
		copy(data[i*p.length:], p.data)
		copy(na[i*p.length:], p.na)
	}

	data = data[:size]
	na = na[:size]

	return InterfacePayload(data, na)
}

func InterfacePayload(data []interface{}, na []bool, options ...Option) Payload {
	length := len(data)
	conf := MergeOptions(options)

	vecNA := make([]bool, length)
	if len(na) > 0 {
		if len(na) == length {
			copy(vecNA, na)
		} else {
			emp := NAPayload(0)
			return emp
		}
	}

	vecData := make([]interface{}, length)
	for i := 0; i < length; i++ {
		if vecNA[i] {
			vecData[i] = nil
		} else {
			vecData[i] = data[i]
		}
	}

	payload := &interfacePayload{
		length:     length,
		data:       vecData,
		printer:    nil,
		convertors: nil,
		DefNAble: DefNAble{
			na: vecNA,
		},
	}

	if conf.HasOption(KeyOptionInterfacePrinterFunc) {
		payload.printer = conf.Value(KeyOptionInterfacePrinterFunc).(InterfacePrinterFunc)
	}

	if conf.HasOption(KeyOptionInterfaceConvertors) {
		payload.convertors = conf.Value(KeyOptionInterfaceConvertors).(*InterfaceConvertors)
	}

	return payload

}

func Interface(data []interface{}, na []bool, options ...Option) Vector {
	return New(InterfacePayload(data, na, options...))
}
