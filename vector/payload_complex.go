package vector

import (
	"logarithmotechnia/embed"
	"math"
	"math/cmplx"
	"strconv"
)

type ComplexPrinter struct {
	Precision int
}

type complexPayload struct {
	length  int
	data    []complex128
	printer ComplexPrinter
	embed.NAble
}

func (p *complexPayload) Type() string {
	return "complex"
}

func (p *complexPayload) Len() int {
	return p.length
}

func (p *complexPayload) Pick(idx int) any {
	return pickValueWithNA(idx, p.data, p.NA, p.length)
}

func (p *complexPayload) Data() []any {
	return dataWithNAToInterfaceArray(p.data, p.NA)
}

func (p *complexPayload) ByIndices(indices []int) Payload {
	data, na := byIndicesWithNA(indices, p.data, p.NA, cmplx.NaN())

	return ComplexPayload(data, na, p.Options()...)
}

func (p *complexPayload) SupportsWhicher(whicher any) bool {
	return supportsWhicherWithNA[complex128](whicher)
}

func (p *complexPayload) Which(whicher any) []bool {
	return whichWithNA(p.data, p.NA, whicher)
}

func (p *complexPayload) Apply(applier any) Payload {
	return applyWithNA(p.data, p.NA, applier, p.Options())
}

func (p *complexPayload) ApplyTo(indices []int, applier any) Payload {
	data, na := applyToWithNA(indices, p.data, p.NA, applier, cmplx.NaN())

	if data == nil {
		return NAPayload(p.length)
	}

	return ComplexPayload(data, na, p.Options()...)
}

func (p *complexPayload) Traverse(traverser any) {
	traverseWithNA(p.data, p.NA, traverser)
}

func (p *complexPayload) SupportsSummarizer(summarizer any) bool {
	return supportsSummarizer[complex128](summarizer)
}

func (p *complexPayload) Summarize(summarizer any) Payload {
	val, na := summarize(p.data, p.NA, summarizer, 0+0i, cmplx.NaN())

	return ComplexPayload([]complex128{val}, []bool{na}, p.Options()...)
}

func (p *complexPayload) Integers() ([]int, []bool) {
	if p.length == 0 {
		return []int{}, []bool{}
	}

	data := make([]int, p.length)
	for i := 0; i < p.length; i++ {
		if p.NA[i] {
			data[i] = 0
		} else {
			data[i] = int(real(p.data[i]))
		}
	}

	na := make([]bool, p.Len())
	copy(na, p.NA)

	return data, na
}

func (p *complexPayload) Floats() ([]float64, []bool) {
	if p.length == 0 {
		return []float64{}, []bool{}
	}

	data := make([]float64, p.length)

	for i := 0; i < p.length; i++ {
		if p.NA[i] {
			data[i] = math.NaN()
		} else {
			data[i] = real(p.data[i])
		}
	}

	na := make([]bool, p.Len())
	copy(na, p.NA)

	return data, na
}

func (p *complexPayload) Complexes() ([]complex128, []bool) {
	if p.length == 0 {
		return []complex128{}, []bool{}
	}

	data := make([]complex128, p.length)
	copy(data, p.data)

	na := make([]bool, p.Len())
	copy(na, p.NA)

	return data, na
}

func (p *complexPayload) Booleans() ([]bool, []bool) {
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

func (p *complexPayload) Strings() ([]string, []bool) {
	if p.length == 0 {
		return []string{}, []bool{}
	}

	data := make([]string, p.length)

	for i := 0; i < p.length; i++ {
		data[i] = p.StrForElem(i + 1)
	}

	na := make([]bool, p.Len())
	copy(na, p.NA)

	return data, na
}

func (p *complexPayload) Anies() ([]any, []bool) {
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
	copy(newNA, p.NA)
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
	data, na := adjustToLesserSizeWithNA(p.data, p.NA, size)

	return ComplexPayload(data, na, p.Options()...)
}

func (p *complexPayload) adjustToBiggerSize(size int) Payload {
	data, na := adjustToBiggerSizeWithNA(p.data, p.NA, p.length, size)

	return ComplexPayload(data, na, p.Options()...)
}

func (p *complexPayload) Options() []Option {
	return []Option{
		ConfOption{keyOptionPrecision, p.printer.Precision},
	}
}

func (p *complexPayload) SetOption(name string, val any) bool {
	switch name {
	case keyOptionPrecision:
		p.printer.Precision = val.(int)
	default:
		return false
	}

	return true
}

func (p *complexPayload) Groups() ([][]int, []any) {
	groups, values := groupsForData(p.data, p.NA)

	return groups, values
}

func (p *complexPayload) StrForElem(idx int) string {
	i := idx - 1

	if p.NA[i] {
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

func (p *complexPayload) Find(needle any) int {
	return find(needle, p.data, p.NA, p.convertComparator)
}

func (p *complexPayload) FindAll(needle any) []int {
	return findAll(needle, p.data, p.NA, p.convertComparator)
}

/* Ordered interface */

func (p *complexPayload) Eq(val any) []bool {
	return eq(val, p.data, p.NA, p.convertComparator)
}

func (p *complexPayload) Neq(val any) []bool {
	return neq(val, p.data, p.NA, p.convertComparator)
}

func (p *complexPayload) convertComparator(val any) (complex128, bool) {
	var v complex128
	ok := true
	switch value := val.(type) {
	case complex128:
		v = value
	case complex64:
		v = complex128(value)
	case float64:
		v = complex(value, 0)
	case float32:
		v = complex(float64(value), 0)
	case int:
		v = complex(float64(value), 0)
	case int64:
		v = complex(float64(value), 0)
	case int32:
		v = complex(float64(value), 0)
	case uint64:
		v = complex(float64(value), 0)
	case uint32:
		v = complex(float64(value), 0)
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
		srcNA = same.NA
	} else if complexable, ok := payload.(Complexable); ok {
		srcData, srcNA = complexable.Complexes()
	} else {
		return p
	}

	dstData := make([]complex128, p.length)
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

	return ComplexPayload(dstData, dstNA, p.Options()...)
}

func (p *complexPayload) IsUnique() []bool {
	booleans := make([]bool, p.length)

	valuesMap := map[complex128]bool{}
	wasNA := false
	wasNaN := false
	wasInf := false
	for i := 0; i < p.length; i++ {
		is := false

		if p.NA[i] {
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

// ComplexPayload creates a payload with complex128 data.
//
// Available options are:
//   - OptionPrecision(precision int) - sets precision for printing payload's values.
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

	payload := &complexPayload{
		length:  length,
		data:    vecData,
		printer: printer,
		NAble: embed.NAble{
			NA: vecNA,
		},
	}
	conf.SetOptions(payload)

	return payload
}

// ComplexWithNA creates a vector with ComplexPayload and allows to set NA-values.
func ComplexWithNA(data []complex128, na []bool, options ...Option) Vector {
	return New(ComplexPayload(data, na, options...), options...)
}

// Complex creates a vector with ComplexPayload.
func Complex(data []complex128, options ...Option) Vector {
	return ComplexWithNA(data, nil, options...)
}
