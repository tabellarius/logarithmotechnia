package vector

import (
	"math"
	"math/cmplx"
	"strconv"
)

type FactorWhicherFunc = func(int, string, bool) bool
type FactorWhicherCompactFunc = func(string, bool) bool
type FactorToStringApplierFunc = func(int, string, bool) (string, bool)
type FactorToStringApplierCompactFunc = func(string, bool) (string, bool)

type factorPayload struct {
	length int
	levels []string
	data   []uint32
}

func (p *factorPayload) Type() string {
	return "factor"
}

func (p *factorPayload) Len() int {
	return p.length
}

func (p *factorPayload) ByIndices(indices []int) Payload {
	data := make([]uint32, 0, len(indices))

	for i, idx := range indices {
		if idx == 0 {
			data[i] = 0
		} else {
			data[i] = p.data[idx-1]
		}
	}

	return factorPayloadFromFactorData(data, p.levels, p.Options()...)
}

func (p *factorPayload) Adjust(size int) Payload {
	if size < p.length {
		return p.adjustToLesserSize(size)
	}

	if size > p.length {
		return p.adjustToBiggerSize(size)
	}

	return p
}

func (p *factorPayload) adjustToLesserSize(size int) Payload {
	data := make([]uint32, size)

	for i := 0; i < size; i++ {
		data[i] = p.data[i]
	}

	return factorPayloadFromFactorData(data, p.levels, p.Options()...)
}

func (p *factorPayload) adjustToBiggerSize(size int) Payload {
	cycles := size / p.length
	if size%p.length > 0 {
		cycles++
	}

	data := make([]uint32, cycles*p.length)

	for i := 0; i < cycles; i++ {
		copy(data[i*p.length:], p.data)
	}

	data = data[:size]

	return factorPayloadFromFactorData(data, p.levels, p.Options()...)
}

func (p *factorPayload) Append(payload Payload) Payload {
	length := p.length + payload.Len()

	curVals, curNA := p.Strings()

	var appVals []string
	var appNA []bool

	if stringable, ok := payload.(Stringable); ok {
		appVals, appNA = stringable.Strings()
	} else {
		appVals, appNA = NAPayload(payload.Len()).(Stringable).Strings()
	}

	newVals := make([]string, length)
	newNA := make([]bool, length)

	copy(newVals, curVals)
	copy(newVals[p.length:], appVals)
	copy(newNA, curNA)
	copy(newNA[p.length:], appNA)

	return FactorPayload(newVals, newNA, p.Options()...)
}

func (p *factorPayload) SupportsWhicher(whicher interface{}) bool {
	if _, ok := whicher.(FactorWhicherFunc); ok {
		return true
	}

	if _, ok := whicher.(FactorWhicherCompactFunc); ok {
		return true
	}

	return false
}

func (p *factorPayload) Which(whicher interface{}) []bool {
	if byFunc, ok := whicher.(FactorWhicherFunc); ok {
		return p.selectByFunc(byFunc)
	}

	if byFunc, ok := whicher.(FactorWhicherCompactFunc); ok {
		return p.selectByCompactFunc(byFunc)
	}

	return make([]bool, p.length)
}

func (p *factorPayload) selectByFunc(byFunc FactorWhicherFunc) []bool {
	booleans := make([]bool, p.length)

	for idx, level := range p.data {
		if byFunc(idx+1, p.levels[level], level == 0) {
			booleans[idx] = true
		}
	}

	return booleans
}

func (p *factorPayload) selectByCompactFunc(byFunc FactorWhicherCompactFunc) []bool {
	booleans := make([]bool, p.length)

	for idx, level := range p.data {
		if byFunc(p.levels[level], level == 0) {
			booleans[idx] = true
		}
	}

	return booleans
}

func (p *factorPayload) SupportsApplier(applier interface{}) bool {
	if _, ok := applier.(FactorToStringApplierFunc); ok {
		return true
	}

	if _, ok := applier.(FactorToStringApplierCompactFunc); ok {
		return true
	}

	return false
}

func (p *factorPayload) Apply(applier interface{}) Payload {
	if applyFunc, ok := applier.(FactorToStringApplierFunc); ok {
		return p.applyToStringByFunc(applyFunc)
	}

	if applyFunc, ok := applier.(FactorToStringApplierCompactFunc); ok {
		return p.applyToStringByCompactFunc(applyFunc)
	}

	return NAPayload(p.length)
}

func (p *factorPayload) applyToStringByFunc(applyFunc StringToStringApplierFunc) Payload {
	data := make([]string, p.length)
	na := make([]bool, p.length)

	for i := 0; i < p.length; i++ {
		dataVal, naVal := applyFunc(i+1, p.levels[p.data[i]], p.data[i] == 0)
		if naVal {
			dataVal = ""
		}
		data[i] = dataVal
		na[i] = naVal
	}

	return StringPayload(data, na, p.Options()...)
}

