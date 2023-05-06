package apply

import (
	"logarithmotechnia/vector"
	"strings"
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
