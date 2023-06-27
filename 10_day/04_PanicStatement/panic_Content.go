package main

import "fmt"

func division(num1, num2 int) {
	if num2 == 0 {
		panic("Cannot divide a number by zero")
	} else {
		var result int = num1 / num2
		fmt.Println(result)
	}
}
func main() {
	//panic:-We use the panic statement to immediately end the execution of the program.
	division(23, 23)
	division(8, 2)
	division(25, 0)
	division(4, 2)
}
