package apply

import (
	"logarithmotechnia/vector"
	"strings"
	"unicode"
)

// Compare returns a vector of integers representing the result of comparing
// each element of the input vector to the given string. The result will be
// -1 if the element is less than the string, 0 if the element is equal to the
// string, and 1 if the element is greater than the string.
func Compare(v vector.Vector, s string) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(val string, na bool) (int, bool) {
		if na {
			return 0, true
		}

		return strings.Compare(val, s), false
	})
}

// Contains returns a vector of booleans representing whether each element of
// the input vector contains the given string.
func Contains(v vector.Vector, s string) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(val string, na bool) (bool, bool) {
		if na {
			return false, true
		}

		return strings.Contains(val, s), false
	})
}

// ContainsAny returns a vector of booleans representing whether each element
// of the input vector contains any of the given characters.
func ContainsAny(v vector.Vector, chars string) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(val string, na bool) (bool, bool) {
		if na {
			return false, true
		}

		return strings.ContainsAny(val, chars), false
	})
}

// ContainsRune returns a vector of booleans representing whether each element
// of the input vector contains the given rune.
func ContainsRune(v vector.Vector, r rune) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(val string, na bool) (bool, bool) {
		if na {
			return false, true
		}

		return strings.ContainsRune(val, r), false
	})
}

// Count returns a vector of integers representing the number of non-overlapping
// instances of the given substring in each element of the input vector.
func Count(v vector.Vector, substr string) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(val string, na bool) (int, bool) {
		if na {
			return 0, true
		}

		return strings.Count(val, substr), false
	})
}

// EqualFold returns a vector of booleans representing whether each element of
// the input vector is equal to the given string, ignoring case.
func EqualFold(v vector.Vector, t string) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(val string, na bool) (bool, bool) {
		if na {
			return false, true
		}

		return strings.EqualFold(val, t), false
	})
}

// Fields returns a vector of vectors of strings representing the words in each
// element of the input vector.
func Fields(v vector.Vector) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(val string, na bool) (vector.Vector, bool) {
		if na {
			return nil, true
		}

		return vector.String(strings.Fields(val)), false
	})
}

// FieldsFunc returns a vector of vectors of strings representing the words in
// each element of the input vector, as split by the given function.
func FieldsFunc(v vector.Vector, f func(rune) bool) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(val string, na bool) (vector.Vector, bool) {
		if na {
			return nil, true
		}

		return vector.String(strings.FieldsFunc(val, f)), false
	})
}

// HasPrefix returns a vector of booleans representing whether each element of
// the input vector has the given prefix.
func HasPrefix(v vector.Vector, prefix string) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(val string, na bool) (bool, bool) {
		if na {
			return false, true
		}

		return strings.HasPrefix(val, prefix), false
	})
}

// HasSuffix returns a vector of booleans representing whether each element of
// the input vector has the given suffix.
func HasSuffix(v vector.Vector, suffix string) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(val string, na bool) (bool, bool) {
		if na {
			return false, true
		}

		return strings.HasSuffix(val, suffix), false
	})
}

// Index returns a vector of integers representing the index of the first
// instance of the given substring in each element of the input vector.
func Index(v vector.Vector, substr string) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(val string, na bool) (int, bool) {
		if na {
			return 0, true
		}

		return strings.Index(val, substr), false
	})
}

// IndexAny returns a vector of integers representing the index of the first
// instance of any of the given characters in each element of the input vector.
func IndexAny(v vector.Vector, chars string) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(val string, na bool) (int, bool) {
		if na {
			return 0, true
		}

		return strings.IndexAny(val, chars), false
	})
}

// IndexByte returns a vector of integers representing the index of the first
// instance of the given byte in each element of the input vector.
func IndexByte(v vector.Vector, c byte) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(val string, na bool) (int, bool) {
		if na {
			return 0, true
		}

		return strings.IndexByte(val, c), false
	})
}

// IndexRune returns a vector of integers representing the index of the first
// instance of the given rune in each element of the input vector.
func IndexRune(v vector.Vector, r rune) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(val string, na bool) (int, bool) {
		if na {
			return 0, true
		}

		return strings.IndexRune(val, r), false
	})
}

