package vector

import (
	"logarithmotechnia/embed"
	"math"
	"math/cmplx"
	"strconv"
)

// integerPayload is a structure, subsisting Integer vectors
type integerPayload struct {
	length int
	data   []int
	embed.NAble
	embed.Arrangeable
}

func (p *integerPayload) Type() string {
	return "integer"
}

func (p *integerPayload) Len() int {
	return p.length
}

func (p *integerPayload) Pick(idx int) any {
	return pickValueWithNA(idx, p.data, p.NA, p.length)
}

func (p *integerPayload) Data() []any {
	return dataWithNAToInterfaceArray(p.data, p.NA)
}

func (p *integerPayload) ByIndices(indices []int) Payload {
	data, na := byIndicesWithNA(indices, p.data, p.NA, 0)

	return IntegerPayload(data, na, p.Options()...)
}

func (p *integerPayload) SupportsWhicher(whicher any) bool {
	return supportsWhicherWithNA[int](whicher)
}

func (p *integerPayload) Which(whicher any) []bool {
	return whichWithNA(p.data, p.NA, whicher)
}

func (p *integerPayload) Apply(applier any) Payload {
	return applyWithNA(p.data, p.NA, applier, p.Options())
}

func (p *integerPayload) Traverse(traverser any) {
	traverseWithNA(p.data, p.NA, traverser)
}

func (p *integerPayload) ApplyTo(indices []int, applier any) Payload {
	data, na := applyToWithNA(indices, p.data, p.NA, applier, 0)

	if data == nil {
		return NAPayload(p.length)
	}

	return IntegerPayload(data, na, p.Options()...)
}

func (p *integerPayload) SupportsSummarizer(summarizer any) bool {
	return supportsSummarizer[int](summarizer)
}

func (p *integerPayload) Summarize(summarizer any) Payload {
	val, na := summarize(p.data, p.NA, summarizer, 0, 0)

	return IntegerPayload([]int{val}, []bool{na}, p.Options()...)
}

func (p *integerPayload) Integers() ([]int, []bool) {
	if p.length == 0 {
		return []int{}, []bool{}
	}

	data := make([]int, p.length)
	copy(data, p.data)

	na := make([]bool, p.Len())
	copy(na, p.NA)

	return data, na
}

func (p *integerPayload) Floats() ([]float64, []bool) {
	if p.length == 0 {
		return []float64{}, []bool{}
	}

	data := make([]float64, p.length)

	for i := 0; i < p.length; i++ {
		if p.NA[i] {
			data[i] = math.NaN()
		} else {
			data[i] = float64(p.data[i])
		}
	}

	na := make([]bool, p.Len())
	copy(na, p.NA)

	return data, na
}

func (p *integerPayload) Complexes() ([]complex128, []bool) {
	if p.length == 0 {
		return []complex128{}, []bool{}
	}

	data := make([]complex128, p.length)
	for i := 0; i < p.length; i++ {
		if p.NA[i] {
			data[i] = cmplx.NaN()
		} else {
			data[i] = complex(float64(p.data[i]), 0)
		}
	}

	na := make([]bool, p.Len())
	copy(na, p.NA)

	return data, na
}

func (p *integerPayload) Booleans() ([]bool, []bool) {
	if p.length == 0 {
		return []bool{}, []bool{}
	}

	data := make([]bool, p.length)

	for i := 0; i < p.length; i++ {
		if p.NA[i] {
			data[i] = false
		} else {
			data[i] = p.data[i] != 0
		}
	}

	na := make([]bool, p.length)
	copy(na, p.NA)

	return data, na
}

func (p *integerPayload) Strings() ([]string, []bool) {
	if p.length == 0 {
		return []string{}, []bool{}
	}

	data := make([]string, p.length)

	for i := 0; i < p.length; i++ {
		if p.NA[i] {
			data[i] = ""
		} else {
			data[i] = strconv.Itoa(p.data[i])
		}
	}

	na := make([]bool, p.Len())
	copy(na, p.NA)

	return data, na
}

func (p *integerPayload) Anies() ([]any, []bool) {
	if p.length == 0 {
		return []any{}, []bool{}
	}

	data := make([]any, p.length)
	for i := 0; i < p.length; i++ {
		if p.NA[i] {
			data[i] = nil
		} else {
			data[i] = p.data[i]
		}
	}

	na := make([]bool, p.length)
	copy(na, p.NA)

	return data, na
}

func (p *integerPayload) Append(payload Payload) Payload {
	length := p.length + payload.Len()

	var vals []int
	var na []bool

	if intable, ok := payload.(Intable); ok {
		vals, na = intable.Integers()
	} else {
		vals, na = NAPayload(payload.Len()).(Intable).Integers()
	}

	newVals := make([]int, length)
	newNA := make([]bool, length)

	copy(newVals, p.data)
	copy(newVals[p.length:], vals)
	copy(newNA, p.NA)
	copy(newNA[p.length:], na)

	return IntegerPayload(newVals, newNA, p.Options()...)
}

