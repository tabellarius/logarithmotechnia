package vector

import (
	"logarithmotechnia/embed"
	"logarithmotechnia/option"
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
	embed.NAble
	embed.Arrangeable
}

// Type returns the type of the payload.
func (p *anyPayload) Type() string {
	return "any"
}

// Len returns the length of the payload.
func (p *anyPayload) Len() int {
	return p.length
}

// Pick returns the value at the given index.
func (p *anyPayload) Pick(idx int) any {
	return PickValueWithNA(idx, p.data, p.NA, p.length)
}

// Data returns the data as a slice of interface{}. If the payload contains NA-values, they will be represented
// as nil.
func (p *anyPayload) Data() []any {
	return DataWithNAToInterfaceArray(p.data, p.NA)
}

// ByIndices returns a new payload with the values at the given indices.
func (p *anyPayload) ByIndices(indices []int) Payload {
	data, na := ByIndicesWithNA(indices, p.data, p.NA, nil)

	return AnyPayload(data, na, p.Options()...)
}

// Find returns the index of the first occurrence of the given value. If the value is not found, -1 is returned.
func (p *anyPayload) Find(needle any) int {
	if p.fn.Eq == nil {
		return 0
	}

	return FindFn(needle, p.data, p.NA, p.convertComparator, p.fn.Eq)
}

// FindAll returns the indices of all occurrences of the given value. If the value is not found, an empty slice is
// returned.
func (p *anyPayload) FindAll(needle any) []int {
	if p.fn.Eq == nil {
		return []int{}
	}

	return FindAllFn(needle, p.data, p.NA, p.convertComparator, p.fn.Eq)
}

// StrForElem returns the string representation of the value at the given index. If the payload contains NA-values,
// they will be represented as "NA".
func (p *anyPayload) StrForElem(idx int) string {
	if p.NA[idx-1] {
		return "NA"
	}

	if p.printer != nil {
		return p.printer(p.data[idx-1])
	}

	return ""
}

// SupportsWhicher returns true if the payload supports the given whicher.
func (p *anyPayload) SupportsWhicher(whicher any) bool {
	return SupportsWhicherWithNA[any](whicher)
}

// Which returns a boolean slice with the same length as the payload. The value at each index is true if the
// whicher returns true for the value at the same index.
func (p *anyPayload) Which(whicher any) []bool {
	return WhichWithNA(p.data, p.NA, whicher)
}

// Apply applies the given applier to each value in the payload. The applier can return a new value and a boolean
// indicating if the value is NA. As a result, a new payload is returned.
func (p *anyPayload) Apply(applier any) Payload {
	return ApplyWithNA(p.data, p.NA, applier, p.Options())
}

// ApplyTo applies the given applier to the values at the given indices.
func (p *anyPayload) ApplyTo(indices []int, applier any) Payload {
	data, na := ApplyToWithNA(indices, p.data, p.NA, applier, nil)

	if data == nil {
		return NAPayload(p.length)
	}

	return AnyPayload(data, na, p.Options()...)
}

// Traverse traverses the payload with the given traverser.
func (p *anyPayload) Traverse(traverser any) {
	TraverseWithNA(p.data, p.NA, traverser)
}

// SupportsSummarizer returns true if the payload supports the given summarizer.
func (p *anyPayload) SupportsSummarizer(summarizer any) bool {
	return SupportsSummarizer[any](summarizer)
}

// Summarize returns a new payload with the result of the summarizer applied to the payload. The new payload will
// contain only one value.
func (p *anyPayload) Summarize(summarizer any) Payload {
	val, na := Summarize(p.data, p.NA, summarizer, nil, nil)

	return AnyPayload([]any{val}, []bool{na}, p.Options()...)
}

func (p *anyPayload) convertComparator(val any) (any, bool) {
	return val, true
}

// Eq returns a boolean slice with the same length as the payload. The value at each index is true if the value at
// the same index is equal to the given value. Payload must have the Eq callback set.
func (p *anyPayload) Eq(val any) []bool {
	if p.fn.Eq == nil {
		return make([]bool, p.length)
	}

	return EqFn(val, p.data, p.NA, p.convertComparator, p.fn.Eq)
}

// Neq returns a boolean slice with the same length as the payload. The value at each index is true if the value at
// the same index is not equal to the given value. Payload must have the Eq callback set.
func (p *anyPayload) Neq(val any) []bool {
	if p.fn.Eq == nil {
		return trueBooleanArr(p.length)
	}

	return NeqFn(val, p.data, p.NA, p.convertComparator, p.fn.Eq)
}

// Gt returns a boolean slice with the same length as the payload. The value at each index is true if the value at
// the same index is greater than the given value. Payload must have the Lt callback set.
func (p *anyPayload) Gt(val any) []bool {
	if p.fn.Lt == nil {
		return make([]bool, p.length)
	}

	return GtFn(val, p.data, p.NA, p.convertComparator, p.fn.Lt)
}

