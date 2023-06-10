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
			param: "12",
			out:   vector.IntegerWithNA([]int{1, 2, -1, 2, 0}, []bool{false, false, true, false, false}),
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

func TestLastIndexFunc(t *testing.T) {
	tests := []struct {
		name string
		in   vector.Vector
		fn   func(rune) bool
		out  vector.Vector
	}{
		{
			name: "string",
			in:   vector.StringWithNA([]string{"ab ab", "def", "", "abab", "ab cd"}, []bool{false, false, true, false, false}),
			fn:   unicode.IsSpace,
			out:  vector.IntegerWithNA([]int{2, -1, -1, -1, 2}, []bool{false, false, true, false, false}),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{11011, 102, 0, 211, 2}, []bool{false, false, true, false, false}),
			fn:   func(r rune) bool { return r == '0' },
			out:  vector.IntegerWithNA([]int{2, 1, 0, -1, -1}, []bool{false, false, true, false, false}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := LastIndexFunc(test.in, test.fn)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("IndexFunc(%v, func) = %v, want %v", test.in, out, test.out)
			}
		})
	}
}

func TestRepeat(t *testing.T) {
	tests := []struct {
		name string
		in   vector.Vector
		n    int
		out  vector.Vector
	}{
		{
			name: "string",
			in:   vector.StringWithNA([]string{"ab ab", "def", "", "abab", "ab cd"}, []bool{false, false, true, false, false}),
			n:    2,
			out:  vector.StringWithNA([]string{"ab abab ab", "defdef", "", "abababab", "ab cdab cd"}, []bool{false, false, true, false, false}),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{11011, 102, 0, 211, 2}, []bool{false, false, true, false, false}),
			n:    3,
			out:  vector.StringWithNA([]string{"110111101111011", "102102102", "", "211211211", "222"}, []bool{false, false, true, false, false}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := Repeat(test.in, test.n)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("Repeat(%v, %v) = %v, want %v", test.in, test.n, out, test.out)
			}
		})
	}
}

func TestReplace(t *testing.T) {
	tests := []struct {
		name string
		in   vector.Vector
		old  string
		new  string
		out  vector.Vector
	}{
		{
			name: "string",
			in:   vector.StringWithNA([]string{"ab ab", "def", "", "abab", "ab cd"}, []bool{false, false, true, false, false}),
			old:  "ab",
			new:  "cd",
			out:  vector.StringWithNA([]string{"cd ab", "def", "", "cdab", "cd cd"}, []bool{false, false, true, false, false}),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{11011, 102, 0, 211, 2}, []bool{false, false, true, false, false}),
			old:  "1",
			new:  "2",
			out:  vector.StringWithNA([]string{"21011", "202", "", "221", "2"}, []bool{false, false, true, false, false}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := Replace(test.in, test.old, test.new, 1)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("Replace(%v, %v, %v) = %v, want %v", test.in, test.old, test.new, out, test.out)
			}
		})
	}
}

func TestReplaceAll(t *testing.T) {
	tests := []struct {
		name string
		in   vector.Vector
		old  string
		new  string
		out  vector.Vector
	}{
		{
			name: "string",
			in:   vector.StringWithNA([]string{"ab ab", "def", "", "abab", "ab cd"}, []bool{false, false, true, false, false}),
			old:  "ab",
			new:  "cd",
			out:  vector.StringWithNA([]string{"cd cd", "def", "", "cdcd", "cd cd"}, []bool{false, false, true, false, false}),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{11011, 102, 0, 211, 2}, []bool{false, false, true, false, false}),
			old:  "1",
			new:  "2",
			out:  vector.StringWithNA([]string{"22022", "202", "", "222", "2"}, []bool{false, false, true, false, false}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := ReplaceAll(test.in, test.old, test.new)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("ReplaceAll(%v, %v, %v) = %v, want %v", test.in, test.old, test.new, out, test.out)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		name string
		in   vector.Vector
		sep  string
		out  vector.Vector
	}{
		{
			name: "string",
			in:   vector.StringWithNA([]string{"ab ab", "def", "", "abab", "ab cd"}, []bool{false, false, true, false, false}),
			sep:  " ",
			out: vector.VectorVector([]vector.Vector{
				vector.String([]string{"ab", "ab"}),
				vector.String([]string{"def"}),
				nil,
				vector.String([]string{"abab"}),
				vector.String([]string{"ab", "cd"}),
			}),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{11011, 102, 0, 211, 2}, []bool{false, false, true, false, false}),
			sep:  "0",
			out: vector.VectorVector([]vector.Vector{
				vector.String([]string{"11", "11"}),
				vector.String([]string{"1", "2"}),
				nil,
				vector.String([]string{"211"}),
				vector.String([]string{"2"}),
			}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := Split(test.in, test.sep)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("Split(%v, %v) = %v, want %v", test.in, test.sep, out, test.out)
			}
		})
	}
}

