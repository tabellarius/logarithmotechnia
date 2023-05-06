package apply

import (
	"logarithmotechnia/vector"
	"testing"
	"unicode"
)

func TestCompare(t *testing.T) {
	tests := []struct {
		name  string
		in    vector.Vector
		param string
		out   vector.Vector
	}{
		{
			name:  "string",
			in:    vector.StringWithNA([]string{"ab", "abc", "", "abc", "abcd"}, []bool{false, false, true, false, false}),
			param: "abc",
			out:   vector.IntegerWithNA([]int{-1, 0, 0, 0, 1}, []bool{false, false, true, false, false}),
		},
		{
			name:  "int",
			in:    vector.IntegerWithNA([]int{1, 11, 0, 11, 2}, []bool{false, false, true, false, false}),
			param: "11",
			out:   vector.IntegerWithNA([]int{-1, 0, 0, 0, 1}, []bool{false, false, true, false, false}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := Compare(test.in, test.param)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("Compare(%v, %v) = %v, want %v", test.in, test.param, out, test.out)
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name  string
		in    vector.Vector
		param string
		out   vector.Vector
	}{
		{
			name:  "string",
			in:    vector.StringWithNA([]string{"ab", "abc", "", "abc", "abcd"}, []bool{false, false, true, false, false}),
			param: "abc",
			out:   vector.BooleanWithNA([]bool{false, true, false, true, true}, []bool{false, false, true, false, false}),
		},
		{
			name:  "int",
			in:    vector.IntegerWithNA([]int{1, 112, 0, 211, 2}, []bool{false, false, true, false, false}),
			param: "11",
			out:   vector.BooleanWithNA([]bool{false, true, false, true, false}, []bool{false, false, true, false, false}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := Contains(test.in, test.param)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("Contains(%v, %v) = %v, want %v", test.in, test.param, out, test.out)
			}
		})
	}
}

func TestContainsAny(t *testing.T) {
	tests := []struct {
		name  string
		in    vector.Vector
		param string
		out   vector.Vector
	}{
		{
			name:  "string",
			in:    vector.StringWithNA([]string{"ab", "def", "", "abc", "abcd"}, []bool{false, false, true, false, false}),
			param: "abc",
			out:   vector.BooleanWithNA([]bool{true, false, false, true, true}, []bool{false, false, true, false, false}),
		},
		{
			name:  "int",
			in:    vector.IntegerWithNA([]int{1, 112, 0, 211, 2}, []bool{false, false, true, false, false}),
			param: "11",
			out:   vector.BooleanWithNA([]bool{true, true, false, true, false}, []bool{false, false, true, false, false}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := ContainsAny(test.in, test.param)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("ContainsAny(%v, %v) = %v, want %v", test.in, test.param, out, test.out)
			}
		})
	}
}

func TestContainsRune(t *testing.T) {
	tests := []struct {
		name  string
		in    vector.Vector
		param rune
		out   vector.Vector
	}{
		{
			name:  "string",
			in:    vector.StringWithNA([]string{"ab", "def", "", "abc", "abcd"}, []bool{false, false, true, false, false}),
			param: 'a',
			out:   vector.BooleanWithNA([]bool{true, false, false, true, true}, []bool{false, false, true, false, false}),
		},
		{
			name:  "int",
			in:    vector.IntegerWithNA([]int{1, 112, 0, 211, 2}, []bool{false, false, true, false, false}),
			param: '1',
			out:   vector.BooleanWithNA([]bool{true, true, false, true, false}, []bool{false, false, true, false, false}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := ContainsRune(test.in, test.param)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("ContainsRune(%v, %v) = %v, want %v", test.in, test.param, out, test.out)
			}
		})
	}
}

