package vector

import (
	"fmt"
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
			vec:       NA(0),
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
			vec:       NA(0),
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
			vec:       NA(0),
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
			vec:       NA(0),
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
			vec:       NA(0),
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
			vec:       NA(0),
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

func TestVector_FromTo(t *testing.T) {
	vec := Integer(
		[]int{1, 2, 3, 4, 5},
		[]bool{false, false, true, false, false},
	)

	testData := []struct {
		name   string
		from   int
		to     int
		out    []int
		outNA  []bool
		length int
	}{
		{
			name:   "FromTo straight",
			from:   2,
			to:     4,
			out:    []int{2, 0, 4},
			outNA:  []bool{false, true, false},
			length: 3,
		},
		{
			name:   "FromTo reverse",
			from:   4,
			to:     2,
			out:    []int{4, 0, 2},
			outNA:  []bool{false, true, false},
			length: 3,
		},
		{
			name:   "FromTo with remove straight",
			from:   -2,
			to:     -4,
			out:    []int{1, 5},
			outNA:  []bool{false, false},
			length: 2,
		},
		{
			name:   "FromTo with remove reverse",
			from:   -4,
			to:     -2,
			out:    []int{1, 5},
			outNA:  []bool{false, false},
			length: 2,
		},
		{
			name:   "FromTo straight with incorrect boundaries",
			from:   0,
			to:     7,
			out:    []int{1, 2, 0, 4, 5},
			outNA:  []bool{false, false, true, false, false},
			length: 5,
		},
		{
			name:   "FromTo full remove with incorrect boundaries",
			from:   0,
			to:     -7,
			out:    []int{},
			outNA:  []bool{},
			length: 0,
		},
		{
			name:   "FromTo different signs",
			from:   -2,
			to:     4,
			out:    []int{},
			outNA:  []bool{},
			length: 0,
		},
		{
			name:   "FromTo different signs reverse",
			from:   2,
			to:     -4,
			out:    []int{},
			outNA:  []bool{},
			length: 0,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			newVec := vec.FromTo(data.from, data.to)
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
	)
	testData := []struct {
		name    string
		whicher interface{}
		out     []int
		outNA   []bool
		length  int
	}{
		{
			name:    "indices",
			whicher: []int{5, 3, 1},
			out:     []int{5, 0, 1},
			outNA:   []bool{false, true, false},
			length:  3,
		},
		{
			name:    "idx",
			whicher: []int{1},
			out:     []int{1},
			outNA:   []bool{false},
			length:  1,
		},
		{
			name:    "booleanPayload",
			whicher: []bool{true, false, true, false, true},
			out:     []int{1, 0, 5},
			outNA:   []bool{false, true, false},
			length:  3,
		},
		{
			name:    "byIntFunc",
			whicher: func(_ int, val int, na bool) bool { return !na && val >= 3 },
			out:     []int{4, 5},
			outNA:   []bool{false, false},
			length:  2,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			newVec := vec.Filter(data.whicher)
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
			vec:   NA(0),
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
			vec:   NA(0),
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
			vec:   NA(0),
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
			vec:    NA(0),
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
			vec:       NA(0),
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
			name:    "zero integerPayload vector",
			vec:     Integer([]int{}, nil),
			isEmpty: true,
		},
		{
			name:    "non-zero integerPayload vector",
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

func TestVector_Clone(t *testing.T) {
	vec := Integer([]int{1, 2, 3, 4, 5}, []bool{false, true, false, true, false}).(*vector)
	newVec := vec.Clone().(*vector)

	if vec.length != newVec.length {
		t.Error(fmt.Sprintf("vec.length (%d) is not equal to newVec.length (%d)", vec.length, newVec.length))
	}

	srcAddr := &(vec.payload.(*integerPayload).data[0])
	newAddr := &(newVec.payload.(*integerPayload).data[0])

	if srcAddr != newAddr {
		t.Error("Payload data was not cloned")
	}

	srcAddrNA := &(vec.payload.(*integerPayload).na[0])
	newAddrNA := &(newVec.payload.(*integerPayload).na[0])

	if srcAddrNA != newAddrNA {
		t.Error("Payload NA data was not cloned")
	}
}

func TestVector_SupportsWhicher(t *testing.T) {
	testData := []struct {
		name            string
		vec             Vector
		whicher         interface{}
		supportsWhicher bool
	}{
		{
			name:            "integerPayload vector + valid whicher",
			vec:             Integer([]int{1, 2, 3}, nil),
			whicher:         func(_ int, val int, _ bool) bool { return val == 1 || val == 3 },
			supportsWhicher: true,
		},
		{
			name:            "integerPayload vector + invalid whicher",
			vec:             Integer([]int{1, 2, 3}, nil),
			whicher:         true,
			supportsWhicher: false,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			supportsWhicher := data.vec.SupportsWhicher(data.whicher)
			if supportsWhicher != data.supportsWhicher {
				t.Error(fmt.Sprintf("data.vec.Which() (%v) is not equal to data.selected (%v)",
					supportsWhicher, data.supportsWhicher))
			}
		})
	}
}

func TestVector_Select(t *testing.T) {
	testData := []struct {
		name     string
		vec      Vector
		whicher  interface{}
		selected []bool
	}{
		{
			name:     "integerPayload vector + valid whicher",
			vec:      Integer([]int{1, 2, 3}, nil),
			whicher:  func(_ int, val int, _ bool) bool { return val == 1 || val == 3 },
			selected: []bool{true, false, true},
		},
		{
			name:     "integerPayload vector + invalid whicher",
			vec:      Integer([]int{1, 2, 3}, nil),
			whicher:  true,
			selected: []bool{false, false, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			selected := data.vec.Which(data.whicher)
			if !reflect.DeepEqual(selected, data.selected) {
				t.Error(fmt.Sprintf("data.vec.Which() (%v) is not equal to data.selected (%v)",
					selected, data.selected))
			}
		})
	}
}

func TestVector_SupportsApplier(t *testing.T) {
	testData := []struct {
		name            string
		vec             Vector
		whicher         interface{}
		supportsApplier bool
	}{
		{
			name:            "integerPayload vector + valid applier",
			vec:             Integer([]int{1, 2, 3}, nil),
			whicher:         func(_ int, val int, na bool) (int, bool) { return 10 * val, na },
			supportsApplier: true,
		},
		{
			name:            "integerPayload vector + invalid applier",
			vec:             Integer([]int{1, 2, 3}, nil),
			whicher:         true,
			supportsApplier: false,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			supportsApplier := data.vec.SupportsApplier(data.whicher)
			if supportsApplier != data.supportsApplier {
				t.Error(fmt.Sprintf("Applier's support (%v) is not equal to expected (%v)",
					supportsApplier, data.supportsApplier))
			}
		})
	}
}

func TestVector_Apply(t *testing.T) {
	testData := []struct {
		name    string
		vec     Vector
		applier interface{}
		dataOut []int
		NAOut   []bool
	}{
		{
			name: "integerPayload vector + valid applier",
			vec:  Integer([]int{1, 2, 3, 4, 5}, nil),
			applier: func(idx int, val int, na bool) (int, bool) {
				if idx == 5 {
					return val, true
				}
				return 10 * val, na
			},
			dataOut: []int{10, 20, 30, 40, 0},
			NAOut:   []bool{false, false, false, false, true},
		},
		{
			name:    "integerPayload vector + invalid applier",
			vec:     Integer([]int{1, 2, 3, 4, 5}, nil),
			applier: true,
			dataOut: []int{0, 0, 0, 0, 0},
			NAOut:   []bool{true, true, true, true, true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			newVec := data.vec.Apply(data.applier)
			integers, na := newVec.Integers()
			if !reflect.DeepEqual(integers, data.dataOut) {
				t.Error(fmt.Sprintf("Integers (%v) is not equal to expected (%v)",
					integers, data.dataOut))
			}
			if !reflect.DeepEqual(na, data.NAOut) {
				t.Error(fmt.Sprintf("NA (%v) is not equal to expected (%v)",
					na, data.NAOut))
			}
		})
	}
}

func TestVector_Append(t *testing.T) {
	vec := Integer([]int{1, 2, 3}, nil)

	testData := []struct {
		name    string
		vec     Vector
		outData []int
		outNA   []bool
	}{
		{
			name:    "boolean",
			vec:     Boolean([]bool{true, true}, []bool{true, false}),
			outData: []int{1, 2, 3, 0, 1},
			outNA:   []bool{false, false, false, true, false},
		},
		{
			name:    "integer",
			vec:     Integer([]int{4, 5}, []bool{true, false}),
			outData: []int{1, 2, 3, 0, 5},
			outNA:   []bool{false, false, false, true, false},
		},
		{
			name:    "empty",
			vec:     Integer([]int{}, []bool{}),
			outData: []int{1, 2, 3},
			outNA:   []bool{false, false, false},
		},
		{
			name:    "na",
			vec:     NA(2),
			outData: []int{1, 2, 3, 0, 0},
			outNA:   []bool{false, false, false, true, true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			outVector := vec.Append(data.vec).(*vector)
			outPayload := outVector.payload.(*integerPayload)

			if outVector.Len() != vec.Len()+data.vec.Len() {
				t.Error(fmt.Sprintf("Output length (%d) does not match expected (%d)",
					outVector.Len(), vec.Len()+data.vec.Len()))
			}
			if !reflect.DeepEqual(data.outData, outPayload.data) {
				t.Error(fmt.Sprintf("Output data (%v) does not match expected (%v)",
					outPayload.data, data.outData))
			}
			if !reflect.DeepEqual(data.outNA, outPayload.na) {
				t.Error(fmt.Sprintf("Output NA (%v) does not match expected (%v)",
					outPayload.na, data.outNA))
			}
		})
	}
}

func TestVector_Find(t *testing.T) {
	vec := Integer([]int{1, 2, 1, 4, 0}, nil)

	testData := []struct {
		name   string
		needle interface{}
		pos    int
	}{
		{"existent", 4, 4},
		{"float64", 2.1, 2},
		{"non-existent", -10, 0},
		{"incorrect type", "true", 0},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			pos := vec.Find(data.needle)

			if pos != data.pos {
				t.Error(fmt.Sprintf("Position (%v) does not match expected (%v)",
					pos, data.pos))
			}
		})
	}
}

func TestVector_FindAll(t *testing.T) {
	vec := Integer([]int{1, 2, 1, 4, 0}, nil)

	testData := []struct {
		name   string
		needle interface{}
		pos    []int
	}{
		{"existent", 1, []int{1, 3}},
		{"float", 1.2, []int{1, 3}},
		{"non-existent", -10, []int{}},
		{"incorrect type", false, []int{}},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			pos := vec.FindAll(data.needle)

			if !reflect.DeepEqual(pos, data.pos) {
				t.Error(fmt.Sprintf("Positions (%v) does not match expected (%v)",
					pos, data.pos))
			}
		})
	}
}
