package apply

import (
	"logarithmotechnia/vector"
	"math"
	"math/cmplx"
	"testing"
)

func TestAbs(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "float",
			in:   vector.FloatWithNA([]float64{1.0, 2.0, math.NaN(), 4.0, 5.0}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Abs(1.0), math.Abs(2.0), math.NaN(), math.Abs(4.0), math.Abs(5.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{1, 2, 0, 4, 5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Abs(1.0), math.Abs(2.0), math.NaN(), math.Abs(4.0), math.Abs(5.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "complex",
			in:   vector.ComplexWithNA([]complex128{1 + 1i, 2 + 2i, cmplx.NaN(), 4 + 4i, 5 + 5i}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{cmplx.Abs(1 + 1i), cmplx.Abs(2 + 2i), math.NaN(), cmplx.Abs(4 + 4i), cmplx.Abs(5 + 5i)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "invalid",
			in:   vector.String([]string{"1", "2", "3", "4", "5"}),
			out:  vector.NA(5),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := Abs(data.in)
			if !vector.CompareVectorsForTest(out, data.out) {
				t.Errorf("Abs(%v) = %v, want %v", data.in, out, data.out)
			}
		})
	}
}

func TestAcos(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "float",
			in:   vector.FloatWithNA([]float64{1.0, 2.0, math.NaN(), 4.0, 5.0}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Acos(1.0), math.Acos(2.0), math.NaN(), math.Acos(4.0), math.Acos(5.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{1, 2, 0, 4, 5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Acos(1.0), math.Acos(2.0), math.NaN(), math.Acos(4.0), math.Acos(5.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "complex",
			in:   vector.ComplexWithNA([]complex128{1 + 1i, 2 + 2i, cmplx.NaN(), 4 + 4i, 5 + 5i}, []bool{false, false, true, false, false}),
			out: vector.ComplexWithNA(
				[]complex128{cmplx.Acos(1 + 1i), cmplx.Acos(2 + 2i), cmplx.NaN(), cmplx.Acos(4 + 4i), cmplx.Acos(5 + 5i)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "invalid",
			in:   vector.String([]string{"1", "2", "3", "4", "5"}),
			out:  vector.NA(5),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := Acos(data.in)
			if !vector.CompareVectorsForTest(out, data.out) {
				t.Errorf("Acos(%v) = %v, want %v", data.in, out, data.out)
			}
		})
	}
}

func TestAcosh(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "float",
			in:   vector.FloatWithNA([]float64{1.0, 2.0, math.NaN(), 4.0, 5.0}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Acosh(1.0), math.Acosh(2.0), math.NaN(), math.Acosh(4.0), math.Acosh(5.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{1, 2, 0, 4, 5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Acosh(1.0), math.Acosh(2.0), math.NaN(), math.Acosh(4.0), math.Acosh(5.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "complex",
			in:   vector.ComplexWithNA([]complex128{1 + 1i, 2 + 2i, cmplx.NaN(), 4 + 4i, 5 + 5i}, []bool{false, false, true, false, false}),
			out: vector.ComplexWithNA(
				[]complex128{cmplx.Acosh(1 + 1i), cmplx.Acosh(2 + 2i), cmplx.NaN(), cmplx.Acosh(4 + 4i), cmplx.Acosh(5 + 5i)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "invalid",
			in:   vector.String([]string{"1", "2", "3", "4", "5"}),
			out:  vector.NA(5),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := Acosh(data.in)
			if !vector.CompareVectorsForTest(out, data.out) {
				t.Errorf("Acosh(%v) = %v, want %v", data.in, out, data.out)
			}
		})
	}
}

func TestAsin(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "float",
			in:   vector.FloatWithNA([]float64{1.0, 2.0, math.NaN(), 4.0, 5.0}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Asin(1.0), math.Asin(2.0), math.NaN(), math.Asin(4.0), math.Asin(5.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{1, 2, 0, 4, 5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Asin(1.0), math.Asin(2.0), math.NaN(), math.Asin(4.0), math.Asin(5.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "complex",
			in:   vector.ComplexWithNA([]complex128{1 + 1i, 2 + 2i, cmplx.NaN(), 4 + 4i, 5 + 5i}, []bool{false, false, true, false, false}),
			out: vector.ComplexWithNA(
				[]complex128{cmplx.Asin(1 + 1i), cmplx.Asin(2 + 2i), cmplx.NaN(), cmplx.Asin(4 + 4i), cmplx.Asin(5 + 5i)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "invalid",
			in:   vector.String([]string{"1", "2", "3", "4", "5"}),
			out:  vector.NA(5),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := Asin(data.in)
			if !vector.CompareVectorsForTest(out, data.out) {
				t.Errorf("Asin(%v) = %v, want %v", data.in, out, data.out)
			}
		})
	}
}

func TestAsinh(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "float",
			in:   vector.FloatWithNA([]float64{1.0, 2.0, math.NaN(), 4.0, 5.0}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Asinh(1.0), math.Asinh(2.0), math.NaN(), math.Asinh(4.0), math.Asinh(5.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{1, 2, 0, 4, 5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Asinh(1.0), math.Asinh(2.0), math.NaN(), math.Asinh(4.0), math.Asinh(5.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "complex",
			in:   vector.ComplexWithNA([]complex128{1 + 1i, 2 + 2i, cmplx.NaN(), 4 + 4i, 5 + 5i}, []bool{false, false, true, false, false}),
			out: vector.ComplexWithNA(
				[]complex128{cmplx.Asinh(1 + 1i), cmplx.Asinh(2 + 2i), cmplx.NaN(), cmplx.Asinh(4 + 4i), cmplx.Asinh(5 + 5i)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "invalid",
			in:   vector.String([]string{"1", "2", "3", "4", "5"}),
			out:  vector.NA(5),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := Asinh(data.in)
			if !vector.CompareVectorsForTest(out, data.out) {
				t.Errorf("Asinh(%v) = %v, want %v", data.in, out, data.out)
			}
		})
	}
}

func TestAtan(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "float",
			in:   vector.FloatWithNA([]float64{1.0, 2.0, math.NaN(), 4.0, 5.0}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Atan(1.0), math.Atan(2.0), math.NaN(), math.Atan(4.0), math.Atan(5.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{1, 2, 0, 4, 5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Atan(1.0), math.Atan(2.0), math.NaN(), math.Atan(4.0), math.Atan(5.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "complex",
			in:   vector.ComplexWithNA([]complex128{1 + 1i, 2 + 2i, cmplx.NaN(), 4 + 4i, 5 + 5i}, []bool{false, false, true, false, false}),
			out: vector.ComplexWithNA(
				[]complex128{cmplx.Atan(1 + 1i), cmplx.Atan(2 + 2i), cmplx.NaN(), cmplx.Atan(4 + 4i), cmplx.Atan(5 + 5i)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "invalid",
			in:   vector.String([]string{"1", "2", "3", "4", "5"}),
			out:  vector.NA(5),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := Atan(data.in)
			if !vector.CompareVectorsForTest(out, data.out) {
				t.Errorf("Atan(%v) = %v, want %v", data.in, out, data.out)
			}
		})
	}
}

func TestAtan2(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "float",
			in:   vector.FloatWithNA([]float64{1.0, 2.0, math.NaN(), 4.0, 5.0}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Atan2(1.0, 0.0), math.Atan2(2.0, 0.0), math.NaN(), math.Atan2(4.0, 0.0), math.Atan2(5.0, 0.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{1, 2, 0, 4, 5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Atan2(1.0, 0.0), math.Atan2(2.0, 0.0), math.NaN(), math.Atan2(4.0, 0.0), math.Atan2(5.0, 0.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "invalid",
			in:   vector.String([]string{"1", "2", "3", "4", "5"}),
			out:  vector.NA(5),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := Atan2(data.in, 0.0)
			if !vector.CompareVectorsForTest(out, data.out) {
				t.Errorf("Atan2(%v) = %v, want %v", data.in, out, data.out)
			}
		})
	}
}

func TestAtanh(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "float",
			in:   vector.FloatWithNA([]float64{1.0, 2.0, math.NaN(), 4.0, 5.0}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Atanh(1.0), math.Atanh(2.0), math.NaN(), math.Atanh(4.0), math.Atanh(5.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{1, 2, 0, 4, 5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Atanh(1.0), math.Atanh(2.0), math.NaN(), math.Atanh(4.0), math.Atanh(5.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "complex",
			in:   vector.ComplexWithNA([]complex128{1 + 1i, 2 + 2i, cmplx.NaN(), 4 + 4i, 5 + 5i}, []bool{false, false, true, false, false}),
			out: vector.ComplexWithNA(
				[]complex128{cmplx.Atanh(1 + 1i), cmplx.Atanh(2 + 2i), cmplx.NaN(), cmplx.Atanh(4 + 4i), cmplx.Atanh(5 + 5i)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "invalid",
			in:   vector.String([]string{"1", "2", "3", "4", "5"}),
			out:  vector.NA(5),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := Atanh(data.in)
			if !vector.CompareVectorsForTest(out, data.out) {
				t.Errorf("Atanh(%v) = %v, want %v", data.in, out, data.out)
			}
		})
	}
}

func TestCbrt(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "float",
			in:   vector.FloatWithNA([]float64{1.0, 2.0, math.NaN(), 4.0, 5.0}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Cbrt(1.0), math.Cbrt(2.0), math.NaN(), math.Cbrt(4.0), math.Cbrt(5.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{1, 2, 0, 4, 5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Cbrt(1.0), math.Cbrt(2.0), math.NaN(), math.Cbrt(4.0), math.Cbrt(5.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "invalid",
			in:   vector.String([]string{"1", "2", "3", "4", "5"}),
			out:  vector.NA(5),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := Cbrt(data.in)
			if !vector.CompareVectorsForTest(out, data.out) {
				t.Errorf("Cbrt(%v) = %v, want %v", data.in, out, data.out)
			}
		})
	}
}

func TestCeil(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "float",
			in:   vector.FloatWithNA([]float64{1.1, 2.2, math.NaN(), 4.4, 5.5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{2.0, 3.0, math.NaN(), 5.0, 6.0},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{1, 2, 0, 4, 5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{1.0, 2.0, math.NaN(), 4.0, 5.0},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "invalid",
			in:   vector.String([]string{"1", "2", "3", "4", "5"}),
			out:  vector.NA(5),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := Ceil(data.in)
			if !vector.CompareVectorsForTest(out, data.out) {
				t.Errorf("Ceil(%v) = %v, want %v", data.in, out, data.out)
			}
		})
	}
}

func TestConj(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "complex",
			in:   vector.ComplexWithNA([]complex128{1 + 1i, 2 + 2i, cmplx.NaN(), 4 + 4i, 5 + 5i}, []bool{false, false, true, false, false}),
			out: vector.ComplexWithNA(
				[]complex128{1 - 1i, 2 - 2i, cmplx.NaN(), 4 - 4i, 5 - 5i},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "invalid",
			in:   vector.String([]string{"1", "2", "3", "4", "5"}),
			out:  vector.NA(5),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := Conj(data.in)
			if !vector.CompareVectorsForTest(out, data.out) {
				t.Errorf("Conj(%v) = %v, want %v", data.in, out, data.out)
			}
		})
	}
}

func TestCopySign(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "float",
			in:   vector.FloatWithNA([]float64{1.1, 2.2, math.NaN(), 4.4, 5.5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{-1.1, -2.2, math.NaN(), -4.4, -5.5},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{1, 2, 0, 4, 5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{-1.0, -2.0, math.NaN(), -4.0, -5.0},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "invalid",
			in:   vector.String([]string{"1", "2", "3", "4", "5"}),
			out:  vector.NA(5),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := CopySign(data.in, -1)
			if !vector.CompareVectorsForTest(out, data.out) {
				t.Errorf("CopySign(%v, -1) = %v, want %v", data.in, out, data.out)
			}
		})
	}
}

func TestCos(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "complex",
			in:   vector.ComplexWithNA([]complex128{1 + 1i, 2 + 2i, cmplx.NaN(), 4 + 4i, 5 + 5i}, []bool{false, false, true, false, false}),
			out: vector.ComplexWithNA(
				[]complex128{cmplx.Cos(1 + 1i), cmplx.Cos(2 + 2i), cmplx.NaN(), cmplx.Cos(4 + 4i), cmplx.Cos(5 + 5i)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "float",
			in:   vector.FloatWithNA([]float64{1.1, 2.2, math.NaN(), 4.4, 5.5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Cos(1.1), math.Cos(2.2), math.NaN(), math.Cos(4.4), math.Cos(5.5)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{1, 2, 0, 4, 5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Cos(1.0), math.Cos(2.0), math.NaN(), math.Cos(4.0), math.Cos(5.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "invalid",
			in:   vector.String([]string{"1", "2", "3", "4", "5"}),
			out:  vector.NA(5),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := Cos(data.in)
			if !vector.CompareVectorsForTest(out, data.out) {
				t.Errorf("Cos(%v) = %v, want %v", data.in, out, data.out)
			}
		})
	}
}

func TestCosh(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "complex",
			in:   vector.ComplexWithNA([]complex128{1 + 1i, 2 + 2i, cmplx.NaN(), 4 + 4i, 5 + 5i}, []bool{false, false, true, false, false}),
			out: vector.ComplexWithNA(
				[]complex128{cmplx.Cosh(1 + 1i), cmplx.Cosh(2 + 2i), cmplx.NaN(), cmplx.Cosh(4 + 4i), cmplx.Cosh(5 + 5i)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "float",
			in:   vector.FloatWithNA([]float64{1.1, 2.2, math.NaN(), 4.4, 5.5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Cosh(1.1), math.Cosh(2.2), math.NaN(), math.Cosh(4.4), math.Cosh(5.5)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{1, 2, 0, 4, 5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Cosh(1.0), math.Cosh(2.0), math.NaN(), math.Cosh(4.0), math.Cosh(5.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "invalid",
			in:   vector.String([]string{"1", "2", "3", "4", "5"}),
			out:  vector.NA(5),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := Cosh(data.in)
			if !vector.CompareVectorsForTest(out, data.out) {
				t.Errorf("Cosh(%v) = %v, want %v", data.in, out, data.out)
			}
		})
	}
}

func TestCot(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "complex",
			in:   vector.ComplexWithNA([]complex128{1 + 1i, 2 + 2i, cmplx.NaN(), 4 + 4i, 5 + 5i}, []bool{false, false, true, false, false}),
			out: vector.ComplexWithNA(
				[]complex128{cmplx.Cot(1 + 1i), cmplx.Cot(2 + 2i), cmplx.NaN(), cmplx.Cot(4 + 4i), cmplx.Cot(5 + 5i)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "invalid",
			in:   vector.String([]string{"1", "2", "3", "4", "5"}),
			out:  vector.NA(5),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := Cot(data.in)
			if !vector.CompareVectorsForTest(out, data.out) {
				t.Errorf("Cot(%v) = %v, want %v", data.in, out, data.out)
			}
		})
	}
}

func TestDim(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "float",
			in:   vector.FloatWithNA([]float64{1.1, 2.2, math.NaN(), 4.4, 5.5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Dim(1.1, 2.0), math.Dim(2.2, 2.0),
					math.NaN(), math.Dim(4.4, 2.0), math.Dim(5.5, 2.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{1, 2, 0, 4, 5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Dim(1.0, 2.0), math.Dim(2.0, 2.0),
					math.NaN(), math.Dim(4.0, 2.0), math.Dim(5.0, 2.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "invalid",
			in:   vector.String([]string{"1", "2", "3", "4", "5"}),
			out:  vector.NA(5),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := Dim(data.in, 2.0)
			if !vector.CompareVectorsForTest(out, data.out) {
				t.Errorf("Dim(%v, 2.0) = %v, want %v", data.in, out, data.out)
			}
		})
	}
}

func TestErf(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "float",
			in:   vector.FloatWithNA([]float64{1.1, 2.2, math.NaN(), 4.4, 5.5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Erf(1.1), math.Erf(2.2), math.NaN(), math.Erf(4.4), math.Erf(5.5)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{1, 2, 0, 4, 5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Erf(1.0), math.Erf(2.0), math.NaN(), math.Erf(4.0), math.Erf(5.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "invalid",
			in:   vector.String([]string{"1", "2", "3", "4", "5"}),
			out:  vector.NA(5),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := Erf(data.in)
			if !vector.CompareVectorsForTest(out, data.out) {
				t.Errorf("Erf(%v) = %v, want %v", data.in, out, data.out)
			}
		})
	}
}

func TestErfc(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "float",
			in:   vector.FloatWithNA([]float64{1.1, 2.2, math.NaN(), 4.4, 5.5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Erfc(1.1), math.Erfc(2.2), math.NaN(), math.Erfc(4.4), math.Erfc(5.5)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{1, 2, 0, 4, 5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Erfc(1.0), math.Erfc(2.0), math.NaN(), math.Erfc(4.0), math.Erfc(5.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "invalid",
			in:   vector.String([]string{"1", "2", "3", "4", "5"}),
			out:  vector.NA(5),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := Erfc(data.in)
			if !vector.CompareVectorsForTest(out, data.out) {
				t.Errorf("Erfc(%v) = %v, want %v", data.in, out, data.out)
			}
		})
	}
}

func TestErfcinv(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "float",
			in:   vector.FloatWithNA([]float64{0.1, 0.2, math.NaN(), 0.4, 0.5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Erfcinv(0.1), math.Erfcinv(0.2), math.NaN(), math.Erfcinv(0.4), math.Erfcinv(0.5)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{1, 2, 0, 4, 5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Erfcinv(1.0), math.Erfcinv(2.0), math.NaN(), math.Erfcinv(4.0), math.Erfcinv(5.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "invalid",
			in:   vector.String([]string{"1", "2", "3", "4", "5"}),
			out:  vector.NA(5),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := Erfcinv(data.in)
			if !vector.CompareVectorsForTest(out, data.out) {
				t.Errorf("Erfcinv(%v) = %v, want %v", data.in, out, data.out)
			}
		})
	}
}

func TestErfinv(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "float",
			in:   vector.FloatWithNA([]float64{0.1, 0.2, math.NaN(), 0.4, 0.5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Erfinv(0.1), math.Erfinv(0.2), math.NaN(), math.Erfinv(0.4), math.Erfinv(0.5)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{1, 2, 0, 4, 5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Erfinv(1.0), math.Erfinv(2.0), math.NaN(), math.Erfinv(4.0), math.Erfinv(5.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "invalid",
			in:   vector.String([]string{"1", "2", "3", "4", "5"}),
			out:  vector.NA(5),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := Erfinv(data.in)
			if !vector.CompareVectorsForTest(out, data.out) {
				t.Errorf("Erfinv(%v) = %v, want %v", data.in, out, data.out)
			}
		})
	}
}

func TestExp(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "complex",
			in:   vector.ComplexWithNA([]complex128{1 + 1i, 2 + 2i, cmplx.NaN(), 4 + 4i, 5 + 5i}, []bool{false, false, true, false, false}),
			out: vector.ComplexWithNA(
				[]complex128{cmplx.Exp(1 + 1i), cmplx.Exp(2 + 2i), cmplx.NaN(), cmplx.Exp(4 + 4i), cmplx.Exp(5 + 5i)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "float",
			in:   vector.FloatWithNA([]float64{1.1, 2.2, math.NaN(), 4.4, 5.5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Exp(1.1), math.Exp(2.2), math.NaN(), math.Exp(4.4), math.Exp(5.5)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{1, 2, 0, 4, 5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Exp(1.0), math.Exp(2.0), math.NaN(), math.Exp(4.0), math.Exp(5.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "invalid",
			in:   vector.String([]string{"1", "2", "3", "4", "5"}),
			out:  vector.NA(5),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := Exp(data.in)
			if !vector.CompareVectorsForTest(out, data.out) {
				t.Errorf("Exp(%v) = %v, want %v", data.in, out, data.out)
			}
		})
	}
}

func TestExp2(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "float",
			in:   vector.FloatWithNA([]float64{1.1, 2.2, math.NaN(), 4.4, 5.5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Exp2(1.1), math.Exp2(2.2), math.NaN(), math.Exp2(4.4), math.Exp2(5.5)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{1, 2, 0, 4, 5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Exp2(1.0), math.Exp2(2.0), math.NaN(), math.Exp2(4.0), math.Exp2(5.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "invalid",
			in:   vector.String([]string{"1", "2", "3", "4", "5"}),
			out:  vector.NA(5),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := Exp2(data.in)
			if !vector.CompareVectorsForTest(out, data.out) {
				t.Errorf("Exp2(%v) = %v, want %v", data.in, out, data.out)
			}
		})
	}
}

func TestExp10(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "float",
			in:   vector.FloatWithNA([]float64{1.1, 2.2, math.NaN(), 4.4, 5.5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Pow(10, 1.1), math.Pow(10, 2.2), math.NaN(), math.Pow(10, 4.4), math.Pow(10, 5.5)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{1, 2, 0, 4, 5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Pow(10, 1.0), math.Pow(10, 2.0), math.NaN(), math.Pow(10, 4.0), math.Pow(10, 5.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "invalid",
			in:   vector.String([]string{"1", "2", "3", "4", "5"}),
			out:  vector.NA(5),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := Exp10(data.in)
			if !vector.CompareVectorsForTest(out, data.out) {
				t.Errorf("Exp10(%v) = %v, want %v", data.in, out, data.out)
			}
		})
	}
}

func TestFloor(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "float",
			in:   vector.FloatWithNA([]float64{1.1, 2.2, math.NaN(), 4.4, 5.5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{1.0, 2.0, math.NaN(), 4.0, 5.0},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{1, 2, 0, 4, 5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{1.0, 2.0, math.NaN(), 4.0, 5.0},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "invalid",
			in:   vector.String([]string{"1", "2", "3", "4", "5"}),
			out:  vector.NA(5),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := Floor(data.in)
			if !vector.CompareVectorsForTest(out, data.out) {
				t.Errorf("Floor(%v) = %v, want %v", data.in, out, data.out)
			}
		})
	}
}

func TestGamma(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "float",
			in:   vector.FloatWithNA([]float64{1.1, 2.2, math.NaN(), 4.4, 5.5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Gamma(1.1), math.Gamma(2.2), math.NaN(), math.Gamma(4.4), math.Gamma(5.5)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{1, 2, 0, 4, 5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Gamma(1.0), math.Gamma(2.0), math.NaN(), math.Gamma(4.0), math.Gamma(5.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "invalid",
			in:   vector.String([]string{"1", "2", "3", "4", "5"}),
			out:  vector.NA(5),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := Gamma(data.in)
			if !vector.CompareVectorsForTest(out, data.out) {
				t.Errorf("Gamma(%v) = %v, want %v", data.in, out, data.out)
			}
		})
	}
}

func TestImag(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "complex",
			in:   vector.ComplexWithNA([]complex128{1 + 1i, 2 + 2i, cmplx.NaN(), 4 + 4i, 5 + 5i}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{imag(1 + 1i), imag(2 + 2i), math.NaN(), imag(4 + 4i), imag(5 + 5i)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "invalid",
			in:   vector.String([]string{"1", "2", "3", "4", "5"}),
			out:  vector.NA(5),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := Imag(data.in)
			if !vector.CompareVectorsForTest(out, data.out) {
				t.Errorf("Imag(%v) = %v, want %v", data.in, out, data.out)
			}
		})
	}
}

func TestLog(t *testing.T) {
	testData := []struct {
		name string
		in   vector.Vector
		out  vector.Vector
	}{
		{
			name: "complex",
			in:   vector.ComplexWithNA([]complex128{1 + 1i, 2 + 2i, cmplx.NaN(), 4 + 4i, 5 + 5i}, []bool{false, false, true, false, false}),
			out: vector.ComplexWithNA(
				[]complex128{cmplx.Log(1 + 1i), cmplx.Log(2 + 2i), cmplx.NaN(), cmplx.Log(4 + 4i), cmplx.Log(5 + 5i)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "float",
			in:   vector.FloatWithNA([]float64{1.1, 2.2, math.NaN(), 4.4, 5.5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Log(1.1), math.Log(2.2), math.NaN(), math.Log(4.4), math.Log(5.5)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "int",
			in:   vector.IntegerWithNA([]int{1, 2, 0, 4, 5}, []bool{false, false, true, false, false}),
			out: vector.FloatWithNA(
				[]float64{math.Log(1.0), math.Log(2.0), math.NaN(), math.Log(4.0), math.Log(5.0)},
				[]bool{false, false, true, false, false},
			),
		},
		{
			name: "invalid",
			in:   vector.String([]string{"1", "2", "3", "4", "5"}),
			out:  vector.NA(5),
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			out := Log(data.in)
			if !vector.CompareVectorsForTest(out, data.out) {
				t.Errorf("Log(%v) = %v, want %v", data.in, out, data.out)
			}
		})
	}
}
