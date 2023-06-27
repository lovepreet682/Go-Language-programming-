package main

import "fmt"

func main() {
	var bill int
	var discount int
	var afterDiscount int
	var saveMoney int
	fmt.Println("Enter bill amount:")
	fmt.Scan(&bill)
	fmt.Println("Enter discount percentage:")
	fmt.Scan(&discount)
	afterDiscount = bill - (bill * discount / 100)
	saveMoney = bill - afterDiscount
	fmt.Println("After discount your bill is: ", afterDiscount)
	fmt.Println("You saved money", saveMoney)

}
