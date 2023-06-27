package main

import "fmt"

//handle the panic attack
func handlePanic() {
	err := recover()

	if err != nil {
		fmt.Println("Error ", err)
	}
}

//panic section in the go lang
func division(num1, num2 int) {

	defer handlePanic()
	if num2 == 0 {
		panic("You can divide the number with zero")
	} else {
		result := num1 / num2
		fmt.Println(result)
	}
}
func main() {
	division(23, 4)
	division(23, 0)
	division(23, 0)
	division(4, 2)
}
