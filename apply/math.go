package apply

import (
	"logarithmotechnia/vector"
	"math"
	"math/cmplx"
)

func Abs(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Abs(val), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Abs(float64(val)), na
		}
	case vector.PayloadTypeComplex:
		fn = func(val complex128, na bool) (float64, bool) {
			return cmplx.Abs(val), na
		}
	default:
		vec := v.AsFloat()
		return Abs(vec)
	}

	return v.Apply(fn)
}

func Acos(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Acos(val), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Acos(float64(val)), na
		}
	case vector.PayloadTypeComplex:
		fn = func(val complex128, na bool) (complex128, bool) {
			return cmplx.Acos(val), na
		}
	default:
		vec := v.AsFloat()
		return Acos(vec)
	}

	return v.Apply(fn)
}

func Acosh(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Acosh(val), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Acosh(float64(val)), na
		}
	case vector.PayloadTypeComplex:
		fn = func(val complex128, na bool) (complex128, bool) {
			return cmplx.Acosh(val), na
		}
	default:
		vec := v.AsFloat()
		return Acosh(vec)
	}

	return v.Apply(fn)
}

func Asin(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Asin(val), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Asin(float64(val)), na
		}
	case vector.PayloadTypeComplex:
		fn = func(val complex128, na bool) (complex128, bool) {
			return cmplx.Asin(val), na
		}
	default:
		vec := v.AsFloat()
		return Asin(vec)
	}

	return v.Apply(fn)
}

func Asinh(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Asinh(val), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Asinh(float64(val)), na
		}
	case vector.PayloadTypeComplex:
		fn = func(val complex128, na bool) (complex128, bool) {
			return cmplx.Asinh(val), na
		}
	default:
		vec := v.AsFloat()
		return Asinh(vec)
	}

	return v.Apply(fn)
}

func Atan(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Atan(val), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Atan(float64(val)), na
		}
	case vector.PayloadTypeComplex:
		fn = func(val complex128, na bool) (complex128, bool) {
			return cmplx.Atan(val), na
		}
	default:
		vec := v.AsFloat()
		return Atan(vec)
	}

	return v.Apply(fn)
}

func Atanh(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Atanh(val), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Atanh(float64(val)), na
		}
	case vector.PayloadTypeComplex:
		fn = func(val complex128, na bool) (complex128, bool) {
			return cmplx.Atanh(val), na
		}
	default:
		vec := v.AsFloat()
		return Atanh(vec)
	}

	return v.Apply(fn)
}

func Atan2(v vector.Vector, x float64) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Atan2(val, x), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Atan2(float64(val), x), na
		}
	default:
		vec := v.AsFloat()
		return Atan2(vec, x)
	}

	return v.Apply(fn)
}

func Cbrt(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Cbrt(val), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Cbrt(float64(val)), na
		}
	default:
		vec := v.AsFloat()
		return Cbrt(vec)
	}

	return v.Apply(fn)
}

func Ceil(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Ceil(val), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Ceil(float64(val)), na
		}
	default:
		vec := v.AsFloat()
		return Ceil(vec)
	}

	return v.Apply(fn)
}

func Conj(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeComplex:
		fn = func(val complex128, na bool) (complex128, bool) {
			return cmplx.Conj(val), na
		}
	default:
		vec := v.AsComplex()
		return Conj(vec)
	}

	return v.Apply(fn)
}

func CopySign(v vector.Vector, x float64) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Copysign(val, x), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Copysign(float64(val), x), na
		}
	default:
		vec := v.AsFloat()
		return CopySign(vec, x)
	}

	return v.Apply(fn)
}

func Cos(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Cos(val), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Cos(float64(val)), na
		}
	case vector.PayloadTypeComplex:
		fn = func(val complex128, na bool) (complex128, bool) {
			return cmplx.Cos(val), na
		}
	default:
		vec := v.AsFloat()
		return Cos(vec)
	}

	return v.Apply(fn)
}

func Cosh(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Cosh(val), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Cosh(float64(val)), na
		}
	case vector.PayloadTypeComplex:
		fn = func(val complex128, na bool) (complex128, bool) {
			return cmplx.Cosh(val), na
		}
	default:
		vec := v.AsFloat()
		return Cosh(vec)
	}

	return v.Apply(fn)
}

