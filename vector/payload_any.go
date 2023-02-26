package vector

import (
	"math"
	"math/cmplx"
	"time"
)

type AnyPrinterFunc = func(any) string

// AnyConvertors holds converter functions which allow to use Integers(), Floats(), Complexes(), Booleans(),
// Strings() and Times() methods.If a converter function is not provided, the corresponding function will return
// values as if the payload contains only NA-values.
type AnyConvertors struct {
	Intabler     func(idx int, val any, na bool) (int, bool)
	Floatabler   func(idx int, val any, na bool) (float64, bool)
	Complexabler func(idx int, val any, na bool) (complex128, bool)
	Boolabler    func(idx int, val any, na bool) (bool, bool)
	Stringabler  func(idx int, val any, na bool) (string, bool)
	Timeabler    func(idx int, val any, na bool) (time.Time, bool)
}

// AnyCallbacks holds four functions necessary to enable full vector functionality.
// Eq enables Equalable functionality. Lt enables Ordered and sorting functionality. HashInt or HashStr enables
// grouping and summarizing.
type AnyCallbacks struct {
	Eq      func(any, any) bool
	Lt      func(any, any) bool
	HashInt func(any) int
	HashStr func(any) string
}

type anyPayload struct {
	length     int
	data       []any
	printer    AnyPrinterFunc
	convertors AnyConvertors
	fn         AnyCallbacks
	DefNAble
	DefArrangeable
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
	if p.fn.Eq == nil {
		return 0
	}

	return findFn(needle, p.data, p.na, p.convertComparator, p.fn.Eq)
}

