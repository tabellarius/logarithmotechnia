package dataframe

import (
	"fmt"
	"math"
	"testing"
)

func TestFromCSVFile(t *testing.T) {
	df, _ := FromCSVFile("/home/noir/projects/logarithmotechnia/csv/persons.csv",
		CSVOptionSeparator(';'),
		CSVOptionSkipFirstLine(true))

	val, frac := math.Modf(16.0)
	fmt.Println(val, frac)
	fmt.Println(df)
}
