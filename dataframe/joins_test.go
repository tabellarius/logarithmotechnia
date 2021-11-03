package dataframe

import (
	"fmt"
	"logarithmotechnia/vector"
	"reflect"
	"testing"
)

func getJoinDataFrames() (*Dataframe, *Dataframe) {
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
		{"Group", vector.String([]string{
			"A", "A", "B", "B", "A", "A", "B", "A", "A", "B", "A",
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
		{"Group", vector.String([]string{
			"A", "B", "A", "B", "B", "A",
		})},
	})

	return employee, department
}

func TestDataframe_InnerJoin(t *testing.T) {
	employee, department := getJoinDataFrames()

	testData := []struct {
		name        string
		joined      *Dataframe
		columnNames []string
		outColumns  []vector.Vector
	}{
		{
			name:        "employee ✕ department",
			joined:      employee.InnerJoin(department, vector.OptionJoinBy("DepType")).Arrange("Name", "Title"),
			columnNames: []string{"Name", "DepType", "Salary", "Group", "DepID", "Title", "Group_1"},
			outColumns: []vector.Vector{
				vector.String([]string{
					"Gera", "Hades", "Jack", "Jane", "Jane", "John", "John", "Marcia", "Marcius", "Robert", "Robert", "Zeus",
				}),
				vector.StringWithNA([]string{
					"sales", "", "production", "research", "research", "research", "research", "production", "production",
					"research", "research", "sales",
				}, []bool{false, true, false, false, false, false, false, false, false, false, false, false}),
				vector.Integer([]int{
					150000, 175000, 80000, 110000, 110000, 120000, 120000, 60000, 90000, 140000, 140000, 225000,
				}),
				vector.String([]string{
					"A", "A", "B", "A", "A", "A", "A", "B", "A", "B", "B", "A",
				}),
				vector.Integer([]int{
					3, 6, 2, 4, 1, 4, 1, 2, 2, 4, 1, 3,
				}),
				vector.String([]string{
					"Sales", "Unknown", "Production", "Laboratory", "R&D", "Laboratory", "R&D", "Production", "Production",
					"Laboratory", "R&D", "Sales",
				}),
				vector.String([]string{
					"A", "A", "B", "B", "A", "B", "A", "B", "B", "B", "A", "A",
				}),
			},
		},
		{
			name:        "department ✕ employee",
			joined:      department.InnerJoin(employee, vector.OptionJoinBy("DepType")).Arrange("Title", "Name"),
			columnNames: []string{"DepID", "Title", "DepType", "Group", "Name", "Salary", "Group_1"},
			outColumns: []vector.Vector{
				vector.Integer([]int{
					4, 4, 4, 2, 2, 2, 1, 1, 1, 3, 3, 6,
				}),
				vector.String([]string{
					"Laboratory", "Laboratory", "Laboratory", "Production", "Production", "Production", "R&D", "R&D",
					"R&D", "Sales", "Sales", "Unknown",
				}),
				vector.StringWithNA([]string{
					"research", "research", "research", "production", "production", "production", "research", "research",
					"research", "sales", "sales", "",
				}, []bool{false, false, false, false, false, false, false, false, false, false, false, true}),
				vector.String([]string{
					"B", "B", "B", "B", "B", "B", "A", "A", "A", "A", "A", "A",
				}),
				vector.String([]string{
					"Jane", "John", "Robert", "Jack", "Marcia", "Marcius", "Jane", "John", "Robert", "Gera", "Zeus", "Hades",
				}),
				vector.Integer([]int{
					110000, 120000, 140000, 80000, 60000, 90000, 110000, 120000, 140000, 150000, 225000, 175000,
				}),
				vector.String([]string{
					"A", "A", "B", "B", "B", "A", "A", "A", "B", "A", "A", "A",
				}),
			},
		},
		{
			name:        "employee ✕ department by group",
			joined:      employee.InnerJoin(department, vector.OptionJoinBy("Group", "DepType")).Arrange("Name", "Title"),
			columnNames: []string{"Name", "DepType", "Salary", "Group", "DepID", "Title"},
			outColumns: []vector.Vector{
				vector.String([]string{
					"Gera", "Hades", "Jack", "Jane", "John", "Marcia", "Robert", "Zeus",
				}),
				vector.StringWithNA([]string{
					"sales", "", "production", "research", "research", "production", "research", "sales",
				}, []bool{false, true, false, false, false, false, false, false}),
				vector.Integer([]int{
					150000, 175000, 80000, 110000, 120000, 60000, 140000, 225000,
				}),
				vector.String([]string{
					"A", "A", "B", "A", "A", "B", "B", "A",
				}),
				vector.Integer([]int{
					3, 6, 2, 1, 1, 2, 4, 3,
				}),
				vector.String([]string{
					"Sales", "Unknown", "Production", "R&D", "R&D", "Production", "Laboratory", "Sales",
				}),
			},
		},
		{
			name:        "department ✕ employee by group",
			joined:      department.InnerJoin(employee, vector.OptionJoinBy("Group", "DepType")).Arrange("Title", "Name"),
			columnNames: []string{"DepID", "Title", "DepType", "Group", "Name", "Salary"},
			outColumns: []vector.Vector{
				vector.Integer([]int{
					4, 2, 2, 1, 1, 3, 3, 6,
				}),
				vector.String([]string{
					"Laboratory", "Production", "Production", "R&D", "R&D", "Sales", "Sales", "Unknown",
				}),
				vector.StringWithNA([]string{
					"research", "production", "production", "research", "research", "sales", "sales", "",
				}, []bool{false, false, false, false, false, false, false, true}),
				vector.String([]string{
					"B", "B", "B", "A", "A", "A", "A", "A",
				}),
				vector.String([]string{
					"Robert", "Jack", "Marcia", "Jane", "John", "Gera", "Zeus", "Hades",
				}),
				vector.Integer([]int{
					140000, 80000, 60000, 110000, 120000, 150000, 225000, 175000,
				}),
			},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if !reflect.DeepEqual(data.joined.columnNames, data.columnNames) {
				t.Error(fmt.Sprintf("Column namess (%v) are not equal to expected (%v)\n",
					data.joined.columnNames, data.columnNames))
			}

			if !vector.CompareVectorArrs(data.joined.columns, data.outColumns) {
				t.Error(fmt.Sprintf("Columns (%v) are not equal to expected (%v)\n",
					data.joined.columns, data.outColumns))
			}
		})
	}
}