func (p *anyPayload) FindAll(needle any) []int {
	if p.fn.Eq == nil {
		return []int{}
	}

	return findAllFn(needle, p.data, p.na, p.convertComparator, p.fn.Eq)
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

func (p *anyPayload) Traverse(traverser any) {
	traverse(p.data, p.na, traverser)
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
	if p.fn.Eq == nil {
		return make([]bool, p.length)
	}

	return eqFn(val, p.data, p.na, p.convertComparator, p.fn.Eq)
}

func (p *anyPayload) Neq(val any) []bool {
	if p.fn.Eq == nil {
		return trueBooleanArr(p.length)
	}

	return neqFn(val, p.data, p.na, p.convertComparator, p.fn.Eq)
}

func (p *anyPayload) Gt(val any) []bool {
	if p.fn.Lt == nil {
		return make([]bool, p.length)
	}

	return gtFn(val, p.data, p.na, p.convertComparator, p.fn.Lt)
}

func (p *anyPayload) Lt(val any) []bool {
	if p.fn.Lt == nil {
		return make([]bool, p.length)
	}

	return ltFn(val, p.data, p.na, p.convertComparator, p.fn.Lt)
}

func (p *anyPayload) Gte(val any) []bool {
	if p.fn.Eq == nil || p.fn.Lt == nil {
		return make([]bool, p.length)
	}

	return gteFn(val, p.data, p.na, p.convertComparator, p.fn.Eq, p.fn.Lt)
}

func (p *anyPayload) Lte(val any) []bool {
	if p.fn.Eq == nil || p.fn.Lt == nil {
		return make([]bool, p.length)
	}

	return lteFn(val, p.data, p.na, p.convertComparator, p.fn.Eq, p.fn.Lt)
}

func (p *anyPayload) IsUnique() []bool {
	if p.fn.Eq == nil || p.length == 0 || p.length == 1 {
		return trueBooleanArr(p.length)
	}

	uniqIdx := make([]int, p.length)

	naIdx := 0
	for i := 0; i < p.length; i++ {
		if p.na[i] {
			if naIdx == 0 {
				naIdx = i
			}
			uniqIdx[i] = naIdx
		}
	}

	for i := 1; i < p.length; i++ {
		uniqIdx[i] = i
		for j := i - 1; j >= 0; j-- {
			if p.fn.Eq(p.data[i], p.data[j]) {
				uniqIdx[i] = j
				break
			}
		}
	}

	booleans := make([]bool, p.length)
	for i := 0; i < p.length; i++ {
		if uniqIdx[i] == i {
			booleans[i] = true
		}
	}

	return booleans
}

func (p *anyPayload) Groups() ([][]int, []any) {
	if p.fn.HashInt != nil {
		return groupsForDataWithHash(p.data, p.na, p.fn.HashInt)
	}

	if p.fn.HashStr != nil {
		return groupsForDataWithHash(p.data, p.na, p.fn.HashStr)
	}

	data := make([]int, p.length)
	for i := 0; i < p.length; i++ {
		data[i] = i
	}

	groups, _ := groupsForData(data, p.na)
	values := make([]any, p.length)
	for i := 0; i < p.length; i++ {
		values[i] = p.data[i]
	}

	return groups, values
}

func (p *anyPayload) Coalesce(payload Payload) Payload {
	if p.length != payload.Len() {
		payload = payload.Adjust(p.length)
	}

	var srcData []any
	var srcNA []bool

	if same, ok := payload.(*anyPayload); ok {
		srcData = same.data
		srcNA = same.na
	} else if intable, ok := payload.(Anyable); ok {
		srcData, srcNA = intable.Anies()
	} else {
		return p
	}

	dstData := make([]any, p.length)
	dstNA := make([]bool, p.length)

	for i := 0; i < p.length; i++ {
		if p.na[i] && !srcNA[i] {
			dstData[i] = srcData[i]
			dstNA[i] = false
		} else {
			dstData[i] = p.data[i]
			dstNA[i] = p.na[i]
		}
	}

	return AnyPayload(dstData, dstNA, p.Options()...)
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
	return []Option{
		ConfOption{keyOptionAnyPrinterFunc, p.printer},
		ConfOption{keyOptionAnyConvertors, p.convertors},
		ConfOption{keyOptionAnyCallbacks, p.fn},
	}
}

func (p *anyPayload) SetOption(name string, val any) bool {
	switch name {
	case keyOptionAnyPrinterFunc:
		p.printer = val.(AnyPrinterFunc)
	case keyOptionAnyConvertors:
		p.convertors = val.(AnyConvertors)
	case keyOptionAnyCallbacks:
		p.fn = val.(AnyCallbacks)
	default:
		return false
	}

	return true
}

// AnyPayload creates a payload with any data.
//
// Available options are:
//   - OptionAnyPrinterFunc(fn AnyPrinterFunc) - sets a function used for printing an element value.
//   - OptionAnyConvertors(convertors AnyConvertors) - sets convertors to integers, floats, strings etc.
//   - OptionAnyCallbacks(callbacks AnyCallbacks) - sets callbacks to enable full vector functionality.
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

	conf.SetOptions(payload)

	fnLess := func(i, j int) bool {
		return i < j
	}
	fnEqual := func(i, j int) bool {
		return i == j
	}

	if payload.fn.Lt != nil && payload.fn.Eq != nil {
		fnLess = func(i, j int) bool {
			return payload.fn.Lt(payload.data[i], payload.data[j])
		}
		fnEqual = func(i, j int) bool {
			return payload.fn.Eq(payload.data[i], payload.data[j])
		}
	}

	payload.DefArrangeable = DefArrangeable{
		Length:   payload.length,
		DefNAble: payload.DefNAble,
		FnLess:   fnLess,
		FnEqual:  fnEqual,
	}

	return payload
}

// AnyWithNA creates a vector with AnyPayload and allows to set NA-values.
func AnyWithNA(data []any, na []bool, options ...Option) Vector {
	return New(AnyPayload(data, na, options...), options...)
}

// Any creates a vector with AnyPayload.
func Any(data []any, options ...Option) Vector {
	return AnyWithNA(data, nil, options...)
}
