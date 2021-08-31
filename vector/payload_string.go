package vector

import (
	"math"
	"math/cmplx"
	"strconv"
)

type stringPayload struct {
	length int
	data   []string
	DefNAble
}

func (p *stringPayload) Type() string {
	return "string"
}

func (p *stringPayload) Len() int {
	return p.length
}

func (p *stringPayload) ByIndices(indices []int) Payload {
	data := make([]string, 0, len(indices))
	na := make([]bool, 0, len(indices))

	for _, idx := range indices {
		data = append(data, p.data[idx-1])
		na = append(na, p.na[idx-1])
	}

	return &stringPayload{
		length: len(data),
		data:   data,
		DefNAble: DefNAble{
			na: na,
		},
	}
}

func (p *stringPayload) SupportsWhicher(whicher interface{}) bool {
	if _, ok := whicher.(func(int, string, bool) bool); ok {
		return true
	}

	if _, ok := whicher.(func(string, bool) bool); ok {
		return true
	}

	return false
}

func (p *stringPayload) Which(whicher interface{}) []bool {
	if byFunc, ok := whicher.(func(int, string, bool) bool); ok {
		return p.selectByFunc(byFunc)
	}

	if byFunc, ok := whicher.(func(string, bool) bool); ok {
		return p.selectByCompactFunc(byFunc)
	}

	return make([]bool, p.length)
}

func (p *stringPayload) selectByFunc(byFunc func(int, string, bool) bool) []bool {
	booleans := make([]bool, p.length)

	for idx, val := range p.data {
		if byFunc(idx+1, val, p.na[idx]) {
			booleans[idx] = true
		}
	}

	return booleans
}

func (p *stringPayload) selectByCompactFunc(byFunc func(string, bool) bool) []bool {
	booleans := make([]bool, p.length)

	for idx, val := range p.data {
		if byFunc(val, p.na[idx]) {
			booleans[idx] = true
		}
	}

	return booleans
}

func (p *stringPayload) SupportsApplier(applier interface{}) bool {
	if _, ok := applier.(func(int, string, bool) (string, bool)); ok {
		return true
	}

	return false
}

func (p *stringPayload) Apply(applier interface{}) Payload {
	var data []string
	var na []bool

	if applyFunc, ok := applier.(func(int, string, bool) (string, bool)); ok {
		data, na = p.applyByFunc(applyFunc)
	} else {
		return NAPayload(p.length)
	}

	return &stringPayload{
		length: p.length,
		data:   data,
		DefNAble: DefNAble{
			na: na,
		},
	}
}

func (p *stringPayload) applyByFunc(applyFunc func(int, string, bool) (string, bool)) ([]string, []bool) {
	data := make([]string, p.length)
	na := make([]bool, p.length)

	for i := 0; i < p.length; i++ {
		dataVal, naVal := applyFunc(i+1, p.data[i], p.na[i])
		if naVal {
			dataVal = ""
		}
		data[i] = dataVal
		na[i] = naVal
	}

	return data, na
}

func (p *stringPayload) SupportsSummarizer(summarizer interface{}) bool {
	if _, ok := summarizer.(func(int, string, string, bool) (string, bool)); ok {
		return true
	}

	return false
}

func (p *stringPayload) Summarize(summarizer interface{}) Payload {
	fn, ok := summarizer.(func(int, string, string, bool) (string, bool))
	if !ok {
		return NAPayload(1)
	}

	val := ""
	na := false
	for i := 0; i < p.length; i++ {
		val, na = fn(i+1, val, p.data[i], p.na[i])

		if na {
			return NAPayload(1)
		}
	}

	return StringPayload([]string{val}, nil)
}

func (p *stringPayload) Integers() ([]int, []bool) {
	if p.length == 0 {
		return []int{}, []bool{}
	}

	data := make([]int, p.length)
	for i := 0; i < p.length; i++ {
		if p.na[i] {
			data[i] = 0
		} else {
			num, err := strconv.ParseFloat(p.data[i], 64)
			if err != nil {
				data[i] = 0
			} else {
				data[i] = int(num)
			}
		}
	}

	na := make([]bool, p.Len())
	copy(na, p.na)

	return data, na
}

func (p *stringPayload) Floats() ([]float64, []bool) {
	if p.length == 0 {
		return []float64{}, []bool{}
	}

	data := make([]float64, p.length)

	for i := 0; i < p.length; i++ {
		if p.na[i] {
			data[i] = math.NaN()
		} else {
			num, err := strconv.ParseFloat(p.data[i], 64)
			if err != nil {
				data[i] = 0
			} else {
				data[i] = num
			}
		}
	}

	na := make([]bool, p.Len())
	copy(na, p.na)

	return data, na
}

func (p *stringPayload) Complexes() ([]complex128, []bool) {
	if p.length == 0 {
		return []complex128{}, []bool{}
	}

	data := make([]complex128, p.length)
	for i := 0; i < p.length; i++ {
		if p.na[i] {
			data[i] = cmplx.NaN()
		} else {
			num, err := strconv.ParseComplex(p.data[i], 128)
			if err != nil {
				data[i] = cmplx.NaN()
			} else {
				data[i] = num
			}
		}
	}

	na := make([]bool, p.Len())
	copy(na, p.na)

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

func (p *stringPayload) Append(vec Vector) Payload {
	length := p.length + vec.Len()

	vals, na := vec.Strings()

	newVals := make([]string, length)
	newNA := make([]bool, length)

	copy(newVals, p.data)
	copy(newVals[p.length:], vals)
	copy(newNA, p.na)
	copy(newNA[p.length:], na)

	return StringPayload(newVals, newNA)
}

func (p *stringPayload) StrForElem(idx int) string {
	if p.na[idx-1] {
		return "NA"
	}

	return p.data[idx-1]
}

func StringPayload(data []string, na []bool) Payload {
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

	return &stringPayload{
		length: length,
		data:   vecData,
		DefNAble: DefNAble{
			na: vecNA,
		},
	}
}

func String(data []string, na []bool, options ...Config) Vector {
	config := mergeConfigs(options)

	return New(StringPayload(data, na), config)
}
