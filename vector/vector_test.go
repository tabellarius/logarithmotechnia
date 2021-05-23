package vector

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math"
	"math/cmplx"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestVector_Len(t *testing.T) {
	testData := []struct {
		in             []int
		expectedLength int
	}{
		{
			in:             []int{1, 2, 3, 4, 5},
			expectedLength: 5,
		},
		{
			in:             []int{1, 2, 3},
			expectedLength: 3,
		},
		{
			in:             []int{1},
			expectedLength: 1,
		},
		{
			in:             []int{},
			expectedLength: 0,
		},
		{
			in:             nil,
			expectedLength: 0,
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			vec := Integer(data.in, nil)
			if vec.Len() != data.expectedLength {
				t.Error(fmt.Sprintf("Length (%d) is not equal to expected (%d)",
					vec.Len(), data.expectedLength))
			}
		})
	}
}

func TestVector_Integers(t *testing.T) {
	testData := []struct {
		vec       Vector
		outValues []int
		outNA     []bool
	}{
		{
			vec:       Integer([]int{1, 2, 3, 4, 5}, []bool{false, false, true, false, false}),
			outValues: []int{1, 2, 0, 4, 5},
			outNA:     []bool{false, false, true, false, false},
		},
		{
			vec:       Empty(),
			outValues: []int{},
			outNA:     []bool{},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			integers, na := data.vec.Integers()
			if !reflect.DeepEqual(integers, data.outValues) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", integers, data.outValues))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", na, data.outNA))
			}
		})
	}
}

func TestVector_Floats(t *testing.T) {
	testData := []struct {
		vec       Vector
		outValues []float64
		outNA     []bool
	}{
		{
			vec:       Integer([]int{1, 2, 3, 4, 5}, []bool{false, false, true, false, false}),
			outValues: []float64{1, 2, math.NaN(), 4, 5},
			outNA:     []bool{false, false, true, false, false},
		},
		{
			vec:       Empty(),
			outValues: []float64{},
			outNA:     []bool{},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			floats, na := data.vec.Floats()
			err := false
			for i, val := range floats {
				if math.IsNaN(val) {
					if !math.IsNaN(data.outValues[i]) {
						err = true
					}
				} else if val != data.outValues[i] {
					err = true
				}
			}
			if err {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", floats, data.outValues))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", na, data.outNA))
			}
		})
	}
}

func TestVector_Complexes(t *testing.T) {
	testData := []struct {
		vec       Vector
		outValues []complex128
		outNA     []bool
	}{
		{
			vec:       Integer([]int{1, 2, 3, 4, 5}, []bool{false, false, true, false, false}),
			outValues: []complex128{1 + 0i, 2 + 0i, cmplx.NaN(), 4 + 0i, 5 + 0i},
			outNA:     []bool{false, false, true, false, false},
		},
		{
			vec:       Empty(),
			outValues: []complex128{},
			outNA:     []bool{},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			complexes, na := data.vec.Complexes()
			correct := true
			for i := 0; i < len(complexes); i++ {
				if cmplx.IsNaN(data.outValues[i]) {
					if !cmplx.IsNaN(complexes[i]) {
						correct = false
					}
				} else if complexes[i] != data.outValues[i] {
					correct = false
				}
			}
			if !correct {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", complexes, data.outValues))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", na, data.outNA))
			}
		})
	}
}

func TestVector_Booleans(t *testing.T) {
	testData := []struct {
		vec       Vector
		outValues []bool
		outNA     []bool
	}{
		{
			vec:       Integer([]int{0, 1, 2, 3, 4, 5}, []bool{false, false, false, true, false, false}),
			outValues: []bool{false, true, true, false, true, true},
			outNA:     []bool{false, false, false, true, false, false},
		},
		{
			vec:       Empty(),
			outValues: []bool{},
			outNA:     []bool{},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			booleans, na := data.vec.Booleans()
			if !reflect.DeepEqual(booleans, data.outValues) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", booleans, data.outValues))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", na, data.outNA))
			}
		})
	}
}

func TestVector_Strings(t *testing.T) {
	testData := []struct {
		vec       Vector
		outValues []string
		outNA     []bool
	}{
		{
			vec:       Integer([]int{1, 2, 3, 4, 5}, []bool{false, false, true, false, false}),
			outValues: []string{"1", "2", "", "4", "5"},
			outNA:     []bool{false, false, true, false, false},
		},
		{
			vec:       Empty(),
			outValues: []string{},
			outNA:     []bool{},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			strings, na := data.vec.Strings()
			if !reflect.DeepEqual(strings, data.outValues) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", strings, data.outValues))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", na, data.outNA))
			}
		})
	}
}

