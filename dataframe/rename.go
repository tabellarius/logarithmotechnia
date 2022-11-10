package dataframe

type Rename struct {
	from string
	to   string
}

func (df *Dataframe) Rename(renames ...Rename) *Dataframe {
	renamesMap := map[string]string{}
	for _, rename := range renames {
		renamesMap[rename.from] = rename.to
	}

	names := make([]string, df.colNum)
	for i, columnName := range df.columnNames {
		if newName, ok := renamesMap[columnName]; ok {
			names[i] = newName
		} else {
			names[i] = columnName
		}
	}
	return New(df.columns, OptionColumnNames(names))
}
