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
		nIndices := []int{}
		vec := df.Cn(group)
		vIndices, _ := vec.Groups()

		for _, gIndices := range vIndices {
			nIndices = append(nIndices, gIndices[0])
		}

		newColumns = append(newColumns, Column{group, vec.ByIndices(nIndices)})
	}

	return New(newColumns, df.Options()...)
}
