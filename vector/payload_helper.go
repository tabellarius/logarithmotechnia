package vector

func pickWithNA[T any](idx int, data []T, na []bool, maxLen int) interface{} {
	if idx < 1 || idx > maxLen {
		return nil
	}

	if na[idx-1] {
		return nil
	}

	return interface{}(data[idx-1])
}

func pick[T any](idx int, data []T, maxLen int) interface{} {
	if idx < 1 || idx > maxLen {
		return nil
	}

	return interface{}(data[idx-1])
}
