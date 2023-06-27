package main

import (
	"fmt"
)

func divide(num1, num2 int) error {
	if num2 == 0 {
		return fmt.Errorf("%d / %d Connot divide the num ", num1, num2)
	} else {
		return nil
	}
}
func main() {
	var num1, num2 int
	fmt.Println("Enter the number one ")
	fmt.Scan(&num1)

	fmt.Println("Enter the number one ")
	fmt.Scan(&num2)

	err := divide(num1, num2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Here Your Output ", num1/num2)
	}
}
