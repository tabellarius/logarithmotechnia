package vector

import (
	"math"
	"math/cmplx"
	"time"
)

type InterfaceWhicherFunc = func(int, interface{}, bool) bool
type InterfaceWhicherCompactFunc = func(interface{}, bool) bool
type InterfaceApplierFunc = func(int, interface{}, bool) (interface{}, bool)
type InterfaceApplierCompactFunc = func(interface{}, bool) (interface{}, bool)
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

func (p *interfacePayload) Pick(idx int) interface{} {
	return pickValueWithNA(idx, p.data, p.na, p.length)
}

func (p *interfacePayload) Data() []interface{} {
	return dataWithNAToInterfaceArray(p.data, p.na)
}

func (p *interfacePayload) ByIndices(indices []int) Payload {
	data, na := byIndices(indices, p.data, p.na, nil)

	return InterfacePayload(data, na, p.Options()...)
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

func (p *interfacePayload) SupportsWhicher(whicher any) bool {
	return supportsWhicher[any](whicher)
}

func (p *interfacePayload) Which(whicher any) []bool {
	return which(p.data, p.na, whicher)
}

func (p *interfacePayload) SupportsApplier(applier any) bool {
	return supportsApplier[any](applier)
}

func (p *interfacePayload) Apply(applier any) Payload {
	data, na := apply(p.data, p.na, applier, 0)

	if data == nil {
		return NAPayload(p.length)
	}

	return InterfacePayload(data, na, p.Options()...)
}

func (p *interfacePayload) ApplyTo(indices []int, applier any) Payload {
	data, na := applyTo(indices, p.data, p.na, applier, nil)

	if data == nil {
		return NAPayload(p.length)
	}

	return InterfacePayload(data, na, p.Options()...)
}

func (p *interfacePayload) SupportsSummarizer(summarizer any) bool {
	return supportsSummarizer[any](summarizer)
}

func (p *interfacePayload) Summarize(summarizer interface{}) Payload {
	val, na := summarize(p.data, p.na, summarizer, nil, nil)

	return InterfacePayload([]interface{}{val}, []bool{na}, p.Options()...)
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

	return InterfacePayload(newVals, newNA, p.Options()...)
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
	data, na := adjustToLesserSizeWithNA(p.data, p.na, size)

	return InterfacePayload(data, na, p.Options()...)
}

func (p *interfacePayload) adjustToBiggerSize(size int) Payload {
	data, na := adjustToBiggerSizeWithNA(p.data, p.na, p.length, size)

	return InterfacePayload(data, na, p.Options()...)
}

func (p *interfacePayload) Options() []Option {
	return []Option{}
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

func InterfaceWithNA(data []interface{}, na []bool, options ...Option) Vector {
	return New(InterfacePayload(data, na, options...), options...)
}

func Interface(data []interface{}, options ...Option) Vector {
	return InterfaceWithNA(data, nil, options...)
}
