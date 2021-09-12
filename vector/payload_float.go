package vector

import (
	"math"
	"math/cmplx"
	"strconv"
)

type FloatWhicherFunc = func(int, float64, bool) bool
type FloatWhicherCompactFunc = func(float64, bool) bool
type FloatToFloatApplierFunc = func(int, float64, bool) (float64, bool)
type FloatToFloatApplierCompactFunc = func(float64, bool) (float64, bool)
type FloatSummarizerFunc = func(int, float64, float64, bool) (float64, bool)

type FloatPrinter struct {
	Precision int
}

type floatPayload struct {
	length  int
	data    []float64
	printer FloatPrinter
	DefNAble
}

func (p *floatPayload) Type() string {
	return "float"
}

func (p *floatPayload) Len() int {
	return p.length
}

func (p *floatPayload) ByIndices(indices []int) Payload {
	data := make([]float64, 0, len(indices))
	na := make([]bool, 0, len(indices))

	for _, idx := range indices {
		data = append(data, p.data[idx-1])
		na = append(na, p.na[idx-1])
	}

	return &floatPayload{
		length:  len(data),
		data:    data,
		printer: p.printer,
		DefNAble: DefNAble{
			na: na,
		},
	}
}

func (p *floatPayload) SupportsWhicher(whicher interface{}) bool {
	if _, ok := whicher.(FloatWhicherFunc); ok {
		return true
	}

	if _, ok := whicher.(FloatWhicherCompactFunc); ok {
		return true
	}

	return false
}

func (p *floatPayload) Which(whicher interface{}) []bool {
	if byFunc, ok := whicher.(FloatWhicherFunc); ok {
		return p.selectByFunc(byFunc)
	}

	if byFunc, ok := whicher.(FloatWhicherCompactFunc); ok {
		return p.selectByCompactFunc(byFunc)
	}

	return make([]bool, p.length)
}

func (p *floatPayload) selectByFunc(byFunc FloatWhicherFunc) []bool {
	booleans := make([]bool, p.length)

	for idx, val := range p.data {
		if byFunc(idx+1, val, p.na[idx]) {
			booleans[idx] = true
		}
	}

	return booleans
}

func (p *floatPayload) selectByCompactFunc(byFunc FloatWhicherCompactFunc) []bool {
	booleans := make([]bool, p.length)

	for idx, val := range p.data {
		if byFunc(val, p.na[idx]) {
			booleans[idx] = true
		}
	}

	return booleans
}

func (p *floatPayload) SupportsApplier(applier interface{}) bool {
	if _, ok := applier.(FloatToFloatApplierFunc); ok {
		return true
	}

	if _, ok := applier.(FloatToFloatApplierCompactFunc); ok {
		return true
	}

	return false
}

func (p *floatPayload) Apply(applier interface{}) Payload {
	if applyFunc, ok := applier.(FloatToFloatApplierFunc); ok {
		return p.applyToFloatByFunc(applyFunc)
	}

	if applyFunc, ok := applier.(FloatToFloatApplierCompactFunc); ok {
		return p.applyToFloatByCompactFunc(applyFunc)
	}

	return NAPayload(p.length)
}

func (p *floatPayload) applyToFloatByFunc(applyFunc FloatToFloatApplierFunc) Payload {
	data := make([]float64, p.length)
	na := make([]bool, p.length)

	for i := 0; i < p.length; i++ {
		dataVal, naVal := applyFunc(i+1, p.data[i], p.na[i])
		if naVal {
			dataVal = math.NaN()
		}
		data[i] = dataVal
		na[i] = naVal
	}

	return FloatPayload(data, na, p.options()...)
}

func (p *floatPayload) applyToFloatByCompactFunc(applyFunc FloatToFloatApplierCompactFunc) Payload {
	data := make([]float64, p.length)
	na := make([]bool, p.length)

	for i := 0; i < p.length; i++ {
		dataVal, naVal := applyFunc(p.data[i], p.na[i])
		if naVal {
			dataVal = math.NaN()
		}
		data[i] = dataVal
		na[i] = naVal
	}

	return FloatPayload(data, na, p.options()...)
}

