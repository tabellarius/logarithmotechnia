package vector

import "fmt"

// Vector is an interface for a different vector types. This structure is similar to R-vectors: it starts from 1,
// allows for an extensive indexing, supports NA-values and named variables
type Vector interface {
	Len() int
	Clone() Vector

	ByIndices(indices []int) Vector
	ByFromTo(from int, to int) []int
	ByBool(booleans []bool) []int
	ByNames(names []string) Vector
	SupportsFilter(selector interface{}) bool
	Filter(filter interface{}) []bool

	IsEmpty() bool

	Nameable
	NAble

	Intable
	Floatable
	Booleable
	Stringable

	Report() Report
	fmt.Stringer
}

type Payload interface {
	Len() int
	ByIndices(indices []int) Payload
	SupportsFilter(filter interface{}) bool
	Filter(filter interface{}) []bool
	StrForElem(idx int) string
}

type Intable interface {
	Integers() ([]int, []bool)
}

type Floatable interface {
	Floats() ([]float64, []bool)
}

type Booleable interface {
	Booleans() ([]bool, []bool)
}

type Stringable interface {
	Strings() ([]string, []bool)
}

// vector holds data and functions shared by all vectors
type vector struct {
	length int
	DefNameable
	payload Payload
	report  Report
}

// Len returns length of vector
func (v *vector) Len() int {
	return v.length
}

func (v *vector) Clone() Vector {
	return &vector{
		length: v.length,
		DefNameable: DefNameable{
			length: v.length,
			names:  v.NamesMap(),
		},
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
		length: newPayload.Len(),
		DefNameable: DefNameable{
			length: newPayload.Len(),
			names:  v.NamesForIndices(selected),
		},
		payload: v.payload.ByIndices(selected),
		report:  Report{},
	}

	return vec
}

func (v *vector) ByBool(booleans []bool) []int {
	if len(booleans) != v.length {
		v.Report().AddError("Number of booleans is not equal to vector's length.")
		return []int{}
	}

	var indices []int
	for index := 0; index < v.length; index++ {
		if booleans[index] == true {
			indices = append(indices, index+1)
		}
	}

	return indices
}

func (v *vector) ByFromTo(from int, to int) []int {
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

/* Selectors */

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

func (v *vector) normalizeFromTo(from, to int) (int, int) {
	if to > v.length {
		to = v.length
	}
	if from < 1 {
		from = 1
	}

	return from, to
}

func (v *vector) ByNames(names []string) Vector {
	indices := make([]int, 0)

	for _, name := range names {
		if idx, ok := v.names[name]; ok {
			indices = append(indices, idx)
		}
	}

	return v.ByIndices(indices)
}

func (v *vector) SupportsFilter(filter interface{}) bool {
	if _, ok := filter.(int); ok {
		return true
	}

	if _, ok := filter.([]int); ok {
		return true
	}

	return v.payload.SupportsFilter(filter)
}

func (v *vector) Filter(filter interface{}) []bool {
	if index, ok := filter.(int); ok {
		return v.selectByIndices([]int{index})
	}

	if indices, ok := filter.([]int); ok {
		return v.selectByIndices(indices)
	}

	if v.payload.SupportsFilter(filter) {
		return v.payload.Filter(filter)
	}

	return []bool{}
}

func (v *vector) selectByIndices(indices []int) []bool {
	booleans := make([]bool, v.length)

	for _, idx := range indices {
		if idx >= 1 && idx <= v.length {
			booleans[idx-1] = true
		}
	}

	return booleans
}

func (v *vector) IsNA() []bool {
	if nable, ok := v.payload.(NAble); ok {
		return nable.IsNA()
	}

	return make([]bool, v.length)
}

/* Not Applicable-related */

func (v *vector) NotNA() []bool {
	if nable, ok := v.payload.(NAble); ok {
		return nable.IsNA()
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

	if v.HasNameFor(idx) {
		str += " (" + v.Name(idx) + ")"
	}

	return str
}

func (v *vector) Strings() ([]string, []bool) {
	panic("implement me")
}

func (v *vector) Floats() ([]float64, []bool) {
	panic("implement me")
}

func (v *vector) Booleans() ([]bool, []bool) {
	panic("implement me")
}

func (v *vector) Integers() ([]int, []bool) {
	panic("implement me")
}

// New creates a vector part of the future vector. This function is used by public functions which create
// typed vectors
func New(payload Payload, options ...Config) Vector {
	config := mergeConfigs(options)

	vec := vector{
		length: payload.Len(),
		DefNameable: DefNameable{
			length: payload.Len(),
			names:  map[string]int{},
		},
		payload: payload,
		report:  Report{},
	}

	if config.NamesMap != nil {
		for name, idx := range config.NamesMap {
			if idx >= 1 && idx <= vec.length {
				vec.names[name] = idx
			}
		}
	}

	return &vec
}

//Empty creates an empty (zero-length) vector
func Empty() Vector {
	return &vector{
		length: 0,
		DefNameable: DefNameable{
			length: 0,
			names:  map[string]int{},
		},
		payload: EmptyPayload(),
		report:  Report{},
	}
}
