package apply

import (
	"logarithmotechnia/vector"
	"strings"
	"unicode"
)

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
