package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

// Function to print the values of a Person struct
func printPersonInfo(p Person) {
	fmt.Println("Name: ", p.name)
	fmt.Println("Age: ", p.age)
	fmt.Println("Job: ", p.job)
	fmt.Println("Salary: ", p.salary)
}

func main() {
	var pers1 Person
	var pers2 Person

	// Pers1 specification
	pers1.name = "Faizan"
	pers1.age = 5
	pers1.job = "Student"
	pers1.salary = 2000

	// Pers2 specification
	pers2.name = "Afraz"
	pers2.age = 24
	pers2.job = "Teacher"
	pers2.salary = 4500

	// Call the function to print Pers1 info
	fmt.Println(" 1:")
	printPersonInfo(pers1)

	// Call the function to print Pers2 info
	fmt.Println("\n2:")
	printPersonInfo(pers2)
}