func (p *factorPayload) applyToStringByCompactFunc(applyFunc StringToStringApplierCompactFunc) Payload {
	data := make([]string, p.length)
	na := make([]bool, p.length)

	for i := 0; i < p.length; i++ {
		dataVal, naVal := applyFunc(p.levels[p.data[i]], p.data[i] == 0)
		if naVal {
			dataVal = ""
		}
		data[i] = dataVal
		na[i] = naVal
	}

	return StringPayload(data, na, p.Options()...)
}

func (p *factorPayload) Integers() ([]int, []bool) {
	if p.length == 0 {
		return []int{}, []bool{}
	}

	data := make([]int, p.length)
	na := make([]bool, p.Len())

	for i := 0; i < p.length; i++ {
		if p.data[i] == 0 {
			data[i] = 0
			na[i] = true
		} else {
			num, err := strconv.ParseFloat(p.levels[p.data[i]], 64)
			if err != nil {
				data[i] = 0
				na[i] = true
			} else {
				data[i] = int(num)
			}
		}
	}

	return data, na
}

func (p *factorPayload) Floats() ([]float64, []bool) {
	if p.length == 0 {
		return []float64{}, []bool{}
	}

	data := make([]float64, p.length)
	na := make([]bool, p.Len())

	for i := 0; i < p.length; i++ {
		if p.data[i] == 0 {
			data[i] = math.NaN()
			na[i] = true
		} else {
			num, err := strconv.ParseFloat(p.levels[p.data[i]], 64)
			if err != nil {
				data[i] = 0
				na[i] = true
			} else {
				data[i] = num
			}
		}
	}

	return data, na
}

func (p *factorPayload) Complexes() ([]complex128, []bool) {
	if p.length == 0 {
		return []complex128{}, []bool{}
	}

	data := make([]complex128, p.length)
	na := make([]bool, p.Len())

	for i := 0; i < p.length; i++ {
		if p.data[i] == 0 {
			data[i] = cmplx.NaN()
			na[i] = true
		} else {
			num, err := strconv.ParseComplex(p.levels[p.data[i]], 128)
			if err != nil {
				data[i] = cmplx.NaN()
				na[i] = true
			} else {
				data[i] = num
			}
		}
	}

	return data, na
}

func (p *factorPayload) Booleans() ([]bool, []bool) {
	if p.length == 0 {
		return []bool{}, []bool{}
	}

	data := make([]bool, p.length)
	na := make([]bool, p.length)

	for i := 0; i < p.length; i++ {
		if p.data[i] == 0 {
			data[i] = false
			na[i] = true
		} else {
			data[i] = p.levels[p.data[i]] != ""
		}
	}

	return data, na
}

func (p *factorPayload) Strings() ([]string, []bool) {
	if p.length == 0 {
		return []string{}, []bool{}
	}

	data := make([]string, p.length)
	na := make([]bool, p.length)

	for i := 0; i < p.length; i++ {
		if p.data[i] == 0 {
			data[i] = ""
			na[i] = true
		} else {
			data[i] = p.levels[p.data[i]]
		}
	}

	return data, na
}

func (p *factorPayload) Interfaces() ([]interface{}, []bool) {
	if p.length == 0 {
		return []interface{}{}, []bool{}
	}

	data := make([]interface{}, p.length)
	na := make([]bool, p.length)

	for i := 0; i < p.length; i++ {
		if p.data[i] == 0 {
			data[i] = nil
			na[i] = true
		} else {
			data[i] = p.data[i]
		}
	}

	return data, na
}

func (p *factorPayload) IsNA() []bool {
	isNA := make([]bool, p.length)

	for i, val := range p.data {
		isNA[i] = val == 0
	}

	return isNA
}

func (p *factorPayload) NotNA() []bool {
	notNA := make([]bool, p.length)

	for i, val := range p.data {
		notNA[i] = val != 0
	}

	return notNA
}

func (p *factorPayload) HasNA() bool {
	for _, val := range p.data {
		if val == 0 {
			return true
		}
	}

	return false
}

func (p *factorPayload) WithNA() []int {
	naIndices := make([]int, 0)

	for i := 0; i < p.length; i++ {
		if p.data[i] == 0 {
			naIndices = append(naIndices, i+1)
		}
	}

	return naIndices
}

func (p *factorPayload) WithoutNA() []int {
	naIndices := make([]int, 0)

	for i := 0; i < p.length; i++ {
		if p.data[i] != 0 {
			naIndices = append(naIndices, i+1)
		}
	}

	return naIndices
}

func (p *factorPayload) Find(needle interface{}) int {
	val, ok := needle.(string)
	if !ok {
		return 0
	}

	if p.length == 0 {
		return 0
	}

	valLevel := uint32(0)
	for i := 1; i < len(p.levels); i++ {
		if val == p.levels[i] {
			valLevel = uint32(i)
		}
	}

	if valLevel == 0 {
		return 0
	}

	for i, level := range p.data {
		if level > 0 && level == valLevel {
			return i + 1
		}
	}

	return 0
}

