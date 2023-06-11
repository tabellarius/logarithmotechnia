package which

import (
	"logarithmotechnia/vector"
	"math"
	"math/cmplx"
)

// IsInf returns a boolean slice indicating whether the corresponding element
// of the input vector is infinite.
func IsInf(v vector.Vector) []bool {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64) bool {
			return math.IsInf(val, 0)
		}
	case vector.PayloadTypeComplex:
		fn = func(val complex128) bool {
			return cmplx.IsInf(val)
		}
	default:
		return make([]bool, v.Len())
	}

	return v.Which(fn)
}

// IsNaN returns a boolean slice indicating whether the corresponding element
// of the input vector is NaN.
func IsNaN(v vector.Vector) []bool {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64) bool {
			return math.IsNaN(val)
		}
	case vector.PayloadTypeComplex:
		fn = func(val complex128) bool {
			return cmplx.IsNaN(val)
		}
	default:
		return make([]bool, v.Len())
	}

	return v.Which(fn)
}

// Signbit returns a boolean slice indicating whether the corresponding element
// of the input vector is negative.
func Signbit(v vector.Vector) []bool {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64) bool {
			return math.Signbit(val)
		}
	case vector.PayloadTypeInteger:
		fn = func(val int) bool {
			return val < 0
		}
	default:
		return make([]bool, v.Len())
	}

	return v.Which(fn)
}
