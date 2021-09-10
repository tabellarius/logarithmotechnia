package dataframe

func strPosInSlice(slice []string, str string) int {
	for i, elem := range slice {
		if str == elem {
			return i
		}
	}

	return -1
}