func Cot(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeComplex:
		fn = func(val complex128, na bool) (complex128, bool) {
			return cmplx.Cot(val), na
		}
	default:
		vec := v.AsComplex()
		return Cot(vec)
	}

	return v.Apply(fn)
}

func Dim(v vector.Vector, x float64) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Dim(val, x), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Dim(float64(val), x), na
		}
	default:
		vec := v.AsFloat()
		return Dim(vec, x)
	}

	return v.Apply(fn)
}

func Erf(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Erf(val), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Erf(float64(val)), na
		}
	default:
		vec := v.AsFloat()
		return Erf(vec)
	}

	return v.Apply(fn)
}

func Erfc(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Erfc(val), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Erfc(float64(val)), na
		}
	default:
		vec := v.AsFloat()
		return Erfc(vec)
	}

	return v.Apply(fn)
}

func Erfcinv(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Erfcinv(val), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Erfcinv(float64(val)), na
		}
	default:
		vec := v.AsFloat()
		return Erfcinv(vec)
	}

	return v.Apply(fn)
}

func Erfinv(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Erfinv(val), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Erfinv(float64(val)), na
		}
	default:
		vec := v.AsFloat()
		return Erfinv(vec)
	}

	return v.Apply(fn)
}

func Exp(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Exp(val), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Exp(float64(val)), na
		}
	case vector.PayloadTypeComplex:
		fn = func(val complex128, na bool) (complex128, bool) {
			return cmplx.Exp(val), na
		}
	default:
		vec := v.AsFloat()
		return Exp(vec)
	}

	return v.Apply(fn)
}

func Exp2(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Exp2(val), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Exp2(float64(val)), na
		}
	default:
		vec := v.AsFloat()
		return Exp2(vec)
	}

	return v.Apply(fn)
}

func Exp10(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Pow(10, val), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Pow(10, float64(val)), na
		}
	default:
		vec := v.AsFloat()
		return Exp10(vec)
	}

	return v.Apply(fn)
}

func Floor(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Floor(val), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Floor(float64(val)), na
		}
	default:
		vec := v.AsFloat()
		return Floor(vec)
	}

	return v.Apply(fn)
}

func Gamma(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Gamma(val), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Gamma(float64(val)), na
		}
	default:
		vec := v.AsFloat()
		return Gamma(vec)
	}

	return v.Apply(fn)
}

func Imag(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeComplex:
		fn = func(val complex128, na bool) (float64, bool) {
			return imag(val), na
		}
	default:
		vec := v.AsComplex()
		return Imag(vec)
	}

	return v.Apply(fn)
}

func IsInf(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (bool, bool) {
			return false, na
		}
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (bool, bool) {
			return math.IsInf(val, 0), na
		}
	case vector.PayloadTypeComplex:
		fn = func(val complex128, na bool) (bool, bool) {
			return cmplx.IsInf(val), na
		}
	default:
		vec := v.AsFloat()
		return IsInf(vec)
	}

	return v.Apply(fn)
}

func IsNaN(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (bool, bool) {
			return false, na
		}
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (bool, bool) {
			return math.IsNaN(val), na
		}
	case vector.PayloadTypeComplex:
		fn = func(val complex128, na bool) (bool, bool) {
			return cmplx.IsNaN(val), na
		}
	default:
		vec := v.AsFloat()
		return IsNaN(vec)
	}

	return v.Apply(fn)
}

func J0(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.J0(float64(val)), na
		}
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.J0(val), na
		}
	default:
		vec := v.AsFloat()
		return J0(vec)
	}

	return v.Apply(fn)
}

func J1(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.J1(float64(val)), na
		}
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.J1(val), na
		}
	default:
		vec := v.AsFloat()
		return J1(vec)
	}

	return v.Apply(fn)
}

func Jn(v vector.Vector, n int) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Jn(n, float64(val)), na
		}
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Jn(n, val), na
		}
	default:
		vec := v.AsFloat()
		return Jn(vec, n)
	}

	return v.Apply(fn)
}

func Log(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Log(val), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Log(float64(val)), na
		}
	case vector.PayloadTypeComplex:
		fn = func(val complex128, na bool) (complex128, bool) {
			return cmplx.Log(val), na
		}
	default:
		vec := v.AsFloat()
		return Log(vec)
	}

	return v.Apply(fn)
}

