package dataframe

import "logarithmotechnia/vector"

type Rename struct {
	Column  string
	NewName string
}

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

	columnsToRelocate := []string{}
	columnsToRename := []Rename{}

	for _, selector := range selectors {
		names := []string{}

		switch selector.(type) {
		case string:
			names = append(names, selector.(string))
		case []string:
			names = append(names, selector.([]string)...)
		case []bool:
			boolSelector, _ := vector.Boolean(selector.([]bool), nil).Adjust(df.Names().Len()).Booleans()
			strings, _ := df.Names().Filter(boolSelector).Strings()
			names = append(names, strings...)
		case Rename:
			names = append(names, selector.(Rename).Column)
			columnsToRename = append(columnsToRename, selector.(Rename))
		}

		for _, name := range names {
			if strPosInSlice(columnsToRelocate, name) == 0 {
				columnsToRelocate = append(columnsToRelocate)
			}
		}
	}

	return nil
}
