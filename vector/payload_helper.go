package vector

func pickValueWithNA[T any](idx int, data []T, na []bool, maxLen int) interface{} {
	if idx < 1 || idx > maxLen {
		return nil
	}

	if na[idx-1] {
		return nil
	}

	return interface{}(data[idx-1])
}

func pickValue[T any](idx int, data []T, maxLen int) interface{} {
	if idx < 1 || idx > maxLen {
		return nil
	}

	return interface{}(data[idx-1])
}

func dataWithNAToInterfaceArray[T any](data []T, na []bool) []interface{} {
	dataLen := len(data)
	outData := make([]interface{}, dataLen)

	for idx, val := range data {
		if na[idx] {
			outData[idx] = nil
		} else {
			outData[idx] = val
		}
	}

	return outData
}

func byIndices[T any](indices []int, srcData []T, srcNA []bool, naDef T) ([]T, []bool) {
	data := make([]T, 0, len(indices))
	na := make([]bool, 0, len(indices))

	for _, idx := range indices {
		if idx == 0 {
			data = append(data, naDef)
			na = append(na, true)
		} else {
			data = append(data, srcData[idx-1])
			na = append(na, srcNA[idx-1])
		}
	}

	return data, na
}

func adjustToLesserSizeWithNA[T any](srcData []T, srcNA []bool, size int) ([]T, []bool) {
	data := make([]T, size)
	na := make([]bool, size)

	copy(data, srcData)
	copy(na, srcNA)

	return data, na
}

func adjustToBiggerSizeWithNA[T any](srcData []T, srcNA []bool, length int, size int) ([]T, []bool) {
	cycles := size / length
	if size%length > 0 {
		cycles++
	}

	data := make([]T, cycles*length)
	na := make([]bool, cycles*length)

	for i := 0; i < cycles; i++ {
		copy(data[i*length:], srcData)
		copy(na[i*length:], srcNA)
	}

	data = data[:size]
	na = na[:size]

	return data, na
}

func adjustToLesserSize[T any](srcData []T, size int) []T {
	data := make([]T, size)

	copy(data, srcData)

	return data
}

func adjustToBiggerSize[T any](srcData []T, size int) []T {
	length := len(srcData)
	cycles := size / length
	if size%length > 0 {
		cycles++
	}

	data := make([]T, cycles*length)

	for i := 0; i < cycles; i++ {
		copy(data[i*length:], srcData)
	}

	data = data[:size]

	return data
}

func applyByFunc[T any](inData []T, inNA []bool, length int,
	applyFunc func(int, T, bool) (T, bool), naDef T) ([]T, []bool) {
	data := make([]T, length)
	na := make([]bool, length)

	for i := 0; i < length; i++ {
		dataVal, naVal := applyFunc(i+1, inData[i], inNA[i])
		if naVal {
			dataVal = naDef
		}
		data[i] = dataVal
		na[i] = naVal
	}

	return data, na
}

func applyByCompactFunc[T any](inData []T, inNA []bool, length int,
	applyFunc func(T, bool) (T, bool), naDef T) ([]T, []bool) {
	data := make([]T, length)
	na := make([]bool, length)

	for i := 0; i < length; i++ {
		dataVal, naVal := applyFunc(inData[i], inNA[i])
		if naVal {
			dataVal = naDef
		}
		data[i] = dataVal
		na[i] = naVal
	}

	return data, na
}

func applyToByFunc[T any](indices []int, inData []T, inNA []bool,
	applyFunc func(int, T, bool) (T, bool), naDef T) ([]T, []bool) {
	length := len(inData)

	data := make([]T, length)
	na := make([]bool, length)

	copy(data, inData)
	copy(na, inNA)

	for _, idx := range indices {
		idx = idx - 1
		dataVal, naVal := applyFunc(idx+1, inData[idx], inNA[idx])
		if naVal {
			dataVal = naDef
		}
		data[idx] = dataVal
		na[idx] = naVal
	}

	return data, na
}

func applyToByCompactFunc[T any](indices []int, inData []T, inNA []bool,
	applyFunc func(T, bool) (T, bool), naDef T) ([]T, []bool) {
	length := len(inData)

	data := make([]T, length)
	na := make([]bool, length)

	copy(data, inData)
	copy(na, inNA)

	for _, idx := range indices {
		idx = idx - 1
		dataVal, naVal := applyFunc(inData[idx], inNA[idx])
		if naVal {
			dataVal = naDef
		}
		data[idx] = dataVal
		na[idx] = naVal
	}

	return data, na
}

func groupsForData[T comparable](srcData []T, srcNA []bool) ([][]int, []interface{}) {
	groupMap := map[T][]int{}
	ordered := []T{}
	na := []int{}

	for i, val := range srcData {
		idx := i + 1

		if srcNA[i] {
			na = append(na, idx)
			continue
		}

		if _, ok := groupMap[val]; !ok {
			groupMap[val] = []int{}
			ordered = append(ordered, val)
		}

		groupMap[val] = append(groupMap[val], idx)
	}

	groups := make([][]int, len(ordered))
	for i, val := range ordered {
		groups[i] = groupMap[val]
	}

	if len(na) > 0 {
		groups = append(groups, na)
	}

	values := make([]interface{}, len(groups))
	for i, val := range ordered {
		values[i] = interface{}(val)
	}
	if len(na) > 0 {
		values[len(values)-1] = nil
	}

	return groups, values
}
