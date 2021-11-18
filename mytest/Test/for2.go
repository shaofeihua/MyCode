package main

import "fmt"

func main() {
	a := 6

	for i := 0; i < 3; i++ {
		a++

		fmt.Println(i, a)
	}

	//fmt.Println(a)
}
