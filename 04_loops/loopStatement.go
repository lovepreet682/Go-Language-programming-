package main

import (
	"fmt"
)

func main() {
	for i := 10; i >= 1; i-- {
		if i%2 == 0 {
			fmt.Println(i)
		}

	}
}