func TestVector_Decimals(t *testing.T) {
	testData := []struct {
		vec       Vector
		outValues []decimal.Decimal
		outNA     []bool
	}{
		{
			vec:       Integer([]int{1, 2, 3, 4, 5}, []bool{false, false, true, false, false}),
			outValues: []decimal.Decimal{decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero},
			outNA:     []bool{true, true, true, true, true},
		},
		{
			vec:       Empty(),
			outValues: []decimal.Decimal{},
			outNA:     []bool{},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			decimals, na := data.vec.Decimals()
			err := false
			for i, val := range decimals {
				if !val.Equals(data.outValues[i]) {
					err = true
				}
			}
			if err {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", decimals, data.outValues))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", na, data.outNA))
			}
		})
	}
}

func TestVector_Times(t *testing.T) {
	testData := []struct {
		vec       Vector
		outValues []time.Time
		outNA     []bool
	}{
		{
			vec:       Integer([]int{1, 2, 3, 4, 5}, []bool{false, false, true, false, false}),
			outValues: []time.Time{{}, {}, {}, {}, {}},
			outNA:     []bool{true, true, true, true, true},
		},
		{
			vec:       Empty(),
			outValues: []time.Time{},
			outNA:     []bool{},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			times, na := data.vec.Times()
			if !reflect.DeepEqual(times, data.outValues) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", times, data.outValues))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", na, data.outNA))
			}
		})
	}
}

func TestVector_ByIndices(t *testing.T) {
	vec := Integer([]int{1, 2, 3, 4, 5}, []bool{false, false, true, false, false})
	testData := []struct {
		name    string
		indices []int
		out     []int
		outNA   []bool
		length  int
	}{
		{
			name:    "normal",
			indices: []int{-1, 0, 4, 3, 2, 6},
			out:     []int{4, 0, 2},
			outNA:   []bool{false, true, false},
			length:  3,
		},
		{
			name:    "empty",
			indices: []int{},
			out:     []int{},
			outNA:   []bool{},
			length:  0,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			newVec := vec.ByIndices(data.indices)
			integers, na := newVec.Integers()
			if !reflect.DeepEqual(integers, data.out) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", integers, data.out))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("Result NA (%v) is not equal to expected NA (%v)", na, data.outNA))
			}
			if newVec.Len() != data.length {
				t.Error(fmt.Sprintf("Result length (%d) is not equal to expected length (%d)",
					newVec.Len(), data.length))
			}
		})
	}
}

func TestVector_ByNames(t *testing.T) {
	vec := Integer(
		[]int{1, 2, 3, 4, 5},
		[]bool{false, false, true, false, false},
		OptionNamesMap(map[string]int{"one": 1, "three": 3, "five": 5}),
	)
	testData := []struct {
		name   string
		names  []string
		out    []int
		outNA  []bool
		length int
	}{
		{
			name:   "all",
			names:  []string{"one", "three", "five"},
			out:    []int{1, 0, 5},
			outNA:  []bool{false, true, false},
			length: 3,
		},
		{
			name:   "with incorrect",
			names:  []string{"zero", "three", "five"},
			out:    []int{0, 5},
			outNA:  []bool{true, false},
			length: 2,
		},
		{
			name:   "empty",
			names:  []string{},
			out:    []int{},
			outNA:  []bool{},
			length: 0,
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			newVec := vec.ByNames(data.names)
			integers, na := newVec.Integers()
			if !reflect.DeepEqual(integers, data.out) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", integers, data.out))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("Result NA (%v) is not equal to expected NA (%v)", na, data.outNA))
			}
			if newVec.Len() != data.length {
				t.Error(fmt.Sprintf("Result length (%d) is not equal to expected length (%d)",
					newVec.Len(), data.length))
			}
		})
	}
}

