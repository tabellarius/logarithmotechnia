package vector

import (
	"golang.org/x/exp/constraints"
	"logarithmotechnia/embed"
)

func genSum[T numeric](data []T, na []bool) (T, bool) {
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

type numeric interface {
	constraints.Integer | constraints.Float | constraints.Complex
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
	na embed.DefNAble,
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

func genProd[T constraints.Integer | constraints.Float | constraints.Complex](data []T, na []bool) (T, bool) {
	var product T
	length := len(data)

	if length == 0 {
		return 0, false
	}

	if length == 1 {
		return data[0], false
	}

	product = data[0]
	for i := 1; i < length; i++ {
		if na[i] {
			return 0, true
		}

		product *= data[i]
	}

	return product, false
}

func genCumSum[T numeric](data []T, na []bool, naDef T) ([]T, []bool) {
	length := len(data)
	if length == 0 {
		return []T{}, []bool{}
	}

	if length == 1 {
		return []T{data[0]}, []bool{na[0]}
	}

	cumSum := make([]T, length)
	cumNA := make([]bool, length)
	copy(cumSum, data)
	copy(cumNA, na)
	isNA := na[0]
	for i := 1; i < length; i++ {
		if isNA {
			cumSum[i] = naDef
			cumNA[i] = true
			continue
		}

		if na[i] {
			cumSum[i] = naDef
			cumNA[i] = true
			isNA = true
			continue
		}

		cumSum[i] = cumSum[i] + cumSum[i-1]
	}

	return cumSum, cumNA
}

func genCumProd[T constraints.Integer | constraints.Float | constraints.Complex](data []T, na []bool, naDef T) ([]T, []bool) {
	length := len(data)
	if length == 0 {
		return []T{}, []bool{}
	}

	if length == 1 {
		return []T{data[0]}, []bool{na[0]}
	}

	cumProd := make([]T, length)
	cumNA := make([]bool, length)
	copy(cumProd, data)
	copy(cumNA, na)
	isNA := na[0]
	for i := 1; i < length; i++ {
		if isNA {
			cumProd[i] = naDef
			cumNA[i] = true
			continue
		}

		if na[i] {
			cumProd[i] = naDef
			cumNA[i] = true
			isNA = true
			continue
		}

		cumProd[i] = cumProd[i] * cumProd[i-1]
	}

	return cumProd, cumNA
}

func genCumMax[T constraints.Ordered](data []T, na []bool, naDef T) ([]T, []bool) {
	length := len(data)
	if length == 0 {
		return []T{}, []bool{}
	}

	if length == 1 {
		return []T{data[0]}, []bool{na[0]}
	}

	cumMax := make([]T, length)
	cumNA := make([]bool, length)
	copy(cumMax, data)
	copy(cumNA, na)
	isNA := na[0]
	for i := 1; i < length; i++ {
		if isNA {
			cumMax[i] = naDef
			cumNA[i] = true
			continue
		}

		if na[i] {
			cumMax[i] = naDef
			cumNA[i] = true
			isNA = true
			continue
		}

		if cumMax[i] < cumMax[i-1] {
			cumMax[i] = cumMax[i-1]
		}
	}

	return cumMax, cumNA
}

func genCumMin[T constraints.Ordered](data []T, na []bool, naDef T) ([]T, []bool) {
	length := len(data)
	if length == 0 {
		return []T{}, []bool{}
	}

	if length == 1 {
		return []T{data[0]}, []bool{na[0]}
	}

	cumMin := make([]T, length)
	cumNA := make([]bool, length)
	copy(cumMin, data)
	copy(cumNA, na)
	isNA := na[0]
	for i := 1; i < length; i++ {
		if isNA {
			cumMin[i] = naDef
			cumNA[i] = true
			continue
		}

		if na[i] {
			cumMin[i] = naDef
			cumNA[i] = true
			isNA = true
			continue
		}

		if cumMin[i] > cumMin[i-1] {
			cumMin[i] = cumMin[i-1]
		}
	}

	return cumMin, cumNA
}
