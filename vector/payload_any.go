package vector

import (
	"math"
	"math/cmplx"
	"time"
)

type AnyPrinterFunc = func(any) string

type AnyConvertors struct {
	Intabler     func(idx int, val any, na bool) (int, bool)
	Floatabler   func(idx int, val any, na bool) (float64, bool)
	Complexabler func(idx int, val any, na bool) (complex128, bool)
	Boolabler    func(idx int, val any, na bool) (bool, bool)
	Stringabler  func(idx int, val any, na bool) (string, bool)
	Timeabler    func(idx int, val any, na bool) (time.Time, bool)
}

type AnyFn struct {
	eq func(any, any) bool
	lt func(any, any) bool
}

type anyPayload struct {
	length     int
	data       []any
	printer    AnyPrinterFunc
	convertors AnyConvertors
	fn         AnyFn
	DefNAble
}

func (p *anyPayload) Type() string {
	return "any"
}

func (p *anyPayload) Len() int {
	return p.length
}

func (p *anyPayload) Pick(idx int) any {
	return pickValueWithNA(idx, p.data, p.na, p.length)
}

func (p *anyPayload) Data() []any {
	return dataWithNAToInterfaceArray(p.data, p.na)
}

func (p *anyPayload) ByIndices(indices []int) Payload {
	data, na := byIndices(indices, p.data, p.na, nil)

	return AnyPayload(data, na, p.Options()...)
}

func (p *anyPayload) Find(needle any) int {
	if p.fn.eq == nil {
		return 0
	}

	return findFn(needle, p.data, p.na, p.convertComparator, p.fn.eq)
}

func (p *anyPayload) FindAll(needle any) []int {
	if p.fn.eq == nil {
		return []int{}
	}

	return findAllFn(needle, p.data, p.na, p.convertComparator, p.fn.eq)
}

func (p *anyPayload) StrForElem(idx int) string {
	if p.na[idx-1] {
		return "NA"
	}

	if p.printer != nil {
		return p.printer(p.data[idx-1])
	}

	return ""
}

func (p *anyPayload) SupportsWhicher(whicher any) bool {
	return supportsWhicher[any](whicher)
}

func (p *anyPayload) Which(whicher any) []bool {
	return which(p.data, p.na, whicher)
}

func (p *anyPayload) Apply(applier any) Payload {
	return apply(p.data, p.na, applier, p.Options())
}

func (p *anyPayload) ApplyTo(indices []int, applier any) Payload {
	data, na := applyTo(indices, p.data, p.na, applier, nil)

	if data == nil {
		return NAPayload(p.length)
	}

	return AnyPayload(data, na, p.Options()...)
}

func (p *anyPayload) SupportsSummarizer(summarizer any) bool {
	return supportsSummarizer[any](summarizer)
}

func (p *anyPayload) Summarize(summarizer any) Payload {
	val, na := summarize(p.data, p.na, summarizer, nil, nil)

	return AnyPayload([]any{val}, []bool{na}, p.Options()...)
}

func (p *anyPayload) convertComparator(val any) (any, bool) {
	return val, true
}

func (p *anyPayload) Eq(val any) []bool {
	if p.fn.eq == nil {
		return make([]bool, p.length)
	}

	return eqFn(val, p.data, p.na, p.convertComparator, p.fn.eq)
}

func (p *anyPayload) Neq(val any) []bool {
	if p.fn.eq == nil {
		return trueBooleanArr(p.length)
	}

	return neqFn(val, p.data, p.na, p.convertComparator, p.fn.eq)
}

func (p *anyPayload) Gt(val any) []bool {
	if p.fn.lt == nil {
		return make([]bool, p.length)
	}

	return gtFn(val, p.data, p.na, p.convertComparator, p.fn.lt)
}

func (p *anyPayload) Lt(val any) []bool {
	if p.fn.lt == nil {
		return make([]bool, p.length)
	}

	return ltFn(val, p.data, p.na, p.convertComparator, p.fn.lt)
}

func (p *anyPayload) Gte(val any) []bool {
	if p.fn.eq == nil || p.fn.lt == nil {
		return make([]bool, p.length)
	}

	return gteFn(val, p.data, p.na, p.convertComparator, p.fn.eq, p.fn.lt)
}

func (p *anyPayload) Lte(val any) []bool {
	if p.fn.eq == nil || p.fn.lt == nil {
		return make([]bool, p.length)
	}

	return lteFn(val, p.data, p.na, p.convertComparator, p.fn.eq, p.fn.lt)
}

func (p *anyPayload) Integers() ([]int, []bool) {
	if p.length == 0 {
		return []int{}, []bool{}
	}

	data := make([]int, p.length)
	na := make([]bool, p.Len())

	for i := 0; i < p.length; i++ {
		val, naVal := 0, true
		if p.convertors.Intabler != nil {
			val, naVal = p.convertors.Intabler(i+1, p.data[i], p.na[i])
		}
		data[i] = val
		na[i] = naVal
	}

	return data, na
}

