package vector

import (
	"math"
	"math/cmplx"
	"time"
)

type InterfaceConvertors struct {
	Intabler     func(idx int, val interface{}, na bool) (int, bool)
	Floatabler   func(idx int, val interface{}, na bool) (float64, bool)
	Complexabler func(idx int, val interface{}, na bool) (complex128, bool)
	Boolabler    func(idx int, val interface{}, na bool) (bool, bool)
	Stringabler  func(idx int, val interface{}, na bool) (string, bool)
	Timeabler    func(idx int, val interface{}, na bool) (time.Time, bool)
}

type interfacePayload struct {
	length     int
	data       []interface{}
	printer    func(payload interface{}) string
	convertors InterfaceConvertors
	DefNAble
}

func (p *interfacePayload) Type() string {
	return "interface"
}

func (p *interfacePayload) Len() int {
	return p.length
}

func (p *interfacePayload) ByIndices(indices []int) Payload {
	data := make([]interface{}, 0, len(indices))
	na := make([]bool, 0, len(indices))

	for _, idx := range indices {
		data = append(data, p.data[idx-1])
		na = append(na, p.na[idx-1])
	}

	return &interfacePayload{
		length:  len(data),
		data:    data,
		printer: p.printer,
		DefNAble: DefNAble{
			na: na,
		},
	}
}

func (p *interfacePayload) StrForElem(idx int) string {
	if p.na[idx-1] {
		return "NA"
	}

	if p.printer != nil {
		return p.printer(p.data[idx-1])
	}

	return ""
}

func (p *interfacePayload) NAPayload() Payload {
	data := make([]interface{}, p.length)
	na := make([]bool, p.length)
	for i := 0; i < p.length; i++ {
		data[i] = nil
		na[i] = true
	}

	return &interfacePayload{
		length:  p.length,
		data:    data,
		printer: p.printer,
		DefNAble: DefNAble{
			na: na,
		},
	}
}

func (p *interfacePayload) SupportsWhicher(whicher interface{}) bool {
	if _, ok := whicher.(func(int, interface{}, bool) bool); ok {
		return true
	}

	return false
}

func (p *interfacePayload) Which(whicher interface{}) []bool {
	if byFunc, ok := whicher.(func(int, interface{}, bool) bool); ok {
		return p.selectByFunc(byFunc)
	}

	return make([]bool, p.length)
}

func (p *interfacePayload) selectByFunc(byFunc func(int, interface{}, bool) bool) []bool {
	booleans := make([]bool, p.length)

	for idx, val := range p.data {
		if byFunc(idx+1, val, p.na[idx]) {
			booleans[idx] = true
		}
	}

	return booleans
}

func (p *interfacePayload) SupportsApplier(applier interface{}) bool {
	if _, ok := applier.(func(int, interface{}, bool) (interface{}, bool)); ok {
		return true
	}

	return false
}

func (p *interfacePayload) Apply(applier interface{}) Payload {
	var data []interface{}
	var na []bool

	if applyFunc, ok := applier.(func(int, interface{}, bool) (interface{}, bool)); ok {
		data, na = p.applyByFunc(applyFunc)
	} else {
		return p.NAPayload()
	}

	return &interfacePayload{
		length:  p.length,
		data:    data,
		printer: p.printer,
		DefNAble: DefNAble{
			na: na,
		},
	}
}

func (p *interfacePayload) applyByFunc(applyFunc func(int, interface{}, bool) (interface{}, bool)) ([]interface{}, []bool) {
	data := make([]interface{}, p.length)
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

func (p *interfacePayload) Integers() ([]int, []bool) {
	if p.length == 0 {
		return []int{}, []bool{}
	}

	data := make([]int, p.length)
	na := make([]bool, p.Len())

	for i := 0; i < p.length; i++ {
		val, naVal := 0, true
		if p.convertors.Intabler != nil {
			val, naVal = p.convertors.Intabler(i+1, p.data[i], p.na[i])
		}
		data[i] = val
		na[i] = naVal
	}

	return data, na
}

func (p *interfacePayload) Floats() ([]float64, []bool) {
	if p.length == 0 {
		return []float64{}, []bool{}
	}

	data := make([]float64, p.length)
	na := make([]bool, p.Len())

	for i := 0; i < p.length; i++ {
		val, naVal := math.NaN(), true
		if p.convertors.Floatabler != nil {
			val, naVal = p.convertors.Floatabler(i+1, p.data[i], p.na[i])
		}
		data[i] = val
		na[i] = naVal
	}

	return data, na
}

func (p *interfacePayload) Complexes() ([]complex128, []bool) {
	if p.length == 0 {
		return []complex128{}, []bool{}
	}

	data := make([]complex128, p.length)
	na := make([]bool, p.Len())

	for i := 0; i < p.length; i++ {
		val, naVal := cmplx.NaN(), true
		if p.convertors.Complexabler != nil {
			val, naVal = p.convertors.Complexabler(i+1, p.data[i], p.na[i])
		}
		data[i] = val
		na[i] = naVal
	}

	return data, na
}

func (p *interfacePayload) Booleans() ([]bool, []bool) {
	if p.length == 0 {
		return []bool{}, []bool{}
	}

	data := make([]bool, p.length)
	na := make([]bool, p.Len())

	for i := 0; i < p.length; i++ {
		val, naVal := false, true
		if p.convertors.Boolabler != nil {
			val, naVal = p.convertors.Boolabler(i+1, p.data[i], p.na[i])
		}
		data[i] = val
		na[i] = naVal
	}

	return data, na
}

func (p *interfacePayload) Strings() ([]string, []bool) {
	if p.length == 0 {
		return []string{}, []bool{}
	}

	data := make([]string, p.length)
	na := make([]bool, p.Len())

	for i := 0; i < p.length; i++ {
		val, naVal := "", true
		if p.convertors.Stringabler != nil {
			val, naVal = p.convertors.Stringabler(i+1, p.data[i], p.na[i])
		}
		data[i] = val
		na[i] = naVal
	}

	return data, na
}

func (p *interfacePayload) Times() ([]time.Time, []bool) {
	if p.length == 0 {
		return []time.Time{}, []bool{}
	}

	data := make([]time.Time, p.length)
	na := make([]bool, p.Len())

	for i := 0; i < p.length; i++ {
		val, naVal := time.Time{}, true
		if p.convertors.Timeabler != nil {
			val, naVal = p.convertors.Timeabler(i+1, p.data[i], p.na[i])
		}
		data[i] = val
		na[i] = naVal
	}

	return data, na
}
