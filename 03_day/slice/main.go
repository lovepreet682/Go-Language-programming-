package main

import "fmt"

func main() {
	fmt.Println("Slice Section :- ")
	student_BCA := []string{"Lovepreet", "Nikhl", "Ankit", "Bobby"}
	fmt.Println(student_BCA)
	fmt.Println()

	student_MCA := []string{"Gagan", "Sahil", "Kuljinder", "Chetan", "Gourav"}
	fmt.Println(student_MCA)
	fmt.Println()

	//append use to add the two and more variable in the single unit
	totalStudent := append(student_BCA, student_MCA...)
	fmt.Print(totalStudent)
}
