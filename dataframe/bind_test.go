package dataframe

import (
	"fmt"
	"logarithmotechnia/vector"
	"testing"
)

func TestDataframe_BindColumns(t *testing.T) {
	df1 := New([]Column{
		{"A", vector.IntegerWithNA([]int{10, 20, 30}, []bool{false, true, true})},
		{"B", vector.Boolean([]bool{true, false, true})},
		{"C", vector.String([]string{"1", "2", "3"})},
	})

	df2 := New([]Column{
		{"E", vector.Boolean([]bool{false, false, true})},
		{"C", vector.String([]string{"one", "two", "three"})},
		{"A", vector.IntegerWithNA([]int{1, 2, 3}, []bool{false, false, true})},
	})

	df := df1.BindRows(df2)
	fmt.Println(df.columns)

}
