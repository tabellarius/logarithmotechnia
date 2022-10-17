package vector

import (
	"math"
	"math/cmplx"
)

type BooleanWhicherFunc = func(int, bool, bool) bool
type BooleanWhicherCompactFunc = func(bool, bool) bool
type BooleanApplierFunc = func(int, bool, bool) (bool, bool)
type BooleanApplierCompactFunc = func(bool, bool) (bool, bool)
type BooleanSummarizerFunc = func(int, bool, bool, bool) (bool, bool)

type booleanPayload struct {
	length int
	data   []bool
	DefNAble
	DefArrangeable
}

func (p *booleanPayload) Type() string {
	return "boolean"
}

func (p *booleanPayload) Len() int {
	return p.length
}

func (p *booleanPayload) Pick(idx int) interface{} {
	return pickValueWithNA(idx, p.data, p.na, p.length)
}

func (p *booleanPayload) Data() []interface{} {
	return dataWithNAToInterfaceArray(p.data, p.na)
}

func (p *booleanPayload) ByIndices(indices []int) Payload {
	data, na := byIndices(indices, p.data, p.na, false)

	return BooleanPayload(data, na, p.Options()...)
}

func (p *booleanPayload) SupportsWhicher(whicher any) bool {
	return supportsWhicher[bool](whicher)
}

func (p *booleanPayload) Which(whicher any) []bool {
	return which(p.data, p.na, whicher)
}

func (p *booleanPayload) SupportsApplier(applier any) bool {
	return supportApplier[bool](applier)
}

func (p *booleanPayload) Apply(applier any) Payload {
	if applyFunc, ok := applier.(BooleanApplierFunc); ok {
		data, na := applyByFunc(p.data, p.na, p.length, applyFunc, false)

		return BooleanPayload(data, na, p.Options()...)
	}

	if applyFunc, ok := applier.(BooleanApplierCompactFunc); ok {
		data, na := applyByCompactFunc(p.data, p.na, p.length, applyFunc, false)

		return BooleanPayload(data, na, p.Options()...)
	}

	return NAPayload(p.length)
}

func (p *booleanPayload) ApplyTo(indices []int, applier interface{}) Payload {
	if applyFunc, ok := applier.(BooleanApplierFunc); ok {
		return p.applyToByFunc(indices, applyFunc)
	}

	if applyFunc, ok := applier.(BooleanApplierCompactFunc); ok {
		return p.applyToByCompactFunc(indices, applyFunc)
	}

	return NAPayload(p.length)
}

func (p *booleanPayload) applyToByFunc(indices []int, applyFunc BooleanApplierFunc) Payload {
	data, na := applyToByFunc(indices, p.data, p.na, applyFunc, false)

	return BooleanPayload(data, na, p.Options()...)
}

func (p *booleanPayload) applyToByCompactFunc(indices []int, applyFunc BooleanApplierCompactFunc) Payload {
	data, na := applyToByCompactFunc(indices, p.data, p.na, applyFunc, false)

	return BooleanPayload(data, na, p.Options()...)
}

func (p *booleanPayload) SupportsSummarizer(summarizer interface{}) bool {
	if _, ok := summarizer.(BooleanSummarizerFunc); ok {
		return true
	}

	return false
}

func (p *booleanPayload) Summarize(summarizer interface{}) Payload {
	fn, ok := summarizer.(BooleanSummarizerFunc)
	if !ok {
		return NAPayload(1)
	}

	val := false
	na := false
	for i := 0; i < p.length; i++ {
		val, na = fn(i+1, val, p.data[i], p.na[i])
		if na {
			return NAPayload(1)
		}
	}

	return BooleanPayload([]bool{val}, nil, p.Options()...)
}

func (p *booleanPayload) Integers() ([]int, []bool) {
	if p.length == 0 {
		return []int{}, []bool{}
	}

	data := make([]int, p.length)
	for i := 0; i < p.length; i++ {
		if p.na[i] {
			data[i] = 0
		} else {
			if p.data[i] {
				data[i] = 1
			} else {
				data[i] = 0
			}
		}
	}

	na := make([]bool, p.Len())
	copy(na, p.na)

	return data, na
}

func (p *booleanPayload) Floats() ([]float64, []bool) {
	if p.length == 0 {
		return []float64{}, []bool{}
	}

	data := make([]float64, p.length)

	for i := 0; i < p.length; i++ {
		if p.na[i] {
			data[i] = math.NaN()
		} else {
			if p.data[i] {
				data[i] = 1
			} else {
				data[i] = 0
			}
		}
	}

	na := make([]bool, p.Len())
	copy(na, p.na)

	return data, na
}

func (p *booleanPayload) Complexes() ([]complex128, []bool) {
	if p.length == 0 {
		return []complex128{}, []bool{}
	}

	data := make([]complex128, p.length)
	for i := 0; i < p.length; i++ {
		if p.na[i] {
			data[i] = cmplx.NaN()
		} else {
			if p.data[i] {
				data[i] = 1
			} else {
				data[i] = 0
			}
		}
	}

	na := make([]bool, p.Len())
	copy(na, p.na)

	return data, na
}

func (p *booleanPayload) Booleans() ([]bool, []bool) {
	if p.length == 0 {
		return []bool{}, []bool{}
	}

	data := make([]bool, p.length)
	copy(data, p.data)

	na := make([]bool, p.length)
	copy(na, p.na)

	return data, na
}

func (p *booleanPayload) Strings() ([]string, []bool) {
	if p.length == 0 {
		return []string{}, []bool{}
	}

	data := make([]string, p.length)

	for i := 0; i < p.length; i++ {
		data[i] = p.StrForElem(i + 1)
	}

	na := make([]bool, p.Len())
	copy(na, p.na)

	return data, na
}

