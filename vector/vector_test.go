package vector

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCommon_newCommon(t *testing.T) {
	vec := newCommon(5)
	if vec.length != 5 {
		t.Error(`newCommon(5): length is not 5`)
	}

	if vec.vec != nil {
		t.Error(`newCommon(5): vec is not nil`)
	}

	if vec.marked != false {
		t.Error(`newCommon(5): marked is not false`)
	}

	if len(vec.report.Errors) != 0 {
		t.Error(`newCommon(5): there are errors in the report`)
	}

	if len(vec.report.Warnings) != 0 {
		t.Error(`newCommon(5): there are warnings in the report`)
	}

	vec = newCommon(0)
	if vec.length != 0 {
		t.Error("newCommon(0): length is not zero")
	}

	vec = newCommon(0)
	if vec.length != 0 {
		t.Error("newCommon(0): length is not zero")
	}

	vec = newCommon(-1)
	if vec.length != 0 {
		t.Error("newCommon(0): length is not zero")
	}
}

func TestCommon_Length(t *testing.T) {
	lengthIn := []int{-10, -1, 0, 1, 5, 12, 25, 100}
	lengthOut := []int{0, 0, 0, 1, 5, 12, 25, 100}

	for i := 0; i < len(lengthIn); i++ {
		vec := newCommon(lengthIn[i])
		if vec.Length() != lengthOut[i] {
			t.Error(fmt.Sprintf("In-length (%d) does not match out-length (%d)", lengthIn[i], lengthOut[i]))
		}
	}
}

func TestCommon_IsEmpty(t *testing.T) {
	lengthIn := []int{-10, -1, 0, 1, 10}
	emptyness := []bool{true, true, true, false, false}

	for i := 0; i < len(lengthIn); i++ {
		vec := newCommon(lengthIn[i])
		if vec.IsEmpty() != emptyness[i] {
			t.Error(fmt.Sprintf("Emptyness of vec(%d) is wrong (%t instead of %t)", lengthIn[i],
				vec.IsEmpty(), emptyness[i]))
		}
	}
}

func TestCommon_ByIndices(t *testing.T) {
	testData := []struct {
		name       string
		indicesIn  []int
		indicesOut []int
		lengthSrc  int
		lengthDst  int
	}{
		{
			name:       "regular",
			indicesIn:  []int{1, 5, 10},
			indicesOut: []int{1, 5, 10},
			lengthSrc:  10,
			lengthDst:  3,
		},
		{
			name:       "regular reversed",
			indicesIn:  []int{10, 5, 1},
			indicesOut: []int{10, 5, 1},
			lengthSrc:  10,
			lengthDst:  3,
		},
		{
			name:       "with negative",
			indicesIn:  []int{-1, 5, 1},
			indicesOut: []int{5, 1},
			lengthSrc:  10,
			lengthDst:  2,
		},
		{
			name:       "with zero",
			indicesIn:  []int{0, 2, 5, 1},
			indicesOut: []int{2, 5, 1},
			lengthSrc:  10,
			lengthDst:  3,
		},
		{
			name:       "out of bounds",
			indicesIn:  []int{11, 2, 5, 1},
			indicesOut: []int{2, 5, 1},
			lengthSrc:  10,
			lengthDst:  3,
		},
	}

	for _, indicesData := range testData {
		t.Run(indicesData.name, func(t *testing.T) {
			vec := newCommon(indicesData.lengthSrc)
			newVec := vec.ByIndices(indicesData.indicesIn).(*common)

			if newVec.length != indicesData.lengthDst {
				t.Error(fmt.Sprintf("newVec's length (%d) is not {%d}", newVec.length, indicesData.lengthDst))
			}

			if len(newVec.selected) != indicesData.lengthDst {
				t.Error(fmt.Sprintf("newVec.selected's length (%d) is not %d", len(newVec.selected),
					indicesData.lengthDst))
			}

			if !reflect.DeepEqual(newVec.selected, indicesData.indicesOut) {
				t.Error(fmt.Sprintf("newVec.selected (%v) is not equal to %v", newVec.selected,
					indicesData.indicesOut))
			}
		})
	}
}