func TestSplitAfter(t *testing.T) {
	tests := []struct {
		name string
		in   vector.Vector
		sep  string
		out  vector.Vector
	}{
		{
			name: "string",
			in:   vector.StringWithNA([]string{"ab ab", "def", "", "abab", "ab cd"}, []bool{false, false, true, false, false}),
			sep:  " ",
			out: vector.VectorVector([]vector.Vector{
				vector.String([]string{"ab ", "ab"}),
				vector.String([]string{"def"}),
				nil,
				vector.String([]string{"abab"}),
				vector.String([]string{"ab ", "cd"}),
			}),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{11011, 102, 0, 211, 2}, []bool{false, false, true, false, false}),
			sep:  "0",
			out: vector.VectorVector([]vector.Vector{
				vector.String([]string{"110", "11"}),
				vector.String([]string{"10", "2"}),
				nil,
				vector.String([]string{"211"}),
				vector.String([]string{"2"}),
			}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := SplitAfter(test.in, test.sep)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("SplitAfter(%v, %v) = %v, want %v", test.in, test.sep, out, test.out)
			}
		})
	}
}

func TestSplitN(t *testing.T) {
	tests := []struct {
		name string
		in   vector.Vector
		sep  string
		n    int
		out  vector.Vector
	}{
		{
			name: "string",
			in:   vector.StringWithNA([]string{"ab ab", "def", "", "abab", "ab cd"}, []bool{false, false, true, false, false}),
			sep:  " ",
			n:    1,
			out: vector.VectorVector([]vector.Vector{
				vector.String([]string{"ab ab"}),
				vector.String([]string{"def"}),
				nil,
				vector.String([]string{"abab"}),
				vector.String([]string{"ab cd"}),
			}),
		},
		{
			name: "string + n = 0",
			in:   vector.StringWithNA([]string{"ab ab", "def", "", "abab", "ab cd"}, []bool{false, false, true, false, false}),
			sep:  " ",
			n:    0,
			out: vector.VectorVector([]vector.Vector{
				vector.String(nil),
				vector.String(nil),
				nil,
				vector.String(nil),
				vector.String(nil),
			}),
		},
		{
			name: "string",
			in:   vector.StringWithNA([]string{"ab ab ab", "def", "", "abab", "ab cd"}, []bool{false, false, true, false, false}),
			sep:  " ",
			n:    2,
			out: vector.VectorVector([]vector.Vector{
				vector.String([]string{"ab", "ab ab"}),
				vector.String([]string{"def"}),
				nil,
				vector.String([]string{"abab"}),
				vector.String([]string{"ab", "cd"}),
			}),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{11011, 102, 0, 211, 2}, []bool{false, false, true, false, false}),
			sep:  "0",
			n:    2,
			out: vector.VectorVector([]vector.Vector{
				vector.String([]string{"11", "11"}),
				vector.String([]string{"1", "2"}),
				nil,
				vector.String([]string{"211"}),
				vector.String([]string{"2"}),
			}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := SplitN(test.in, test.sep, test.n)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("SplitN(%v, %v, %v) = %v, want %v", test.in, test.sep, test.n, out, test.out)
			}
		})
	}
}