func (p *floatPayload) SupportsSummarizer(summarizer interface{}) bool {
	if _, ok := summarizer.(FloatSummarizerFunc); ok {
		return true
	}

	return false
}

func (p *floatPayload) Summarize(summarizer interface{}) Payload {
	fn, ok := summarizer.(FloatSummarizerFunc)
	if !ok {
		return NAPayload(1)
	}

	val := 0.0
	na := false
	for i := 0; i < p.length; i++ {
		val, na = fn(i+1, val, p.data[i], p.na[i])

		if na {
			return NAPayload(1)
		}
	}

	return FloatPayload([]float64{val}, nil, p.options()...)
}

func (p *floatPayload) Integers() ([]int, []bool) {
	if p.length == 0 {
		return []int{}, []bool{}
	}

	data := make([]int, p.length)
	for i := 0; i < p.length; i++ {
		if p.na[i] {
			data[i] = 0
		} else {
			data[i] = int(p.data[i])
		}
	}

	na := make([]bool, p.Len())
	copy(na, p.na)

	return data, na
}

func (p *floatPayload) Floats() ([]float64, []bool) {
	if p.length == 0 {
		return []float64{}, []bool{}
	}

	data := make([]float64, p.length)
	copy(data, p.data)

	na := make([]bool, p.Len())
	copy(na, p.na)

	return data, na
}

func (p *floatPayload) Complexes() ([]complex128, []bool) {
	if p.length == 0 {
		return []complex128{}, []bool{}
	}

	data := make([]complex128, p.length)
	for i := 0; i < p.length; i++ {
		if p.na[i] {
			data[i] = cmplx.NaN()
		} else {
			if math.IsNaN(p.data[i]) {
				data[i] = cmplx.NaN()
			} else {
				data[i] = complex(p.data[i], 0)
			}
		}
	}

	na := make([]bool, p.Len())
	copy(na, p.na)

	return data, na
}

