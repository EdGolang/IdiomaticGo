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

type IEmployeRolePossesor interface {
	GetRole() EmployeeRole
}

type IPrintableEmployee interface {	
	Print()
}

type Employee struct {
	Role EmployeeRole
}
func (e Employee) GetRole() EmployeeRole {
	return e.Role
}
func (e Employee) Print() {
	fmt.Println("Not a valid role type")
}

type DeveloperType struct {
	Employee
	Language string
}
func (d DeveloperType) Print() {
	fmt.Println("Developer codes in " + d.Language)
}

type TesterType struct {
	Employee
	Framework string
}
func (t TesterType) Print() {
	fmt.Println("Tester tests with " + t.Framework)
}

type ProjectManagerType struct {
	Employee
	Methodology string
}
func (p ProjectManagerType) Print() {
	fmt.Println("PM delivers using " + p.Methodology)
}

var employees = map[int]interface{} {
	1: DeveloperType {
		Employee : Employee{
			Role: Developer,
		},
		Language : "Go",
	},
	
	2: TesterType {
		Employee : Employee{
			Role: Tester,
		},
		Framework: "Selenium",
	},
	
	3: ProjectManagerType {
		Employee : Employee{
			Role: ProjectManager,
		},
		Methodology: "Agile",
	},
}

func NewEmployee (employeeData interface{}) IPrintableEmployee {
	employee := employeeData.(IEmployeRolePossesor)
	
	switch (employee.GetRole()) {
	case Developer:
		developer := &DeveloperType{}
		*developer = employeeData.(DeveloperType)
		return developer
		
	case Tester:
		return employeeData.(TesterType)

	case ProjectManager:
		return employeeData.(ProjectManagerType)

	default:
		return Employee{}
	}
}

func LoadEmployee (employeeId int) IPrintableEmployee {
	employeeData := employees[employeeId]
	return NewEmployee(employeeData)
}

func main() {
	employee := LoadEmployee(1)
	employee.Print()
}