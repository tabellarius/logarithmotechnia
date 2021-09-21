package vector

import (
	"logarithmotechnia/util"
	"time"
)

// Vector is an interface for a different vector types. This structure is similar to R-vectors: it starts from 1,
// allows for an extensive indexing, supports IsNA-values and named variables
type Vector interface {
	Type() string
	Len() int
	Payload() Payload
	Clone() Vector

	ByIndices(indices []int) Vector
	FromTo(from, to int) Vector
	Filter(whicher interface{}) Vector
	SupportsWhicher(whicher interface{}) bool
	Which(whicher interface{}) []bool
	SupportsApplier(applier interface{}) bool
	Apply(applier interface{}) Vector
	Append(vec Vector) Vector
	Adjust(size int) Vector

	IsEmpty() bool

	NAble

	Intable
	Floatable
	Boolable
	Stringable
	Complexable
	Timeable
	Interfaceable
	AsInteger() Vector
	AsFloat(options ...Option) Vector
	AsComplex(options ...Option) Vector
	AsBoolean() Vector
	AsString() Vector
	AsTime() Vector
	AsInterface() Vector
	Transform(fn TransformFunc) Vector

	Finder
	Has(interface{}) bool
	Comparable
	Arrangeable

	Report() Report
}

type Payload interface {
	Type() string
	Len() int
	ByIndices(indices []int) Payload
	StrForElem(idx int) string
	Append(payload Payload) Payload
	Adjust(size int) Payload
	Options() []Option
}

type Whichable interface {
	SupportsWhicher(whicher interface{}) bool
	Which(whicher interface{}) []bool
}

type Appliable interface {
	SupportsApplier(applier interface{}) bool
	Apply(applier interface{}) Payload
}

type Summarizable interface {
	SupportsSummarizer(summarizer interface{}) bool
	Summarize(summarizer interface{}) Payload
}

type Intable interface {
	Integers() ([]int, []bool)
}

type Floatable interface {
	Floats() ([]float64, []bool)
}

type Boolable interface {
	Booleans() ([]bool, []bool)
}

type Stringable interface {
	Strings() ([]string, []bool)
}

type Complexable interface {
	Complexes() ([]complex128, []bool)
}

type Timeable interface {
	Times() ([]time.Time, []bool)
}

type Interfaceable interface {
	Interfaces() ([]interface{}, []bool)
}

type TransformFunc func([]interface{}, []bool) Payload

type Configurable interface {
	Options() []Option
}

type Finder interface {
	Find(interface{}) int
	FindAll(interface{}) []int
}

type Comparable interface {
	Eq(interface{}) []bool
	Neq(interface{}) []bool
	Gt(interface{}) []bool
	Lt(interface{}) []bool
	Gte(interface{}) []bool
	Lte(interface{}) []bool
}

type Arrangeable interface {
	SortedIndices() []int
	SortedIndicesWithRanks() ([]int, []int)
}

// vector holds data and functions shared by all vectors
type vector struct {
	length  int
	payload Payload
	report  Report
}

func (v *vector) Type() string {
	return v.payload.Type()
}

// Len returns length of vector
func (v *vector) Len() int {
	return v.length
}

func (v *vector) Payload() Payload {
	return v.payload
}

func (v *vector) Clone() Vector {
	return &vector{
		length:  v.length,
		payload: v.payload,
		report:  v.report.Copy(),
	}
}

func (v *vector) ByIndices(indices []int) Vector {
	var selected []int

	for _, index := range indices {
		if index >= 1 && index <= v.length {
			selected = append(selected, index)
		}
	}

	newPayload := v.payload.ByIndices(selected)
	vec := &vector{
		length:  newPayload.Len(),
		payload: v.payload.ByIndices(selected),
		report:  Report{},
	}

	return vec
}

func (v *vector) normalizeFromTo(from, to int) (int, int) {
	if to > v.length {
		to = v.length
	}
	if from < 1 {
		from = 1
	}

	return from, to
}

func (v *vector) FromTo(from, to int) Vector {
	return v.ByIndices(v.filterByFromTo(from, to))
}

func (v *vector) Filter(whicher interface{}) Vector {
	if index, ok := whicher.(int); ok {
		return v.ByIndices([]int{index})
	}

	if indices, ok := whicher.([]int); ok {
		return v.ByIndices(indices)
	}

	if booleans, ok := whicher.([]bool); ok {
		return v.ByIndices(v.filterByBooleans(booleans))
	}

	if v.SupportsWhicher(whicher) {
		return v.ByIndices(v.filterByBooleans(v.Which(whicher)))
	}

	return NA(0)
}

func (v *vector) SupportsWhicher(whicher interface{}) bool {
	payload, ok := v.payload.(Whichable)
	if ok {
		return payload.SupportsWhicher(whicher)
	}

	return false
}

