package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operation string

	fmt.Println("Welcome to Our Calculator")

	fmt.Println("Enter number ")
	fmt.Scan(&num1)

	fmt.Println("Enter any operator which you want to perform (+ || - || * || / )")
	fmt.Scan(&operation)

	fmt.Println("Enter number ")
	fmt.Scan(&num2)

	var result float64

	switch operation {
	case "+":
		result = num1 + num2

	case "-":
		result = num1 - num2

	case "*":
		result = num1 * num2

	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Error: Division by zero is not allowed......")
			return
		}

	default:
		{
			fmt.Println("Please enter the correct data")
		}
	}
	fmt.Println("result is Here ", result)

}
