package main

import (
	"fmt"
)

type employee struct {
	name     string
	salary   int
	position string
}

type company struct {
	companyName string
	employees   []employee
}

func main() {
	// Add three employees using the employee struct
	emp1 := employee{"Imran", 80000, "Tecaher"}
	emp2 := employee{"Sana", 75000, " Officer"}
	emp3 := employee{"Sara", 90000, " Scientist"}

	// Create an array of employees
	employees := []employee{emp1, emp2, emp3}

	// Create a company struct and add values to it
	companyDetails := company{"Systems", employees}

	// Print company details
	fmt.Println("Company Name:", companyDetails.companyName)
	fmt.Println("Employees:")

	for _, emp := range companyDetails.employees {
		fmt.Println("Name:", emp.name)
		fmt.Println("Salary:", emp.salary)
		fmt.Println("Position:", emp.position)
		fmt.Println()
	}
}