// LastIndex returns a vector of integers representing the index of the last
// instance of the given substring in each element of the input vector.
func LastIndex(v vector.Vector, substr string) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(s string, na bool) (int, bool) {
		if na {
			return 0, true
		}

		return strings.LastIndex(s, substr), false
	})
}

// LastIndexAny returns a vector of integers representing the index of the last
// instance of any of the given characters in each element of the input vector.
func LastIndexAny(v vector.Vector, chars string) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(s string, na bool) (int, bool) {
		if na {
			return 0, true
		}

		return strings.LastIndexAny(s, chars), false
	})
}

// LastIndexByte returns a vector of integers representing the index of the last
// instance of the given byte in each element of the input vector.
func LastIndexByte(v vector.Vector, c byte) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(s string, na bool) (int, bool) {
		if na {
			return 0, true
		}

		return strings.LastIndexByte(s, c), false
	})
}

// LastIndexFunc returns a vector of integers representing the index of the last
// instance of the given rune in each element of the input vector.
func LastIndexFunc(v vector.Vector, f func(rune) bool) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(s string, na bool) (int, bool) {
		if na {
			return 0, true
		}

		return strings.LastIndexFunc(s, f), false
	})
}

// Repeat returns a vector of strings representing the input vector repeated
// count times.
func Repeat(v vector.Vector, count int) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(s string, na bool) (string, bool) {
		if na {
			return "", true
		}

		return strings.Repeat(s, count), false
	})
}

// Replace returns a vector of strings representing the input vector with the
// first n non-overlapping instances of old replaced by new.
func Replace(v vector.Vector, old, new string, n int) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(s string, na bool) (string, bool) {
		if na {
			return "", true
		}

		return strings.Replace(s, old, new, n), false
	})
}

// ReplaceAll returns a vector of strings representing the input vector with all
// non-overlapping instances of old replaced by new.
func ReplaceAll(v vector.Vector, old, new string) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(s string, na bool) (string, bool) {
		if na {
			return "", true
		}

		return strings.ReplaceAll(s, old, new), false
	})
}

// Split returns a vector of vectors representing the input vector split into
// substrings separated by the given separator.
func Split(v vector.Vector, sep string) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(s string, na bool) (vector.Vector, bool) {
		if na {
			return nil, true
		}

		return vector.String(strings.Split(s, sep)), false
	})
}

// SplitAfter returns a vector of vectors representing the input vector split
// into substrings separated by the given separator, with the separator included
// in each substring.
func SplitAfter(v vector.Vector, sep string) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(s string, na bool) (vector.Vector, bool) {
		if na {
			return nil, true
		}

		return vector.String(strings.SplitAfter(s, sep)), false
	})
}

// SplitAfterN returns a vector of vectors representing the input vector split
// into substrings separated by the given separator, with the separator included
// in each substring, with the input vector split at most n times.
func SplitAfterN(v vector.Vector, sep string, n int) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(s string, na bool) (vector.Vector, bool) {
		if na {
			return nil, true
		}

		return vector.String(strings.SplitAfterN(s, sep, n)), false
	})
}

// SplitN returns a vector of vectors representing the input vector split into
// substrings separated by the given separator, with the input vector split at
// most n times.
func SplitN(v vector.Vector, sep string, n int) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(s string, na bool) (vector.Vector, bool) {
		if na {
			return nil, true
		}

		return vector.String(strings.SplitN(s, sep, n)), false
	})
}

// ToLower returns a vector of strings representing the input vector with all
// Unicode letters mapped to their lower case.
func ToLower(v vector.Vector) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(s string, na bool) (string, bool) {
		if na {
			return "", true
		}

		return strings.ToLower(s), false
	})
}

// ToLowerSpecial returns a vector of strings representing the input vector with
// all Unicode letters mapped to their lower case, giving priority to the
// special casing rules.
func ToLowerSpecial(v vector.Vector, c unicode.SpecialCase) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(s string, na bool) (string, bool) {
		if na {
			return "", true
		}

		return strings.ToLowerSpecial(c, s), false
	})
}

