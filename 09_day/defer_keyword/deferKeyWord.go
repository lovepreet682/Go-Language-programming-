package main

import "fmt"

func Greeting(msg string) {
	fmt.Println("Good Morning", msg)
}

func looping() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
func main() {
	defer looping()
	defer Greeting("Lovepreet")
	fmt.Println("Exato")
}