func Log10(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Log10(val), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Log10(float64(val)), na
		}
	case vector.PayloadTypeComplex:
		fn = func(val complex128, na bool) (complex128, bool) {
			return cmplx.Log10(val), na
		}
	default:
		vec := v.AsFloat()
		return Log10(vec)
	}

	return v.Apply(fn)
}

func Log2(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Log2(val), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Log2(float64(val)), na
		}
	default:
		vec := v.AsFloat()
		return Log2(vec)
	}

	return v.Apply(fn)
}

func Phase(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeComplex:
		fn = func(val complex128, na bool) (float64, bool) {
			return cmplx.Phase(val), na
		}
	default:
		vec := v.AsComplex()
		return Phase(vec)
	}

	return v.Apply(fn)
}

func Pow(v vector.Vector, p float64) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Pow(val, p), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Pow(float64(val), p), na
		}
	case vector.PayloadTypeComplex:
		fn = func(val complex128, na bool) (complex128, bool) {
			return cmplx.Pow(val, complex(p, 0)), na
		}
	default:
		vec := v.AsFloat()
		return Pow(vec, p)
	}

	return v.Apply(fn)
}

func Round(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (int, bool) {
			if na {
				return 0, true
			}

			if math.IsInf(val, 0) {
				return 0, true
			}

			return int(math.Round(val)), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (int, bool) {
			if na {
				return 0, true
			}

			return val, na
		}
	default:
		vec := v.AsFloat()
		return Round(vec)
	}

	return v.Apply(fn)
}

func RoundToEven(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (int, bool) {
			if na {
				return 0, true
			}

			if math.IsInf(val, 0) {
				return 0, true
			}

			return int(math.RoundToEven(val)), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (int, bool) {
			if na {
				return 0, true
			}

			return val, na
		}
	default:
		vec := v.AsFloat()
		return RoundToEven(vec)
	}

	return v.Apply(fn)
}

func Signbit(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (bool, bool) {
			return math.Signbit(val), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (bool, bool) {
			return val < 0, na
		}
	default:
		vec := v.AsFloat()
		return Signbit(vec)
	}

	return v.Apply(fn)
}

func Sin(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Sin(val), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Sin(float64(val)), na
		}
	case vector.PayloadTypeComplex:
		fn = func(val complex128, na bool) (complex128, bool) {
			return cmplx.Sin(val), na
		}
	default:
		vec := v.AsFloat()
		return Sin(vec)
	}

	return v.Apply(fn)
}

func Sinh(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Sinh(val), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Sinh(float64(val)), na
		}
	case vector.PayloadTypeComplex:
		fn = func(val complex128, na bool) (complex128, bool) {
			return cmplx.Sinh(val), na
		}
	default:
		vec := v.AsFloat()
		return Sinh(vec)
	}

	return v.Apply(fn)
}

func Sqrt(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			if na {
				return val, true
			}

			if val < 0 {
				return math.NaN(), true
			}

			return math.Sqrt(val), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			if na {
				return 0, true
			}

			if val < 0 {
				return 0, true
			}

			return math.Sqrt(float64(val)), na
		}
	case vector.PayloadTypeComplex:
		fn = func(val complex128, na bool) (complex128, bool) {
			return cmplx.Sqrt(val), na
		}
	default:
		vec := v.AsFloat()
		return Sqrt(vec)
	}

	return v.Apply(fn)
}

func Tan(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Tan(val), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Tan(float64(val)), na
		}
	case vector.PayloadTypeComplex:
		fn = func(val complex128, na bool) (complex128, bool) {
			return cmplx.Tan(val), na
		}
	default:
		vec := v.AsFloat()
		return Tan(vec)
	}

	return v.Apply(fn)
}

func Tanh(v vector.Vector) vector.Vector {
	var fn any

	switch v.Type() {
	case vector.PayloadTypeFloat:
		fn = func(val float64, na bool) (float64, bool) {
			return math.Tanh(val), na
		}
	case vector.PayloadTypeInteger:
		fn = func(val int, na bool) (float64, bool) {
			return math.Tanh(float64(val)), na
		}
	case vector.PayloadTypeComplex:
		fn = func(val complex128, na bool) (complex128, bool) {
			return cmplx.Tanh(val), na
		}
	default:
		vec := v.AsFloat()
		return Tanh(vec)
	}

	return v.Apply(fn)
}