// ToTitle returns a vector of strings representing the input vector with all
// Unicode letters mapped to their title case.
func ToTitle(v vector.Vector) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(s string, na bool) (string, bool) {
		if na {
			return "", true
		}

		return strings.ToTitle(s), false
	})
}

// ToTitleSpecial returns a vector of strings representing the input vector with
// all Unicode letters mapped to their title case, giving priority to the
// special casing rules.
func ToTitleSpecial(v vector.Vector, c unicode.SpecialCase) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(s string, na bool) (string, bool) {
		if na {
			return "", true
		}

		return strings.ToTitleSpecial(c, s), false
	})
}

// ToUpper returns a vector of strings representing the input vector with all
// Unicode letters mapped to their upper case.
func ToUpper(v vector.Vector) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(s string, na bool) (string, bool) {
		if na {
			return "", true
		}

		return strings.ToUpper(s), false
	})
}

// ToUpperSpecial returns a vector of strings representing the input vector with
// all Unicode letters mapped to their upper case, giving priority to the
// special casing rules.
func ToUpperSpecial(v vector.Vector, c unicode.SpecialCase) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(s string, na bool) (string, bool) {
		if na {
			return "", true
		}

		return strings.ToUpperSpecial(c, s), false
	})
}

// Trim returns a vector of strings representing the input vector with all
// leading and trailing Unicode code points contained in cutset removed.
func Trim(v vector.Vector, cutset string) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(s string, na bool) (string, bool) {
		if na {
			return "", true
		}

		return strings.Trim(s, cutset), false
	})
}

// TrimFunc returns a vector of strings representing the input vector with all
// leading and trailing Unicode code points, for which the function returns true,
// removed.
func TrimFunc(v vector.Vector, f func(rune) bool) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(s string, na bool) (string, bool) {
		if na {
			return "", true
		}

		return strings.TrimFunc(s, f), false
	})
}

// TrimLeft returns a vector of strings representing the input vector with all
// leading Unicode code points contained in cutset removed.
func TrimLeft(v vector.Vector, cutset string) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(s string, na bool) (string, bool) {
		if na {
			return "", true
		}
		return strings.TrimLeft(s, cutset), false
	})
}

// TrimLeftFunc returns a vector of strings representing the input vector with all
// leading Unicode code points, for which the function returns true, removed.
func TrimLeftFunc(v vector.Vector, f func(rune) bool) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(s string, na bool) (string, bool) {
		if na {
			return "", true
		}
		return strings.TrimLeftFunc(s, f), false
	})
}

// TrimPrefix returns a vector of strings representing the input vector with all
// leading instances of prefix removed.
func TrimPrefix(v vector.Vector, prefix string) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(s string, na bool) (string, bool) {
		if na {
			return "", true
		}
		return strings.TrimPrefix(s, prefix), false
	})
}

// TrimRight returns a vector of strings representing the input vector with all
// trailing Unicode code points contained in cutset removed.
func TrimRight(v vector.Vector, cutset string) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(s string, na bool) (string, bool) {
		if na {
			return "", true
		}
		return strings.TrimRight(s, cutset), false
	})
}

// TrimRightFunc returns a vector of strings representing the input vector with all
// trailing Unicode code points, for which the function returns true, removed.
func TrimRightFunc(v vector.Vector, f func(rune) bool) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}

	return vec.Apply(func(s string, na bool) (string, bool) {
		if na {
			return "", true
		}
		return strings.TrimRightFunc(s, f), false
	})
}

// TrimSpace returns a vector of strings representing the input vector with all
// leading and trailing white space removed, as defined by Unicode.
func TrimSpace(v vector.Vector) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}
	return vec.Apply(func(s string, na bool) (string, bool) {
		if na {
			return "", true
		}
		return strings.TrimSpace(s), false
	})
}

// TrimSuffix returns a vector of strings representing the input vector with all
// trailing instances of suffix removed.
func TrimSuffix(v vector.Vector, suffix string) vector.Vector {
	vec := v
	if v.Type() != vector.PayloadTypeString {
		vec = v.AsString()
	}
	return vec.Apply(func(s string, na bool) (string, bool) {
		if na {
			return "", true
		}
		return strings.TrimSuffix(s, suffix), false
	})
}
