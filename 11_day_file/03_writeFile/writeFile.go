package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	message := "Hello Im adding the some information in the file"
	err := ioutil.WriteFile("03_writeFile/firstFile.txt", []byte(message), 0644)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Text has been added in the file")
}
