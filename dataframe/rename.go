package dataframe

type Rename struct {
	from string
	to   string
}

// Rename allows renameing of the dataframe columns.
// Possible parameters are:
//   - Rename struct
//   - []string slice of two elements (from, two)
//   - array of []string slice of two elements (from, two)
func (df *Dataframe) Rename(renames ...any) *Dataframe {
	renamesMap := map[string]string{}
	for _, rename := range renames {
		switch r := rename.(type) {
		case Rename:
			renamesMap[r.from] = r.to
		case []Rename:
			for i := 0; i < len(r); i++ {
				renamesMap[r[i].from] = r[i].to
			}
		case []string:
			if len(r) == 2 {
				renamesMap[r[0]] = r[1]
			}
		case [][]string:
			for i := 0; i < len(r); i++ {
				if len(r[i]) == 2 {
					renamesMap[r[i][0]] = r[i][1]
				}
			}
		}
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
