package vector

import (
	"math"
	"strconv"
)

const maxIntPrint = 5

// integer is a structure, subsisting integer vectors
type integer struct {
	length int
	data   []int
	DefNA
}

func (p *integer) Len() int {
	return p.length
}

func (p *integer) ByIndices(indices []int) Payload {
	data := make([]int, 1, len(indices)+1)
	na := make([]bool, 1, len(indices)+1)

	for _, idx := range indices {
		data = append(data, p.data[idx])
		na = append(na, p.na[idx])
	}

	length := len(data) - 1

	return &integer{
		length: length,
		data:   data,
		DefNA: DefNA{
			length: length,
			na:     na,
		},
	}
}

func (p *integer) SupportsFilter(selector interface{}) bool {
	if _, ok := selector.([]int); ok {
		return true
	}

	if _, ok := selector.(func(int, int, bool) bool); ok {
		return true
	}

	return false
}

func (p *integer) Filter(filter interface{}) []bool {
	if byFunc, ok := filter.(func(int, int, bool) bool); ok {
		return p.selectByFunc(byFunc)
	}

	return make([]bool, p.length)
}

func (p *integer) selectByFunc(byFunc func(int, int, bool) bool) []bool {
	booleans := make([]bool, p.length)

	for idx, val := range p.data {
		if byFunc(idx, val, p.na[idx]) {
			booleans[idx-1] = true
		}
	}

	return booleans
}

func (p *integer) Integers() ([]int, []bool) {
	if p.length == 0 {
		return []int{}, []bool{}
	}

	data := make([]int, p.length)
	for i := 1; i <= p.length; i++ {
		if p.na[i] {
			data[i-1] = 0
		} else {
			data[i-1] = p.data[i]
		}
	}

	na := make([]bool, p.Len())
	copy(na, p.na[1:])

	return data, na
}

func (p *integer) Floats() ([]float64, []bool) {
	if p.length == 0 {
		return []float64{}, nil
	}

	data := make([]float64, p.length)

	for i := 1; i <= p.length; i++ {
		if p.na[i] {
			data[i-1] = math.NaN()
		} else {
			data[i-1] = float64(p.data[i])
		}
	}

	na := make([]bool, p.Len())
	copy(na, p.na[1:])

	return data, na
}

func (p *integer) Booleans() ([]bool, []bool) {
	if p.length == 0 {
		return []bool{}, nil
	}

	data := make([]bool, p.length)

	for i := 1; i <= p.length; i++ {
		if p.na[i] {
			data[i-1] = false
		} else {
			data[i-1] = p.data[i] != 0
		}
	}

	na := make([]bool, p.Len())
	copy(na, p.na[1:])

	return data, na
}

func (p *integer) Strings() ([]string, []bool) {
	if p.length == 0 {
		return []string{}, nil
	}

	data := make([]string, p.length)

	for i := 1; i <= p.length; i++ {
		if p.na[i] {
			data[i-1] = ""
		} else {
			data[i-1] = strconv.Itoa(p.data[i])
		}
	}

	na := make([]bool, p.Len())
	copy(na, p.na[1:])

	return data, na
}

func (p *integer) StrForElem(idx int) string {
	if p.na[idx] {
		return "NA"
	}

	return strconv.Itoa(p.data[idx])
}

func Integer(data []int, na []bool, options ...Config) Vector {
	length := len(data)

	vecData := make([]int, length+1)
	if length > 0 {
		copy(vecData[1:], data)
	}

	vecNA := make([]bool, length+1)
	if len(na) > 0 {
		if len(na) == length {
			copy(vecNA[1:], na)
		} else {
			emp := Empty()
			emp.Report().AddError("integerPayload(): data length is not equal to na's length")
			return emp
		}
	}

	payload := &integer{
		length: length,
		data:   vecData,
		DefNA: DefNA{
			length: length,
			na:     vecNA,
		},
	}

	return New(payload, options...)
}