func TestCommon_ByFromTo(t *testing.T) {
	testData := []struct {
		name      string
		from      int
		to        int
		lengthSrc int
		lengthDst int
		selected  []int
	}{
		{
			name:      "full",
			from:      1,
			to:        10,
			lengthSrc: 10,
			lengthDst: 10,
			selected:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:      "full reverse",
			from:      10,
			to:        1,
			lengthSrc: 10,
			lengthDst: 10,
			selected:  []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			name:      "full + out of bounds",
			from:      1,
			to:        12,
			lengthSrc: 10,
			lengthDst: 10,
			selected:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:      "full reverse + out of bounds",
			from:      12,
			to:        1,
			lengthSrc: 10,
			lengthDst: 10,
			selected:  []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			name:      "partial",
			from:      3,
			to:        5,
			lengthSrc: 10,
			lengthDst: 3,
			selected:  []int{3, 4, 5},
		},
		{
			name:      "partial reverse",
			from:      5,
			to:        3,
			lengthSrc: 10,
			lengthDst: 3,
			selected:  []int{5, 4, 3},
		},
		{
			name:      "partial + removal",
			from:      -4,
			to:        -7,
			lengthSrc: 10,
			lengthDst: 6,
			selected:  []int{1, 2, 3, 8, 9, 10},
		},
		{
			name:      "partial + removal reverse",
			from:      -7,
			to:        -4,
			lengthSrc: 10,
			lengthDst: 6,
			selected:  []int{1, 2, 3, 8, 9, 10},
		},
		{
			name:      "partial + zero",
			from:      0,
			to:        5,
			lengthSrc: 10,
			lengthDst: 5,
			selected:  []int{1, 2, 3, 4, 5},
		},
		{
			name:      "partial + zero + removal reverse",
			from:      -5,
			to:        0,
			lengthSrc: 10,
			lengthDst: 5,
			selected:  []int{6, 7, 8, 9, 10},
		},
		{
			name:      "zero",
			from:      0,
			to:        0,
			lengthSrc: 10,
			lengthDst: 0,
			selected:  []int{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			vec := newCommon(data.lengthSrc)
			newVec := vec.ByFromTo(data.from, data.to).(*common)

			if newVec.length != data.lengthDst {
				t.Error(fmt.Sprintf("newVec's length (%d) is not {%d}", newVec.length, data.lengthDst))
			}

			if len(newVec.selected) != data.lengthDst {
				t.Error(fmt.Sprintf("newVec.selected's length (%d) is not %d", len(newVec.selected),
					data.lengthDst))
			}

			if !reflect.DeepEqual(newVec.selected, data.selected) {
				t.Error(fmt.Sprintf("newVec.selected (%v) is not equal to vec.selected %v", newVec.selected,
					data.selected))
			}
		})
	}
}

func TestCommon_ByBool(t *testing.T) {
	testData := []struct {
		name     string
		booleans []bool
		length   int
		selected []int
		emptyVec bool
	}{
		{
			name:     "full",
			booleans: []bool{true, true, true, true, true},
			length:   5,
			selected: []int{1, 2, 3, 4, 5},
			emptyVec: false,
		},
		{
			name:     "partial",
			booleans: []bool{false, true, false, true, false},
			length:   2,
			selected: []int{2, 4},
			emptyVec: false,
		},
		{
			name:     "none",
			booleans: []bool{false, false, false, false, false},
			length:   0,
			selected: []int{},
			emptyVec: false,
		},
		{
			name:     "incorrect length",
			booleans: []bool{true, true, true, true},
			length:   0,
			selected: []int{},
			emptyVec: true,
		},
	}

	vec := newCommon(5)
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			newVec := vec.ByBool(data.booleans)
			if data.emptyVec {
				if _, ok := newVec.(*empty); !ok {
					t.Error("Returned vector is not empty type")
				}
			} else {
				newCom := newVec.(*common)
				if newCom.length != data.length {
					t.Error(fmt.Sprintf("Output length (%d) is not %d", newCom.length, data.length))
				}
				if !reflect.DeepEqual(newCom.selected, data.selected) {
					t.Error(fmt.Sprintf("newVec.selected (%v) is not equal to %v", newCom.selected,
						data.selected))
				}
			}
		})
	}
}

func TestCommon_Clone(t *testing.T) {
	vec := newCommon(10)
	newVec := vec.Clone().(*common)

	if newVec.vec != nil {
		t.Error("newVec.vec is not nil")
	}
	if newVec.marked == false {
		t.Error("newVec.marked is false")
	}
	if newVec.length != vec.length {
		t.Error("newVec.length is not equal to vec.length")
	}
	if !reflect.DeepEqual(newVec.selected, vec.selected) {
		t.Error(fmt.Sprintf("newVec.selected (%v) is not equal to vec.selected %v", newVec.selected,
			vec.selected))
	}
}

func TestCommon_Mark(t *testing.T) {
	vec := newCommon(10)
	vec.Mark()
	if !vec.marked {
		t.Error("vec.marked is not true")
	}
}

func TestCommon_Marked(t *testing.T) {
	vec := newCommon(10)
	if vec.Marked() {
		t.Error("vec.marked is true for new vector")
	}

	vec.marked = true
	if !vec.Marked() {
		t.Error("vec was not marked")
	}
}

func TestCommon_Refresh(t *testing.T) {
	vec := newCommon(10)
	vec.marked = true
	vec.Refresh()
	if vec.marked {
		t.Error("vec.marked is not false")
	}
}

func TestCommon_Report(t *testing.T) {
	vec := newCommon(10)
	vec.report = Report{
		Errors:   []string{"Error"},
		Warnings: []string{"Warning"},
	}
	if !reflect.DeepEqual(vec.report, vec.Report()) {
		t.Error("Strange report")
	}
}