func (v *vector) Which(whicher interface{}) []bool {
	payload, ok := v.payload.(Whichable)
	if ok && payload.SupportsWhicher(whicher) {
		return payload.Which(whicher)
	}

	return make([]bool, v.length)
}

func (v *vector) SupportsApplier(applier interface{}) bool {
	payload, ok := v.payload.(Appliable)
	if ok {
		return payload.SupportsApplier(applier)
	}

	return false
}

func (v *vector) Apply(applier interface{}) Vector {
	payload, ok := v.payload.(Appliable)
	var newP Payload
	if ok && payload.SupportsApplier(applier) {
		newP = payload.Apply(applier)
	} else {
		newP = NAPayload(v.payload.Len())
	}

	newV := v.Clone().(*vector)
	newV.payload = newP

	return newV
}

func (v *vector) filterByBooleans(booleans []bool) []int {
	return util.ToIndices(v.length, booleans)
}

func (v *vector) filterByFromTo(from int, to int) []int {
	/* from and to have different signs */
	if from*to < 0 {
		v.Report().AddError("From and to can not have different signs.")
		return []int{}
	}

	var indices []int
	if from == 0 && to == 0 {
		indices = []int{}
	} else if from > 0 && from > to {
		indices = v.byFromToReverse(to, from)
	} else if from <= 0 && to <= 0 {
		from *= -1
		to *= -1
		if from > to {
			from, to = to, from
		}
		indices = v.byFromToWithRemove(from, to)
	} else {
		indices = v.byFromToRegular(from, to)
	}

	return indices
}

func (v *vector) byFromToRegular(from, to int) []int {
	from, to = v.normalizeFromTo(from, to)

	indices := make([]int, to-from+1)
	index := 0
	for idx := from; idx <= to; idx++ {
		indices[index] = idx
		index++
	}

	return indices
}

func (v *vector) byFromToReverse(from, to int) []int {
	from, to = v.normalizeFromTo(from, to)

	indices := make([]int, to-from+1)
	index := 0
	for idx := to; idx >= from; idx-- {
		indices[index] = idx
		index++
	}

	return indices
}

func (v *vector) byFromToWithRemove(from, to int) []int {
	from, to = v.normalizeFromTo(from, to)

	indices := make([]int, from-1+v.length-to)
	index := 0
	for idx := 1; idx < from; idx++ {
		indices[index] = idx
		index++
	}
	for idx := to + 1; idx <= v.length; idx++ {
		indices[index] = idx
		index++
	}

	return indices
}

func (v *vector) Append(vec Vector) Vector {
	newPayload := v.payload.Append(vec.Payload())

	return New(newPayload)
}

func (v *vector) Adjust(size int) Vector {
	newPayload := v.payload.Adjust(size)

	return New(newPayload)
}

func (v *vector) IsNA() []bool {
	if nable, ok := v.payload.(NAble); ok {
		return nable.IsNA()
	}

	return make([]bool, v.length)
}

func (v *vector) NotNA() []bool {
	if nable, ok := v.payload.(NAble); ok {
		return nable.NotNA()
	}

	notNA := make([]bool, v.length)
	for i := 0; i < v.length; i++ {
		notNA[i] = true
	}

	return notNA
}

/* Not Applicable-related */

func (v *vector) HasNA() bool {
	if nable, ok := v.payload.(NAble); ok {
		return nable.HasNA()
	}

	return false
}

func (v *vector) WithNA() []int {
	if nable, ok := v.payload.(NAble); ok {
		return nable.WithNA()
	}

	return []int{}
}

func (v *vector) WithoutNA() []int {
	if nable, ok := v.payload.(NAble); ok {
		return nable.WithoutNA()
	}

	return []int{}
}

func (v *vector) IsEmpty() bool {
	return v.length == 0
}

func (v *vector) Report() Report {
	return v.report
}

func (v *vector) String() string {
	str := "["

	if v.length > 0 {
		str += v.strForElem(1)
	}
	if v.length > 1 {
		for i := 2; i <= v.length; i++ {
			if i <= maxIntPrint {
				str += ", " + v.strForElem(i)
			} else {
				str += ", ..."
				break
			}
		}
	}

	str += "]"

	return str
}

func (v *vector) strForElem(idx int) string {
	str := v.payload.StrForElem(idx)

	return str
}

func (v *vector) Strings() ([]string, []bool) {
	if payload, ok := v.payload.(Stringable); ok {
		return payload.Strings()
	}

	return NA(v.length).Strings()
}

func (v *vector) Floats() ([]float64, []bool) {
	if payload, ok := v.payload.(Floatable); ok {
		return payload.Floats()
	}

	return NA(v.length).Floats()
}

func (v *vector) Complexes() ([]complex128, []bool) {
	if payload, ok := v.payload.(Complexable); ok {
		return payload.Complexes()
	}

	return NA(v.length).Complexes()
}

func (v *vector) Booleans() ([]bool, []bool) {
	if payload, ok := v.payload.(Boolable); ok {
		return payload.Booleans()
	}

	return NA(v.length).Booleans()
}