func TestCount(t *testing.T) {
	tests := []struct {
		name  string
		in    vector.Vector
		param string
		out   vector.Vector
	}{
		{
			name:  "string",
			in:    vector.StringWithNA([]string{"ab", "def", "", "abab", "abcd"}, []bool{false, false, true, false, false}),
			param: "ab",
			out:   vector.IntegerWithNA([]int{1, 0, 0, 2, 1}, []bool{false, false, true, false, false}),
		},
		{
			name:  "int",
			in:    vector.IntegerWithNA([]int{1, 112, 0, 11011, 2}, []bool{false, false, true, false, false}),
			param: "11",
			out:   vector.IntegerWithNA([]int{0, 1, 0, 2, 0}, []bool{false, false, true, false, false}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := Count(test.in, test.param)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("Count(%v, %v) = %v, want %v", test.in, test.param, out, test.out)
			}
		})
	}
}

func TestEqualFold(t *testing.T) {
	tests := []struct {
		name  string
		in    vector.Vector
		param string
		out   vector.Vector
	}{
		{
			name:  "string",
			in:    vector.StringWithNA([]string{"ab", "def", "", "abab", "ab"}, []bool{false, false, true, false, false}),
			param: "AB",
			out:   vector.BooleanWithNA([]bool{true, false, false, false, true}, []bool{false, false, true, false, false}),
		},
		{
			name:  "int",
			in:    vector.IntegerWithNA([]int{11, 22, 0, 22, 11}, []bool{false, false, true, false, false}),
			param: "11",
			out:   vector.BooleanWithNA([]bool{true, false, false, false, true}, []bool{false, false, true, false, false}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := EqualFold(test.in, test.param)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("Count(%v, %v) = %v, want %v", test.in, test.param, out, test.out)
			}
		})
	}
}

func TestHasPrefix(t *testing.T) {
	tests := []struct {
		name  string
		in    vector.Vector
		param string
		out   vector.Vector
	}{
		{
			name:  "string",
			in:    vector.StringWithNA([]string{"ab", "def", "", "abab", "abcd"}, []bool{false, false, true, false, false}),
			param: "ab",
			out:   vector.BooleanWithNA([]bool{true, false, false, true, true}, []bool{false, false, true, false, false}),
		},
		{
			name:  "int",
			in:    vector.IntegerWithNA([]int{11, 112, 0, 211, 2}, []bool{false, false, true, false, false}),
			param: "11",
			out:   vector.BooleanWithNA([]bool{true, true, false, false, false}, []bool{false, false, true, false, false}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := HasPrefix(test.in, test.param)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("HasPrefix(%v, %v) = %v, want %v", test.in, test.param, out, test.out)
			}
		})
	}
}

func TestHasSuffix(t *testing.T) {
	tests := []struct {
		name  string
		in    vector.Vector
		param string
		out   vector.Vector
	}{
		{
			name:  "string",
			in:    vector.StringWithNA([]string{"ab", "def", "", "abab", "abcd"}, []bool{false, false, true, false, false}),
			param: "ab",
			out:   vector.BooleanWithNA([]bool{true, false, false, true, false}, []bool{false, false, true, false, false}),
		},
		{
			name:  "int",
			in:    vector.IntegerWithNA([]int{11, 112, 0, 211, 2}, []bool{false, false, true, false, false}),
			param: "11",
			out:   vector.BooleanWithNA([]bool{true, false, false, true, false}, []bool{false, false, true, false, false}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := HasSuffix(test.in, test.param)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("HasSuffix(%v, %v) = %v, want %v", test.in, test.param, out, test.out)
			}
		})
	}
}

func TestIndex(t *testing.T) {
	tests := []struct {
		name  string
		in    vector.Vector
		param string
		out   vector.Vector
	}{
		{
			name:  "string",
			in:    vector.StringWithNA([]string{"ab", "def", "", "abab", "abcd"}, []bool{false, false, true, false, false}),
			param: "ab",
			out:   vector.IntegerWithNA([]int{0, -1, -1, 0, 0}, []bool{false, false, true, false, false}),
		},
		{
			name:  "int",
			in:    vector.IntegerWithNA([]int{11, 112, 0, 211, 2}, []bool{false, false, true, false, false}),
			param: "11",
			out:   vector.IntegerWithNA([]int{0, 0, -1, 1, -1}, []bool{false, false, true, false, false}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := Index(test.in, test.param)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("Index(%v, %v) = %v, want %v", test.in, test.param, out, test.out)
			}
		})
	}
}

