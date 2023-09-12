package main

import (
	"crypto/sha256"
	"fmt"
)

type Student struct {
	rollNo   int
	name     string
	address  string
	subjects []string // Array of subjects
}

type StudentList []Student

func (s Student) Print(no int) {

	fmt.Println(" rollno   ", s.rollNo)
	fmt.Println(" name     ", s.name)
	fmt.Println(" address  ", s.address)
	fmt.Println(" subjects ", s.subjects)
	fmt.Println("Hash:            ", calculateHash(s))
	fmt.Println()
}

func calculateHash(s Student) string {
	data := fmt.Sprintf("%d%s%s%s%v", s.rollNo, s.name, s.address, s.subjects)
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

func main() {
	// Create a list of students with subjects
	students := StudentList{
		{
			rollNo:   1,
			name:     "Faizan",
			address:  "Islamabad",
			subjects: []string{"phy", "cal", "pk study"},
		},
		{
			rollNo:   2,
			name:     "imran",
			address:  "lahore",
			subjects: []string{"urdu", "sceine", "discrete"},
		},
	}

	// Iterate through the list and print each student
	for i, student := range students {
		student.Print(i)
	}
}
