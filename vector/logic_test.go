package vector

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAnd(t *testing.T) {
	testData := []struct {
		name     string
		booleans [][]bool
		result   []bool
	}{
		{
			name:     "empty",
			booleans: [][]bool{},
			result:   []bool{},
		},
		{
			name: "one argument",
			booleans: [][]bool{
				{true, false, true, false, true},
			},
			result: []bool{true, false, true, false, true},
		},
		{
			name: "shorter",
			booleans: [][]bool{
				{true, false, true, false, true},
				{true, true, false},
			},
			result: []bool{true, false, false, false, false},
		},
		{
			name: "longer",
			booleans: [][]bool{
				{true, false, true, false, true},
				{true, true, false, false, true, true, false},
			},
			result: []bool{true, false, false, false, true},
		},
		{
			name: "equal",
			booleans: [][]bool{
				{true, false, true, false, true},
				{true, true, false, false, true},
			},
			result: []bool{true, false, false, false, true},
		},
		{
			name: "multiple",
			booleans: [][]bool{
				{true, false, true, false, true},
				{true, true, false, false, true},
				{true, false, true},
				{true, true, true, true, false},
			},
			result: []bool{true, false, false, false, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			result := And(data.booleans...)

			if !reflect.DeepEqual(result, data.result) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)",
					result, data.result))
			}
		})
	}
}

func TestOr(t *testing.T) {
	testData := []struct {
		name     string
		booleans [][]bool
		result   []bool
	}{
		{
			name:     "empty",
			booleans: [][]bool{},
			result:   []bool{},
		},
		{
			name: "one argument",
			booleans: [][]bool{
				{true, false, true, false, true},
			},
			result: []bool{true, false, true, false, true},
		},
		{
			name: "shorter",
			booleans: [][]bool{
				{true, false, true, false, true},
				{true, true, false},
			},
			result: []bool{true, true, true, false, true},
		},
		{
			name: "longer",
			booleans: [][]bool{
				{true, false, true, false, true},
				{true, true, false, false, true, true, false},
			},
			result: []bool{true, true, true, false, true},
		},
		{
			name: "equal",
			booleans: [][]bool{
				{true, false, true, false, true},
				{true, true, false, false, true},
			},
			result: []bool{true, true, true, false, true},
		},
		{
			name: "multiple",
			booleans: [][]bool{
				{true, false, true, false, true},
				{true, true, false, false, true},
				{true, false, true},
				{true, true, true, true, false},
			},
			result: []bool{true, true, true, true, true},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			result := Or(data.booleans...)

			if !reflect.DeepEqual(result, data.result) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)",
					result, data.result))
			}
		})
	}
}

func TestXor(t *testing.T) {
	testData := []struct {
		name     string
		booleans [][]bool
		result   []bool
	}{
		{
			name:     "empty",
			booleans: [][]bool{},
			result:   []bool{},
		},
		{
			name: "one argument",
			booleans: [][]bool{
				{true, false, true, false, true},
			},
			result: []bool{true, false, true, false, true},
		},
		{
			name: "shorter",
			booleans: [][]bool{
				{true, false, true, false, true},
				{true, true, false},
			},
			result: []bool{false, true, true, false, true},
		},
		{
			name: "longer",
			booleans: [][]bool{
				{true, false, true, false, true},
				{true, true, false, false, true, true, false},
			},
			result: []bool{false, true, true, false, false},
		},
		{
			name: "equal",
			booleans: [][]bool{
				{true, false, true, false, true},
				{true, true, false, false, true},
			},
			result: []bool{false, true, true, false, false},
		},
		{
			name: "multiple",
			booleans: [][]bool{
				{true, false, true, false, true},
				{true, true, false, false, true},
				{true, false, true},
				{true, true, true, true, false},
			},
			result: []bool{false, false, true, true, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			result := Xor(data.booleans...)

			if !reflect.DeepEqual(result, data.result) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)",
					result, data.result))
			}
		})
	}
}

func TestNot(t *testing.T) {
	testData := []struct {
		name     string
		booleans []bool
		result   []bool
	}{
		{
			name:     "empty",
			booleans: []bool{},
			result:   []bool{},
		},
		{
			name:     "mixed",
			booleans: []bool{true, false, true, false, true},
			result:   []bool{false, true, false, true, false},
		},
		{
			name:     "all true",
			booleans: []bool{true, true, true, true, true},
			result:   []bool{false, false, false, false, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			result := Not(data.booleans)

			if !reflect.DeepEqual(result, data.result) {
				t.Error(fmt.Sprintf("Result (%v) is not equal to expected (%v)",
					result, data.result))
			}
		})
	}
}
