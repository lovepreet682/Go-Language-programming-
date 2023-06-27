package main

import (
	"fmt"
)

func main() {

	//first way to declered the variable
	var num int = 10

	//Second way to declered the variable
	var num2 = 29

	//Third way to declered the variable
	num3 := 23

	fmt.Println(num)
	fmt.Println(num2)
	fmt.Println(num3)

	//Constants are the fixed values that cannot be changed once declared
	const name = "Lovepreet"
	fmt.Println(name)

	//Creating Multiple Variables at Once
	var number, city = "Chandigarh", 100
	fmt.Println(number, city)

	//use of printf
	city1 := "Noida"
	fmt.Printf("my number is %d and the city is %s", num, city1)
}
