package util

import (
	"math"
	"math/cmplx"
)

func EqualFloatArrays(arr1, arr2 []float64) bool {
	if (arr1 == nil) != (arr2 == nil) {
		return false
	}

	length := len(arr1)
	if length != len(arr2) {
		return false
	}

	for i := 0; i < length; i++ {
		if math.IsNaN(arr1[i]) {
			if !math.IsNaN(arr2[i]) {
				return false
			}
		} else {
			if arr1[i] != arr2[i] {
				return false
			}
		}
	}

	return true
}

func EqualComplexArrays(arr1, arr2 []complex128) bool {
	if (arr1 == nil) != (arr2 == nil) {
		return false
	}

	length := len(arr1)
	if length != len(arr2) {
		return false
	}

	for i := 0; i < length; i++ {
		if cmplx.IsNaN(arr1[i]) {
			if !cmplx.IsNaN(arr2[i]) {
				return false
			}
		} else {
			if arr1[i] != arr2[i] {
				return false
			}
		}
	}

	return true
}
