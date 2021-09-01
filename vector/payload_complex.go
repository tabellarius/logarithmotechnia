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

	return ComplexPayload(data, na, p.options()...)
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

	return ComplexPayload(data, na, p.options()...)
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

	return ComplexPayload(data, na, p.options()...)
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

	return ComplexPayload([]complex128{val}, nil, p.options()...)
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

func (p *complexPayload) Append(vec Vector) Payload {
	length := p.length + vec.Len()

	vals, na := vec.Complexes()

	newVals := make([]complex128, length)
	newNA := make([]bool, length)

	copy(newVals, p.data)
	copy(newVals[p.length:], vals)
	copy(newNA, p.na)
	copy(newNA[p.length:], na)

	return ComplexPayload(newVals, newNA, p.options()...)
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

func (p *complexPayload) options() []Option {
	return []Option{
		OptionPrecision(p.printer.Precision),
	}
}

func ComplexPayload(data []complex128, na []bool, options ...Option) Payload {
	length := len(data)
	conf := mergeOptions(options)

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

	if conf.HasOption(OPTION_PRECISION) {
		printer.Precision = conf.Value(OPTION_PRECISION).(int)
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

func Complex(data []complex128, na []bool, options ...Option) Vector {
	return New(ComplexPayload(data, na, options...))
}
