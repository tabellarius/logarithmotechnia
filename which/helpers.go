package which

func trueBooleanArr(size int) []bool {
	booleans := make([]bool, size)

	for i := 0; i < size; i++ {
		booleans[i] = true
	}

	return booleans
}

func invertBooleanArr(arr []bool) []bool {
	booleans := make([]bool, len(arr))

	for i, val := range arr {
		booleans[i] = !val
	}

	return booleans
}
