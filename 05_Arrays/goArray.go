package main

import (
	"fmt"
)

func main() {

	//int
	var arrayData = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arrayData)

	//String
	language := []string{"Go", "Java", "C++", "PHP", "Node js"}
	fmt.Println(language)

	//change the value of the language
	language[3] = "Spring boot"
	fmt.Println(language)

	//find the length of the array
	length := len(language)
	fmt.Println("Length of the data is ", length)

	//looping through the array
	for i := 0; i < len(language); i++ {
		fmt.Println(language[i])
	}

	//take the input from the user
	fmt.Println("How Many values you want to print: -")
	var size int
	fmt.Scanln(&size)

	var arr = make([]string, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter the %dth value: ", i)
		fmt.Scan(&arr[i])
	}
	fmt.Println("Values successfully Saved ")
	fmt.Print("Here your output ", arr)

}
