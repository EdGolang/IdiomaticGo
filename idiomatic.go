package main

import (
	"fmt"
)

type EmployeeRole int
const (
	Developer EmployeeRole = iota
	Tester
	ProjectManager
)

type IEmployeeRolePossesor interface {
	GetRole() EmployeeRole
}

type IPrintableEmployee interface {
	GetPrintString() interface{}
}

type IEmployee interface {
	IEmployeeRolePossesor
	IPrintableEmployee
}

type EmployeeType struct {
	Role EmployeeRole
}
func (e EmployeeType) GetRole() EmployeeRole {
	return e.Role
}

type DeveloperType struct {
	EmployeeType
	Language string
}
func (d DeveloperType) GetPrintString() interface{} {
	return "Developer likes to develop in " + d.Language
}

type TesterType struct {
	EmployeeType
	Framework string
}
func (t TesterType) GetPrintString() interface{} {
	return "Tester tests with " + t.Framework
}

type ProjectManagerType struct {
	EmployeeType
	Methodology string
}
func (p ProjectManagerType) GetPrintString() interface{} {
	return "PM delivers using " + p.Methodology
}

var employees = map[int]interface{} {
	1: DeveloperType {
		EmployeeType : EmployeeType {
			Role: Developer,
		},
		Language : "Go",
	},
	
	2: TesterType {
		EmployeeType : EmployeeType {
			Role: Tester,
		},
		Framework: "Selenium",
	},
	
	3: ProjectManagerType {
		EmployeeType : EmployeeType {
			Role: ProjectManager,
		},
		Methodology: "Agile",
	},
}

func main() {
	employee := employees[1].(IEmployee)
	fmt.Printf("%s\n", employee.GetPrintString())
}