func TestIndexAny(t *testing.T) {
	tests := []struct {
		name  string
		in    vector.Vector
		param string
		out   vector.Vector
	}{
		{
			name:  "string",
			in:    vector.StringWithNA([]string{"ab", "def", "", "abab", "abcd"}, []bool{false, false, true, false, false}),
			param: "ab",
			out:   vector.IntegerWithNA([]int{0, -1, -1, 0, 0}, []bool{false, false, true, false, false}),
		},
		{
			name:  "int",
			in:    vector.IntegerWithNA([]int{11, 112, 0, 211, 2}, []bool{false, false, true, false, false}),
			param: "11",
			out:   vector.IntegerWithNA([]int{0, 0, -1, 1, -1}, []bool{false, false, true, false, false}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := IndexAny(test.in, test.param)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("IndexAny(%v, %v) = %v, want %v", test.in, test.param, out, test.out)
			}
		})
	}
}

func TestIndexByte(t *testing.T) {
	tests := []struct {
		name  string
		in    vector.Vector
		param byte
		out   vector.Vector
	}{
		{
			name:  "string",
			in:    vector.StringWithNA([]string{"ab", "def", "", "abab", "abcd"}, []bool{false, false, true, false, false}),
			param: 'a',
			out:   vector.IntegerWithNA([]int{0, -1, -1, 0, 0}, []bool{false, false, true, false, false}),
		},
		{
			name:  "int",
			in:    vector.IntegerWithNA([]int{11, 112, 0, 211, 2}, []bool{false, false, true, false, false}),
			param: '1',
			out:   vector.IntegerWithNA([]int{0, 0, -1, 1, -1}, []bool{false, false, true, false, false}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := IndexByte(test.in, test.param)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("IndexByte(%v, %v) = %v, want %v", test.in, test.param, out, test.out)
			}
		})
	}
}

func TestIndexRune(t *testing.T) {
	tests := []struct {
		name  string
		in    vector.Vector
		param rune
		out   vector.Vector
	}{
		{
			name:  "string",
			in:    vector.StringWithNA([]string{"ab", "def", "", "abab", "abcd"}, []bool{false, false, true, false, false}),
			param: 'a',
			out:   vector.IntegerWithNA([]int{0, -1, -1, 0, 0}, []bool{false, false, true, false, false}),
		},
		{
			name:  "int",
			in:    vector.IntegerWithNA([]int{11, 112, 0, 211, 2}, []bool{false, false, true, false, false}),
			param: '1',
			out:   vector.IntegerWithNA([]int{0, 0, -1, 1, -1}, []bool{false, false, true, false, false}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := IndexRune(test.in, test.param)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("IndexRune(%v, %v) = %v, want %v", test.in, test.param, out, test.out)
			}
		})
	}
}

func TestFields(t *testing.T) {
	tests := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "string",
			in:   vector.StringWithNA([]string{"ab bc bc ab", "de fg", "", "ab ab", "abcd"}, []bool{false, false, true, false, false}),
			out: vector.VectorVector([]vector.Vector{
				vector.String([]string{"ab", "bc", "bc", "ab"}),
				vector.String([]string{"de", "fg"}),
				nil,
				vector.String([]string{"ab", "ab"}),
				vector.String([]string{"abcd"}),
			}),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{11, 11021, 0, 201, 2}, []bool{false, false, true, false, false}),
			out: vector.VectorVector([]vector.Vector{
				vector.String([]string{"11"}),
				vector.String([]string{"11021"}),
				nil,
				vector.String([]string{"201"}),
				vector.String([]string{"2"}),
			}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := Fields(test.in)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("Fields(%v) = %v, want %v", test.in, out, test.out)
			}
		})
	}
}

