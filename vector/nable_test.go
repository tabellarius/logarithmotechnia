package vector

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestDefNAble_NA(t *testing.T) {
	testData := []struct {
		name string
		in   []bool
		out  []bool
	}{
		{
			in:  []bool{true, true, false, false, false, true},
			out: []bool{true, true, false, false, false, true},
		},
		{
			in:  []bool{true, true, false},
			out: []bool{true, true, false},
		},
		{
			in:  []bool{},
			out: []bool{},
		},
		{
			in:  nil,
			out: []bool{},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			nable := DefNAble{
				na: data.in,
			}
			na := nable.IsNA()
			if !reflect.DeepEqual(na, data.out) {
				t.Error(fmt.Sprintf("Value IsNA(%v) is not equal to out(%v)", na, data.out))
			}
		})
	}
}

func TestDefNAble_NotNA(t *testing.T) {
	testData := []struct {
		name string
		in   []bool
		out  []bool
	}{
		{
			in:  []bool{true, true, false, false, false, true},
			out: []bool{false, false, true, true, true, false},
		},
		{
			in:  []bool{true, true, false},
			out: []bool{false, false, true},
		},
		{
			in:  []bool{},
			out: []bool{},
		},
		{
			in:  nil,
			out: []bool{},
		},
	}

	for i, data := range testData {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			nable := DefNAble{
				na: data.in,
			}
			na := nable.NotNA()
			if !reflect.DeepEqual(na, data.out) {
				t.Error(fmt.Sprintf("Value IsNA(%v) is not equal to out(%v)", na, data.out))
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
		{
			name:     "empty",
			na:       []bool{},
			expected: false,
		},
	}

	for index, data := range testData {
		t.Run(fmt.Sprintf("Data %d", index), func(t *testing.T) {
			nable := DefNAble{
				na: data.na,
			}
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
			nable := DefNAble{
				na: data.na,
			}
			withNA := nable.WithNA()
			if !reflect.DeepEqual(withNA, data.expected) {
				t.Error(fmt.Sprintf("nable.OnlyNa() %v is not equal to data.expected %v", withNA,
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
			nable := DefNAble{
				na: data.na,
			}
			withoutNA := nable.WithoutNA()
			if !reflect.DeepEqual(withoutNA, data.expected) {
				t.Error(fmt.Sprintf("nable.OnlyNa() %v is not equal to data.expected %v", withoutNA,
					data.expected))
			}
		})
	}
}
