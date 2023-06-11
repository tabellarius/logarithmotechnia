package which

import (
	"logarithmotechnia/vector"
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  []bool
	}{
		{
			name: "string",
			in:   vector.String([]string{"foo", "bar", "baz", "qux", "quux", "corge", "grault"}),
			out:  []bool{false, false, false, true, true, false, false},
		},
		{
			name: "int",
			in:   vector.Integer([]int{1, 2, 3, 4, 5, 6}),
			out:  []bool{false, false, false, false, false, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := Contains(data.in, "qu")
			if !reflect.DeepEqual(out, data.out) {
				t.Errorf("Expected %v, got %v", data.out, out)
			}
		})
	}
}

func TestContainsAny(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  []bool
	}{
		{
			name: "string",
			in:   vector.String([]string{"foo", "qar", "buz", "qux", "quux", "corge", "grault"}),
			out:  []bool{false, true, true, true, true, false, true},
		},
		{
			name: "int",
			in:   vector.Integer([]int{1, 2, 3, 4, 5, 6}),
			out:  []bool{false, false, false, false, false, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := ContainsAny(data.in, "qu")
			if !reflect.DeepEqual(out, data.out) {
				t.Errorf("Expected %v, got %v", data.out, out)
			}
		})
	}
}

func TestContainsRune(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  []bool
	}{
		{
			name: "string",
			in:   vector.String([]string{"foo", "bar", "baz", "qux", "quux", "corge", "grault"}),
			out:  []bool{false, false, false, true, true, false, false},
		},
		{
			name: "int",
			in:   vector.Integer([]int{1, 2, 3, 4, 5, 6}),
			out:  []bool{false, false, false, false, false, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := ContainsRune(data.in, 'q')
			if !reflect.DeepEqual(out, data.out) {
				t.Errorf("Expected %v, got %v", data.out, out)
			}
		})
	}
}

func TestHasPrefix(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  []bool
	}{
		{
			name: "string",
			in:   vector.String([]string{"foo", "bar", "baz", "qux", "quux", "corge", "grault"}),
			out:  []bool{false, false, false, true, true, false, false},
		},
		{
			name: "int",
			in:   vector.Integer([]int{1, 2, 3, 4, 5, 6}),
			out:  []bool{false, false, false, false, false, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := HasPrefix(data.in, "qu")
			if !reflect.DeepEqual(out, data.out) {
				t.Errorf("Expected %v, got %v", data.out, out)
			}
		})
	}
}

func TestHasSuffix(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  []bool
	}{
		{
			name: "string",
			in:   vector.String([]string{"foo", "bar", "baz", "qux", "quux", "corge", "grault"}),
			out:  []bool{false, false, false, false, false, false, true},
		},
		{
			name: "int",
			in:   vector.Integer([]int{1, 2, 3, 4, 5, 6}),
			out:  []bool{false, false, false, false, false, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := HasSuffix(data.in, "lt")
			if !reflect.DeepEqual(out, data.out) {
				t.Errorf("Expected %v, got %v", data.out, out)
			}
		})
	}
}

func TestEqualFold(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  []bool
	}{
		{
			name: "string",
			in:   vector.String([]string{"foo", "bar", "baz", "qu", "qU", "corge", "grault"}),
			out:  []bool{false, false, false, true, true, false, false},
		},
		{
			name: "int",
			in:   vector.Integer([]int{1, 2, 3, 4, 5, 6}),
			out:  []bool{false, false, false, false, false, false},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := EqualFold(data.in, "Qu")
			if !reflect.DeepEqual(out, data.out) {
				t.Errorf("Expected %v, got %v", data.out, out)
			}
		})
	}
}
