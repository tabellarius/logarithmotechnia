package dataframe

import (
	"fmt"
	"logarithmotechnia/vector"
	"testing"
)

func TestDataframe_Summarize(t *testing.T) {
	df := New([]Column{
		{"A", vector.Integer([]int{100, 200, 200, 30, 30, 120, 140, 70})},
		{"B", vector.IntegerWithNA([]int{100, 100, 40, 30, 40, 80, 90, 110},
			[]bool{false, true, true, true, false, false, false, false})},
		{"C", vector.Boolean([]bool{true, false, false, false, true, false, true, false})},
		{"D", vector.String([]string{"A", "B", "C", "A", "B", "D", "D", "D"})},
	})

	groupedByD := df.GroupBy("D")

	testData := []struct {
		name        string
		groupedDf   *Dataframe
		summarizers []interface{}
		vecs        []vector.Vector
		columnNames []string
	}{
		{
			name:      "vectors",
			groupedDf: groupedByD,
			summarizers: []interface{}{
				groupedByD.Cn("A").Sum(),
				groupedByD.Cn("B").Sum(),
			},
			vecs: []vector.Vector{
				vector.Integer([]int{130, 230, 200, 330}),
				vector.IntegerWithNA([]int{0, 0, 0, 280}, []bool{true, true, true, false}),
				vector.String([]string{"A", "B", "C", "D"}),
			},
			columnNames: []string{"A", "B", "D"},
		},
	}

	_ = testData

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			sumDf := data.groupedDf.Summarize(data.summarizers...)

			if !vector.CompareVectorArrs(sumDf.columns, data.vecs) {
				t.Error(fmt.Sprintf("sumDf columns (%v) are not equal to expected (%v)",
					sumDf.columns, data.vecs))
			}

		})
	}

	df = df.GroupBy("D")

	summedDf := df.Summarize(
		df.Cn("A").Sum(),
		df.Cn("B").Sum(),
	)

	_ = summedDf

	fmt.Println(summedDf)
}
