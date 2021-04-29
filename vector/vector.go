package vector

// Vector is an interface for a different vector types. This structure is similar to R-vectors: it starts from 1,
// allows for an extensive indexing, supports NA-values and named variables
type Vector interface {
	Clone() Vector
	Length() int
	ByIndices(indices []int) Vector
	ByFromTo(from int, to int) Vector
	IsEmpty() bool
}

type Reporter interface {
	Report() Report
}

type Intable interface {
	Integers() []int
}

type Floatable interface {
	Floats() []float64
}

type Booleable interface {
	Booleans() []bool
}

type Stringable interface {
	Strings() []string
}

type Markable interface {
	Marked() bool
	Mark()
}

type Refreshable interface {
	Refresh()
}

// common holds data and functions shared by all vectors
type common struct {
	vec      Vector
	length   int
	marked   bool
	report   Report
	selected []int
}

func (v *common) IsEmpty() bool {
	return v.length == 0
}

func (v *common) ByIndices(indices []int) Vector {
	selected := []int{}

	newIndex := 0
	for _, index := range indices {
		if index >= 1 && index <= v.length {
			selected = append(selected, index)
			newIndex++
		}
	}

	vec := newCommon(len(selected))
	vec.selected = selected

	return &vec
}

func (v *common) ByFromTo(from int, to int) Vector {
	/* from and to have different signs */
	if from*to < 0 {
		emp := Empty()
		emp.Report().AddError("From and to can not have different signs.")
		return emp
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

	if v.vec != nil {
		return v.vec.ByIndices(indices)
	} else {
		return v.ByIndices(indices)
	}
}

func (v *common) byFromToRegular(from, to int) []int {
	from, to = v.normalizeFromTo(from, to)

	indices := make([]int, to-from+1)
	index := 0
	for idx := from; idx <= to; idx++ {
		indices[index] = idx
		index++
	}

	return indices
}

func (v *common) byFromToReverse(from, to int) []int {
	from, to = v.normalizeFromTo(from, to)

	indices := make([]int, to-from+1)
	index := 0
	for idx := to; idx >= from; idx-- {
		indices[index] = idx
		index++
	}

	return indices
}

func (v *common) byFromToWithRemove(from, to int) []int {
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

func (v *common) normalizeFromTo(from, to int) (int, int) {
	if to > v.length {
		to = v.length
	}
	if from < 1 {
		from = 1
	}

	return from, to
}

func (v *common) Clone() Vector {
	v.marked = true

	return &common{
		vec:      nil,
		length:   v.length,
		marked:   true,
		report:   CopyReport(v.report),
		selected: v.selected,
	}
}

func (v *common) Marked() bool {
	return v.marked
}

func (v *common) Mark() {
	v.marked = true
}

func (v *common) Refresh() {
	if !v.marked || v.length == 0 {
		return
	}

	v.marked = false
}

func (v *common) Report() Report {
	return v.report
}

// Length returns length of vector
func (v *common) Length() int {
	return v.length
}

// newCommon creates a common part of the future vector. This function is used by public function which create
// typed vectors
func newCommon(length int) common {
	if length < 0 {
		length = 0
	}

	vec := common{
		vec:      nil,
		length:   length,
		marked:   false,
		report:   Report{},
		selected: []int{},
	}

	return vec
}

func newNamesAndNAble(vec Vector, config Config) (DefNameable, DefNAble) {
	nameable := DefNameable{
		vec:   vec,
		names: map[string]int{},
	}

	if config.NamesMap != nil && len(config.NamesMap) > 0 {
		for name, idx := range config.NamesMap {
			nameable.names[name] = idx
		}
	}

	na := DefNAble{
		vec: vec,
		na:  make([]bool, vec.Length()+1),
	}

	if config.NA != nil {
		if len(config.NA) == vec.Length() {
			copy(na.na[1:], config.NA)
		} else if reporter, ok := vec.(Reporter); ok {
			reporter.Report().AddWarning("Size of NA array must be equal to vector's length.")
		}
	}

	return nameable, na
}
