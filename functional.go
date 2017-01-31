package main

import (
	"errors"
	"fmt"
)

var employees = map[int]interface{} {
	1: map[string]string {
		"Role": "D",
		"Language" : "Go",
	},
	
	2: map[string]string {
		"Role": "T",
		"Framework" : "Selenium",
	},
	
	3: map[string]string {
		"Role": "P",
		"Methodology": "Agile",
	},
}

type EmployeeRole int
const (
	Developer EmployeeRole = iota
	Tester
	ProjectManager
)

var roleMap = map[string]EmployeeRole {
	"D" : Developer,
	"T" : Tester,
	"P" : ProjectManager,
}

type PrintFunction func(map[string]string)

var printFunctionMap = map[EmployeeRole]PrintFunction {
	Developer: PrintDeveloper,
	Tester: PrintTester,
	ProjectManager: PrintProjectManager,
}

func PrintDeveloper(userData map[string]string) {
	fmt.Println("Developer likes to code in " + userData["Language"])
}

func PrintTester(userData map[string]string) {
	fmt.Println("Tester likes to test with " + userData["Framework"])
}

func PrintProjectManager(userData map[string]string) {
	fmt.Println("Project manager like to deliver using " + userData["Methodology"]);
}

type LoadUserFunction func(int) (map[string]string, error)
func LoadUser(id int) (map[string]string, error) {
	if employee, validId := employees[id]; !validId {
		return nil, errors.New("Invalid employee id")
	} else {
		employeeMap := make(map[string]string, 2)
		for k, v := range employee.(map[string]string) {
			employeeMap[k] = v
		}
		
		return employeeMap, nil
	}
}

type PrintUserFunction func (map[string]string) ()
func PrintUser(employee map[string]string) {
	roleAbbreviation := employee["Role"]
	role := roleMap[roleAbbreviation]
	printFunctionMap[role](employee)
}

func LoadAndPrint(employeeId int, loadUser LoadUserFunction, printUser PrintUserFunction) {
	if employee, loadError := loadUser(employeeId); loadError != nil {
		fmt.Println(loadError.Error())
	} else {
		printUser(employee)
	}
}

func main() {
	LoadAndPrint(1, LoadUser, PrintUser)
}