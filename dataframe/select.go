package dataframe

type FromToColNames struct {
	from string
	to   string
}

type FromToColIndices struct {
	from int
	to   int
}

func (df *Dataframe) Select(selectors ...interface{}) *Dataframe {
	colNames := make([]string, df.colNum)
	copy(colNames, df.config.columnNames)

	for _, selector := range selectors {
		switch selector.(type) {
		case string:
			colNames = df.selectByName(colNames, selector.(string))
		case []string:
			colNames = df.selectByNames(colNames, selector.([]string))
		case int:
			colNames = df.selectByIndex(colNames, selector.(int))
		case []int:
			colNames = df.selectByIndices(colNames, selector.([]int))
		case FromToColNames:
			colNames = df.selectByFromToColNames(colNames, selector.(FromToColNames))
		case FromToColIndices:
		}
	}

	return df
}

func (df *Dataframe) selectByName(colNames []string, name string) []string {
	remove := false

	if name[0] == '-' {
		remove = true
		name = name[1:]
	}

	if remove {
		pos := strPosInSlice(colNames, name)
		if pos != -1 {
			return append(colNames[:pos], colNames[:pos+1]...)
		}
	} else {
		if strPosInSlice(colNames, name) == -1 && df.HasColumn(name) {
			return append(colNames, name)
		}
	}

	return colNames
}

func (df *Dataframe) selectByNames(colNames []string, names []string) []string {
	for _, name := range names {
		colNames = df.selectByName(colNames, name)
	}

	return colNames
}

func (df *Dataframe) selectByIndex(colNames []string, index int) []string {
	if index >= 1 && index <= df.colNum {
		colNames = df.selectByName(colNames, df.config.columnNames[index-1])
	}

	return colNames
}

func (df *Dataframe) selectByIndices(colNames []string, indices []int) []string {
	for _, index := range indices {
		colNames = df.selectByIndex(colNames, index)
	}

	return colNames
}

func (df *Dataframe) selectByFromToColNames(colNames []string, fromTo FromToColNames) []string {
	fromIndex := -1
	toIndex := -1

	fromIndex = strPosInSlice(df.config.columnNames, fromTo.from)
	if fromIndex == -1 {
		return colNames
	}

	toIndex = strPosInSlice(df.config.columnNames, fromTo.to)
	if toIndex == -1 {
		return colNames
	}

	inc := 1
	if toIndex < fromIndex {
		inc = -1
	}

	for i := fromIndex; i <= toIndex; i = i + inc {
		colNames = df.selectByName(colNames, df.config.columnNames[i-1])
	}

	return colNames
}

func (df *Dataframe) selectByFromToIndices(colNames []string, fromTo FromToColIndices) []string {
	if !df.IsValidColumnIndex(fromTo.from) || !df.IsValidColumnIndex(fromTo.to) {
		return colNames
	}

	inc := 1
	if fromTo.to < fromTo.from {
		inc = -1
	}

	for i := fromTo.from; i <= fromTo.to; i = i + inc {
		colNames = df.selectByName(colNames, df.config.columnNames[i-1])
	}

	return colNames
}

func strPosInSlice(slice []string, str string) int {
	for i, elem := range slice {
		if str == elem {
			return i
		}
	}

	return -1
}
