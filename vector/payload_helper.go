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

func adjustToLesserSize[T any](srcData []T, srcNA []bool, size int) ([]T, []bool) {
	data := make([]T, size)
	na := make([]bool, size)

	copy(data, srcData)
	copy(na, srcNA)

	return data, na
}

func adjustToBiggerSize[T any](srcData []T, srcNA []bool, length int, size int) ([]T, []bool) {
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

func applyToByFunc[T any](whicher []bool, inData []T, inNA []bool, length int,
	applyFunc func(int, T, bool) (T, bool), naDef T) ([]T, []bool) {
	data := make([]T, length)
	na := make([]bool, length)

	for i := 0; i < length; i++ {
		if !whicher[i] {
			continue
		}

		dataVal, naVal := applyFunc(i+1, inData[i], inNA[i])
		if naVal {
			dataVal = naDef
		}
		data[i] = dataVal
		na[i] = naVal
	}

	return data, na
}

func applyToByCompactFunc[T any](whicher []bool, inData []T, inNA []bool, length int,
	applyFunc func(T, bool) (T, bool), naDef T) ([]T, []bool) {
	data := make([]T, length)
	na := make([]bool, length)

	for i := 0; i < length; i++ {
		if !whicher[i] {
			continue
		}

		dataVal, naVal := applyFunc(inData[i], inNA[i])
		if naVal {
			dataVal = naDef
		}
		data[i] = dataVal
		na[i] = naVal
	}

	return data, na
}
