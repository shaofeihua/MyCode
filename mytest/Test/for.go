package main

import "fmt"

func main() {
	a := 8

	for {
		a++
		if a > 10 {
			break
		}
	}
	fmt.Println(a)
}
