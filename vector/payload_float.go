package vector

import (
	"math"
	"math/cmplx"
	"strconv"
)

type FloatWhicherFunc = func(int, float64, bool) bool
type FloatWhicherCompactFunc = func(float64, bool) bool
type FloatApplierFunc = func(int, float64, bool) (float64, bool)
type FloatApplierCompactFunc = func(float64, bool) (float64, bool)
type FloatSummarizerFunc = func(int, float64, float64, bool) (float64, bool)

type FloatPrinter struct {
	Precision int
}

type floatPayload struct {
	length  int
	data    []float64
	printer FloatPrinter
	DefNAble
	DefArrangeable
}

func (p *floatPayload) Type() string {
	return "float"
}

func (p *floatPayload) Len() int {
	return p.length
}

func (p *floatPayload) Pick(idx int) interface{} {
	return pickValueWithNA(idx, p.data, p.na, p.length)
}

func (p *floatPayload) Data() []interface{} {
	return dataWithNAToInterfaceArray(p.data, p.na)
}

func (p *floatPayload) ByIndices(indices []int) Payload {
	data, na := byIndices(indices, p.data, p.na, math.NaN())

	return FloatPayload(data, na, p.Options()...)
}

func (p *floatPayload) SupportsWhicher(whicher any) bool {
	return supportsWhicher[float64](whicher)
}

func (p *floatPayload) Which(whicher any) []bool {
	return which(p.data, p.na, whicher)
}

func (p *floatPayload) SupportsApplier(applier any) bool {
	return supportsApplier[float64](applier)
}

func (p *floatPayload) Apply(applier any) Payload {
	data, na := apply(p.data, p.na, applier, math.NaN())

	if data == nil {
		return NAPayload(p.length)
	}

	return FloatPayload(data, na, p.Options()...)
}

func (p *floatPayload) ApplyTo(indices []int, applier any) Payload {
	data, na := applyTo(indices, p.data, p.na, applier, math.NaN())

	if data == nil {
		return NAPayload(p.length)
	}

	return FloatPayload(data, na, p.Options()...)
}

func (p *floatPayload) SupportsSummarizer(summarizer any) bool {
	return supportsSummarizer[float64](summarizer)
}

func (p *floatPayload) Summarize(summarizer interface{}) Payload {
	val, na := summarize(p.data, p.na, summarizer, 0.0, math.NaN())

	return FloatPayload([]float64{val}, []bool{na}, p.Options()...)
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

	return FloatPayload(newVals, newNA, p.Options()...)
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
	data, na := adjustToLesserSizeWithNA(p.data, p.na, size)

	return FloatPayload(data, na, p.Options()...)
}

func (p *floatPayload) adjustToBiggerSize(size int) Payload {
	data, na := adjustToBiggerSizeWithNA(p.data, p.na, p.length, size)

	return FloatPayload(data, na, p.Options()...)
}

/* Finder interface */

func (p *floatPayload) Find(needle any) int {
	return find(needle, p.data, p.na, p.convertComparator)
}

func (p *floatPayload) FindAll(needle interface{}) []int {
	return findAll(needle, p.data, p.na, p.convertComparator)
}

func (p *floatPayload) Eq(val interface{}) []bool {
	return eq(val, p.data, p.na, p.convertComparator)
}

func (p *floatPayload) Neq(val interface{}) []bool {
	return neq(val, p.data, p.na, p.convertComparator)
}

func (p *floatPayload) Gt(val interface{}) []bool {
	return gt(val, p.data, p.na, p.convertComparator)
}

func (p *floatPayload) Lt(val interface{}) []bool {
	return lt(val, p.data, p.na, p.convertComparator)
}

func (p *floatPayload) Gte(val interface{}) []bool {
	return gte(val, p.data, p.na, p.convertComparator)
}

func (p *floatPayload) Lte(val interface{}) []bool {
	return lte(val, p.data, p.na, p.convertComparator)
}

func (p *floatPayload) convertComparator(val interface{}) (float64, bool) {
	var v float64
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
	default:
		ok = false
	}

	return v, ok
}

func (p *floatPayload) Groups() ([][]int, []interface{}) {
	groups, values := groupsForData(p.data, p.na)

	return groups, values
}

func (p *floatPayload) IsUnique() []bool {
	booleans := make([]bool, p.length)

	valuesMap := map[float64]bool{}
	wasNA := false
	wasNaN := false
	wasInfPlus := false
	wasInfMinus := false
	for i := 0; i < p.length; i++ {
		is := false

		if p.na[i] {
			if !wasNA {
				is = true
				wasNA = true
			}
		} else if math.IsNaN(p.data[i]) {
			if !wasNaN {
				is = true
				wasNaN = true
			}
		} else if math.IsInf(p.data[i], 1) {
			if !wasInfPlus {
				is = true
				wasInfPlus = true
			}
		} else if math.IsInf(p.data[i], -1) {
			if !wasInfMinus {
				is = true
				wasInfMinus = true
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

func (p *floatPayload) Coalesce(payload Payload) Payload {
	if p.length != payload.Len() {
		payload = payload.Adjust(p.length)
	}

	var srcData []float64
	var srcNA []bool

	if same, ok := payload.(*floatPayload); ok {
		srcData = same.data
		srcNA = same.na
	} else if floatable, ok := payload.(Floatable); ok {
		srcData, srcNA = floatable.Floats()
	} else {
		return p
	}

	dstData := make([]float64, p.length)
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

	return FloatPayload(dstData, dstNA, p.Options()...)
}

func (p *floatPayload) Options() []Option {
	return []Option{
		OptionPrecision(p.printer.Precision),
	}
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

	payload := &floatPayload{
		length:  length,
		data:    vecData,
		printer: printer,
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

func FloatWithNA(data []float64, na []bool, options ...Option) Vector {
	return New(FloatPayload(data, na, options...), options...)
}

func Float(data []float64, options ...Option) Vector {
	return FloatWithNA(data, nil, options...)
}
