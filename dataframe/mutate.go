package dataframe

import (
	"logarithmotechnia/vector"
)

func (df *Dataframe) Mutate(columns []Column, options ...vector.Option) *Dataframe {
	conf := vector.MergeOptions(options)

	afterColumnIndex := df.colNum

	if conf.HasOption(KeyOptionAfterColumn) {
		pos := df.Names().Find(conf.Value(KeyOptionAfterColumn))
		if pos > 0 {
			afterColumnIndex = pos
		}
	}

	if conf.HasOption(KeyOptionBeforeColumn) {
		pos := df.Names().Find(conf.Value(KeyOptionBeforeColumn))
		if pos > 0 {
			afterColumnIndex = pos - 1
		}
	}

	columnMap := map[string]vector.Vector{}
	for i, column := range df.columns {
		columnMap[df.columnNames[i]] = column
	}

	uniqueNewNames := []string{}
	for _, column := range columns {
		if _, ok := columnMap[column.name]; !ok {
			uniqueNewNames = append(uniqueNewNames, column.name)
		}
		columnMap[column.name] = column.vector
	}

	newNames := []string{}
	newNames = append(newNames, df.columnNames[:afterColumnIndex]...)
	newNames = append(newNames, uniqueNewNames...)
	newNames = append(newNames, df.columnNames[afterColumnIndex:]...)

	newColumns := []vector.Vector{}
	for _, name := range newNames {
		newColumns = append(newColumns, columnMap[name])
	}

	return New(newColumns, OptionColumnNames(newNames))
}