func (p *booleanPayload) Interfaces() ([]interface{}, []bool) {
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

	na := make([]bool, p.Len())
	copy(na, p.na)

	return data, na
}

func (p *booleanPayload) Append(payload Payload) Payload {
	length := p.length + payload.Len()

	var vals []bool
	var na []bool

	if boolable, ok := payload.(Boolable); ok {
		vals, na = boolable.Booleans()
	} else {
		vals, na = NAPayload(payload.Len()).(Boolable).Booleans()
	}

	newVals := make([]bool, length)
	newNA := make([]bool, length)

	copy(newVals, p.data)
	copy(newVals[p.length:], vals)
	copy(newNA, p.na)
	copy(newNA[p.length:], na)

	return BooleanPayload(newVals, newNA)
}

func (p *booleanPayload) Adjust(size int) Payload {
	if size < p.length {
		return p.adjustToLesserSize(size)
	}

	if size > p.length {
		return p.adjustToBiggerSize(size)
	}

	return p
}

func (p *booleanPayload) adjustToLesserSize(size int) Payload {
	data, na := adjustToLesserSizeWithNA(p.data, p.na, size)

	return BooleanPayload(data, na)
}

func (p *booleanPayload) adjustToBiggerSize(size int) Payload {
	data, na := adjustToBiggerSizeWithNA(p.data, p.na, p.length, size)

	return BooleanPayload(data, na)
}

func (p *booleanPayload) Groups() ([][]int, []interface{}) {
	groups, values := groupsForData(p.data, p.na)

	return groups, values
}

func (p *booleanPayload) StrForElem(idx int) string {
	if p.na[idx-1] {
		return "NA"
	}

	if p.data[idx-1] {
		return "true"
	}

	return "false"
}

func (p *booleanPayload) Find(needle interface{}) int {
	val, ok := needle.(bool)
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

func (p *booleanPayload) FindAll(needle interface{}) []int {
	val, ok := needle.(bool)
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

/* Finder interface */

func (p *booleanPayload) Eq(val interface{}) []bool {
	cmp := make([]bool, p.length)

	v, ok := val.(bool)
	if !ok {
		return cmp
	}

	for i, datum := range p.data {
		if p.na[i] {
			cmp[i] = false
		} else {
			cmp[i] = datum == v
		}
	}

	return cmp
}

/* Comparable interface */

func (p *booleanPayload) Neq(val interface{}) []bool {
	cmp := make([]bool, p.length)
	v, ok := val.(bool)

	if !ok {
		for i := range p.data {
			cmp[i] = true
		}

		return cmp
	}

	for i, datum := range p.data {
		if p.na[i] {
			cmp[i] = true
		} else {
			cmp[i] = datum != v
		}
	}

	return cmp
}

func (p *booleanPayload) Gt(val interface{}) []bool {
	cmp := make([]bool, p.length)

	v, ok := val.(bool)
	if !ok {
		return cmp
	}

	for i, datum := range p.data {
		if p.na[i] {
			cmp[i] = false
		} else {
			cmp[i] = v == true && datum == false
		}
	}

	return cmp
}

func (p *booleanPayload) Lt(val interface{}) []bool {
	cmp := make([]bool, p.length)

	v, ok := val.(bool)
	if !ok {
		return cmp
	}

	for i, datum := range p.data {
		if p.na[i] {
			cmp[i] = false
		} else {
			cmp[i] = v == false && datum == true
		}
	}

	return cmp
}

func (p *booleanPayload) Gte(val interface{}) []bool {
	cmp := make([]bool, p.length)

	v, ok := val.(bool)
	if !ok {
		return cmp
	}

	for i, datum := range p.data {
		if p.na[i] {
			cmp[i] = false
		} else {
			cmp[i] = !(v == false && datum == true)
		}
	}

	return cmp
}

func (p *booleanPayload) Lte(val interface{}) []bool {
	cmp := make([]bool, p.length)

	v, ok := val.(bool)
	if !ok {
		return cmp
	}

	for i, datum := range p.data {
		if p.na[i] {
			cmp[i] = false
		} else {
			cmp[i] = !(v == true && datum == false)
		}
	}

	return cmp
}

func (p *booleanPayload) IsUnique() []bool {
	booleans := make([]bool, p.length)

	valuesMap := map[bool]bool{}
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

func (p *booleanPayload) Options() []Option {
	return []Option{}
}

func (p *booleanPayload) Coalesce(payload Payload) Payload {
	if p.length != payload.Len() {
		payload = payload.Adjust(p.length)
	}

	var srcData []bool
	var srcNA []bool

	if same, ok := payload.(*booleanPayload); ok {
		srcData = same.data
		srcNA = same.na
	} else if boolable, ok := payload.(Boolable); ok {
		srcData, srcNA = boolable.Booleans()
	} else {
		return p
	}

	dstData := make([]bool, p.length)
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

	return BooleanPayload(dstData, dstNA, p.Options()...)
}

func BooleanPayload(data []bool, na []bool, _ ...Option) Payload {
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

	vecData := make([]bool, length)
	for i := 0; i < length; i++ {
		if vecNA[i] {
			vecData[i] = false
		} else {
			vecData[i] = data[i]
		}
	}

	payload := &booleanPayload{
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
			return !payload.data[i] && payload.data[j]
		},
		FnEqual: func(i, j int) bool {
			return payload.data[i] == payload.data[j]
		},
	}

	return payload
}

func BooleanWithNA(data []bool, na []bool, options ...Option) Vector {
	return New(BooleanPayload(data, na, options...), options...)
}

func Boolean(data []bool, options ...Option) Vector {
	return BooleanWithNA(data, nil, options...)
}
