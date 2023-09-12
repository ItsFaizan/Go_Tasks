package main

import (
	"fmt"
)

type Student struct {
	rollNo  int
	name    string
	address string
}

type StudentList []Student

func (s Student) Print(no int) {

	fmt.Println("student rollno   ", s.rollNo)
	fmt.Println("student name     ", s.name)
	fmt.Println("student address  ", s.address)
	fmt.Println()
}

func main() {
	// Create a list of students
	students := StudentList{
		{1, "faizan", "Islamabad"},
		{2, "Imran", "Karachi"},
	}

	// Iterate through the list and print each student
	for i, student := range students {
		student.Print(i)
	}
}
