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
		switch rename.(type) {
		case Rename:
			renamesMap[rename.(Rename).from] = rename.(Rename).to
		case []Rename:
			arrRename := rename.([]Rename)
			for i := 0; i < len(arrRename); i++ {
				renamesMap[arrRename[i].from] = arrRename[i].to
			}
		case []string:
			strRename := rename.([]string)
			if len(strRename) == 2 {
				renamesMap[strRename[0]] = strRename[1]
			}
		case [][]string:
			strArrRename := rename.([][]string)
			for i := 0; i < len(strArrRename); i++ {
				if len(strArrRename[i]) == 2 {
					renamesMap[strArrRename[i][0]] = strArrRename[i][1]
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