func TestSplitAfterN(t *testing.T) {
	tests := []struct {
		name string
		in   vector.Vector
		sep  string
		n    int
		out  vector.Vector
	}{
		{
			name: "string + n = 1",
			in:   vector.StringWithNA([]string{"ab ab", "def", "", "abab", "ab cd"}, []bool{false, false, true, false, false}),
			sep:  " ",
			n:    1,
			out: vector.VectorVector([]vector.Vector{
				vector.String([]string{"ab ab"}),
				vector.String([]string{"def"}),
				nil,
				vector.String([]string{"abab"}),
				vector.String([]string{"ab cd"}),
			}),
		},
		{
			name: "string + n = 0",
			in:   vector.StringWithNA([]string{"ab ab", "def", "", "abab", "ab cd"}, []bool{false, false, true, false, false}),
			sep:  " ",
			n:    0,
			out: vector.VectorVector([]vector.Vector{
				vector.String(nil),
				vector.String(nil),
				nil,
				vector.String(nil),
				vector.String(nil),
			}),
		},
		{
			name: "string + n = 2",
			in:   vector.StringWithNA([]string{"ab ab ab", "def", "", "abab", "ab cd"}, []bool{false, false, true, false, false}),
			sep:  " ",
			n:    2,
			out: vector.VectorVector([]vector.Vector{
				vector.String([]string{"ab ", "ab ab"}),
				vector.String([]string{"def"}),
				nil,
				vector.String([]string{"abab"}),
				vector.String([]string{"ab ", "cd"}),
			}),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{11011, 102, 0, 211, 2}, []bool{false, false, true, false, false}),
			sep:  "0",
			n:    2,
			out: vector.VectorVector([]vector.Vector{
				vector.String([]string{"110", "11"}),
				vector.String([]string{"10", "2"}),
				nil,
				vector.String([]string{"211"}),
				vector.String([]string{"2"}),
			}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := SplitAfterN(test.in, test.sep, test.n)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("SplitN(%v, %v, %v) = %v, want %v", test.in, test.sep, test.n, out, test.out)
			}
		})
	}
}

func TestToLower(t *testing.T) {
	tests := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "string",
			in:   vector.String([]string{"aB aЩ", "DEF", "", "ABab", "ΧΞΙ"}),
			out:  vector.String([]string{"ab aщ", "def", "", "abab", "χξι"}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := ToLower(test.in)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("ToLower(%v) = %v, want %v", test.in, out, test.out)
			}
		})
	}
}

func TestToLowerSpecial(t *testing.T) {
	tests := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "string",
			in:   vector.String([]string{"aB aЩ", "DEF", "", "ABab", "Önnek İş"}),
			out:  vector.String([]string{"ab aщ", "def", "", "abab", "önnek iş"}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := ToLowerSpecial(test.in, unicode.TurkishCase)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("ToLowerSpecial(%v, %v) = %v, want %v", test.in, "el", out, test.out)
			}
		})
	}
}

func TestToUpper(t *testing.T) {
	tests := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "string",
			in:   vector.String([]string{"aB aщ", "DEF", "", "ABab", "ΧξΙ"}),
			out:  vector.String([]string{"AB AЩ", "DEF", "", "ABAB", "ΧΞΙ"}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := ToUpper(test.in)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("ToUpper(%v) = %v, want %v", test.in, out, test.out)
			}
		})
	}
}

func TestToUpperSpecial(t *testing.T) {
	tests := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "string",
			in:   vector.String([]string{"aB aщ", "DEF", "", "ABab", "önnek iş"}),
			out:  vector.String([]string{"AB AЩ", "DEF", "", "ABAB", "ÖNNEK İŞ"}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := ToUpperSpecial(test.in, unicode.TurkishCase)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("ToUpperSpecial(%v, %v) = %v, want %v", test.in, "el", out, test.out)
			}
		})
	}
}

func TestToTitle(t *testing.T) {
	tests := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "string",
			in:   vector.String([]string{"aB aщ", "ǳef", "", "ABab", "ΧξΙ"}),
			out:  vector.String([]string{"AB AЩ", "ǲEF", "", "ABAB", "ΧΞΙ"}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := ToTitle(test.in)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("ToTitle(%v) = %v, want %v", test.in, out, test.out)
			}
		})
	}
}

func TestToTitleSpecial(t *testing.T) {
	tests := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "string",
			in:   vector.String([]string{"aB aщ", "ǳef", "", "ABab", "önnek iş"}),
			out:  vector.String([]string{"AB AЩ", "ǲEF", "", "ABAB", "ÖNNEK İŞ"}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := ToTitleSpecial(test.in, unicode.TurkishCase)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("ToTitleSpecial(%v, %v) = %v, want %v", test.in, "el", out, test.out)
			}
		})
	}
}

func TestTrim(t *testing.T) {
	tests := []struct {
		name string
		in   vector.Vector
		cut  string
		out  vector.Vector
	}{
		{
			name: "string",
			in:   vector.String([]string{"ab ab", "def", "", "abab", "ab cd"}),
			cut:  "ab",
			out:  vector.String([]string{" ", "def", "", "", " cd"}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := Trim(test.in, test.cut)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("Trim(%v, %v) = %v, want %v", test.in, test.cut, out, test.out)
			}
		})
	}
}

