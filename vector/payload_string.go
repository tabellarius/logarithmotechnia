package vector

import (
	"math"
	"math/cmplx"
	"strconv"
)

type StringWhicherFunc = func(int, string, bool) bool
type StringWhicherCompactFunc = func(string, bool) bool
type StringApplierFunc = func(int, string, bool) (string, bool)
type StringApplierCompactFunc = func(string, bool) (string, bool)
type StringSummarizerFunc = func(int, string, string, bool) (string, bool)

type stringPayload struct {
	length int
	data   []string
	DefNAble
	DefArrangeable
}

func (p *stringPayload) Type() string {
	return "string"
}

func (p *stringPayload) Len() int {
	return p.length
}

func (p *stringPayload) Pick(idx int) interface{} {
	return pickValueWithNA(idx, p.data, p.na, p.length)
}

func (p *stringPayload) Data() []interface{} {
	return dataWithNAToInterfaceArray(p.data, p.na)
}

func (p *stringPayload) ByIndices(indices []int) Payload {
	data, na := byIndices(indices, p.data, p.na, "")

	return StringPayload(data, na, p.Options()...)
}

func (p *stringPayload) SupportsWhicher(whicher any) bool {
	return supportsWhicher[string](whicher)
}

func (p *stringPayload) Which(whicher any) []bool {
	return which(p.data, p.na, whicher)
}

func (p *stringPayload) SupportsApplier(applier any) bool {
	return supportsApplier[string](applier)
}

func (p *stringPayload) Apply(applier any) Payload {
	data, na := apply(p.data, p.na, applier, "")

	if data == nil {
		return NAPayload(p.length)
	}

	return StringPayload(data, na, p.Options()...)
}

func (p *stringPayload) ApplyTo(indices []int, applier any) Payload {
	data, na := applyTo(indices, p.data, p.na, applier, "")

	if data == nil {
		return NAPayload(p.length)
	}

	return StringPayload(data, na, p.Options()...)
}

func (p *stringPayload) SupportsSummarizer(summarizer any) bool {
	return supportsSummarizer[string](summarizer)
}

func (p *stringPayload) Summarize(summarizer interface{}) Payload {
	val, na := summarize(p.data, p.na, summarizer, "", "")

	return StringPayload([]string{val}, []bool{na}, p.Options()...)
}

