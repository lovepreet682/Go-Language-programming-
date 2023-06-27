package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("02_createFile/firstFile.txt")

	if err != nil {
		fmt.Println("Files not created ", err)
	}
	fmt.Println("File is created successfully")
	defer file.Close()
}
