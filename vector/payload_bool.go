package vector

import (
	"math"
	"math/cmplx"
)

type booleanPayload struct {
	length int
	data   []bool
	DefNAble
}

func (p *booleanPayload) Type() string {
	return "boolean"
}

func (p *booleanPayload) Len() int {
	return p.length
}

func (p *booleanPayload) ByIndices(indices []int) Payload {
	data := make([]bool, 0, len(indices))
	na := make([]bool, 0, len(indices))

	for _, idx := range indices {
		data = append(data, p.data[idx-1])
		na = append(na, p.na[idx-1])
	}

	return &booleanPayload{
		length: len(data),
		data:   data,
		DefNAble: DefNAble{
			na: na,
		},
	}
}

func (p *booleanPayload) SupportsWhicher(whicher interface{}) bool {
	if _, ok := whicher.(func(int, bool, bool) bool); ok {
		return true
	}

	return false
}

func (p *booleanPayload) Which(whicher interface{}) []bool {
	if byFunc, ok := whicher.(func(int, bool, bool) bool); ok {
		return p.selectByFunc(byFunc)
	}

	return make([]bool, p.length)
}

func (p *booleanPayload) selectByFunc(byFunc func(int, bool, bool) bool) []bool {
	booleans := make([]bool, p.length)

	for idx, val := range p.data {
		if byFunc(idx+1, val, p.na[idx]) {
			booleans[idx] = true
		}
	}

	return booleans
}

func (p *booleanPayload) SupportsApplier(applier interface{}) bool {
	if _, ok := applier.(func(int, bool, bool) (bool, bool)); ok {
		return true
	}

	return false
}

func (p *booleanPayload) Apply(applier interface{}) Payload {
	var data []bool
	var na []bool

	if applyFunc, ok := applier.(func(int, bool, bool) (bool, bool)); ok {
		data, na = p.applyByFunc(applyFunc)
	} else {
		return NAPayload(p.length)
	}

	return &booleanPayload{
		length: p.length,
		data:   data,
		DefNAble: DefNAble{
			na: na,
		},
	}
}

func (p *booleanPayload) applyByFunc(applyFunc func(int, bool, bool) (bool, bool)) ([]bool, []bool) {
	data := make([]bool, p.length)
	na := make([]bool, p.length)

	for i := 0; i < p.length; i++ {
		dataVal, naVal := applyFunc(i+1, p.data[i], p.na[i])
		if naVal {
			dataVal = false
		}
		data[i] = dataVal
		na[i] = naVal
	}

	return data, na
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
		data[i] = interface{}(p.data[i])
	}

	na := make([]bool, p.length)
	copy(na, p.na)

	return data, na
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

func Boolean(data []bool, na []bool, options ...Config) Vector {
	config := mergeConfigs(options)

	length := len(data)

	vecNA := make([]bool, length)
	if len(na) > 0 {
		if len(na) == length {
			copy(vecNA, na)
		} else {
			emp := Empty()
			emp.Report().AddError("Boolean(): data length is not equal to na's length")
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

	return New(payload, config)
}
