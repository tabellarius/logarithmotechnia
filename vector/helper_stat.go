package vector

import "golang.org/x/exp/constraints"

func genSum[T constraints.Integer | constraints.Float | constraints.Complex](data []T, na []bool) (T, bool) {
	var sum T

	for i, val := range data {
		if na[i] {
			return 0, true
		}
		sum += val
	}

	return sum, false
}

func genMin[T constraints.Ordered](data []T, na []bool) (T, bool) {
	var min T
	if len(data) == 0 {
		return min, true
	}

	min = data[0]
	for i := 1; i < len(data); i++ {
		if na[i] {
			return min, true
		}

		if data[i] < min {
			min = data[i]
		}
	}

	return min, false
}

func genMax[T constraints.Ordered](data []T, na []bool) (T, bool) {
	var max T
	if len(data) == 0 {
		return max, true
	}

	max = data[0]
	for i := 1; i < len(data); i++ {
		if na[i] {
			return max, true
		}

		if data[i] > max {
			max = data[i]
		}
	}

	return max, false
}
