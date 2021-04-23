package vector

// Vector is an interface for a different vector types. This structure is similar to R-vectors: it starts from 1,
// allows for an extensive indexing, supports NA-values and named variables
type Vector interface {
	Clone() Vector
	Length() int
	ByIndex(indices []int) Vector
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
	return v.length > 0
}

func (v *common) ByIndex(indices []int) Vector {
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

	if from > 0 && from < to {
		return v.byFromToReverse(to, from)
	}

	if from < 0 && to < 0 {
		return v.byFromToWithRemove(from, to)
	}

	return v.byFromToRegular(from, to)
}

func (v *common) byFromToRegular(from, to int) Vector {
	if to > v.length {
		to = v.length
	}
	if from < 1 {
		from = 1
	}

	indices := make([]int, from-to+1)
	index := 0
	for idx := from; idx <= to; idx++ {
		indices[index] = idx
		index++
	}

	return v.vec.ByIndex(indices)
}

func (v *common) byFromToReverse(from, to int) Vector {

}

func (v *common) byFromToWithRemove(from, to int) Vector {

}

func (v *common) Clone() Vector {
	v.marked = true

	return &common{
		length: v.length,
		marked: true,
		report: CopyReport(v.report),
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
func newCommon(length int, options ...Config) common {
	vec := common{
		length:   length,
		marked:   false,
		report:   Report{},
		selected: []int{},
	}

	return vec
}

type Config struct {
	Names    []string
	NamesMap map[int]string
	NA       []bool
	NAMap    map[int]bool
}

func NA(na []bool) Config {
	return Config{NA: na}
}

func Names(names []string) Config {
	return Config{Names: names}
}

func NamesMap(namesMap map[int]string) Config {
	return Config{NamesMap: namesMap}
}

func mergeConfigs(configs []Config) Config {
	config := Config{}

	for _, c := range configs {
		if c.Names != nil {
			config.Names = c.Names
		}
		if c.NamesMap != nil {
			config.NamesMap = c.NamesMap
		}
		if c.NA != nil {
			config.NA = c.NA
		}
		if c.NAMap != nil {
			config.NAMap = c.NAMap
		}
	}

	return config
}
