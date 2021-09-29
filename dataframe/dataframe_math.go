package dataframe

import "logarithmotechnia/vector"

func (df *Dataframe) Sum() *Dataframe {
	newColumns := make([]vector.Vector, df.colNum)
	for i, column := range df.columns {
		newColumns[i] = column.Sum()
	}

	return New(newColumns, df.OptionsWithNames()...)
}
