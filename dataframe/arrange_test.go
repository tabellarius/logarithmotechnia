package dataframe

import (
	"fmt"
	"logarithmotechnia/vector"
	"testing"
)

func TestDataframe_Arrange(t *testing.T) {
	df := New([]Column{
		{"activity", vector.Boolean([]bool{true, false, true, false, true, false, true, false, true,
			false, true, false, true, false}, nil)},
		{"type", vector.Integer([]int{1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3}, nil)},
		{"salary", vector.Integer([]int{100, 300, 50, 50, 350, 45, 120, 225, 60, 30, 220, 70, 180, 35, 110}, nil)},
	})
	newDf := df.Arrange("type", "salary", vector.OptionArrangeReverse(false))

	fmt.Println(df.columns)
	fmt.Println(newDf.columns)
	fmt.Println(newDf.columnNames)
}
