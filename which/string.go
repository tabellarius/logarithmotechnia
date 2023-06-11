package which

import (
	"logarithmotechnia/vector"
	"strings"
)

// Contains returns a boolean slice indicating whether the corresponding element
// of the input vector contains the given substring.
func Contains(v vector.Vector, s string) []bool {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		return make([]bool, v.Len())
	}

	return vec.Which(func(val string) bool {
		return strings.Contains(val, s)
	})
}

// ContainsAny returns a boolean slice indicating whether the corresponding element
// of the input vector contains any of the given characters.
func ContainsAny(v vector.Vector, s string) []bool {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		return make([]bool, v.Len())
	}

	return vec.Which(func(val string) bool {
		return strings.ContainsAny(val, s)
	})
}

// ContainsRune returns a boolean slice indicating whether the corresponding element
// of the input vector contains the given rune.
func ContainsRune(v vector.Vector, r rune) []bool {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		return make([]bool, v.Len())
	}

	return vec.Which(func(val string) bool {
		return strings.ContainsRune(val, r)
	})
}

// HasPrefix returns a boolean slice indicating whether the corresponding element
// of the input vector has the given prefix.
func HasPrefix(v vector.Vector, prefix string) []bool {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		return make([]bool, v.Len())
	}

	return vec.Which(func(val string) bool {
		return strings.HasPrefix(val, prefix)
	})
}

// HasSuffix returns a boolean slice indicating whether the corresponding element
// of the input vector has the given suffix.
func HasSuffix(v vector.Vector, suffix string) []bool {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		return make([]bool, v.Len())
	}

	return vec.Which(func(val string) bool {
		return strings.HasSuffix(val, suffix)
	})
}

// EqualFold returns a boolean slice indicating whether the corresponding element
// of the input vector is equal to the given string, ignoring case.
func EqualFold(v vector.Vector, t string) []bool {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		return make([]bool, v.Len())
	}

	return vec.Which(func(val string) bool {
		return strings.EqualFold(val, t)
	})
}