// Lt returns a boolean slice with the same length as the payload. The value at each index is true if the value at
// the same index is less than the given value. Payload must have the Lt callback set.
func (p *anyPayload) Lt(val any) []bool {
	if p.fn.Lt == nil {
		return make([]bool, p.length)
	}

	return LtFn(val, p.data, p.NA, p.convertComparator, p.fn.Lt)
}

// Gte returns a boolean slice with the same length as the payload. The value at each index is true if the value at
// the same index is greater than or equal to the given value. Payload must have the Eq and Lt callbacks set.
func (p *anyPayload) Gte(val any) []bool {
	if p.fn.Eq == nil || p.fn.Lt == nil {
		return make([]bool, p.length)
	}

	return GteFn(val, p.data, p.NA, p.convertComparator, p.fn.Eq, p.fn.Lt)
}

// Lte returns a boolean slice with the same length as the payload. The value at each index is true if the value at
// the same index is less than or equal to the given value. Payload must have the Eq and Lt callbacks set.
func (p *anyPayload) Lte(val any) []bool {
	if p.fn.Eq == nil || p.fn.Lt == nil {
		return make([]bool, p.length)
	}

	return LteFn(val, p.data, p.NA, p.convertComparator, p.fn.Eq, p.fn.Lt)
}

// IsUnique returns a boolean slice with the same length as the payload. The value at each index is true if the
// value at the same index has been not seen before. Payload must have the Eq callback set.
func (p *anyPayload) IsUnique() []bool {
	if p.fn.Eq == nil || p.length == 0 || p.length == 1 {
		return trueBooleanArr(p.length)
	}

	uniqIdx := make([]int, p.length)

	naIdx := 0
	for i := 0; i < p.length; i++ {
		if p.NA[i] {
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

// Groups returns a slice of group indices and a slice of values. The payload must have the
// HashInt or HashStr callback set.
func (p *anyPayload) Groups() ([][]int, []any) {
	if p.fn.HashInt != nil {
		return GroupsForDataWithHash(p.data, p.NA, p.fn.HashInt)
	}

	if p.fn.HashStr != nil {
		return GroupsForDataWithHash(p.data, p.NA, p.fn.HashStr)
	}

	data := make([]int, p.length)
	for i := 0; i < p.length; i++ {
		data[i] = i
	}

	groups, _ := GroupsForData(data, p.NA)
	values := make([]any, p.length)
	for i := 0; i < p.length; i++ {
		values[i] = p.data[i]
	}

	return groups, values
}

// Coalesce returns a new payload with the same length as the receiver. The value at each index is the value at the
// same index in the receiver if the receiver is not NA, otherwise the value at the same index in the given payload.
func (p *anyPayload) Coalesce(payload Payload) Payload {
	if p.length != payload.Len() {
		payload = payload.Adjust(p.length)
	}

	var srcData []any
	var srcNA []bool

	if same, ok := payload.(*anyPayload); ok {
		srcData = same.data
		srcNA = same.NA
	} else if intable, ok := payload.(Anyable); ok {
		srcData, srcNA = intable.Anies()
	} else {
		return p
	}

	dstData := make([]any, p.length)
	dstNA := make([]bool, p.length)

	for i := 0; i < p.length; i++ {
		if p.NA[i] && !srcNA[i] {
			dstData[i] = srcData[i]
			dstNA[i] = false
		} else {
			dstData[i] = p.data[i]
			dstNA[i] = p.NA[i]
		}
	}

	return AnyPayload(dstData, dstNA, p.Options()...)
}

// Integers converts the payload's data to a slice of ints and a slice of bools for na-values. The payload must
// have the Intabler callback set.
func (p *anyPayload) Integers() ([]int, []bool) {
	if p.length == 0 {
		return []int{}, []bool{}
	}

	data := make([]int, p.length)
	na := make([]bool, p.Len())

	for i := 0; i < p.length; i++ {
		val, naVal := 0, true
		if p.convertors.Intabler != nil {
			val, naVal = p.convertors.Intabler(i+1, p.data[i], p.NA[i])
		}
		data[i] = val
		na[i] = naVal
	}

	return data, na
}

// Floats converts the payload's data to a slice of float64s and a slice of bools for na-values. The payload must
// have the Floatabler callback set.
func (p *anyPayload) Floats() ([]float64, []bool) {
	if p.length == 0 {
		return []float64{}, []bool{}
	}

	data := make([]float64, p.length)
	na := make([]bool, p.Len())

	for i := 0; i < p.length; i++ {
		val, naVal := math.NaN(), true
		if p.convertors.Floatabler != nil {
			val, naVal = p.convertors.Floatabler(i+1, p.data[i], p.NA[i])
		}
		data[i] = val
		na[i] = naVal
	}

	return data, na
}

// Complexes converts the payload's data to a slice of complex128s and a slice of bools for na-values. The payload must
// have the Complexabler callback set.
func (p *anyPayload) Complexes() ([]complex128, []bool) {
	if p.length == 0 {
		return []complex128{}, []bool{}
	}

	data := make([]complex128, p.length)
	na := make([]bool, p.Len())

	for i := 0; i < p.length; i++ {
		val, naVal := cmplx.NaN(), true
		if p.convertors.Complexabler != nil {
			val, naVal = p.convertors.Complexabler(i+1, p.data[i], p.NA[i])
		}
		data[i] = val
		na[i] = naVal
	}

	return data, na
}

// Booleans converts the payload's data to a slice of bools and a slice of bools for na-values. The payload must
// have the Boolabler callback set.
func (p *anyPayload) Booleans() ([]bool, []bool) {
	if p.length == 0 {
		return []bool{}, []bool{}
	}

	data := make([]bool, p.length)
	na := make([]bool, p.Len())

	for i := 0; i < p.length; i++ {
		val, naVal := false, true
		if p.convertors.Boolabler != nil {
			val, naVal = p.convertors.Boolabler(i+1, p.data[i], p.NA[i])
		}
		data[i] = val
		na[i] = naVal
	}

	return data, na
}

// Strings converts the payload's data to a slice of strings and a slice of bools for na-values. The payload must
// have the Stringabler callback set.
func (p *anyPayload) Strings() ([]string, []bool) {
	if p.length == 0 {
		return []string{}, []bool{}
	}

	data := make([]string, p.length)
	na := make([]bool, p.Len())

	for i := 0; i < p.length; i++ {
		val, naVal := "", true
		if p.convertors.Stringabler != nil {
			val, naVal = p.convertors.Stringabler(i+1, p.data[i], p.NA[i])
		}
		data[i] = val
		na[i] = naVal
	}

	return data, na
}

// Times converts the payload's data to a slice of time.Time and a slice of bools for na-values. The payload must
// have the Timeabler callback set.
func (p *anyPayload) Times() ([]time.Time, []bool) {
	if p.length == 0 {
		return []time.Time{}, []bool{}
	}

	data := make([]time.Time, p.length)
	na := make([]bool, p.Len())

	for i := 0; i < p.length; i++ {
		val, naVal := time.Time{}, true
		if p.convertors.Timeabler != nil {
			val, naVal = p.convertors.Timeabler(i+1, p.data[i], p.NA[i])
		}
		data[i] = val
		na[i] = naVal
	}

	return data, na
}

// Anies returns the payload's data as a slice of any and a slice of bools for na-values.
func (p *anyPayload) Anies() ([]any, []bool) {
	if p.length == 0 {
		return []any{}, []bool{}
	}

	data := make([]any, p.length)
	copy(data, p.data)

	na := make([]bool, p.length)
	copy(na, p.NA)

	return data, na
}

// Append appends a payload to the current payload.
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
	copy(newNA, p.NA)
	copy(newNA[p.length:], na)

	return AnyPayload(newVals, newNA, p.Options()...)
}

// Adjust adjusts the payload's size.
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
	data, na := AdjustToLesserSizeWithNA(p.data, p.NA, size)

	return AnyPayload(data, na, p.Options()...)
}

func (p *anyPayload) adjustToBiggerSize(size int) Payload {
	data, na := AdjustToBiggerSizeWithNA(p.data, p.NA, p.length, size)

	return AnyPayload(data, na, p.Options()...)
}

// Options returns the payload's options.
func (p *anyPayload) Options() []option.Option {
	return []option.Option{
		ConfOption{keyOptionAnyPrinterFunc, p.printer},
		ConfOption{keyOptionAnyConvertors, p.convertors},
		ConfOption{keyOptionAnyCallbacks, p.fn},
	}
}

// SetOption sets an option.
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
func AnyPayload(data []any, na []bool, options ...option.Option) Payload {
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
		NAble: embed.NAble{
			NA: vecNA,
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

	payload.Arrangeable = embed.Arrangeable{
		Length:  payload.length,
		NAble:   payload.NAble,
		FnLess:  fnLess,
		FnEqual: fnEqual,
	}

	return payload
}

// AnyWithNA creates a vector with AnyPayload and allows to set NA-values.
func AnyWithNA(data []any, na []bool, options ...option.Option) Vector {
	return New(AnyPayload(data, na, options...), options...)
}

// Any creates a vector with AnyPayload.
func Any(data []any, options ...option.Option) Vector {
	return AnyWithNA(data, nil, options...)
}