func (p *factorPayload) FindAll(needle interface{}) []int {
	val, ok := needle.(string)
	if !ok {
		return []int{}
	}

	if p.length == 0 {
		return []int{}
	}

	valLevel := uint32(p.Level(val))

	if valLevel == 0 {
		return []int{}
	}

	found := []int{}
	for i, level := range p.data {
		if level > 0 && level == valLevel {
			found = append(found, i+1)
		}
	}

	return found
}

func (p *factorPayload) Eq(val interface{}) []bool {
	cmp := make([]bool, p.length)

	v, ok := p.convertComparator(val)
	if !ok {
		return cmp
	}

	valLevel := uint32(p.Level(v))

	for i, level := range p.data {
		if level == 0 {
			cmp[i] = false
		} else {
			cmp[i] = level == valLevel
		}
	}

	return cmp
}

func (p *factorPayload) Neq(val interface{}) []bool {
	cmp := make([]bool, p.length)

	v, ok := p.convertComparator(val)
	if !ok {
		for i := range p.data {
			cmp[i] = true
		}

		return cmp
	}

	valLevel := uint32(p.Level(v))

	for i, level := range p.data {
		if level == 0 {
			cmp[i] = false
		} else {
			cmp[i] = level != valLevel
		}
	}

	return cmp
}

func (p *factorPayload) Gt(val interface{}) []bool {
	cmp := make([]bool, p.length)

	v, ok := p.convertComparator(val)
	if !ok {
		return cmp
	}

	for i, level := range p.data {
		if level == 0 {
			cmp[i] = false
		} else {
			cmp[i] = p.levels[level] > v
		}
	}

	return cmp
}

func (p *factorPayload) Lt(val interface{}) []bool {
	cmp := make([]bool, p.length)

	v, ok := p.convertComparator(val)
	if !ok {
		return cmp
	}

	for i, level := range p.data {
		if level == 0 {
			cmp[i] = false
		} else {
			cmp[i] = p.levels[level] < v
		}
	}

	return cmp
}

func (p *factorPayload) Gte(val interface{}) []bool {
	cmp := make([]bool, p.length)

	v, ok := p.convertComparator(val)
	if !ok {
		return cmp
	}

	for i, level := range p.data {
		if level == 0 {
			cmp[i] = false
		} else {
			cmp[i] = p.levels[level] >= v
		}
	}

	return cmp
}

func (p *factorPayload) Lte(val interface{}) []bool {
	cmp := make([]bool, p.length)

	v, ok := p.convertComparator(val)
	if !ok {
		return cmp
	}

	for i, level := range p.data {
		if level == 0 {
			cmp[i] = false
		} else {
			cmp[i] = p.levels[level] <= v
		}
	}

	return cmp
}

func (p *factorPayload) convertComparator(val interface{}) (string, bool) {
	var v string
	ok := true
	switch val.(type) {
	case string:
		v = val.(string)
	default:
		ok = false
	}

	return v, ok
}

func (p *factorPayload) Groups() ([][]int, []interface{}) {
	panic("implement me")
}

func (p *factorPayload) SortedIndices() []int {
	panic("implement me")
}

func (p *factorPayload) SortedIndicesWithRanks() ([]int, []int) {
	panic("implement me")
}

func (p *factorPayload) IsUnique() []bool {
	panic("implement me")
}

func (p *factorPayload) Coalesce(payload Payload) Payload {
	panic("implement me")
}

func (p *factorPayload) Level(val string) int {
	level := 0

	for i := 1; i < len(p.levels); i++ {
		if val == p.levels[i] {
			level = i
		}
	}

	return level
}

func (p *factorPayload) StrForElem(idx int) string {
	if p.data[idx-1] == 0 {
		return "NA"
	}

	return "\"" + p.levels[p.data[idx-1]] + "\""
}

func (p *factorPayload) Options() []Option {
	return []Option{}
}

func factorPayloadFromFactorData(data []uint32, levels []string, _ ...Option) Payload {
	payload := &factorPayload{
		length: len(data),
		levels: levels,
		data:   data,
	}

	return payload
}

func FactorPayload(data []string, na []bool, options ...Option) Payload {
	length := len(data)

	if na != nil && len(na) != length {
		emp := NAPayload(0)
		return emp
	}

	vecData := make([]uint32, length)
	vecLevels := make([]string, 1)
	levelsMap := map[string]uint32{}

	for i := 0; i < length; i++ {
		if na != nil && na[i] {
			vecData[i] = 0
		} else {
			if level, ok := levelsMap[data[i]]; ok {
				vecData[i] = level
			} else {
				newLevel := uint32(len(vecLevels))
				vecLevels = append(vecLevels, data[i])
				levelsMap[data[i]] = newLevel
				vecData[i] = newLevel
			}
		}
	}

	return factorPayloadFromFactorData(vecData, vecLevels, options...)
}