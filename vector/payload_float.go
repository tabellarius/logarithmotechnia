package vector

import (
	"math"
	"math/cmplx"
	"strconv"
)

type FloatPrinter struct {
	Precision int
}

type floatPayload struct {
	length  int
	data    []float64
	printer FloatPrinter
	DefNAble
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
		length: len(data),
		data:   data,
		DefNAble: DefNAble{
			na: na,
		},
	}
}

func (p *floatPayload) SupportsSelector(selector interface{}) bool {
	if _, ok := selector.(func(int, float64, bool) bool); ok {
		return true
	}

	return false
}

func (p *floatPayload) Select(selector interface{}) []bool {
	if byFunc, ok := selector.(func(int, float64, bool) bool); ok {
		return p.selectByFunc(byFunc)
	}

	return make([]bool, p.length)
}

func (p *floatPayload) selectByFunc(byFunc func(int, float64, bool) bool) []bool {
	booleans := make([]bool, p.length)

	for idx, val := range p.data {
		if byFunc(idx+1, val, p.na[idx]) {
			booleans[idx] = true
		}
	}

	return booleans
}

func (p *floatPayload) SupportsApplier(applier interface{}) bool {
	if _, ok := applier.(func(int, float64, bool) (float64, bool)); ok {
		return true
	}

	return false
}

func (p *floatPayload) Apply(applier interface{}) Payload {
	var data []float64
	var na []bool

	if applyFunc, ok := applier.(func(int, float64, bool) (float64, bool)); ok {
		data, na = p.applyByFunc(applyFunc)
	} else {
		return p.NAPayload()
	}

	return &floatPayload{
		length:  p.length,
		data:    data,
		printer: p.printer,
		DefNAble: DefNAble{
			na: na,
		},
	}
}

func (p *floatPayload) applyByFunc(applyFunc func(int, float64, bool) (float64, bool)) ([]float64, []bool) {
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

	return data, na
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
		return []float64{}, nil
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
		return []bool{}, nil
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
		return []string{}, nil
	}

	data := make([]string, p.length)

	for i := 0; i < p.length; i++ {
		data[i] = p.StrForElem(i + 1)
	}

	na := make([]bool, p.Len())
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

func (p *floatPayload) NAPayload() Payload {
	data := make([]float64, p.length)
	na := make([]bool, p.length)
	for i := 0; i < p.length; i++ {
		data[i] = math.NaN()
		na[i] = true
	}

	return &floatPayload{
		length:  p.length,
		data:    data,
		printer: p.printer,
		DefNAble: DefNAble{
			na: na,
		},
	}
}

func Float(data []float64, na []bool, options ...Config) Vector {
	config := mergeConfigs(options)

	length := len(data)

	vecNA := make([]bool, length)
	if len(na) > 0 {
		if len(na) == length {
			copy(vecNA, na)
		} else {
			emp := Empty()
			emp.Report().AddError("Float(): data length is not equal to na's length")
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

	printer := FloatPrinter{Precision: 3}
	if config.FloatPrinter != nil {
		printer = *config.FloatPrinter
	}

	payload := &floatPayload{
		length:  length,
		data:    vecData,
		printer: printer,
		DefNAble: DefNAble{
			na: vecNA,
		},
	}

	return New(payload, config)
}
