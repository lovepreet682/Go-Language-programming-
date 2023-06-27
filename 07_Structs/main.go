package main

import "fmt"

//crate the Structs in golang
type User struct {
	Name   string
	Email  string
	Age    int
	City   string
	Salary float32
}

func main() {
	fmt.Println("Here Your Details ")

	//create the variable
	love := User{
		"Lovepreet",
		"lovepreet@gmail.com",
		23,
		"Jalandhar",
		3200.0,
	}

	//create second user details
	abhishek := User{
		"abhishek",
		"abhishek@gmail.com",
		24,
		"Lucknow",
		3200.0,
	}
	fmt.Println(love)

	//print specific data
	fmt.Println("Name of User ", abhishek.Name, " And Email is  ", abhishek.Email)
}
