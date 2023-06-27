package main

import "fmt"

func main() {
	subjectMarks := map[string]int{
		"Java":   80,
		"PHP":    92,
		"Golang": 88,
	}
	fmt.Println(subjectMarks)

	//access specific data
	fmt.Println(subjectMarks["Java"])

	//delete the value
	delete(subjectMarks, "PHP")
	fmt.Println(subjectMarks)
}
