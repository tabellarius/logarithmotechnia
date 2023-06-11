package apply

import (
	"logarithmotechnia/vector"
	"math"
	"math/cmplx"
)

// Abs returns the absolute value of each element in the vector.
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

// Acos returns the arccosine of each element in the vector.
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

// Acosh returns the hyperbolic arccosine of each element in the vector.
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

// Asin returns the arcsine of each element in the vector.
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

// Asinh returns the hyperbolic arcsine of each element in the vector.
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

// Atan returns the arctangent of each element in the vector.
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

// Atanh returns the hyperbolic arctangent of each element in the vector.
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

// Atan2 returns the arctangent of each element in the vector.
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

// Cbrt returns the cube root of each element in the vector.
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

// Ceil returns the ceiling of each element in the vector.
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

// Conj returns the complex conjugate of each element in the vector.
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

// CopySign returns a vector with the magnitude of v and the sign of x.
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

// Cos returns the cosine of each element in the vector.
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

// Cosh returns the hyperbolic cosine of each element in the vector.
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

// Cot returns the cotangent of each element in the vector.
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

// Dim returns a vector with the maximum of x-y or 0
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

// Erf returns the error function of each element in the vector.
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

// Erfc returns the complementary error function of each element in the vector.
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

// Erfcinv returns the inverse of the complementary error function of each element in the vector.
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

// Erfinv returns the inverse of the error function of each element in the vector.
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

// Exp returns the exponential of each element in the vector.
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

// Exp2 returns the base-2 exponential of each element in the vector.
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

// Exp2 returns the base-10 exponential of each element in the vector.
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

// Floor returns the greatest integer value less than or equal to each element in the vector.
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

// Gamma returns the gamma function of each element in the vector.
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

// Imag returns the imaginary part of each element in the vector.
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

// IsInf returns whether each element in the vector is positive or negative infinity in boolean vector format.
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

// IsNaN returns whether each element in the vector is NaN in boolean vector format.
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

// J0 returns the Bessel function of the first kind of order 0 of each element in the vector.
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

// J1 returns the Bessel function of the first kind of order 1 of each element in the vector.
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

// Jn returns the Bessel function of the first kind of order n of each element in the vector.
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

// Log returns the natural logarithm of each element in the vector.
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

// Log10 returns the base-10 logarithm of each element in the vector.
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

// Log2 returns the base-2 logarithm of each element in the vector.
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

// Phase returns the phase (also known as the argument) of each element in the vector.
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

// Pow returns the element-wise power of the vector.
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

// Round returns the nearest integer, rounding half away from zero.
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

// RoundToEven returns the nearest integer, rounding ties to even.
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

// Signbit returns whether the sign bit is set for each element in the vector.
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

// Sin returns the sine of each element in the vector.
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

// Sinh returns the hyperbolic sine of each element in the vector.
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

// Sqrt returns the square root of each element in the vector.
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

// Tan returns the tangent of each element in the vector.
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

// Tanh returns the hyperbolic tangent of each element in the vector.
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
