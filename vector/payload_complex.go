package vector

import (
	"math"
	"math/cmplx"
	"strconv"
)

type ComplexWhicherFunc = func(int, complex128, bool) bool
type ComplexWhicherCompactFunc = func(complex128, bool) bool
type ComplexToComplexApplierFunc = func(int, complex128, bool) (complex128, bool)
type ComplexToComplexApplierCompactFunc = func(complex128, bool) (complex128, bool)
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

func (p *complexPayload) ByIndices(indices []int) Payload {
	data := make([]complex128, 0, len(indices))
	na := make([]bool, 0, len(indices))

	for _, idx := range indices {
		data = append(data, p.data[idx-1])
		na = append(na, p.na[idx-1])
	}

	return ComplexPayload(data, na, p.Options()...)
}

func (p *complexPayload) SupportsWhicher(whicher interface{}) bool {
	if _, ok := whicher.(ComplexWhicherFunc); ok {
		return true
	}

	if _, ok := whicher.(ComplexWhicherCompactFunc); ok {
		return true
	}

	return false
}

func (p *complexPayload) Which(whicher interface{}) []bool {
	if byFunc, ok := whicher.(ComplexWhicherFunc); ok {
		return p.selectByFunc(byFunc)
	}

	if byFunc, ok := whicher.(ComplexWhicherCompactFunc); ok {
		return p.selectByCompactFunc(byFunc)
	}

	return make([]bool, p.length)
}

func (p *complexPayload) selectByFunc(byFunc ComplexWhicherFunc) []bool {
	booleans := make([]bool, p.length)

	for idx, val := range p.data {
		if byFunc(idx+1, val, p.na[idx]) {
			booleans[idx] = true
		}
	}

	return booleans
}

func (p *complexPayload) selectByCompactFunc(byFunc ComplexWhicherCompactFunc) []bool {
	booleans := make([]bool, p.length)

	for idx, val := range p.data {
		if byFunc(val, p.na[idx]) {
			booleans[idx] = true
		}
	}

	return booleans
}

func (p *complexPayload) SupportsApplier(applier interface{}) bool {
	if _, ok := applier.(ComplexToComplexApplierFunc); ok {
		return true
	}

	if _, ok := applier.(ComplexToComplexApplierCompactFunc); ok {
		return true
	}

	return false
}

func (p *complexPayload) Apply(applier interface{}) Payload {
	if applyFunc, ok := applier.(ComplexToComplexApplierFunc); ok {
		return p.applyToComplexByFunc(applyFunc)
	}

	if applyFunc, ok := applier.(ComplexToComplexApplierCompactFunc); ok {
		return p.applyToComplexByCompactFunc(applyFunc)
	}

	return NAPayload(p.length)

}

func (p *complexPayload) applyToComplexByFunc(applyFunc ComplexToComplexApplierFunc) Payload {
	data := make([]complex128, p.length)
	na := make([]bool, p.length)

	for i := 0; i < p.length; i++ {
		dataVal, naVal := applyFunc(i+1, p.data[i], p.na[i])
		if naVal {
			dataVal = cmplx.NaN()
		}
		data[i] = dataVal
		na[i] = naVal
	}

	return ComplexPayload(data, na, p.Options()...)
}

func (p *complexPayload) applyToComplexByCompactFunc(applyFunc ComplexToComplexApplierCompactFunc) Payload {
	data := make([]complex128, p.length)
	na := make([]bool, p.length)

	for i := 0; i < p.length; i++ {
		dataVal, naVal := applyFunc(p.data[i], p.na[i])
		if naVal {
			dataVal = cmplx.NaN()
		}
		data[i] = dataVal
		na[i] = naVal
	}

	return ComplexPayload(data, na, p.Options()...)
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
	data := make([]complex128, size)
	na := make([]bool, size)

	copy(data, p.data)
	copy(na, p.na)

	return ComplexPayload(data, na, p.Options()...)
}

func (p *complexPayload) adjustToBiggerSize(size int) Payload {
	cycles := size / p.length
	if size%p.length > 0 {
		cycles++
	}

	data := make([]complex128, cycles*p.length)
	na := make([]bool, cycles*p.length)

	for i := 0; i < cycles; i++ {
		copy(data[i*p.length:], p.data)
		copy(na[i*p.length:], p.na)
	}

	data = data[:size]
	na = na[:size]

	return ComplexPayload(data, na, p.Options()...)
}

func (p *complexPayload) Options() []Option {
	return []Option{
		OptionPrecision(p.printer.Precision),
	}
}

func (p *complexPayload) Groups() ([][]int, []interface{}) {
	groupMap := map[complex128][]int{}
	ordered := []complex128{}
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
		if val == datum {
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
		if val == datum {
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

func ComplexWithNA(data []complex128, na []bool, options ...Option) Vector {
	return New(ComplexPayload(data, na, options...), options...)
}

func Complex(data []complex128, options ...Option) Vector {
	return ComplexWithNA(data, nil, options...)
}
