package vector

import (
	"math"
	"math/cmplx"
	"strconv"
)

type ComplexWhicherFunc = func(int, complex128, bool) bool
type ComplexWhicherCompactFunc = func(complex128, bool) bool
type ComplexApplierFunc = func(int, complex128, bool) (complex128, bool)
type ComplexApplierCompactFunc = func(complex128, bool) (complex128, bool)
type ComplexSummarizerFunc = func(int, complex128, complex128, bool) (complex128, bool)

type ComplexPrinter struct {
	Precision int
}

type complexPayload struct {
	length  int
	data    []complex128
	printer ComplexPrinter
	DefNAble
}

func (p *complexPayload) Type() string {
	return "complex"
}

func (p *complexPayload) Len() int {
	return p.length
}

func (p *complexPayload) Pick(idx int) interface{} {
	return pickValueWithNA(idx, p.data, p.na, p.length)
}

func (p *complexPayload) Data() []interface{} {
	return dataWithNAToInterfaceArray(p.data, p.na)
}

func (p *complexPayload) ByIndices(indices []int) Payload {
	data, na := byIndices(indices, p.data, p.na, cmplx.NaN())

	return ComplexPayload(data, na, p.Options()...)
}

func (p *complexPayload) SupportsWhicher(whicher any) bool {
	return supportsWhicher[complex128](whicher)
}

func (p *complexPayload) Which(whicher interface{}) []bool {
	if byFunc, ok := whicher.(ComplexWhicherFunc); ok {
		return selectByFunc(p.data, p.na, byFunc)
	}

	if byFunc, ok := whicher.(ComplexWhicherCompactFunc); ok {
		return selectByCompactFunc(p.data, p.na, byFunc)
	}

	return make([]bool, p.length)
}

func (p *complexPayload) SupportsApplier(applier any) bool {
	return supportApplier[complex128](applier)
}

func (p *complexPayload) Apply(applier interface{}) Payload {
	if applyFunc, ok := applier.(ComplexApplierFunc); ok {
		data, na := applyByFunc(p.data, p.na, p.length, applyFunc, cmplx.NaN())

		return ComplexPayload(data, na, p.Options()...)
	}

	if applyFunc, ok := applier.(ComplexApplierCompactFunc); ok {
		data, na := applyByCompactFunc(p.data, p.na, p.length, applyFunc, cmplx.NaN())

		return ComplexPayload(data, na, p.Options()...)
	}

	return NAPayload(p.length)

}

func (p *complexPayload) ApplyTo(indices []int, applier interface{}) Payload {
	//TODO implement me
	panic("implement me")
}

func (p *complexPayload) SupportsSummarizer(summarizer interface{}) bool {
	if _, ok := summarizer.(ComplexSummarizerFunc); ok {
		return true
	}

	return false
}

func (p *complexPayload) Summarize(summarizer interface{}) Payload {
	fn, ok := summarizer.(ComplexSummarizerFunc)
	if !ok {
		return NAPayload(1)
	}

	val := 0 + 0i
	na := false
	for i := 0; i < p.length; i++ {
		val, na = fn(i+1, val, p.data[i], p.na[i])

		if na {
			return NAPayload(1)
		}
	}

	return ComplexPayload([]complex128{val}, nil, p.Options()...)
}

func (p *complexPayload) Integers() ([]int, []bool) {
	if p.length == 0 {
		return []int{}, []bool{}
	}

	data := make([]int, p.length)
	for i := 0; i < p.length; i++ {
		if p.na[i] {
			data[i] = 0
		} else {
			data[i] = int(real(p.data[i]))
		}
	}

	na := make([]bool, p.Len())
	copy(na, p.na)

	return data, na
}

func (p *complexPayload) Floats() ([]float64, []bool) {
	if p.length == 0 {
		return []float64{}, []bool{}
	}

	data := make([]float64, p.length)

	for i := 0; i < p.length; i++ {
		if p.na[i] {
			data[i] = math.NaN()
		} else {
			data[i] = real(p.data[i])
		}
	}

	na := make([]bool, p.Len())
	copy(na, p.na)

	return data, na
}

func (p *complexPayload) Complexes() ([]complex128, []bool) {
	if p.length == 0 {
		return []complex128{}, []bool{}
	}

	data := make([]complex128, p.length)
	copy(data, p.data)

	na := make([]bool, p.Len())
	copy(na, p.na)

	return data, na
}

func (p *complexPayload) Booleans() ([]bool, []bool) {
	if p.length == 0 {
		return []bool{}, []bool{}
	}

	data := make([]bool, p.length)

	for i := 0; i < p.length; i++ {
		if p.na[i] {
			data[i] = false
		} else {
			data[i] = p.data[i] != 0
		}
	}

	na := make([]bool, p.length)
	copy(na, p.na)

	return data, na
}

