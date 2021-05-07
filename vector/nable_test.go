package vector

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDefNAble_Refresh(t *testing.T) {
	vec1 := New(5)
	vec2 := New(5)
	nable1 := DefNAble{
		vec: &vec1,
		na:  []bool{false, true, false, false, false, true},
	}
	nable2 := DefNAble{
		vec: &vec2,
		na:  nable1.na,
	}
	nable2.Refresh()
	if &nable1.na[0] == &nable2.na[0] {
		t.Error("nable2 was not refreshed")
	}
	if !reflect.DeepEqual(nable1.na, nable2.na) {
		t.Error("booleans are not equal for nable1 and nable2")
	}
}

func TestDefNAble_NA(t *testing.T) {
	vec := New(5)
	nable := DefNAble{
		vec: &vec,
		na:  []bool{false, true, false, false, false, true},
	}
	na := nable.NA()
	if len(na) != vec.length {
		t.Error("Output vector's length is wrong")
	}
	if &na[0] == &nable.na[1] {
		t.Error("Output vector is not a copy")
	}
	if !reflect.DeepEqual(na, nable.na[1:]) {
		t.Error("Output vector is not equal")
	}
}

func TestDefNAble_IsNA(t *testing.T) {
	vec := New(5)
	nable := DefNAble{
		vec: &vec,
		na:  []bool{true, true, false, false, false, true},
	}
	testData := []struct {
		idx      int
		expected bool
	}{
		{
			idx:      -1,
			expected: false,
		},
		{
			idx:      0,
			expected: false,
		},
		{
			idx:      1,
			expected: true,
		},
		{
			idx:      3,
			expected: false,
		},
		{
			idx:      5,
			expected: true,
		},
		{
			idx:      6,
			expected: false,
		},
	}

	for _, data := range testData {
		t.Run(fmt.Sprintf("Index_%d", data.idx), func(t *testing.T) {
			isNA := nable.IsNA(data.idx)
			if isNA != data.expected {
				t.Error(fmt.Sprintf("Value IsNA(%t) is not equal to expected(%t)", isNA, data.expected))
			}
		})
	}
}

func TestDefNAble_SetNA(t *testing.T) {
	testData := []struct {
		name     string
		na       []bool
		expected []bool
	}{
		{
			name:     "without na",
			na:       []bool{false, false, false, false, false},
			expected: []bool{false, false, false, false, false, false},
		},
		{
			name:     "with na",
			na:       []bool{true, false, true, false, true},
			expected: []bool{false, true, false, true, false, true},
		},
		{
			name:     "incorrect size",
			na:       []bool{true, false, true, false, true, false},
			expected: []bool{false, false, false, false, false, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			vec := New(5)
			nable := newDefaultNAble(&vec)
			nable.SetNA(data.na)
			if !reflect.DeepEqual(nable.na, data.expected) {
				t.Error("nable.na is not equal to data.expected")
			}
		})
	}
}

func TestDefNAble_HasNA(t *testing.T) {
	testData := []struct {
		name     string
		na       []bool
		expected bool
	}{
		{
			name:     "no NA (1)",
			na:       []bool{false, false, false, false, false},
			expected: false,
		},
		{
			name:     "no NA (2)",
			na:       []bool{false, false, true, false, false},
			expected: true,
		},
		{
			name:     "with NA",
			na:       []bool{true, false, true, false, true},
			expected: true,
		},
	}

	vec := New(5)
	for index, data := range testData {
		t.Run(fmt.Sprintf("Data %d", index), func(t *testing.T) {
			nable := newDefaultNAble(&vec)
			nable.SetNA(data.na)
			if nable.HasNA() != data.expected {
				t.Error(fmt.Sprintf("nable.HasNA() (%t) is not equal to data.expected (%t)", nable.HasNA(),
					data.expected))
			}
		})
	}
}

func TestDefNAble_OnlyNA(t *testing.T) {
	testData := []struct {
		name     string
		na       []bool
		expected []int
	}{
		{
			name:     "without na",
			na:       []bool{false, false, false, false, false},
			expected: []int{},
		},
		{
			name:     "with na#1",
			na:       []bool{true, false, true, false, true},
			expected: []int{1, 3, 5},
		},
		{
			name:     "with na#2",
			na:       []bool{false, false, false, true, true},
			expected: []int{4, 5},
		},
		{
			name:     "with na#3",
			na:       []bool{false, false, false, false, true},
			expected: []int{5},
		},
		{
			name:     "all na",
			na:       []bool{true, true, true, true, true},
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			vec := New(5)
			nable := newDefaultNAble(&vec)
			nable.SetNA(data.na)
			if !reflect.DeepEqual(nable.OnlyNA(), data.expected) {
				t.Error(fmt.Sprintf("nable.OnlyNa() %v is not equal to data.expected %v", nable.OnlyNA(),
					data.expected))
			}
		})
	}
}

func TestDefNAble_WithoutNA(t *testing.T) {
	testData := []struct {
		name     string
		na       []bool
		expected []int
	}{
		{
			name:     "without na",
			na:       []bool{false, false, false, false, false},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "with na#1",
			na:       []bool{true, false, true, false, true},
			expected: []int{2, 4},
		},
		{
			name:     "with na#2",
			na:       []bool{false, false, false, true, true},
			expected: []int{1, 2, 3},
		},
		{
			name:     "with na#3",
			na:       []bool{false, false, false, false, true},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "all na",
			na:       []bool{true, true, true, true, true},
			expected: []int{},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			vec := New(5)
			nable := newDefaultNAble(&vec)
			nable.SetNA(data.na)
			if !reflect.DeepEqual(nable.WithoutNA(), data.expected) {
				t.Error(fmt.Sprintf("nable.OnlyNa() %v is not equal to data.expected %v", nable.OnlyNA(),
					data.expected))
			}
		})
	}
}
