package main

import "fmt"

func main() {
	if 3 < 4 {
		fmt.Println("I love you")
	}

	if a := 1; a == 10 {
		fmt.Println("a = 10")
	} else if a > 10 {
		fmt.Println("a > 10")
	} else if a < 10 {
		fmt.Println("a < 10")
	}
}
