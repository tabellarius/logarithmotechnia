package dataframe

func (df *Dataframe) Summarize(columns []Column) *Dataframe {
	if !df.IsGrouped() {
		return df
	}

	newColumns := []Column{}
	for _, column := range columns {
		newColumns = append(newColumns, column)
	}

	for _, group := range df.GroupedBy() {
		vec := df.Cn(group)

		newColumns = append(newColumns, Column{group, vec.ByIndices(vec.GroupFirstElements())})
	}

	return New(newColumns, df.Options()...)
}
