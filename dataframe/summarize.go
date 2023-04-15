package dataframe

import "logarithmotechnia/vector"

// Summarize allows using aggregation functions on grouped arrays using the columns from this arrays.
// Example:
//
//	groupedDf := df.GroupBy("Category")
//	aggregatedDf := groupedDf.Summarize(groupedDf.Cn("Price").Sum(), groupedDf.Cn("Capacity").Sum())
func (df *Dataframe) Summarize(columns ...any) *Dataframe {
	if !df.IsGrouped() {
		return df
	}

	newColumns := []Column{}
	for _, column := range columns {
		switch c := column.(type) {
		case vector.Vector:
			newColumns = append(newColumns, Column{
				Name:   c.Name(),
				Vector: c,
			})
		case []vector.Vector:
			for _, columnVec := range c {
				newColumns = append(newColumns, Column{
					Name:   columnVec.Name(),
					Vector: columnVec,
				})
			}
		case Column:
			newColumns = append(newColumns, c)
		case []Column:
			for _, columnCol := range c {
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
