package vector

import (
	"fmt"
	"strconv"
)

// IntegerVector is an integer vector
type IntegerVector interface {
	Vector
	Intable
	Floatable
	Booleable
	Stringable
	fmt.Stringer
}

const maxIntPrint = 5

// integer is a structure, subsisting integer vectors
type integer struct {
	common
	data []int
}

func (v *integer) Integers() []int {
	if v.length == 0 {
		return []int{}
	}

	arr := make([]int, v.length)
	copy(arr, v.data[1:])

	return arr
}

func (v *integer) Floats() []float64 {
	if v.length == 0 {
		return []float64{}
	}

	arr := make([]float64, v.length)

	for i := 1; i <= v.length; i++ {
		arr[i-1] = float64(v.data[i])
	}

	return arr
}

func (v *integer) Booleans() []bool {
	if v.length == 0 {
		return []bool{}
	}

	arr := make([]bool, v.length)

	for i := 1; i <= v.length; i++ {
		arr[i-1] = v.data[i] != 0
	}

	return arr
}

func (v *integer) Strings() []string {
	if v.length == 0 {
		return []string{}
	}

	arr := make([]string, v.length)

	for i := 1; i <= v.length; i++ {
		arr[i-1] = strconv.Itoa(v.data[i])
	}

	return arr
}

func (v *integer) String() string {
	str := "["

	if v.length > 0 {
		str += strconv.Itoa(v.data[1])
	}
	if v.length > 1 {
		for i := 2; i <= v.length; i++ {
			if i <= maxIntPrint {
				str += ", " + strconv.Itoa(v.data[i])
				if _, ok := v.names[i]; ok {
					str += " (" + v.names[i] + ")"
				}
			} else {
				str += ", ..."
				break
			}
		}
	}

	str += "]"

	return str
}

func (v *integer) Clone() Vector {
	com := v.common.Clone()
	c := com.(*common)

	vec := &integer{
		common: *c,
		data:   v.data,
	}

	return vec
}

func (v *integer) Refresh() {
	v.common.Refresh()

	data := make([]int, len(v.data))
	copy(data, v.data)
}

// NewIntegerVector creates a new integer vector
func NewIntegerVector(data []int, options ...Config) IntegerVector {
	if data == nil {
		data = make([]int, 0)
	}
	length := len(data)
	com := newCommon(length, options...)

	vecData := make([]int, length+1)
	if length > 0 {
		copy(vecData[1:], data)
	}

	vector := &integer{
		common: com,
		data:   vecData,
	}

	return vector
}