func TestVector_Filter(t *testing.T) {
	vec := Integer(
		[]int{1, 2, 3, 4, 5},
		[]bool{false, false, true, false, false},
		OptionNamesMap(map[string]int{"one": 1, "three": 3, "five": 5}),
	)
	testData := []struct {
		name     string
		selector interface{}
		out      []int
		outNA    []bool
		length   int
	}{
		{
			name:     "indices",
			selector: []int{5, 3, 1},
			out:      []int{5, 0, 1},
			outNA:    []bool{false, true, false},
			length:   3,
		},
		{
			name:     "idx",
			selector: []int{1},
			out:      []int{1},
			outNA:    []bool{false},
			length:   1,
		},
		{
			name:     "FromTo straight",
			selector: FromTo{2, 4},
			out:      []int{2, 0, 4},
			outNA:    []bool{false, true, false},
			length:   3,
		},
		{
			name:     "FromTo reverse",
			selector: FromTo{4, 2},
			out:      []int{4, 0, 2},
			outNA:    []bool{false, true, false},
			length:   3,
		},
		{
			name:     "FromTo with remove straight",
			selector: FromTo{-2, -4},
			out:      []int{1, 5},
			outNA:    []bool{false, false},
			length:   2,
		},
		{
			name:     "FromTo with remove reverse",
			selector: FromTo{-4, -2},
			out:      []int{1, 5},
			outNA:    []bool{false, false},
			length:   2,
		},
		{
			name:     "FromTo straight with incorrect boundaries",
			selector: FromTo{0, 7},
			out:      []int{1, 2, 0, 4, 5},
			outNA:    []bool{false, false, true, false, false},
			length:   5,
		},
		{
			name:     "FromTo full remove with incorrect boundaries",
			selector: FromTo{0, -7},
			out:      []int{},
			outNA:    []bool{},
			length:   0,
		},
		{
			name:     "FromTo different signs",
			selector: FromTo{-2, 4},
			out:      []int{},
			outNA:    []bool{},
			length:   0,
		},
		{
			name:     "FromTo different signs reverse",
			selector: FromTo{2, -4},
			out:      []int{},
			outNA:    []bool{},
			length:   0,
		},
		{
			name:     "boolean",
			selector: []bool{true, false, true, false, true},
			out:      []int{1, 0, 5},
			outNA:    []bool{false, true, false},
			length:   3,
		},
		{
			name:     "string",
			selector: "one",
			out:      []int{1},
			outNA:    []bool{false},
			length:   1,
		},
		{
			name:     "[]string",
			selector: []string{"one", "three"},
			out:      []int{1, 0},
			outNA:    []bool{false, true},
			length:   2,
		},
		{
			name:     "byIntFunc",
			selector: func(_ int, val int, na bool) bool { return !na && val >= 3 },
			out:      []int{4, 5},
			outNA:    []bool{false, false},
			length:   2,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			newVec := vec.Filter(data.selector)
			integers, na := newVec.Integers()
			if !reflect.DeepEqual(integers, data.out) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)", integers, data.out))
			}
			if !reflect.DeepEqual(na, data.outNA) {
				t.Error(fmt.Sprintf("Result NA (%v) is not equal to expected NA (%v)", na, data.outNA))
			}
			if newVec.Len() != data.length {
				t.Error(fmt.Sprintf("Result length (%d) is not equal to expected length (%d)",
					newVec.Len(), data.length))
			}
		})
	}
}

func TestVector_IsNA(t *testing.T) {
	testData := []struct {
		name  string
		vec   Vector
		notNA []bool
	}{
		{
			name:  "with NA",
			vec:   Integer([]int{1, 2, 3}, []bool{false, false, false}),
			notNA: []bool{false, false, false},
		},
		{
			name:  "without NA",
			vec:   Integer([]int{1, 2, 3}, []bool{false, true, false}),
			notNA: []bool{false, true, false},
		},
		{
			name:  "empty",
			vec:   Empty(),
			notNA: []bool{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			_ = data
		})
	}
}

func TestVector_NotNA(t *testing.T) {
	testData := []struct {
		name  string
		vec   Vector
		notNA []bool
	}{
		{
			name:  "with NA",
			vec:   Integer([]int{1, 2, 3}, []bool{false, false, false}),
			notNA: []bool{true, true, true},
		},
		{
			name:  "without NA",
			vec:   Integer([]int{1, 2, 3}, []bool{false, true, false}),
			notNA: []bool{true, false, true},
		},
		{
			name:  "empty",
			vec:   Empty(),
			notNA: []bool{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			_ = data
		})
	}
}

func TestVector_HasNA(t *testing.T) {
	testData := []struct {
		name  string
		vec   Vector
		hasNA bool
	}{
		{
			name:  "regular with nil NA",
			vec:   Integer([]int{1, 2, 3}, nil),
			hasNA: false,
		},
		{
			name:  "regular without NA",
			vec:   Integer([]int{1, 2, 3}, []bool{false, false, false}),
			hasNA: false,
		},
		{
			name:  "regular with NA",
			vec:   Integer([]int{1, 2, 3}, []bool{false, true, false}),
			hasNA: true,
		},
		{
			name:  "Empty",
			vec:   Empty(),
			hasNA: false,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if data.vec.HasNA() != data.hasNA {
				t.Error(fmt.Sprintf("data.vec.HasNA() (%t) is not equal to data.hasNA (%t)",
					data.vec.HasNA(), data.hasNA))
			}
		})
	}
}

