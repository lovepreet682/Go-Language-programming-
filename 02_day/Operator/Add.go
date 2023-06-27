package main

import (
	"fmt"
)

func main() {

	var (
		sum1 = 100 + 50
		sum2 = sum1 + 250
		sum3 = sum2 + sum2
	)
	fmt.Println(sum3)

	var (
		sum4 = 500 - 50
		sum5 = sum4 - 250
		sum6 = sum5 - sum4
	)

	fmt.Println(sum6)
}
