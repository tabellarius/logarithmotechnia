package dataframe

import (
	"logarithmotechnia/vector"
)

func (df *Dataframe) Relocate(arguments ...interface{}) *Dataframe {
	selectors := []interface{}{}
	options := []vector.Option{}

	for _, argument := range arguments {
		switch argument.(type) {
		case vector.Option:
			options = append(options, argument.(vector.Option))
		default:
			selectors = append(selectors, argument)
		}
	}

	curNames, _ := df.Names().Strings()
	columnsToRelocate := []string{}

	for _, selector := range selectors {
		names := []string{}

		switch selector.(type) {
		case string:
			names = append(names, selector.(string))
		case []string:
			names = append(names, selector.([]string)...)
		case []bool:
			boolSelector, _ := vector.BooleanWithNA(selector.([]bool), nil).Adjust(df.Names().Len()).Booleans()
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
	if conf.HasOption(vector.KeyOptionBeforeColumn) {
		pos = strPosInSlice(curNames, conf.Value(vector.KeyOptionBeforeColumn).(string))
		if pos != -1 {
			insertPosition = pos
		}
	}

	if conf.HasOption(vector.KeyOptionAfterColumn) {
		pos = strPosInSlice(curNames, conf.Value(vector.KeyOptionAfterColumn).(string))
		if pos != -1 {
			insertPosition = pos + 1
		}
	}

	selectColumns = append(selectColumns, curNames[:insertPosition]...)
	selectColumns = append(selectColumns, columnsToRelocate...)
	selectColumns = append(selectColumns, curNames[insertPosition:]...)

	return df.Select(selectColumns)
}
