package main

import "fmt"

func main() {
	var a int = 1

	switch a {
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	default:
		fmt.Println("None")
	}
}
