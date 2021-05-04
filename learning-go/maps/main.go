package main

import "fmt"

func main() {
	var employees = map[int]string{
		1001: "John",
		1002: "Steve",
		1003: "Maria",
		1004: "Ahmet",
		1005: "Faruk",
		1006: "Furkan",
	}

	printEmployee(employees, 1001)
	printEmployee(employees, 1010)
	deleteEmployee(employees, 1006)
	printEmployeeGet(employees)
	if isEmployeeExists(employees, 1002) {
		fmt.Println("EmployeeId 1002 found")
	}
}

func printEmployee(employees map[int]string, employeeId int) {
	if name, ok := employees[employeeId]; ok {
		fmt.Printf("name = %s, ok = %v\n", name, ok)
	} else {
		fmt.Printf("EmployeeId %d not found\n", employeeId)
	}
}

func printEmployeeGet(employees map[int]string) {

	fmt.Println(employees)

}

func isEmployeeExists(employees map[int]string, employeeId int) bool {
	_, ok := employees[employeeId]
	return ok
}
func deleteEmployee(employees map[int]string, employeeId int) {

	delete(employees, employeeId)
	fmt.Printf("\nid: %d", employeeId)

}
