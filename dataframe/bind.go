package dataframe

import (
	"logarithmotechnia/vector"
)

// BindColumns binds other dataframes to the dataframe as new columns. If number of rows differs the resulting
// dataframe will have the length of the longest of the all dataframes and missing values will be filled by NA.
func (df *Dataframe) BindColumns(arguments ...any) *Dataframe {
	dataframes := []*Dataframe{}
	for _, arg := range arguments {
		switch arg.(type) {
		case *Dataframe:
			dataframes = append(dataframes, arg.(*Dataframe))
		case []*Dataframe:
			dataframes = append(dataframes, arg.([]*Dataframe)...)
		}
	}

	startDf := df
	for _, dataframe := range dataframes {
		startDf = bindTwoDataframesByColumns(startDf, dataframe)
	}

	return startDf
}

func bindTwoDataframesByColumns(src, app *Dataframe) *Dataframe {
	columns := make([]vector.Vector, src.colNum+app.colNum)
	columnNames := make([]string, src.colNum+app.colNum)

	idx := 0
	for i := 0; i < src.colNum; i++ {
		columns[idx] = src.columns[i]
		columnNames[idx] = src.columnNames[i]
		idx++
	}

	for i := 0; i < app.colNum; i++ {
		columns[idx] = app.columns[i]
		columnNames[idx] = app.columnNames[i]
		idx++
	}

	options := src.Options()
	options = append(options, OptionColumnNames(columnNames))

	return New(columns, options...)
}

// BindRows adds rows of other dataframes to the dataframe. If a column does not present in the dataframe, it will be
// added into resulting one. Missing values will be filled with NA.
func (df *Dataframe) BindRows(arguments ...any) *Dataframe {
	dataframes := []*Dataframe{}
	for _, arg := range arguments {
		switch arg.(type) {
		case *Dataframe:
			dataframes = append(dataframes, arg.(*Dataframe))
		case []*Dataframe:
			dataframes = append(dataframes, arg.([]*Dataframe)...)
		}
	}

	startDf := df
	for _, dataframe := range dataframes {
		startDf = bindTwoDataFramesByRows(startDf, dataframe)
	}

	return startDf
}

func bindTwoDataFramesByRows(src, app *Dataframe) *Dataframe {
	columns := make([]vector.Vector, src.colNum)
	columnNames := make([]string, src.colNum)

	for i := 0; i < src.colNum; i++ {
		var tmpColumn vector.Vector
		if app.Names().Has(src.columnNames[i]) {
			tmpColumn = app.Select(src.columnNames[i]).Ci(1)
		} else {
			tmpColumn = vector.NA(app.rowNum)
		}
		columns[i] = src.columns[i].Append(tmpColumn)
		columnNames[i] = src.columnNames[i]
	}

	options := src.Options()
	options = append(options, OptionColumnNames(columnNames))

	return New(columns, options...)
}
