package main

import "fmt"

//create a method
func greetings() {
	fmt.Println("Welcome to the function concept")
}

//adding two number using function
func addNum(a, b int) int {
	return a + b
}

//print Loop 1 to 10 using function
func loopSection() {
	for i := 0; i <= 4; i++ {
		fmt.Println(i)
	}
}

//print msg to by using name
func printName(name string) {
	fmt.Println("name ", name)
}

func main() {
	fmt.Println("Good Morning")
	greetings()

	//print the add Function value
	fmt.Println("Adding two number ", addNum(3, 5))

	//print the loop
	loopSection()

	//print the msg
	name := "Love Rehan"
	fmt.Println(name)

}