func TestTrimLeft(t *testing.T) {
	tests := []struct {
		name string
		in   vector.Vector
		cut  string
		out  vector.Vector
	}{
		{
			name: "string",
			in:   vector.String([]string{"ab ab", "def", "", "abab", "ab cd"}),
			cut:  "ab",
			out:  vector.String([]string{" ab", "def", "", "", " cd"}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := TrimLeft(test.in, test.cut)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("TrimLeft(%v, %v) = %v, want %v", test.in, test.cut, out, test.out)
			}
		})
	}
}

func TestTrimRight(t *testing.T) {
	tests := []struct {
		name string
		in   vector.Vector
		cut  string
		out  vector.Vector
	}{
		{
			name: "string",
			in:   vector.String([]string{"ab ab", "def", "", "abab", "ab cd"}),
			cut:  "ab",
			out:  vector.String([]string{"ab ", "def", "", "", "ab cd"}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := TrimRight(test.in, test.cut)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("TrimRight(%v, %v) = %v, want %v", test.in, test.cut, out, test.out)
			}
		})
	}
}

func TestTrimSpace(t *testing.T) {
	tests := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "string",
			in:   vector.String([]string{" ab ab ", "def", "", " abab", "ab cd "}),
			out:  vector.String([]string{"ab ab", "def", "", "abab", "ab cd"}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := TrimSpace(test.in)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("TrimSpace(%v) = %v, want %v", test.in, out, test.out)
			}
		})
	}
}

func TestTrimPrefix(t *testing.T) {
	tests := []struct {
		name string
		in   vector.Vector
		cut  string
		out  vector.Vector
	}{
		{
			name: "string",
			in:   vector.String([]string{"abab ab", "def", "", "abab", "ab cd"}),
			cut:  "ab",
			out:  vector.String([]string{"ab ab", "def", "", "ab", " cd"}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := TrimPrefix(test.in, test.cut)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("TrimPrefix(%v, %v) = %v, want %v", test.in, test.cut, out, test.out)
			}
		})
	}
}

func TestTrimSuffix(t *testing.T) {
	tests := []struct {
		name string
		in   vector.Vector
		cut  string
		out  vector.Vector
	}{
		{
			name: "string",
			in:   vector.String([]string{"ab abab", "def", "", "abab", "ab cd"}),
			cut:  "ab",
			out:  vector.String([]string{"ab ab", "def", "", "ab", "ab cd"}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := TrimSuffix(test.in, test.cut)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("TrimSuffix(%v, %v) = %v, want %v", test.in, test.cut, out, test.out)
			}
		})
	}
}

func TestTrimFunc(t *testing.T) {
	tests := []struct {
		name string
		in   vector.Vector
		f    func(rune) bool
		out  vector.Vector
	}{
		{
			name: "string",
			in:   vector.String([]string{"ab ab", "def", "", "abab", "ab cd"}),
			f:    func(r rune) bool { return r == 'a' || r == 'b' },
			out:  vector.String([]string{" ", "def", "", "", " cd"}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := TrimFunc(test.in, test.f)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("TrimFunc(%v, fn) = %v, want %v", test.in, out, test.out)
			}
		})
	}
}

func TestTrimLeftFunc(t *testing.T) {
	tests := []struct {
		name string
		in   vector.Vector
		f    func(rune) bool
		out  vector.Vector
	}{
		{
			name: "string",
			in:   vector.String([]string{"ab ab", "def", "", "abab", "ab cd"}),
			f:    func(r rune) bool { return r == 'a' || r == 'b' },
			out:  vector.String([]string{" ab", "def", "", "", " cd"}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := TrimLeftFunc(test.in, test.f)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("TrimLeftFunc(%v, fn) = %v, want %v", test.in, out, test.out)
			}
		})
	}
}

func TestTrimRightFunc(t *testing.T) {
	tests := []struct {
		name string
		in   vector.Vector
		f    func(rune) bool
		out  vector.Vector
	}{
		{
			name: "string",
			in:   vector.String([]string{"ab ab", "def", "", "abab", "ab cd"}),
			f:    func(r rune) bool { return r == 'a' || r == 'b' },
			out:  vector.String([]string{"ab ", "def", "", "", "ab cd"}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := TrimRightFunc(test.in, test.f)
			if !vector.CompareVectorsForTest(out, test.out) {
				t.Errorf("TrimRightFunc(%v, fn) = %v, want %v", test.in, out, test.out)
			}
		})
	}
}
