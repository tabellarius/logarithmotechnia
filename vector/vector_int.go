package vector

import (
	"fmt"
	"strconv"
)

// IntegerVector is an integer vector
type IntegerVector interface {
	Vector
	Nameable
	NAble
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
	DefNameable
	DefNAble
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

func (v *integer) String() string {
	str := "["

	if v.length > 0 {
		str += v.strForElem(1)
	}
	if v.length > 1 {
		for i := 2; i <= v.length; i++ {
			if i <= maxIntPrint {
				str += ", " + v.strForElem(i)
			} else {
				str += ", ..."
				break
			}
		}
	}

	str += "]"

	return str
}

func (v *integer) strForElem(idx int) string {
	str := strconv.Itoa(v.data[idx])
	if v.DefNameable.HasNameFor(idx) {
		str += " (" + v.DefNameable.Name(idx) + ")"
	}
	return str
}

// NewIntegerVector creates a new integer vector
func NewIntegerVector(data []int, options ...Config) IntegerVector {
	config := mergeConfigs(options)

	length := len(data)
	com := newCommon(length)

	vecData := make([]int, length+1)
	if length > 0 {
		copy(vecData[1:], data)
	}

	vector := &integer{
		common: com,
		data:   vecData,
	}

	com.vec = vector

	nameable, nable := newNamesAndNAble(vector, config)
	vector.DefNameable = nameable
	vector.DefNAble = nable

	return vector
}