func (p *integerPayload) Adjust(size int) Payload {
	if size < p.length {
		return p.adjustToLesserSize(size)
	}

	if size > p.length {
		return p.adjustToBiggerSize(size)
	}

	return p
}

func (p *integerPayload) adjustToLesserSize(size int) Payload {
	data, na := adjustToLesserSizeWithNA(p.data, p.NA, size)

	return IntegerPayload(data, na, p.Options()...)
}

func (p *integerPayload) adjustToBiggerSize(size int) Payload {
	data, na := adjustToBiggerSizeWithNA(p.data, p.NA, p.length, size)

	return IntegerPayload(data, na, p.Options()...)
}

func (p *integerPayload) Groups() ([][]int, []any) {
	groups, values := groupsForData(p.data, p.NA)

	return groups, values
}

func (p *integerPayload) StrForElem(idx int) string {
	if p.NA[idx-1] {
		return "NA"
	}

	return strconv.Itoa(p.data[idx-1])
}

func (p *integerPayload) Find(needle any) int {
	return find(needle, p.data, p.NA, p.convertComparator)
}

/* Finder interface */

func (p *integerPayload) FindAll(needle any) []int {
	return findAll(needle, p.data, p.NA, p.convertComparator)
}

func (p *integerPayload) Eq(val any) []bool {
	return eq(val, p.data, p.NA, p.convertComparator)
}

func (p *integerPayload) Neq(val any) []bool {
	return neq(val, p.data, p.NA, p.convertComparator)
}

func (p *integerPayload) Gt(val any) []bool {
	return gt(val, p.data, p.NA, p.convertComparator)
}

func (p *integerPayload) Lt(val any) []bool {
	return lt(val, p.data, p.NA, p.convertComparator)
}

func (p *integerPayload) Gte(val any) []bool {
	return gte(val, p.data, p.NA, p.convertComparator)
}

func (p *integerPayload) Lte(val any) []bool {
	return lte(val, p.data, p.NA, p.convertComparator)
}

func (p *integerPayload) convertComparator(val any) (int, bool) {
	var v int
	ok := true
	switch value := val.(type) {
	case int:
		v = value
	case int64:
		v = int(value)
	case int32:
		v = int(value)
	case uint64:
		v = int(value)
	case uint32:
		v = int(value)
	case complex128:
		ip := imag(value)
		rp, fp := math.Modf(real(value))
		if ip == 0 && fp == 0 {
			v = int(rp)
		} else {
			ok = false
		}
	case complex64:
		ip := imag(value)
		rp, fp := math.Modf(float64(real(value)))
		if ip == 0 && fp == 0 {
			v = int(rp)
		} else {
			ok = false
		}
	case float64:
		rp, fp := math.Modf(value)
		if fp == 0 {
			v = int(rp)
		} else {
			ok = false
		}
	case float32:
		rp, fp := math.Modf(float64(value))
		if fp == 0 {
			v = int(rp)
		} else {
			ok = false
		}
	default:
		ok = false
	}

	return v, ok
}

func (p *integerPayload) IsUnique() []bool {
	booleans := make([]bool, p.length)

	valuesMap := map[int]bool{}
	wasNA := false
	for i := 0; i < p.length; i++ {
		is := false

		if p.NA[i] {
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

func (p *integerPayload) Coalesce(payload Payload) Payload {
	if p.length != payload.Len() {
		payload = payload.Adjust(p.length)
	}

	var srcData []int
	var srcNA []bool

	if same, ok := payload.(*integerPayload); ok {
		srcData = same.data
		srcNA = same.NA
	} else if intable, ok := payload.(Intable); ok {
		srcData, srcNA = intable.Integers()
	} else {
		return p
	}

	dstData := make([]int, p.length)
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

	return IntegerPayload(dstData, dstNA, p.Options()...)
}

func (p *integerPayload) Options() []Option {
	return []Option{}
}

func (p *integerPayload) SetOption(string, any) bool {
	return false
}

// IntegerPayload creates a payload with integer data.
func IntegerPayload(data []int, na []bool, _ ...Option) Payload {
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

	vecData := make([]int, length)
	for i := 0; i < length; i++ {
		if vecNA[i] {
			vecData[i] = 0
		} else {
			vecData[i] = data[i]
		}
	}

	payload := &integerPayload{
		length: length,
		data:   vecData,
		NAble: embed.NAble{
			NA: vecNA,
		},
	}

	payload.Arrangeable = embed.Arrangeable{
		Length: payload.length,
		NAble:  payload.NAble,
		FnLess: func(i, j int) bool {
			return payload.data[i] < payload.data[j]
		},
		FnEqual: func(i, j int) bool {
			return payload.data[i] == payload.data[j]
		},
	}

	return payload
}

// IntegerWithNA creates a vector with IntegerPayload and allows to set NA-values.
func IntegerWithNA(data []int, na []bool, options ...Option) Vector {
	return New(IntegerPayload(data, na, options...), options...)
}

// Integer creates a vector with IntegerPayload.
func Integer(data []int, options ...Option) Vector {
	return IntegerWithNA(data, nil, options...)
}