func TestVector_WithNA(t *testing.T) {
	testData := []struct {
		name   string
		vec    Vector
		withNA []int
	}{
		{
			name:   "regular with nil NA",
			vec:    Integer([]int{1, 2, 3}, nil),
			withNA: []int{},
		},
		{
			name:   "regular without NA",
			vec:    Integer([]int{1, 2, 3}, []bool{false, false, false}),
			withNA: []int{},
		},
		{
			name:   "regular with NA",
			vec:    Integer([]int{1, 2, 3}, []bool{false, true, true}),
			withNA: []int{2, 3},
		},
		{
			name:   "Empty",
			vec:    Empty(),
			withNA: []int{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if !reflect.DeepEqual(data.vec.WithNA(), data.withNA) {
				t.Error(fmt.Sprintf("data.vec.WithNA() (%v) is not equal to data.WithNA (%v)",
					data.vec.WithNA(), data.withNA))
			}
		})
	}
}

func TestVector_WithoutNA(t *testing.T) {
	testData := []struct {
		name      string
		vec       Vector
		withoutNA []int
	}{
		{
			name:      "regular with nil NA",
			vec:       Integer([]int{1, 2, 3}, nil),
			withoutNA: []int{1, 2, 3},
		},
		{
			name:      "regular without NA",
			vec:       Integer([]int{1, 2, 3}, []bool{false, false, false}),
			withoutNA: []int{1, 2, 3},
		},
		{
			name:      "regular with NA",
			vec:       Integer([]int{1, 2, 3}, []bool{false, true, true}),
			withoutNA: []int{1},
		},
		{
			name:      "Empty",
			vec:       Empty(),
			withoutNA: []int{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if !reflect.DeepEqual(data.vec.WithoutNA(), data.withoutNA) {
				t.Error(fmt.Sprintf("data.vec.WithoutNA() (%v) is not equal to data.WithoutNA (%v)",
					data.vec.WithoutNA(), data.withoutNA))
			}
		})
	}
}

func TestVector_IsEmpty(t *testing.T) {
	testData := []struct {
		name    string
		vec     Vector
		isEmpty bool
	}{
		{
			name:    "zero integer vector",
			vec:     Integer([]int{}, nil),
			isEmpty: true,
		},
		{
			name:    "non-zero integer vector",
			vec:     Integer([]int{1, 2, 3}, nil),
			isEmpty: false,
		},
		{
			name:    "empty vector",
			vec:     Integer([]int{}, nil),
			isEmpty: true,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if data.vec.IsEmpty() != data.isEmpty {
				t.Error(fmt.Sprintf("data.vec.IsEmpty() (%t) is not equal to data.isEmpty (%t)",
					data.vec.IsEmpty(), data.isEmpty))
			}
		})
	}
}

func TestVector_SupportsSelector(t *testing.T) {
	testData := []struct {
		name             string
		vec              Vector
		selector         interface{}
		supportsSelector bool
	}{
		{
			name:             "integer vector + valid selector",
			vec:              Integer([]int{1, 2, 3}, nil),
			selector:         func(_ int, val int, _ bool) bool { return val == 1 || val == 3 },
			supportsSelector: true,
		},
		{
			name:             "integer vector + invalid selector",
			vec:              Integer([]int{1, 2, 3}, nil),
			selector:         true,
			supportsSelector: false,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			supportsSelector := data.vec.SupportsSelector(data.selector)
			if supportsSelector != data.supportsSelector {
				t.Error(fmt.Sprintf("data.vec.Select() (%v) is not equal to data.selected (%v)",
					supportsSelector, data.supportsSelector))
			}
		})
	}
}

func TestVector_Select(t *testing.T) {
	testData := []struct {
		name     string
		vec      Vector
		selector interface{}
		selected []bool
	}{
		{
			name:     "integer vector + valid selector",
			vec:      Integer([]int{1, 2, 3}, nil),
			selector: func(_ int, val int, _ bool) bool { return val == 1 || val == 3 },
			selected: []bool{true, false, true},
		},
		{
			name:     "integer vector + invalid selector",
			vec:      Integer([]int{1, 2, 3}, nil),
			selector: true,
			selected: []bool{false, false, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			selected := data.vec.Select(data.selector)
			if !reflect.DeepEqual(selected, data.selected) {
				t.Error(fmt.Sprintf("data.vec.Select() (%v) is not equal to data.selected (%v)",
					selected, data.selected))
			}
		})
	}
}
