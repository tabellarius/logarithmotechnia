package dataframe

import (
	"fmt"
	"logarithmotechnia/vector"
	"testing"
)

func TestDataframe_IsGrouped(t *testing.T) {
	df := New([]Column{
		{"A", vector.Integer([]int{100, 200, 200, 30, 30})},
		{"B", vector.IntegerWithNA([]int{100, 100, 40, 30, 40}, []bool{false, true, true, true, false})},
		{"C", vector.Boolean([]bool{true, false, true, false, true})},
		{"D", vector.String([]string{"1", "2", "3", "4", "5"})},
	})

	testData := []struct {
		name      string
		df        *Dataframe
		isGrouped bool
	}{
		{
			name:      "non-grouped",
			df:        df,
			isGrouped: false,
		},
		{
			name:      "grouped",
			df:        df.GroupBy("A"),
			isGrouped: true,
		},
		{
			name:      "grouped by two",
			df:        df.GroupBy("A", "B"),
			isGrouped: true,
		},
		{
			name:      "grouped by incorrect",
			df:        df.GroupBy("F"),
			isGrouped: false,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			isGrouped := data.df.IsGrouped()
			if isGrouped != data.isGrouped {
				t.Error(fmt.Sprintf("isGrouped (%v) is not equal to expected (%v)",
					isGrouped, data.isGrouped))
			}
		})
	}
}

func TestDataframe_GroupedBy(t *testing.T) {
	df := New([]Column{
		{"A", vector.Integer([]int{100, 200, 200, 30, 30})},
		{"B", vector.IntegerWithNA([]int{100, 100, 40, 30, 40}, []bool{false, true, true, true, false})},
		{"C", vector.Boolean([]bool{true, false, true, false, true})},
		{"D", vector.String([]string{"1", "2", "3", "4", "5"})},
	})

	testData := []struct {
		name      string
		df        *Dataframe
		groupedBy []string
	}{
		{
			name:      "non-grouped",
			df:        df,
			groupedBy: []string{},
		},
		{
			name:      "grouped",
			df:        df.GroupBy("A"),
			groupedBy: []string{"A"},
		},
		{
			name:      "grouped by two",
			df:        df.GroupBy("A", "B"),
			groupedBy: []string{"A", "B"},
		},
		{
			name:      "grouped by incorrect",
			df:        df.GroupBy("F"),
			groupedBy: []string{},
		},
		{
			name:      "grouped by mixed",
			df:        df.GroupBy("C", "F", "A"),
			groupedBy: []string{"C", "A"},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			groupedBy := data.df.GroupedBy()
			for i, group := range groupedBy {
				if group != data.df.groupedBy[i] {
					t.Error(fmt.Sprintf("groupedBy (%v) is not equal to expected (%v)",
						groupedBy, data.groupedBy))
				}
			}
		})
	}
}

func TestDataframe_GroupBy(t *testing.T) {
	df := New([]Column{
		{"A", vector.Integer([]int{100, 200, 200, 30, 30})},
		{"B", vector.IntegerWithNA([]int{100, 100, 40, 30, 40}, []bool{false, true, true, true, false})},
		{"C", vector.Boolean([]bool{true, false, true, false, true})},
		{"D", vector.String([]string{"1", "2", "3", "4", "5"})},
	})

	testData := []struct {
		name      string
		df        *Dataframe
		groupedBy []string
		isGrouped bool
	}{
		{
			name:      "non-grouped",
			df:        df,
			groupedBy: []string{},
			isGrouped: false,
		},
		{
			name:      "grouped",
			df:        df.GroupBy("A"),
			groupedBy: []string{"A"},
			isGrouped: true,
		},
		{
			name:      "grouped by two",
			df:        df.GroupBy("A", "B"),
			groupedBy: []string{"A", "B"},
			isGrouped: true,
		},
		{
			name:      "grouped by array",
			df:        df.GroupBy([]string{"A", "B"}),
			groupedBy: []string{"A", "B"},
			isGrouped: true,
		},
		{
			name:      "grouped by incorrect",
			df:        df.GroupBy("F"),
			groupedBy: []string{},
			isGrouped: false,
		},
		{
			name:      "grouped by mixed",
			df:        df.GroupBy("C", "F", "A"),
			groupedBy: []string{"C", "A"},
			isGrouped: true,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {

			groupedBy := data.df.GroupedBy()
			for i, group := range groupedBy {
				if group != data.df.groupedBy[i] {
					t.Error(fmt.Sprintf("groupedBy (%v) is not equal to expected (%v)",
						groupedBy, data.groupedBy))
				}
			}

			isGrouped := data.df.IsGrouped()
			if isGrouped != data.isGrouped {
				t.Error(fmt.Sprintf("isGrouped (%v) is not equal to expected (%v)",
					isGrouped, data.isGrouped))
			}
		})
	}
}
