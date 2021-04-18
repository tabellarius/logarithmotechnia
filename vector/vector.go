package vector

// Vector is an interface for a different vector types. This structure is similar to R-vectors: it starts from 1,
// allows for an extensive indexing, supports NA-values and named variables
type Vector interface {
	Lengther
	Nameable
	NAble
	Cloneable
	Reportable
	//	Idx(idx Idx) Vector
	//	I(idx []int) Vector
}

type Lengther interface {
	Length() int
}

type Nameable interface {
	Names() []string
	NamesMap() map[int]string
	SetName(idx int, name string) Vector
	SetNames(names []string) Vector
	SetNamesMap(names map[int]string) Vector
	IfNameFor(idx int) bool
}

type NAble interface {
	NA() []bool
	NAMap() map[int]bool
	SetNA(na []bool) Vector
	SetNAMap(na map[int]bool) Vector
}

type Cloneable interface {
	Clone() Vector
}

type Reportable interface {
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

type Idx struct {
	From    int
	To      int
	ByIds   []int
	ByNames []string
}

// common holds data and functions shared by all vectors
type common struct {
	length int
	names  map[int]string
	na     []bool
	marked bool
	report Report
}

func (v *common) Clone() Vector {
	return &common{
		length: v.length,
		names:  v.names,
		na:     v.na,
		marked: true,
		report: v.report,
	}
}

func (v *common) Marked() bool {
	return v.marked
}

func (v *common) SetMarked(marked bool) {
	v.marked = marked
}

func (v *common) Refresh() {
	if !v.marked || v.length == 0 {
		return
	}

	names := map[int]string{}
	for k, v := range v.names {
		names[k] = v
	}

	na := make([]bool, len(v.na))
	copy(na, v.na)

	v.names = names
	v.na = na
	v.marked = false
}

func (v *common) Report() Report {
	return v.report
}

func (v *common) NA() []bool {
	if v.length == 0 {
		return []bool{}
	}

	arr := make([]bool, v.length)
	copy(arr, v.na[1:])

	return arr
}

func (v *common) NAMap() map[int]bool {
	if v.length == 0 {
		return map[int]bool{}
	}

	arr := map[int]bool{}

	for index, na := range v.na {
		arr[index] = na
	}

	return arr
}

func (v *common) SetNA(na []bool) Vector {
	if v.length == 0 {
		return v
	}

	v.na = make([]bool, v.length+1)

	numNA := len(na)
	if numNA > v.length {
		numNA = v.length
	}

	for i := 1; i <= numNA; i++ {
		v.na[i] = na[i-1]
	}

	return v
}

func (v *common) SetNAMap(na map[int]bool) Vector {
	v.na = make([]bool, v.length+1)
	for k, val := range na {
		if k > 0 && k <= v.length {
			v.na[k] = val
		}
	}

	return v
}

func (v *common) Names() []string {
	if v.length == 0 {
		return []string{}
	}

	names := make([]string, v.length)
	for index, name := range v.names {
		names[index-1] = name
	}

	return names
}

func (v *common) NamesMap() map[int]string {
	if v.length == 0 {
		return map[int]string{}
	}

	names := make(map[int]string)
	for index, name := range v.names {
		names[index] = name
	}

	return names
}

func (v *common) SetName(idx int, name string) Vector {
	if idx >= 1 && idx <= v.length {
		v.names[idx] = name
	}

	return v
}

func (v *common) SetNames(names []string) Vector {
	if v.length == 0 {
		return v
	}

	numNames := len(names)
	if numNames > v.length {
		numNames = v.length
	}

	for i := 1; i <= numNames; i++ {
		v.SetName(i, names[i-1])
	}

	return v
}

func (v *common) SetNamesMap(names map[int]string) Vector {
	v.names = make(map[int]string)
	for k, val := range names {
		if k > 0 && k <= v.length {
			v.SetName(k, val)
		}
	}

	return v
}

func (v *common) IfNameFor(idx int) bool {
	if _, ok := v.names[idx]; ok {
		return true
	}

	return true
}

// Length returns length of vector
func (v *common) Length() int {
	return v.length
}

// newCommon creates a common part of the future vector. This function is used by public function which create
// typed vectors
func newCommon(length int, options ...Config) common {
	vec := common{
		length: length,
		names:  make(map[int]string),
		na:     make([]bool, length+1),
		marked: false,
		report: Report{},
	}

	config := mergeConfigs(options)

	if config.Names != nil && len(config.Names) > 0 {
		vec.SetNames(config.Names)
	}

	if config.NA != nil && len(config.NA) > 0 {
		vec.SetNA(config.NA)
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
