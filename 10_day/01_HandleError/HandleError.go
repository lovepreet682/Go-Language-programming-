package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("handle the Error in my Program")
	message := "Loverehan"

	//create the custom msg
	myErrorMsg := errors.New("Wrong Name using")

	if message != "GO" {
		fmt.Println(myErrorMsg)
	}
}
