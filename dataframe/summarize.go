package dataframe

import "logarithmotechnia/vector"

func (df *Dataframe) Summarize(columns ...interface{}) *Dataframe {
	if !df.IsGrouped() {
		return df
	}

	newColumns := []Column{}
	for _, column := range columns {
		switch column.(type) {
		case vector.Vector:
			newColumns = append(newColumns, Column{
				name:   column.(vector.Vector).Name(),
				vector: column.(vector.Vector),
			})
		case []vector.Vector:
			for _, columnVec := range column.([]vector.Vector) {
				newColumns = append(newColumns, Column{
					name:   columnVec.Name(),
					vector: columnVec,
				})
			}
		case Column:
			newColumns = append(newColumns, column.(Column))
		case []Column:
			for _, columnCol := range column.([]Column) {
				newColumns = append(newColumns, columnCol)
			}
		}

	}

	for _, group := range df.GroupedBy() {
		vec := df.Cn(group)

		newColumns = append(newColumns, Column{group, vec.ByIndices(vec.GroupFirstElements())})
	}

	return New(newColumns, df.Options()...)
}