func TestFieldsFunc(t *testing.T) {
	tests := []struct {
		name string
		in   vector.Vector
		fn   func(rune) bool
		out  vector.Vector
	}{
		{
			name: "string",
			in:   vector.StringWithNA([]string{"ab bc bc ab", "de fg", "", "ab ab", "abcd"}, []bool{false, false, true, false, false}),
			fn:   unicode.IsSpace,
			out: vector.VectorVector([]vector.Vector{
				vector.String([]string{"ab", "bc", "bc", "ab"}),
				vector.String([]string{"de", "fg"}),
				nil,
				vector.String([]string{"ab", "ab"}),
				vector.String([]string{"abcd"}),
			}),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{11, 11021, 0, 201, 2}, []bool{false, false, true, false, false}),
			fn:   func(r rune) bool { return r == '0' },
			out: vector.VectorVector([]vector.Vector{
				vector.String([]string{"11"}),
				vector.String([]string{"11", "21"}),
				nil,
				vector.String([]string{"2", "1"}),
				vector.String([]string{"2"}),
			}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := FieldsFunc(test.in, test.fn)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("FieldsFunc(%v, func) = %v, want %v", test.in, out, test.out)
			}
		})
	}
}

func TestLastIndex(t *testing.T) {
	tests := []struct {
		name  string
		in    vector.Vector
		param string
		out   vector.Vector
	}{
		{
			name:  "string",
			in:    vector.StringWithNA([]string{"ab", "def", "", "abab", "abcd"}, []bool{false, false, true, false, false}),
			param: "ab",
			out:   vector.IntegerWithNA([]int{0, -1, -1, 2, 0}, []bool{false, false, true, false, false}),
		},
		{
			name:  "int",
			in:    vector.IntegerWithNA([]int{11, 112, 0, 211, 2}, []bool{false, false, true, false, false}),
			param: "11",
			out:   vector.IntegerWithNA([]int{0, 0, -1, 1, -1}, []bool{false, false, true, false, false}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := LastIndex(test.in, test.param)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("Index(%v, %v) = %v, want %v", test.in, test.param, out, test.out)
			}
		})
	}
}

func TestLastIndexByte(t *testing.T) {
	tests := []struct {
		name  string
		in    vector.Vector
		param byte
		out   vector.Vector
	}{
		{
			name:  "string",
			in:    vector.StringWithNA([]string{"ab", "def", "", "abab", "abcd"}, []bool{false, false, true, false, false}),
			param: 'a',
			out:   vector.IntegerWithNA([]int{0, -1, -1, 2, 0}, []bool{false, false, true, false, false}),
		},
		{
			name:  "int",
			in:    vector.IntegerWithNA([]int{11, 112, 0, 211, 2}, []bool{false, false, true, false, false}),
			param: '1',
			out:   vector.IntegerWithNA([]int{1, 1, -1, 2, -1}, []bool{false, false, true, false, false}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := LastIndexByte(test.in, test.param)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("IndexByte(%v, %v) = %v, want %v", test.in, test.param, out, test.out)
			}
		})
	}
}

func TestLastIndexAny(t *testing.T) {
	tests := []struct {
		name  string
		in    vector.Vector
		param string
		out   vector.Vector
	}{
		{
			name:  "string",
			in:    vector.StringWithNA([]string{"ab", "def", "", "abab", "abcd"}, []bool{false, false, true, false, false}),
			param: "ab",
			out:   vector.IntegerWithNA([]int{1, -1, -1, 3, 1}, []bool{false, false, true, false, false}),
		},
		{
			name:  "int",
			in:    vector.IntegerWithNA([]int{11, 112, 0, 211, 2}, []bool{false, false, true, false, false}),
			param: "11",
			out:   vector.IntegerWithNA([]int{0, 0, -1, 1, -1}, []bool{false, false, true, false, false}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := LastIndexAny(test.in, test.param)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("IndexAny(%v, %v) = %v, want %v", test.in, test.param, out, test.out)
			}
		})
	}
}
