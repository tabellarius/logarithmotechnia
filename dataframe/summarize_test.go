package dataframe

import (
	"fmt"
	"logarithmotechnia/vector"
	"testing"
)

func TestDataframe_Summarize(t *testing.T) {
	df := New([]Column{
		{"A", vector.Integer([]int{100, 200, 200, 30, 30})},
		{"B", vector.IntegerWithNA([]int{100, 100, 40, 30, 40}, []bool{false, true, true, true, false})},
		{"C", vector.Boolean([]bool{true, false, false, false, true})},
		{"D", vector.String([]string{"A", "B", "C", "A", "B"})},
	})

	df = df.GroupBy("C", "D")

	summedDf := df.Summarize([]Column{
		{"A", df.Cn("A").Sum()},
		{"B", df.Cn("B").Sum()},
	})

	fmt.Println(summedDf)
}
