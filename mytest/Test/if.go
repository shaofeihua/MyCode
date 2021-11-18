package main

import "fmt"

func main() {
	a := 5

	if a := 1; a > 0 {
		a++
		fmt.Println(&a)
	}

	fmt.Println(a)
}
