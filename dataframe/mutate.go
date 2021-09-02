package dataframe

import (
	"logarithmotechnia/vector"
)

func (df *Dataframe) Mutate(columns map[string]vector.Vector, options ...vector.Option) *Dataframe {
	/*
		conf := vector.MergeOptions(options)

		beforeColumnIndex := -1
		afterColumnIndex := df.colNum-1

		if conf.HasOption(vector.KeyOptionAfterColumn) {
			afterColumn := conf.Value(vector.KeyOptionAfterColumn)
		}
	*/
	return nil
}

func (df *Dataframe) Transmute(map[string]vector.Vector) *Dataframe {
	panic("implement me")
}