func (p *anyPayload) Floats() ([]float64, []bool) {
	if p.length == 0 {
		return []float64{}, []bool{}
	}

	data := make([]float64, p.length)
	na := make([]bool, p.Len())

	for i := 0; i < p.length; i++ {
		val, naVal := math.NaN(), true
		if p.convertors.Floatabler != nil {
			val, naVal = p.convertors.Floatabler(i+1, p.data[i], p.na[i])
		}
		data[i] = val
		na[i] = naVal
	}

	return data, na
}

func (p *anyPayload) Complexes() ([]complex128, []bool) {
	if p.length == 0 {
		return []complex128{}, []bool{}
	}

	data := make([]complex128, p.length)
	na := make([]bool, p.Len())

	for i := 0; i < p.length; i++ {
		val, naVal := cmplx.NaN(), true
		if p.convertors.Complexabler != nil {
			val, naVal = p.convertors.Complexabler(i+1, p.data[i], p.na[i])
		}
		data[i] = val
		na[i] = naVal
	}

	return data, na
}

func (p *anyPayload) Booleans() ([]bool, []bool) {
	if p.length == 0 {
		return []bool{}, []bool{}
	}

	data := make([]bool, p.length)
	na := make([]bool, p.Len())

	for i := 0; i < p.length; i++ {
		val, naVal := false, true
		if p.convertors.Boolabler != nil {
			val, naVal = p.convertors.Boolabler(i+1, p.data[i], p.na[i])
		}
		data[i] = val
		na[i] = naVal
	}

	return data, na
}

func (p *anyPayload) Strings() ([]string, []bool) {
	if p.length == 0 {
		return []string{}, []bool{}
	}

	data := make([]string, p.length)
	na := make([]bool, p.Len())

	for i := 0; i < p.length; i++ {
		val, naVal := "", true
		if p.convertors.Stringabler != nil {
			val, naVal = p.convertors.Stringabler(i+1, p.data[i], p.na[i])
		}
		data[i] = val
		na[i] = naVal
	}

	return data, na
}

func (p *anyPayload) Times() ([]time.Time, []bool) {
	if p.length == 0 {
		return []time.Time{}, []bool{}
	}

	data := make([]time.Time, p.length)
	na := make([]bool, p.Len())

	for i := 0; i < p.length; i++ {
		val, naVal := time.Time{}, true
		if p.convertors.Timeabler != nil {
			val, naVal = p.convertors.Timeabler(i+1, p.data[i], p.na[i])
		}
		data[i] = val
		na[i] = naVal
	}

	return data, na
}

func (p *anyPayload) Anies() ([]any, []bool) {
	if p.length == 0 {
		return []any{}, []bool{}
	}

	data := make([]any, p.length)
	copy(data, p.data)

	na := make([]bool, p.length)
	copy(na, p.na)

	return data, na
}

func (p *anyPayload) Append(payload Payload) Payload {
	length := p.length + payload.Len()

	var vals []any
	var na []bool

	if interfaceable, ok := payload.(Anyable); ok {
		vals, na = interfaceable.Anies()
	} else {
		vals, na = NAPayload(payload.Len()).(Anyable).Anies()
	}

	newVals := make([]any, length)
	newNA := make([]bool, length)

	copy(newVals, p.data)
	copy(newVals[p.length:], vals)
	copy(newNA, p.na)
	copy(newNA[p.length:], na)

	return AnyPayload(newVals, newNA, p.Options()...)
}

func (p *anyPayload) Adjust(size int) Payload {
	if size < p.length {
		return p.adjustToLesserSize(size)
	}

	if size > p.length {
		return p.adjustToBiggerSize(size)
	}

	return p
}

func (p *anyPayload) adjustToLesserSize(size int) Payload {
	data, na := adjustToLesserSizeWithNA(p.data, p.na, size)

	return AnyPayload(data, na, p.Options()...)
}

func (p *anyPayload) adjustToBiggerSize(size int) Payload {
	data, na := adjustToBiggerSizeWithNA(p.data, p.na, p.length, size)

	return AnyPayload(data, na, p.Options()...)
}

func (p *anyPayload) Options() []Option {
	return []Option{}
}

func AnyPayload(data []any, na []bool, options ...Option) Payload {
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

	vecData := make([]any, length)
	for i := 0; i < length; i++ {
		if vecNA[i] {
			vecData[i] = nil
		} else {
			vecData[i] = data[i]
		}
	}

	payload := &anyPayload{
		length:  length,
		data:    vecData,
		printer: nil,
		DefNAble: DefNAble{
			na: vecNA,
		},
	}

	if conf.HasOption(KeyOptionAnyPrinterFunc) {
		payload.printer = conf.Value(KeyOptionAnyPrinterFunc).(AnyPrinterFunc)
	}

	if conf.HasOption(KeyOptionAnyConvertors) {
		payload.convertors = conf.Value(KeyOptionAnyConvertors).(AnyConvertors)
	}

	if conf.HasOption(KeyOptionAnyCallbacks) {
		payload.fn = conf.Value(KeyOptionAnyCallbacks).(AnyFn)
	}

	return payload

}

func AnyWithNA(data []any, na []bool, options ...Option) Vector {
	return New(AnyPayload(data, na, options...), options...)
}

func Any(data []any, options ...Option) Vector {
	return AnyWithNA(data, nil, options...)
}
