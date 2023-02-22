package vector

import (
	"logarithmotechnia/util"
	"time"
)

const maxPrintElements = 15

// Vector is an interface for a different vector types. This structure is similar to R-vectors: it starts from 1,
// allows for an extensive indexing, supports IsNA-values and named variables
type Vector interface {
	Name() string
	SetName(name string) Vector

	Type() string
	Len() int
	Payload() Payload
	Clone() Vector

	ByIndices(indices []int) Vector
	FromTo(from, to int) Vector
	Filter(whicher any) Vector
	SupportsWhicher(whicher any) bool
	Which(whicher any) []bool
	Apply(applier any) Vector
	ApplyTo(whicher any, applier any) Vector
	Traverse(traverser any)
	Append(vec Vector) Vector
	Adjust(size int) Vector
	Pick(idx int) any
	Data() []any

	Groups() ([][]int, []any)
	Ungroup() Vector
	IsGrouped() bool
	GroupByIndices(index GroupIndex) Vector
	GroupVectors() []Vector
	GroupFirstElements() []int

	IsEmpty() bool

	NAble

	Intable
	Floatable
	Boolable
	Stringable
	Complexable
	Timeable
	Anyable
	AsInteger(options ...Option) Vector
	AsFloat(options ...Option) Vector
	AsComplex(options ...Option) Vector
	AsBoolean(options ...Option) Vector
	AsString(options ...Option) Vector
	AsTime(options ...Option) Vector
	AsAny(options ...Option) Vector

	Finder
	Has(any) bool
	Comparable
	Ordered
	Arrangeable

	IsUniquer
	Unique() Vector

	Coalesce(...Vector) Vector

	Odd() []bool
	Even() []bool
	Nth(int) []bool

	Options() []Option
	SetOption(Option) bool

	Arithmetics
	Statistics
}

type Payload interface {
	Type() string
	Len() int
	ByIndices(indices []int) Payload
	StrForElem(idx int) string
	Append(payload Payload) Payload
	Adjust(size int) Payload
	Options() []Option
	SetOption(string, any) bool
	Pick(idx int) any
	Data() []any
}

type Whichable interface {
	SupportsWhicher(whicher any) bool
	Which(whicher any) []bool
}

type Appliable interface {
	Apply(applier any) Payload
}

type AppliableTo interface {
	Whichable
	ApplyTo(indices []int, applier any) Payload
}

type Traversable interface {
	Traverse(traverser any)
}

