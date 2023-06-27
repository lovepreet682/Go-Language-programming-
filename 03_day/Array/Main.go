package main

import "fmt"

func main() {
	//decleared the array
	//var array_name=[length]datatype{values}
	var name = []string{"Love Rehan", "Gagan", "Sahil", "Kuljinder", "Chetan"}
	fmt.Println(name)

	//another way to use and declared the array
	num := [5]int{2, 3, 5, 6, 7}
	fmt.Println(num)

	//if i want to print specific data
	fmt.Println(name[2])

	//change the data
	name[2] = "Sahil Dhimaan"
	fmt.Println(name)

	//find the length
	fmt.Println("Total length of the array ", len(name))

	//Array initalization
	arr1 := [5]int{}
	fmt.Println(arr1)

	//for each in the array
	fmt.Println("Here is the List of name :- ")
	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
	}

}
