package main

import "fmt"

func main() {
	//take the input from the user
	var num1, num2 int
	fmt.Print("Enter the First Number:- ")
	fmt.Scanln(&num1)

	fmt.Print("Enter the Second number:- ")
	fmt.Scanln(&num2)

	fmt.Println("The sum of two number is ", num1+num2)
}
