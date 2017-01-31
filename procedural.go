package main

import (
	"fmt"
	"errors"
)

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

func LoadDeveloper(id int) (map[string]string, error) {
	return employees[id].(map[string]string), nil
}

func LoadTester(id int) (map[string]string, error) {
	return employees[id].(map[string]string), nil
}

func LoadProjectManager(id int) (map[string]string, error) {
	return employees[id].(map[string]string), nil
}

func LoadUser(id int) (map[string]string, error) {
	if employee, validId := employees[id]; !validId {
		return nil, errors.New("Invalid employee id")
	} else {
		roleAbbreviation := employee.(map[string]string)["Role"]

		switch roleMap[roleAbbreviation] {
		case Developer:
			return LoadDeveloper(id)
			
		case Tester:
			return LoadTester(id)
			
		case ProjectManager:
			return LoadProjectManager(id)
			
		default:
			return nil, errors.New("User loaded but role not catered for");
		}
	}
}

func PrintUser(user map[string]string) {
	switch roleMap[user["Role"]] {
	case Developer:
		fmt.Println("Developer codes in " + user["Language"])
		break;

	case Tester:
		fmt.Println("Tester tests with " + user["Framework"])
		break;

	case ProjectManager:
		fmt.Println("PM delivers using " + user["Methodology"])
		break;

	default:
		fmt.Println("Role not catered for");
		break;
	}
}

func main() {
	if userData, loadError := LoadUser(1); loadError == nil {
		PrintUser(userData)
	}
}