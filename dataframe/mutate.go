package dataframe

import (
	"logarithmotechnia/vector"
)

// Mutate transforms a dataframe by adding new columns or changing new ones.
// This function accepts Column, []Column, vector.Vector, []vector.Vector, Option and []Option.
// Vectors must have a name.
// Possible options are:
//   - OptionAfterColumn("name")
//   - OptionBeforeColumn("name")
func (df *Dataframe) Mutate(arguments ...any) *Dataframe {
	columns := []Column{}
	options := []Option{}

	for _, arg := range arguments {
		switch val := arg.(type) {
		case Column:
			columns = append(columns, val)
		case []Column:
			columns = append(columns, val...)
		case vector.Vector:
			if val.Name() != "" {
				columns = append(columns, Column{val.Name(), val})
			}
		case []vector.Vector:
			for _, v := range val {
				if v.Name() != "" {
					columns = append(columns, Column{v.Name(), v})
				}
			}
		case Option:
			options = append(options, val)
		case []Option:
			options = append(options, val...)
		}
	}

	conf := MergeOptions(options)

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
		if _, ok := columnMap[column.Name]; !ok {
			uniqueNewNames = append(uniqueNewNames, column.Name)
		}
		columnMap[column.Name] = column.Vector
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