func (v *vector) Integers() ([]int, []bool) {
	if payload, ok := v.payload.(Intable); ok {
		return payload.Integers()
	}

	return NA(v.length).Integers()
}

func (v *vector) Times() ([]time.Time, []bool) {
	if payload, ok := v.payload.(Timeable); ok {
		return payload.Times()
	}

	return NA(v.length).Times()
}

func (v *vector) Interfaces() ([]interface{}, []bool) {
	if payload, ok := v.payload.(Interfaceable); ok {
		return payload.Interfaces()
	}

	return NA(v.length).Interfaces()
}

func (v *vector) AsInteger() Vector {
	if payload, ok := v.payload.(Intable); ok {
		values, na := payload.Integers()

		return Integer(values, na)
	}

	return NA(v.length)
}

func (v *vector) AsFloat(options ...Option) Vector {
	if payload, ok := v.payload.(Floatable); ok {
		values, na := payload.Floats()

		return Float(values, na, options...)
	}

	return NA(v.length)
}

func (v *vector) AsComplex(options ...Option) Vector {
	if payload, ok := v.payload.(Complexable); ok {
		values, na := payload.Complexes()

		return Complex(values, na, options...)
	}

	return NA(v.length)
}

func (v *vector) AsBoolean() Vector {
	if payload, ok := v.payload.(Boolable); ok {
		values, na := payload.Booleans()

		return Boolean(values, na)
	}

	return NA(v.length)
}

func (v *vector) AsString() Vector {
	if payload, ok := v.payload.(Stringable); ok {
		values, na := payload.Strings()

		return String(values, na)
	}

	return NA(v.length)
}

func (v *vector) AsTime() Vector {
	if payload, ok := v.payload.(Timeable); ok {
		values, na := payload.Times()

		return Time(values, na)
	}

	return NA(v.length)
}

func (v *vector) AsInterface() Vector {
	if payload, ok := v.payload.(Interfaceable); ok {
		values, na := payload.Interfaces()

		return Interface(values, na)
	}

	return NA(v.length)
}

func (v *vector) Transform(fn TransformFunc) Vector {
	if interfaceable, ok := v.Payload().(Interfaceable); ok {
		values, na := interfaceable.Interfaces()
		payload := fn(values, na)

		return New(payload)
	}

	return NA(v.length)
}

/* Finder interface */

func (v *vector) Find(needle interface{}) int {
	if finder, ok := v.payload.(Finder); ok {
		return finder.Find(needle)
	}

	return 0
}

func (v *vector) FindAll(needle interface{}) []int {
	if finder, ok := v.payload.(Finder); ok {
		return finder.FindAll(needle)
	}

	return []int{}
}

func (v *vector) Has(needle interface{}) bool {
	if finder, ok := v.payload.(Finder); ok {
		return finder.Find(needle) > 0
	}

	return false
}

/* Comparable interface */

func (v *vector) Eq(val interface{}) []bool {
	if comparable, ok := v.payload.(Comparable); ok {
		return comparable.Eq(val)
	}

	return make([]bool, v.length)
}

func (v *vector) Neq(val interface{}) []bool {
	if comparable, ok := v.payload.(Comparable); ok {
		return comparable.Neq(val)
	}

	cmp := make([]bool, v.length)
	for i := range cmp {
		cmp[i] = true
	}

	return cmp
}

func (v *vector) Gt(val interface{}) []bool {
	if comparable, ok := v.payload.(Comparable); ok {
		return comparable.Gt(val)
	}

	return make([]bool, v.length)
}

func (v *vector) Lt(val interface{}) []bool {
	if comparable, ok := v.payload.(Comparable); ok {
		return comparable.Lt(val)
	}

	return make([]bool, v.length)
}

func (v *vector) Gte(val interface{}) []bool {
	if comparable, ok := v.payload.(Comparable); ok {
		return comparable.Gte(val)
	}

	return make([]bool, v.length)
}

func (v *vector) Lte(val interface{}) []bool {
	if comparable, ok := v.payload.(Comparable); ok {
		return comparable.Lte(val)
	}

	return make([]bool, v.length)
}

/* Arrangeable interface */

func (v *vector) SortedIndices() []int {
	if arrangeable, ok := v.payload.(Arrangeable); ok {
		return arrangeable.SortedIndices()
	}

	return indicesArray(v.length)
}

func (v *vector) SortedIndicesWithRanks() ([]int, []int) {
	if arrangeable, ok := v.payload.(Arrangeable); ok {
		return arrangeable.SortedIndicesWithRanks()
	}

	indices := indicesArray(v.length)

	return indices, indices
}

// New creates a vector part of the future vector. This function is used by public functions which create
// typed vectors
func New(payload Payload) Vector {
	vec := vector{
		length:  payload.Len(),
		payload: payload,
		report:  Report{},
	}

	return &vec
}
