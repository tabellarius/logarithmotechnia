package vector

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDefNameable_SetName(t *testing.T) {
	testData := []struct {
		name string
		idx  int
		set  bool
	}{
		{"", 1, false},
		{"zero", 0, false},
		{"negative", -1, false},
		{"out of bounds", 6, false},
		{"regular I", 1, true},
		{"regular V", 5, true},
	}

	for _, data := range testData {
		t.Run(fmt.Sprintf(`[%d]"%s"`, data.idx, data.name), func(t *testing.T) {
			vec := newCommon(5)
			nm := newDefNameable(&vec)
			nm.SetName(data.idx, data.name)
			idx, ok := nm.names[data.name]
			if ok != data.set {
				t.Error(fmt.Sprintf("names[%s] was not set", data.name))
			}
			if ok && idx != data.idx {
				t.Error(fmt.Sprintf("names[%s] is not %d", data.name, data.idx))
			}
		})
	}
}

func TestDefNameable_SetNames(t *testing.T) {
	testData := []struct {
		name  string
		names []string
		set   []bool
	}{
		{
			name:  "full",
			names: []string{"one", "two", "three", "four", "five"},
		},
		{
			name:  "partial",
			names: []string{"one", "two", "three"},
		},
		{
			name:  "overfill",
			names: []string{"one", "two", "three", "four", "five", "six", "seven"},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			vec := newCommon(5)
			nm := newDefNameable(&vec)
			nm.SetNames(data.names)
			for i := 1; i <= len(data.names); i++ {
				nameIndex := i - 1
				idx, ok := nm.names[data.names[nameIndex]]
				if ok {
					if i > nm.vec.Length() {
						t.Error(fmt.Sprintf("names[%s] (%d-th) can not be set on this vector",
							data.names[nameIndex], i))
					} else if idx != i {
						t.Error(fmt.Sprintf("names[%s] is not %d but %d", data.names[nameIndex], i, idx))
					}
				} else {
					if i <= nm.vec.Length() {
						t.Error(fmt.Sprintf("names[%s] was not set", data.names[nameIndex]))
					}
				}
			}
		})
	}
}

func TestDefNameable_SetNamesMap(t *testing.T) {
	testData := []struct {
		name     string
		namesMap map[string]int
		expected map[string]int
	}{
		{
			name:     "correct",
			namesMap: map[string]int{"one": 1, "two": 2, "four": 4},
			expected: map[string]int{"one": 1, "two": 2, "four": 4},
		},
		{
			name:     "with incorrect fields",
			namesMap: map[string]int{"minus one": -1, "zero": 0, "two": 2, "four": 4, "seven": 7},
			expected: map[string]int{"two": 2, "four": 4},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			vec := newCommon(5)
			nm := newDefNameable(&vec)
			nm.SetNamesMap(data.namesMap)
			if !reflect.DeepEqual(nm.names, data.expected) {
				t.Error("Expected map is not equal to result map")
			}
		})
	}
}

func TestDefNameable_HasName(t *testing.T) {
	vec := newCommon(5)
	nm := newDefNameable(&vec)
	nm.names = map[string]int{"one": 1}
	if !nm.HasName("one") {
		t.Error("Set name is not reported")
	}
	if nm.HasName("none") {
		t.Error("Not set name is reported")
	}
}

func TestDefNameable_HasNameFor(t *testing.T) {
	vec := newCommon(5)
	nm := newDefNameable(&vec)
	nm.names = map[string]int{"one": 1, "two": 2, "four": 4}
	testData := []struct {
		idx int
		set bool
	}{
		{0, false}, {1, true}, {2, true},
		{3, false}, {4, true}, {5, false},
	}

	for _, data := range testData {
		t.Run(fmt.Sprintf("Index_%d", data.idx), func(t *testing.T) {
			if nm.HasNameFor(data.idx) != data.set {
				t.Error(fmt.Sprintf("HasNameFor(%d) is not %t", data.idx, data.set))
			}
		})
	}
}

func TestDefNameable_Names(t *testing.T) {
	testData := []struct {
		name     string
		namesMap map[string]int
		expected []string
	}{
		{
			name:     "one to three",
			namesMap: map[string]int{"one": 1, "two": 2, "three": 3},
			expected: []string{"one", "two", "three", "", ""},
		},
		{
			name:     "two to four",
			namesMap: map[string]int{"two": 2, "three": 3, "four": 4},
			expected: []string{"", "two", "three", "four", ""},
		},
		{
			name:     "full",
			namesMap: map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5},
			expected: []string{"one", "two", "three", "four", "five"},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			vec := newCommon(5)
			nm := newDefNameable(&vec)
			nm.names = data.namesMap
			names := nm.Names()
			if !reflect.DeepEqual(names, data.expected) {
				t.Error(fmt.Sprintf("Expected string array (%v) is not equal to result (%v)",
					names, data.expected))
			}
		})
	}
}

func TestDefNameable_NamesMap(t *testing.T) {
	testData := []struct {
		name     string
		namesMap map[string]int
		expected map[string]int
	}{
		{
			name:     "one to three",
			namesMap: map[string]int{"one": 1, "two": 2, "three": 3},
			expected: map[string]int{"one": 1, "two": 2, "three": 3},
		},
		{
			name:     "two to four",
			namesMap: map[string]int{"two": 2, "three": 3, "four": 4},
			expected: map[string]int{"two": 2, "three": 3, "four": 4},
		},
		{
			name:     "full",
			namesMap: map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5},
			expected: map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			vec := newCommon(5)
			nm := newDefNameable(&vec)
			nm.names = data.namesMap
			names := nm.NamesMap()
			if reflect.ValueOf(names).Pointer() == reflect.ValueOf(nm.names).Pointer() {
				t.Error(fmt.Sprintf("Returned map is not a copy"))
			}
			if !reflect.DeepEqual(names, data.expected) {
				t.Error(fmt.Sprintf("Expected string map (%v) is not equal to result (%v)",
					names, data.expected))
			}
		})
	}
}
