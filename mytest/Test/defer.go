package main

import (
	"fmt"
)

func main() {
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")

	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}