func (p *floatPayload) Booleans() ([]bool, []bool) {
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

func (p *floatPayload) Strings() ([]string, []bool) {
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

func (p *floatPayload) Interfaces() ([]interface{}, []bool) {
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

func (p *floatPayload) StrForElem(idx int) string {
	i := idx - 1

	if p.na[i] {
		return "NA"
	}

	if math.IsInf(p.data[i], +1) {
		return "+Inf"
	}

	if math.IsInf(p.data[i], -1) {
		return "-Inf"
	}

	if math.IsNaN(p.data[i]) {
		return "NaN"
	}

	return strconv.FormatFloat(p.data[i], 'f', p.printer.Precision, 64)
}

func (p *floatPayload) Append(payload Payload) Payload {
	length := p.length + payload.Len()

	var vals []float64
	var na []bool

	if floatable, ok := payload.(Floatable); ok {
		vals, na = floatable.Floats()
	} else {
		vals, na = NAPayload(payload.Len()).(Floatable).Floats()
	}

	newVals := make([]float64, length)
	newNA := make([]bool, length)

	copy(newVals, p.data)
	copy(newVals[p.length:], vals)
	copy(newNA, p.na)
	copy(newNA[p.length:], na)

	return FloatPayload(newVals, newNA, p.options()...)
}

func (p *floatPayload) Adjust(size int) Payload {
	if size < p.length {
		return p.adjustToLesserSize(size)
	}

	if size > p.length {
		return p.adjustToBiggerSize(size)
	}

	return p
}

func (p *floatPayload) adjustToLesserSize(size int) Payload {
	data := make([]float64, size)
	na := make([]bool, size)

	copy(data, p.data)
	copy(na, p.na)

	return FloatPayload(data, na, p.options()...)
}

func (p *floatPayload) adjustToBiggerSize(size int) Payload {
	cycles := size / p.length
	if size%p.length > 0 {
		cycles++
	}

	data := make([]float64, cycles*p.length)
	na := make([]bool, cycles*p.length)

	for i := 0; i < cycles; i++ {
		copy(data[i*p.length:], p.data)
		copy(na[i*p.length:], p.na)
	}

	data = data[:size]
	na = na[:size]

	return FloatPayload(data, na, p.options()...)
}

func (p *floatPayload) options() []Option {
	return []Option{
		OptionPrecision(p.printer.Precision),
	}
}

/* Finder interface */

func (p *floatPayload) Find(needle interface{}) int {
	var val float64

	switch v := needle.(type) {
	case float64:
		val = v
	case int:
		val = float64(v)
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

func (p *floatPayload) FindAll(needle interface{}) []int {
	var val float64

	switch v := needle.(type) {
	case float64:
		val = v
	case int:
		val = float64(v)
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

func (p *floatPayload) Eq(val interface{}) []bool {
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

func (p *floatPayload) Neq(val interface{}) []bool {
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

func (p *floatPayload) Gt(val interface{}) []bool {
	cmp := make([]bool, p.length)

	v, ok := p.convertComparator(val)
	if !ok {
		return cmp
	}

	for i, datum := range p.data {
		if p.na[i] {
			cmp[i] = false
		} else {
			cmp[i] = datum > v
		}
	}

	return cmp
}

func (p *floatPayload) Lt(val interface{}) []bool {
	cmp := make([]bool, p.length)

	v, ok := p.convertComparator(val)
	if !ok {
		return cmp
	}

	for i, datum := range p.data {
		if p.na[i] {
			cmp[i] = false
		} else {
			cmp[i] = datum < v
		}
	}

	return cmp
}

func (p *floatPayload) Gte(val interface{}) []bool {
	cmp := make([]bool, p.length)

	v, ok := p.convertComparator(val)
	if !ok {
		return cmp
	}

	for i, datum := range p.data {
		if p.na[i] {
			cmp[i] = false
		} else {
			cmp[i] = datum >= v
		}
	}

	return cmp
}

func (p *floatPayload) Lte(val interface{}) []bool {
	cmp := make([]bool, p.length)

	v, ok := p.convertComparator(val)
	if !ok {
		return cmp
	}

	for i, datum := range p.data {
		if p.na[i] {
			cmp[i] = false
		} else {
			cmp[i] = datum <= v
		}
	}

	return cmp
}

func (p *floatPayload) convertComparator(val interface{}) (float64, bool) {
	var v float64
	var err error
	ok := true
	switch val.(type) {
	case complex128:
		ip := imag(val.(complex128))
		if ip == 0 {
			v = real(val.(complex128))
		} else {
			ok = false
		}
	case complex64:
		ip := imag(val.(complex64))
		if ip == 0 {
			v = float64(real(val.(complex64)))
		} else {
			ok = false
		}
	case float64:
		v = val.(float64)
	case float32:
		v = float64(val.(float32))
	case int:
		v = float64(val.(int))
	case int64:
		v = float64(val.(int64))
	case int32:
		v = float64(val.(int32))
	case uint64:
		v = float64(val.(uint64))
	case uint32:
		v = float64(val.(uint32))
	case string:
		v, err = strconv.ParseFloat(val.(string), 64)
		if err != nil {
			ok = false
		}
	default:
		ok = false
	}

	return v, ok
}

func FloatPayload(data []float64, na []bool, options ...Option) Payload {
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

	vecData := make([]float64, length)
	for i := 0; i < length; i++ {
		if vecNA[i] {
			vecData[i] = math.NaN()
		} else {
			vecData[i] = data[i]
		}
	}

	printer := FloatPrinter{
		Precision: 3,
	}

	if conf.HasOption(KeyOptionPrecision) {
		printer.Precision = conf.Value(KeyOptionPrecision).(int)
	}

	return &floatPayload{
		length:  length,
		data:    vecData,
		printer: printer,
		DefNAble: DefNAble{
			na: vecNA,
		},
	}
}

func Float(data []float64, na []bool, options ...Option) Vector {
	return New(FloatPayload(data, na, options...))
}
