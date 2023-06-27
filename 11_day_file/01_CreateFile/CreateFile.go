package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello")
	content := "This is the my first text file"
	file, err := os.Create("01_CreateFile/GoFile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is:- ", length)
	defer file.Close()

}
