package main

import (
	"fmt"
	"os"
)

func main() {
	oldName := "03_writeFile/firstFile.txt"
	newFile := "03_writeFile/thirdFile.txt"

	err := os.Rename(oldName, newFile)

	if err != nil {
		fmt.Print("Error remaing File ", err)
	}
	fmt.Println("File renamed successfully.")
}