func (p *complexPayload) Strings() ([]string, []bool) {
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

func (p *complexPayload) Interfaces() ([]interface{}, []bool) {
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

func (p *complexPayload) Append(payload Payload) Payload {
	length := p.length + payload.Len()

	var vals []complex128
	var na []bool

	if complexable, ok := payload.(Complexable); ok {
		vals, na = complexable.Complexes()
	} else {
		vals, na = NAPayload(payload.Len()).(Complexable).Complexes()
	}

	newVals := make([]complex128, length)
	newNA := make([]bool, length)

	copy(newVals, p.data)
	copy(newVals[p.length:], vals)
	copy(newNA, p.na)
	copy(newNA[p.length:], na)

	return ComplexPayload(newVals, newNA, p.Options()...)
}

func (p *complexPayload) Adjust(size int) Payload {
	if size < p.length {
		return p.adjustToLesserSize(size)
	}

	if size > p.length {
		return p.adjustToBiggerSize(size)
	}

	return p
}

func (p *complexPayload) adjustToLesserSize(size int) Payload {
	data, na := adjustToLesserSizeWithNA(p.data, p.na, size)

	return ComplexPayload(data, na, p.Options()...)
}

func (p *complexPayload) adjustToBiggerSize(size int) Payload {
	data, na := adjustToBiggerSizeWithNA(p.data, p.na, p.length, size)

	return ComplexPayload(data, na, p.Options()...)
}

func (p *complexPayload) Options() []Option {
	return []Option{
		OptionPrecision(p.printer.Precision),
	}
}

func (p *complexPayload) Groups() ([][]int, []interface{}) {
	groups, values := groupsForData(p.data, p.na)

	return groups, values
}

func (p *complexPayload) StrForElem(idx int) string {
	i := idx - 1

	if p.na[i] {
		return "NA"
	}

	if cmplx.IsInf(p.data[i]) {
		return "Inf"
	}

	if cmplx.IsNaN(p.data[i]) {
		return "NaN"
	}

	return strconv.FormatComplex(p.data[i], 'f', p.printer.Precision, 128)
}

/* Finder interface */

func (p *complexPayload) Find(needle interface{}) int {
	var val complex128

	switch v := needle.(type) {
	case complex128:
		val = v
	case float64:
		val = complex(v, 0)
	case int:
		val = complex(float64(v), 0)
	default:
		return 0
	}

	for i, datum := range p.data {
		if !p.na[i] && val == datum {
			return i + 1
		}
	}

	return 0
}

func (p *complexPayload) FindAll(needle interface{}) []int {
	var val complex128

	switch v := needle.(type) {
	case complex128:
		val = v
	case float64:
		val = complex(v, 0)
	case int:
		val = complex(float64(v), 0)
	default:
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

/* Comparable interface */

func (p *complexPayload) Eq(val interface{}) []bool {
	cmp := make([]bool, p.length)

	v, ok := p.convertComparator(val)
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

func (p *complexPayload) Neq(val interface{}) []bool {
	cmp := make([]bool, p.length)

	v, ok := p.convertComparator(val)
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

func (p *complexPayload) Gt(interface{}) []bool {
	cmp := make([]bool, p.length)

	return cmp
}

func (p *complexPayload) Lt(interface{}) []bool {
	cmp := make([]bool, p.length)

	return cmp
}

func (p *complexPayload) Gte(interface{}) []bool {
	cmp := make([]bool, p.length)

	return cmp
}

func (p *complexPayload) Lte(interface{}) []bool {
	cmp := make([]bool, p.length)

	return cmp
}

func (p *complexPayload) convertComparator(val interface{}) (complex128, bool) {
	var v complex128
	ok := true
	switch val.(type) {
	case complex128:
		v = val.(complex128)
	case complex64:
		v = complex128(val.(complex64))
	case float64:
		v = complex(val.(float64), 0)
	case float32:
		v = complex(float64(val.(float32)), 0)
	case int:
		v = complex(float64(val.(int)), 0)
	case int64:
		v = complex(float64(val.(int64)), 0)
	case int32:
		v = complex(float64(val.(int32)), 0)
	case uint64:
		v = complex(float64(val.(uint64)), 0)
	case uint32:
		v = complex(float64(val.(uint32)), 0)
	default:
		ok = false
	}

	return v, ok
}

func (p *complexPayload) Coalesce(payload Payload) Payload {
	if p.length != payload.Len() {
		payload = payload.Adjust(p.length)
	}

	var srcData []complex128
	var srcNA []bool

	if same, ok := payload.(*complexPayload); ok {
		srcData = same.data
		srcNA = same.na
	} else if complexable, ok := payload.(Complexable); ok {
		srcData, srcNA = complexable.Complexes()
	} else {
		return p
	}

	dstData := make([]complex128, p.length)
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

	return ComplexPayload(dstData, dstNA, p.Options()...)
}

func ComplexPayload(data []complex128, na []bool, options ...Option) Payload {
	length := len(data)
	conf := MergeOptions(options)

	vecNA := make([]bool, length)
	if len(na) > 0 {
		if len(na) == length {
			copy(vecNA, na)
		} else {
			emp := NAPayload(0)
			return emp
		}
	}

	vecData := make([]complex128, length)
	for i := 0; i < length; i++ {
		if vecNA[i] {
			vecData[i] = cmplx.NaN()
		} else {
			vecData[i] = data[i]
		}
	}

	printer := ComplexPrinter{
		Precision: 3,
	}

	if conf.HasOption(KeyOptionPrecision) {
		printer.Precision = conf.Value(KeyOptionPrecision).(int)
	}

	return &complexPayload{
		length:  length,
		data:    vecData,
		printer: printer,
		DefNAble: DefNAble{
			na: vecNA,
		},
	}
}

func (p *complexPayload) IsUnique() []bool {
	booleans := make([]bool, p.length)

	valuesMap := map[complex128]bool{}
	wasNA := false
	wasNaN := false
	wasInf := false
	for i := 0; i < p.length; i++ {
		is := false

		if p.na[i] {
			if !wasNA {
				is = true
				wasNA = true
			}
		} else if cmplx.IsNaN(p.data[i]) {
			if !wasNaN {
				is = true
				wasNaN = true
			}
		} else if cmplx.IsInf(p.data[i]) {
			if !wasInf {
				is = true
				wasInf = true
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

func ComplexWithNA(data []complex128, na []bool, options ...Option) Vector {
	return New(ComplexPayload(data, na, options...), options...)
}

func Complex(data []complex128, options ...Option) Vector {
	return ComplexWithNA(data, nil, options...)
}
