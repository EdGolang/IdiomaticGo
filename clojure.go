package main

import (
	"errors"
	"fmt"
)

type UserType map[string]string

var employees = map[int]interface{} {
	1: UserType {
		"Role": "D",
		"Language" : "Go",
	},
	
	2: UserType {
		"Role": "T",
		"Framework" : "Selenium",
	},
	
	3: UserType {
		"Role": "P",
		"Methodology": "Agile",
	},
}

type UserLoadResult struct {
	User UserType
	Error error
}

type UserLoadResultCreator func () (UserLoadResult)

func ComputePrintString(functionToCall UserLoadResultCreator) UserLoadResult {
	result := functionToCall()
	if (result.Error != nil) {
		fmt.Println(result.Error.Error())
	}

	switch (result.User["Role"]) {
	case "D":
		result.User["PrintString"] = "Developer likes to code in " + result.User["Language"]
		break;

	case "T":
		result.User["PrintString"] = "Tester likes to test with " + result.User["Framework"]
		break;

	case "P":
		result.User["PrintString"] = "Project manager like to delivery using " + result.User["Methodology"]
		break;

	default:
		break;
	}
	
	return result		
}

func LoadUser(id int) UserLoadResult {
	employee, found := employees[id]
	
	mapToUserLoadResult := func () UserLoadResult {
		var invalidIdError error
		if (!found) {
			invalidIdError = errors.New("No user with the given Id could be found")
		}
		
		return UserLoadResult {
			User: employee.(UserType),
			Error: invalidIdError,
		}
	}

	return ComputePrintString(mapToUserLoadResult)
}

func main() {
	outputChannel := make(chan UserLoadResult)

	go func() {
		outputChannel <- LoadUser(1)
	}()

	func(loadUserResult UserLoadResult) {
		fmt.Println(loadUserResult.User["PrintString"])		
	}(<- outputChannel)
}