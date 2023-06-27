package main

import (
	"fmt"
)

func main() {
	var age int
	var name string

	fmt.Println("Enter your Name ")
	fmt.Scanf("%s", &name)

	fmt.Println("Enter your age ")
	fmt.Scan(&age)

	if age >= 18 {
		fmt.Printf("Yes! %s Your can give the vote", name)
	} else {
		fmt.Printf("Sorry! %s Your are not eligible for the vote", name)
	}
}
