package dataframe

import (
	"fmt"
	"logarithmotechnia/vector"
	"testing"
)

func TestDataframe_InnerJoin(t *testing.T) {
	employee := New([]Column{
		{"Name", vector.String([]string{
			"John", "Jane", "Jack", "Robert", "Marcius", "Catullus", "Marcia", "Gera", "Zeus", "Hephaestus", "Hades",
		})},
		{"DepType", vector.StringWithNA([]string{
			"research", "research", "production", "research", "production", "logistics", "production", "sales", "sales", "factory", "",
		}, []bool{false, false, false, false, false, false, false, false, false, false, true})},
		{"Salary", vector.Integer([]int{
			120000, 110000, 80000, 140000, 90000, 100000, 60000, 150000, 225000, 150000, 175000,
		})},
	})

	department := New([]Column{
		{"DepID", vector.Integer([]int{1, 2, 3, 4, 5, 6})},
		{"Title", vector.String([]string{
			"R&D", "Production", "Sales", "Laboratory", "Warehouse", "Unknown",
		})},
		{"DepType", vector.StringWithNA([]string{
			"research", "production", "sales", "research", "wares", "",
		}, []bool{false, false, false, false, false, true})},
	})

	joined := employee.InnerJoin(department, vector.OptionJoinBy("DepType")).Arrange("Name", "Title")
	fmt.Println(joined)

	//	employee.InnerJoin(department, vector.OptionJoinBy("DepartmentType", "DepartmentId"))
}
