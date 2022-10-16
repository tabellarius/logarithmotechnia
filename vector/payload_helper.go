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