func TestDataframe_LeftJoin(t *testing.T) {
	employee, department := getJoinDataFrames()

	testData := []struct {
		name        string
		joined      *Dataframe
		columnNames []string
		outColumns  []vector.Vector
	}{
		{
			name:        "employee ✕ department",
			joined:      employee.LeftJoin(department, vector.OptionJoinBy("DepType")).Arrange("Name", "Title"),
			columnNames: []string{"Name", "DepType", "Salary", "Group", "DepID", "Title", "Group_1"},
			outColumns: []vector.Vector{
				vector.String([]string{
					"Catullus", "Gera", "Hades", "Hephaestus", "Jack", "Jane", "Jane", "John", "John", "Marcia",
					"Marcius", "Robert", "Robert", "Zeus",
				}),
				vector.StringWithNA([]string{
					"logistics", "sales", "", "factory", "production", "research", "research", "research", "research",
					"production", "production", "research", "research", "sales",
				}, []bool{false, false, true, false, false, false, false, false, false, false, false, false, false, false}),
				vector.Integer([]int{
					100000, 150000, 175000, 150000, 80000, 110000, 110000, 120000, 120000, 60000, 90000, 140000, 140000,
					225000,
				}),
				vector.String([]string{
					"A", "A", "A", "B", "B", "A", "A", "A", "A", "B", "A", "B", "B", "A",
				}),
				vector.IntegerWithNA([]int{
					0, 3, 6, 0, 2, 4, 1, 4, 1, 2, 2, 4, 1, 3,
				}, []bool{true, false, false, true, false, false, false, false, false, false, false, false, false, false}),
				vector.StringWithNA([]string{
					"", "Sales", "Unknown", "", "Production", "Laboratory", "R&D", "Laboratory", "R&D", "Production",
					"Production", "Laboratory", "R&D", "Sales",
				}, []bool{true, false, false, true, false, false, false, false, false, false, false, false, false, false}),
				vector.StringWithNA([]string{
					"", "A", "A", "", "B", "B", "A", "B", "A", "B", "B", "B", "A", "A",
				}, []bool{true, false, false, true, false, false, false, false, false, false, false, false, false, false}),
			},
		},
		{
			name:        "department ✕ employee",
			joined:      department.LeftJoin(employee, vector.OptionJoinBy("DepType")).Arrange("Title", "Name"),
			columnNames: []string{},
			outColumns: []vector.Vector{
				vector.Integer([]int{}),
				vector.String([]string{}),
				vector.String([]string{}),
				vector.Integer([]int{}),
				vector.String([]string{}),
				vector.Integer([]int{}),
				vector.String([]string{}),
			},
		},
		{
			name:        "employee ✕ department by group",
			joined:      employee.LeftJoin(department, vector.OptionJoinBy("DepType")).Arrange("Name", "Title"),
			columnNames: []string{},
			outColumns: []vector.Vector{
				vector.String([]string{}),
				vector.String([]string{}),
				vector.Integer([]int{}),
				vector.String([]string{}),
				vector.Integer([]int{}),
				vector.String([]string{}),
			},
		},
		{
			name:        "department ✕ employee by group",
			joined:      employee.LeftJoin(department, vector.OptionJoinBy("DepType")).Arrange("Title", "Name"),
			columnNames: []string{},
			outColumns: []vector.Vector{
				vector.Integer([]int{}),
				vector.String([]string{}),
				vector.String([]string{}),
				vector.String([]string{}),
				vector.String([]string{}),
				vector.Integer([]int{}),
			},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if !reflect.DeepEqual(data.joined.columnNames, data.columnNames) {
				t.Error(fmt.Sprintf("Column namess (%v) are not equal to expected (%v)\n",
					data.joined.columnNames, data.columnNames))
			}

			if !vector.CompareVectorArrs(data.joined.columns, data.outColumns) {
				t.Error(fmt.Sprintf("Columns (%v) are not equal to expected (%v)\n",
					data.joined.columns, data.outColumns))
			}
		})
	}
}
