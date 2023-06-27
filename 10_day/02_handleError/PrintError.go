package main

import (
	"errors"
	"fmt"
)

func handleErrorMsg(name string) error {
	//define the error
	errorMsg := errors.New("Invalid Name")

	//check the error
	if name != "loverehan" {
		return errorMsg
	} else {
		return nil
	}
}
func main() {
	fmt.Println("Let's Check the Error")
	name := "loverehan"
	errorMsg := handleErrorMsg(name)

	if errorMsg != nil {
		fmt.Println(errorMsg)
	} else {
		fmt.Println("hello", name)
	}
}