func (p *stringPayload) Integers() ([]int, []bool) {
	if p.length == 0 {
		return []int{}, []bool{}
	}

	data := make([]int, p.length)
	na := make([]bool, p.Len())
	copy(na, p.na)

	for i := 0; i < p.length; i++ {
		if p.na[i] {
			data[i] = 0
		} else {
			num, err := strconv.Atoi(p.data[i])
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

func (p *stringPayload) Floats() ([]float64, []bool) {
	if p.length == 0 {
		return []float64{}, []bool{}
	}

	data := make([]float64, p.length)
	na := make([]bool, p.Len())
	copy(na, p.na)

	for i := 0; i < p.length; i++ {
		if p.na[i] {
			data[i] = math.NaN()
		} else {
			num, err := strconv.ParseFloat(p.data[i], 64)
			if err != nil {
				data[i] = math.NaN()
				na[i] = true
			} else {
				data[i] = num
			}
		}
	}

	return data, na
}

func (p *stringPayload) Complexes() ([]complex128, []bool) {
	if p.length == 0 {
		return []complex128{}, []bool{}
	}

	data := make([]complex128, p.length)
	na := make([]bool, p.Len())
	copy(na, p.na)

	for i := 0; i < p.length; i++ {
		if p.na[i] {
			data[i] = cmplx.NaN()
		} else {
			num, err := strconv.ParseComplex(p.data[i], 128)
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

func (p *stringPayload) Booleans() ([]bool, []bool) {
	if p.length == 0 {
		return []bool{}, []bool{}
	}

	data := make([]bool, p.length)

	for i := 0; i < p.length; i++ {
		if p.na[i] {
			data[i] = false
		} else {
			data[i] = p.data[i] != ""
		}
	}

	na := make([]bool, p.length)
	copy(na, p.na)

	return data, na
}

func (p *stringPayload) Strings() ([]string, []bool) {
	if p.length == 0 {
		return []string{}, []bool{}
	}

	data := make([]string, p.length)
	copy(data, p.data)

	na := make([]bool, p.Len())
	copy(na, p.na)

	return data, na
}

func (p *stringPayload) Interfaces() ([]interface{}, []bool) {
	if p.length == 0 {
		return []interface{}{}, []bool{}
	}

	data := make([]interface{}, p.length)
	for i := 0; i < p.length; i++ {
		if p.na[i] {
			data[i] = nil
		} else {
			data[i] = p.data[i]
		}
	}

	na := make([]bool, p.length)
	copy(na, p.na)

	return data, na
}

func (p *stringPayload) Append(payload Payload) Payload {
	length := p.length + payload.Len()

	var vals []string
	var na []bool

	if stringable, ok := payload.(Stringable); ok {
		vals, na = stringable.Strings()
	} else {
		vals, na = NAPayload(payload.Len()).(Stringable).Strings()
	}

	newVals := make([]string, length)
	newNA := make([]bool, length)

	copy(newVals, p.data)
	copy(newVals[p.length:], vals)
	copy(newNA, p.na)
	copy(newNA[p.length:], na)

	return StringPayload(newVals, newNA, p.Options()...)
}

func (p *stringPayload) Groups() ([][]int, []interface{}) {
	groupMap := map[string][]int{}
	ordered := []string{}
	na := []int{}

	for i, val := range p.data {
		idx := i + 1

		if p.na[i] {
			na = append(na, idx)
			continue
		}

		if _, ok := groupMap[val]; !ok {
			groupMap[val] = []int{}
			ordered = append(ordered, val)
		}

		groupMap[val] = append(groupMap[val], idx)
	}

	groups := make([][]int, len(ordered))
	for i, val := range ordered {
		groups[i] = groupMap[val]
	}

	if len(na) > 0 {
		groups = append(groups, na)
	}

	values := make([]interface{}, len(groups))
	for i, val := range ordered {
		values[i] = interface{}(val)
	}
	if len(na) > 0 {
		values[len(values)-1] = nil
	}

	return groups, values
}

func (p *stringPayload) StrForElem(idx int) string {
	if p.na[idx-1] {
		return "NA"
	}

	return "\"" + p.data[idx-1] + "\""
}

func (p *stringPayload) Adjust(size int) Payload {
	if size < p.length {
		return p.adjustToLesserSize(size)
	}

	if size > p.length {
		return p.adjustToBiggerSize(size)
	}

	return p
}

func (p *stringPayload) adjustToLesserSize(size int) Payload {
	data := make([]string, size)
	na := make([]bool, size)

	copy(data, p.data)
	copy(na, p.na)

	return StringPayload(data, na, p.Options()...)
}

func (p *stringPayload) adjustToBiggerSize(size int) Payload {
	cycles := size / p.length
	if size%p.length > 0 {
		cycles++
	}

	data := make([]string, cycles*p.length)
	na := make([]bool, cycles*p.length)

	for i := 0; i < cycles; i++ {
		copy(data[i*p.length:], p.data)
		copy(na[i*p.length:], p.na)
	}

	data = data[:size]
	na = na[:size]

	return StringPayload(data, na, p.Options()...)
}

/* Finder interface */

func (p *stringPayload) Find(needle interface{}) int {
	val, ok := needle.(string)
	if !ok {
		return 0
	}

	for i, datum := range p.data {
		if !p.na[i] && val == datum {
			return i + 1
		}
	}

	return 0
}

func (p *stringPayload) FindAll(needle interface{}) []int {
	val, ok := needle.(string)
	if !ok {
		return []int{}
	}

	found := []int{}
	for i, datum := range p.data {
		if !p.na[i] && val == datum {
			found = append(found, i+1)
		}
	}

	return found
}

func (p *stringPayload) Eq(val interface{}) []bool {
	return eq(val, p.data, p.na, p.convertComparator)
}

func (p *stringPayload) Neq(val interface{}) []bool {
	return neq(val, p.data, p.na, p.convertComparator)
}

func (p *stringPayload) Gt(val interface{}) []bool {
	return gt(val, p.data, p.na, p.convertComparator)
}

func (p *stringPayload) Lt(val interface{}) []bool {
	return lt(val, p.data, p.na, p.convertComparator)
}

func (p *stringPayload) Gte(val interface{}) []bool {
	return gte(val, p.data, p.na, p.convertComparator)
}

func (p *stringPayload) Lte(val interface{}) []bool {
	return lte(val, p.data, p.na, p.convertComparator)
}

func (p *stringPayload) convertComparator(val interface{}) (string, bool) {
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

func (p *stringPayload) IsUnique() []bool {
	booleans := make([]bool, p.length)

	valuesMap := map[string]bool{}
	wasNA := false
	for i := 0; i < p.length; i++ {
		is := false

		if p.na[i] {
			if !wasNA {
				is = true
				wasNA = true
			}
		} else {
			if _, ok := valuesMap[p.data[i]]; !ok {
				is = true
				valuesMap[p.data[i]] = true
			}
		}

		booleans[i] = is
	}

	return booleans
}

func (p *stringPayload) Coalesce(payload Payload) Payload {
	if p.length != payload.Len() {
		payload = payload.Adjust(p.length)
	}

	var srcData []string
	var srcNA []bool

	if same, ok := payload.(*stringPayload); ok {
		srcData = same.data
		srcNA = same.na
	} else if stringable, ok := payload.(Stringable); ok {
		srcData, srcNA = stringable.Strings()
	} else {
		return p
	}

	dstData := make([]string, p.length)
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

	return StringPayload(dstData, dstNA, p.Options()...)
}

func (p *stringPayload) Options() []Option {
	return []Option{}
}

func StringPayload(data []string, na []bool, _ ...Option) Payload {
	length := len(data)

	vecNA := make([]bool, length)
	if len(na) > 0 {
		if len(na) == length {
			copy(vecNA, na)
		} else {
			emp := NAPayload(0)
			return emp
		}
	}

	vecData := make([]string, length)
	for i := 0; i < length; i++ {
		if vecNA[i] {
			vecData[i] = ""
		} else {
			vecData[i] = data[i]
		}
	}

	payload := &stringPayload{
		length: length,
		data:   vecData,
		DefNAble: DefNAble{
			na: vecNA,
		},
	}

	payload.DefArrangeable = DefArrangeable{
		Length:   payload.length,
		DefNAble: payload.DefNAble,
		FnLess: func(i, j int) bool {
			return payload.data[i] < payload.data[j]
		},
		FnEqual: func(i, j int) bool {
			return payload.data[i] == payload.data[j]
		},
	}

	return payload
}

func StringWithNA(data []string, na []bool, options ...Option) Vector {
	return New(StringPayload(data, na, options...), options...)
}

func String(data []string, options ...Option) Vector {
	return StringWithNA(data, nil, options...)
}
