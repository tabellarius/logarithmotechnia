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

type calculable interface {
	constraints.Integer | constraints.Float
}

func genMean[T calculable](data []T, na []bool) (float64, bool) {
	var sum float64
	length := len(data)

	if length == 0 {
		return 0, false
	}

	for i := 0; i < length; i++ {
		if na[i] {
			return 0, true
		}

		sum += float64(data[i])
	}

	return sum / float64(length), false
}

func genMedian[T constraints.Integer | constraints.Float](
	data []T,
	na DefNAble,
	sorter func() []int,
) (T, bool) {
	var median T
	length := len(data)

	if length == 0 || na.HasNA() {
		return 0, true
	}

	if length == 1 {
		return data[0], false
	}

	sortedIndices := sorter()
	if length%2 == 0 {
		median = (data[sortedIndices[length/2-1]] + data[sortedIndices[length/2]]) / 2
	} else {
		median = data[sortedIndices[length/2]]
	}

	return median, false
}
