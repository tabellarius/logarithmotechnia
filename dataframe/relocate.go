package dataframe

import (
	"logarithmotechnia/vector"
)

// Relocate re-orders columns.
// There are two types of arguments: column selectors and options.
//
// Possible selectors:
//   - a column name
//   - an array of column names
//   - a column index
//   - an array of column indices
//
// Options:
//   - OptionBeforeColumn(name string)
//   - OptionAfterColumn(name string)
func (df *Dataframe) Relocate(arguments ...any) *Dataframe {
	selectors := []any{}
	options := []vector.Option{}

	for _, argument := range arguments {
		switch arg := argument.(type) {
		case vector.Option:
			options = append(options, arg)
		default:
			selectors = append(selectors, arg)
		}
	}

	curNames, _ := df.Names().Strings()
	columnsToRelocate := []string{}

	for _, selector := range selectors {
		names := []string{}

		switch val := selector.(type) {
		case string:
			names = append(names, val)
		case []string:
			names = append(names, val...)
		case int:
			if df.IsValidColumnIndex(val) {
				names = append(names, df.columnNames[val-1])
			}
		case []int:
			for _, index := range val {
				if df.IsValidColumnIndex(index) {
					names = append(names, df.columnNames[index-1])
				}
			}
		case []bool:
			boolSelector, _ := vector.BooleanWithNA(val, nil).Adjust(df.Names().Len()).Booleans()
			strings, _ := df.Names().Filter(boolSelector).Strings()
			names = append(names, strings...)
		}

		for _, name := range names {
			pos := strPosInSlice(curNames, name)
			if pos != -1 {
				columnsToRelocate = append(columnsToRelocate, name)
				curNames = append(curNames[:pos], curNames[pos+1:]...)
			}
		}
	}

	conf := vector.MergeOptions(options)
	selectColumns := []string{}
	insertPosition := len(curNames)

	var pos int
	if conf.HasOption(KeyOptionBeforeColumn) {
		pos = strPosInSlice(curNames, conf.Value(KeyOptionBeforeColumn).(string))
		if pos != -1 {
			insertPosition = pos
		}
	}

	if conf.HasOption(KeyOptionAfterColumn) {
		pos = strPosInSlice(curNames, conf.Value(KeyOptionAfterColumn).(string))
		if pos != -1 {
			insertPosition = pos + 1
		}
	}

	selectColumns = append(selectColumns, curNames[:insertPosition]...)
	selectColumns = append(selectColumns, columnsToRelocate...)
	selectColumns = append(selectColumns, curNames[insertPosition:]...)

	return df.Select(selectColumns)
}