type Summarizable interface {
	SupportsSummarizer(summarizer any) bool
	Summarize(summarizer any) Payload
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

type Anyable interface {
	Anies() ([]any, []bool)
}

type Configurable interface {
	Options() []Option
}

type Finder interface {
	Find(any) int
	FindAll(any) []int
}

type Comparable interface {
	Eq(any) []bool
	Neq(any) []bool
}

type Ordered interface {
	Gt(any) []bool
	Lt(any) []bool
	Gte(any) []bool
	Lte(any) []bool
}

type Arrangeable interface {
	SortedIndices() []int
	SortedIndicesWithRanks() ([]int, []int)
}

type Grouper interface {
	Groups() ([][]int, []any)
}

type IsUniquer interface {
	IsUnique() []bool
}

type Coalescer interface {
	Coalesce(Payload) Payload
}

type vectorOptions struct {
	maxPrintElements int
}

// vector holds data and functions shared by all vectors
type vector struct {
	name       string
	length     int
	payload    Payload
	groupIndex GroupIndex
	options    vectorOptions
}

func (v *vector) Name() string {
	return v.name
}

func (v *vector) SetName(name string) Vector {
	v.name = name

	return v
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
	vec := New(v.payload, v.Options()...)
	vec.(*vector).groupIndex = v.groupIndex

	return vec
}

func (v *vector) ByIndices(indices []int) Vector {
	var selected []int

	for _, index := range indices {
		if index >= 0 && index <= v.length {
			selected = append(selected, index)
		}
	}

	newPayload := v.payload.ByIndices(selected)

	return New(newPayload, v.Options()...)
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

func (v *vector) Filter(whicher any) Vector {
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

func (v *vector) SupportsWhicher(whicher any) bool {
	payload, ok := v.payload.(Whichable)
	if ok {
		return payload.SupportsWhicher(whicher)
	}

	return false
}

func (v *vector) Which(whicher any) []bool {
	payload, ok := v.payload.(Whichable)
	if ok && payload.SupportsWhicher(whicher) {
		return payload.Which(whicher)
	}

	return make([]bool, v.length)
}

func (v *vector) Apply(applier any) Vector {
	payload, ok := v.payload.(Appliable)
	if !ok {
		return NA(v.Len())
	}

	newPayload := payload.Apply(applier)

	return New(newPayload, v.Options()...)
}

func (v *vector) ApplyTo(whicher any, applier any) Vector {
	payload, ok := v.payload.(AppliableTo)
	if !ok {
		return NA(v.length)
	}

	indices := []int{}

	whBool, ok := whicher.([]bool)
	processed := false
	if ok {
		indices = util.ToIndices(v.length, whBool)
		processed = true
	}

	whIdx, ok := whicher.([]int)
	if ok {
		indices = v.applyToAdjustIndicesWhicher(whIdx)
		processed = true
	}

	if !processed {
		indices = util.ToIndices(v.length, v.Which(whicher))
	}

	newPayload := payload.ApplyTo(indices, applier)

	return New(newPayload, v.Options()...)
}

func (v *vector) applyToAdjustIndicesWhicher(whicher []int) []int {
	indices := make([]int, 0)

	for _, index := range whicher {
		if index > 0 && index <= v.length {
			indices = append(indices, index)
		}
	}

	return indices
}

func (v *vector) filterByBooleans(booleans []bool) []int {
	return util.ToIndices(v.length, booleans)
}

func (v *vector) filterByFromTo(from int, to int) []int {
	/* from and to have different signs */
	if from*to < 0 {
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

func (v *vector) Traverse(traverser any) {
	if payload, ok := v.payload.(Traversable); ok {
		payload.Traverse(traverser)
	}
}

func (v *vector) Append(vec Vector) Vector {
	newPayload := v.payload.Append(vec.Payload())

	return New(newPayload, v.Options()...)
}

func (v *vector) Adjust(size int) Vector {
	newPayload := v.payload.Adjust(size)

	return New(newPayload, v.Options()...)
}

func (v *vector) Pick(idx int) any {
	return v.payload.Pick(idx)
}

func (v *vector) Data() []any {
	return v.payload.Data()
}

func (v *vector) Groups() ([][]int, []any) {
	if groupper, ok := v.payload.(Grouper); ok {
		return groupper.Groups()
	}

	return [][]int{incIndices(indicesArray(v.length))}, []any{nil}
}

func (v *vector) IsGrouped() bool {
	return v.groupIndex != nil
}

func (v *vector) GroupByIndices(groups GroupIndex) Vector {
	if len(groups) == 0 {
		return v
	}

	newVec := New(v.payload, v.Options()...).(*vector)
	newVec.groupIndex = groups

	return newVec
}

func (v *vector) GroupVectors() []Vector {
	if !v.IsGrouped() {
		return nil
	}

	vectors := make([]Vector, len(v.groupIndex))
	for i, indices := range v.groupIndex {
		vectors[i] = v.ByIndices(indices)
	}

	return vectors
}

func (v *vector) GroupFirstElements() []int {
	indices := []int{}

	if v.IsGrouped() {
		if v.Len() > 0 {
			indices = v.groupIndex.FirstElements()
		}
	} else {
		indices = append(indices, 1)
	}

	return indices
}

func (v *vector) Ungroup() Vector {
	if !v.IsGrouped() {
		return v
	}

	newVec := New(v.payload, v.Options()...).(*vector)

	return newVec
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

func (v *vector) HasNA() bool {
	if nable, ok := v.payload.(NAble); ok {
		return nable.HasNA()
	}

	return false
}

/* Not Applicable-related */

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

func (v *vector) String() string {
	str := "[(" + v.Type() + ")]"

	if v.length > 0 {
		str += v.strForElem(1)
	}
	if v.length > 1 {
		for i := 2; i <= v.length; i++ {
			if i <= v.options.maxPrintElements {
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

func (v *vector) Anies() ([]any, []bool) {
	if payload, ok := v.payload.(Anyable); ok {
		return payload.Anies()
	}

	return NA(v.length).Anies()
}

func (v *vector) AsInteger(options ...Option) Vector {
	if payload, ok := v.payload.(Intable); ok {
		values, na := payload.Integers()

		return IntegerWithNA(values, na, options...)
	}

	return NA(v.length)
}

func (v *vector) AsFloat(options ...Option) Vector {
	if payload, ok := v.payload.(Floatable); ok {
		values, na := payload.Floats()

		return FloatWithNA(values, na, options...)
	}

	return NA(v.length)
}

func (v *vector) AsComplex(options ...Option) Vector {
	if payload, ok := v.payload.(Complexable); ok {
		values, na := payload.Complexes()

		return ComplexWithNA(values, na, options...)
	}

	return NA(v.length)
}

func (v *vector) AsBoolean(options ...Option) Vector {
	if payload, ok := v.payload.(Boolable); ok {
		values, na := payload.Booleans()

		return BooleanWithNA(values, na, options...)
	}

	return NA(v.length)
}

func (v *vector) AsString(options ...Option) Vector {
	if payload, ok := v.payload.(Stringable); ok {
		values, na := payload.Strings()

		return StringWithNA(values, na, options...)
	}

	return NA(v.length)
}

func (v *vector) AsTime(options ...Option) Vector {
	if payload, ok := v.payload.(Timeable); ok {
		values, na := payload.Times()

		return TimeWithNA(values, na, options...)
	}

	return NA(v.length)
}

func (v *vector) AsAny(options ...Option) Vector {
	if payload, ok := v.payload.(Anyable); ok {
		values, na := payload.Anies()

		return AnyWithNA(values, na, options...)
	}

	return NA(v.length)
}

func (v *vector) Find(needle any) int {
	if finder, ok := v.payload.(Finder); ok {
		return finder.Find(needle)
	}

	return 0
}

/* Finder interface */

func (v *vector) FindAll(needle any) []int {
	if finder, ok := v.payload.(Finder); ok {
		return finder.FindAll(needle)
	}

	return []int{}
}

func (v *vector) Has(needle any) bool {
	if finder, ok := v.payload.(Finder); ok {
		return finder.Find(needle) > 0
	}

	return false
}

/* Comparable interface */

func (v *vector) Eq(val any) []bool {
	if comparee, ok := v.payload.(Comparable); ok {
		return comparee.Eq(val)
	}

	return make([]bool, v.length)
}

func (v *vector) Neq(val any) []bool {
	if comparee, ok := v.payload.(Comparable); ok {
		return comparee.Neq(val)
	}

	cmp := make([]bool, v.length)
	for i := range cmp {
		cmp[i] = true
	}

	return cmp
}

/* Ordered interface */

func (v *vector) Gt(val any) []bool {
	if comparee, ok := v.payload.(Ordered); ok {
		return comparee.Gt(val)
	}

	return make([]bool, v.length)
}

func (v *vector) Lt(val any) []bool {
	if comparee, ok := v.payload.(Ordered); ok {
		return comparee.Lt(val)
	}

	return make([]bool, v.length)
}

func (v *vector) Gte(val any) []bool {
	if comparee, ok := v.payload.(Ordered); ok {
		return comparee.Gte(val)
	}

	return make([]bool, v.length)
}

func (v *vector) Lte(val any) []bool {
	if comparee, ok := v.payload.(Ordered); ok {
		return comparee.Lte(val)
	}

	return make([]bool, v.length)
}

func (v *vector) SortedIndices() []int {
	if arrangeable, ok := v.payload.(Arrangeable); ok {
		return arrangeable.SortedIndices()
	}

	return indicesArray(v.length)
}

/* Arrangeable interface */

func (v *vector) SortedIndicesWithRanks() ([]int, []int) {
	if arrangeable, ok := v.payload.(Arrangeable); ok {
		return arrangeable.SortedIndicesWithRanks()
	}

	indices := indicesArray(v.length)

	return indices, indices
}

func (v *vector) Unique() Vector {
	if uniquer, ok := v.payload.(IsUniquer); ok {
		return v.Filter(uniquer.IsUnique())
	}

	return v
}

func (v *vector) IsUnique() []bool {
	if uniquer, ok := v.payload.(IsUniquer); ok {
		return uniquer.IsUnique()
	}

	return trueBooleanArr(v.length)
}

func (v *vector) Coalesce(vectors ...Vector) Vector {
	if len(vectors) == 0 {
		return v
	}

	coalescer, ok := v.payload.(Coalescer)
	if !ok {
		return v
	}

	var payload Payload
	for _, v := range vectors {
		payload = coalescer.Coalesce(v.Payload())
		coalescer, ok = payload.(Coalescer)
		if !ok {
			break
		}
	}

	return New(payload, v.Options()...)
}

func (v *vector) Even() []bool {
	booleans := make([]bool, v.length)

	for i := 1; i <= v.length; i++ {
		if i%2 == 0 {
			booleans[i-1] = true
		}
	}

	return booleans
}

func (v *vector) Odd() []bool {
	booleans := make([]bool, v.length)

	for i := 1; i <= v.length; i++ {
		if i%2 == 1 {
			booleans[i-1] = true
		}
	}

	return booleans
}

func (v *vector) Nth(nth int) []bool {
	booleans := make([]bool, v.length)

	for i := 1; i <= v.length; i++ {
		if i%nth == 0 {
			booleans[i-1] = true
		}
	}

	return booleans
}

func (v *vector) Options() []Option {
	return []Option{
		OptionVectorName(v.name),
	}
}

func (v *vector) SetOption(option Option) bool {
	if option.Key() == keyOptionVectorName {
		v.name = option.Value().(string)

		return true
	}

	if option.Key() == keyOptionMaxPrintElements {
		v.options.maxPrintElements = option.Value().(int)

		return true
	}

	return v.payload.SetOption(option.Key(), option.Value())
}

// New creates a vector part of the future vector. This function is used by public functions which create
// typed vectors
func New(payload Payload, options ...Option) Vector {
	vec := vector{
		length:  payload.Len(),
		payload: payload,
		options: vectorOptions{
			maxPrintElements: maxPrintElements,
		},
	}

	for _, option := range options {
		if option == nil {
			continue
		}

		vec.SetOption(option)
	}

	return &vec
}
