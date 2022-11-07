package vector

import "golang.org/x/exp/constraints"

func genSum[T constraints.Integer | constraints.Float | constraints.Complex](data []T, na []bool) (T, bool) {
	var sum T
	isNA := false
	for i, val := range data {
		if na[i] {
			sum = 0
			isNA = true
			break
		}

		sum += val
	}

	return sum, isNA
}
