package which

import (
	"logarithmotechnia/vector"
	"math"
	"math/cmplx"
)

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
