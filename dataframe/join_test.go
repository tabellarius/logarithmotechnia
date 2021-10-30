package dataframe

import (
	"logarithmotechnia/vector"
	"testing"
)

func TestDataframe_InnerJoin(t *testing.T) {
	employee := New([]Column{
		{"Name", vector.String([]string{"Jim", "John", "Maria", "Dick", "Marcus", "Anna", "Fiona"})},
		{"DepartmentId", vector.Integer([]int{2, 3, 2, 5, 4, 1, 1})},
		{"DepartmentType",
			vector.String([]string{"engineering", "marketing", "engineering", "research", "research", "sales", "sales"})},
	})

	department := New([]Column{
		{"DepartmentId", vector.Integer([]int{1, 2, 3, 4, 5, 6})},
		{"Title", vector.String([]string{"Sales", "Engineering", "Marketing", "R&D", "Laboratory", "Warehouse"})},
		{"DepartmentType", vector.String([]string{"sales", "engineering", "marketing", "research", "research", "wares"})},
	})

	employee.InnerJoin(department, vector.OptionJoinBy("DepartmentType"))
	//	employee.InnerJoin(department, vector.OptionJoinBy("DepartmentType", "DepartmentId"))
}
