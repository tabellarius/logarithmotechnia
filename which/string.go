package which

import (
	"logarithmotechnia/vector"
	"strings"
)

func Contains(v vector.Vector, s string) []bool {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		return make([]bool, v.Len())
	}

	return vec.Which(func(val string) bool {
		return strings.Contains(val, s)
	})
}

func ContainsAny(v vector.Vector, s string) []bool {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		return make([]bool, v.Len())
	}

	return vec.Which(func(val string) bool {
		return strings.ContainsAny(val, s)
	})
}

func ContainsRune(v vector.Vector, r rune) []bool {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		return make([]bool, v.Len())
	}

	return vec.Which(func(val string) bool {
		return strings.ContainsRune(val, r)
	})
}

func HasPrefix(v vector.Vector, prefix string) []bool {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		return make([]bool, v.Len())
	}

	return vec.Which(func(val string) bool {
		return strings.HasPrefix(val, prefix)
	})
}

func HasSuffix(v vector.Vector, suffix string) []bool {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		return make([]bool, v.Len())
	}

	return vec.Which(func(val string) bool {
		return strings.HasSuffix(val, suffix)
	})
}

func EqualFold(v vector.Vector, t string) []bool {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		return make([]bool, v.Len())
	}

	return vec.Which(func(val string) bool {
		return strings.EqualFold(val, t)
	})
}
