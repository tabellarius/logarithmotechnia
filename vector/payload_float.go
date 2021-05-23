package vector

import (
	"math"
	"strconv"
)

type float struct {
	length int
	data   []float64
	DefNAble
}

func (p *float) ByIndices(indices []int) Payload {
	data := make([]float64, 0, len(indices))
	na := make([]bool, 0, len(indices))

	for _, idx := range indices {
		data = append(data, p.data[idx-1])
		na = append(na, p.na[idx-1])
	}

	return &float{
		length: len(data),
		data:   data,
		DefNAble: DefNAble{
			na: na,
		},
	}
}

func (p *float) SupportsSelector(selector interface{}) bool {
	if _, ok := selector.(func(int, float64, bool) bool); ok {
		return true
	}

	return false
}

func (p *float) Select(selector interface{}) []bool {
	if byFunc, ok := selector.(func(int, float64, bool) bool); ok {
		return p.selectByFunc(byFunc)
	}

	return make([]bool, p.length)
}

func (p *float) selectByFunc(byFunc func(int, float64, bool) bool) []bool {
	booleans := make([]bool, p.length)

	for idx, val := range p.data {
		if byFunc(idx+1, val, p.na[idx]) {
			booleans[idx] = true
		}
	}

	return booleans
}

func (p *float) Len() int {
	return p.length
}

func (p *float) Integers() ([]int, []bool) {
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

func (p *float) Floats() ([]float64, []bool) {
	if p.length == 0 {
		return []float64{}, nil
	}

	data := make([]float64, p.length)

	for i := 0; i < p.length; i++ {
		if p.na[i] {
			data[i] = math.NaN()
		} else {
			data[i] = p.data[i]
		}
	}

	na := make([]bool, p.Len())
	copy(na, p.na)

	return data, na
}

func (p *float) Complexes() ([]complex128, []bool) {
	if p.length == 0 {
		return []complex128{}, []bool{}
	}

	data := make([]complex128, p.length)
	for i := 0; i < p.length; i++ {
		if p.na[i] {
			data[i] = 0
		} else {
			data[i] = complex(p.data[i], 0)
		}
	}

	na := make([]bool, p.Len())
	copy(na, p.na)

	return data, na
}

func (p *float) Booleans() ([]bool, []bool) {
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

func (p *float) Strings() ([]string, []bool) {
	if p.length == 0 {
		return []string{}, nil
	}

	data := make([]string, p.length)

	for i := 0; i < p.length; i++ {
		if p.na[i] {
			data[i] = ""
		} else {
			data[i] = p.elemToStr(i)
		}
	}

	na := make([]bool, p.Len())
	copy(na, p.na)

	return data, na
}

func (p *float) elemToStr(i int) string {
	if math.IsNaN(p.data[i]) {
		return "NaN"
	}

	return strconv.FormatFloat(p.data[i], 'f', 3, 64)
}

func (p *float) StrForElem(idx int) string {
	if p.na[idx-1] {
		return "NA"
	}

	return p.elemToStr(idx - 1)
}

func Float(data []float64, na []bool, options ...Config) Vector {
	length := len(data)

	vecData := make([]float64, length)
	if length > 0 {
		copy(vecData, data)
	}

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

	payload := &float{
		length: length,
		data:   vecData,
		DefNAble: DefNAble{
			na: vecNA,
		},
	}

	return New(payload, options...)
}
