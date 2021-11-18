package main

import (
	"fmt"
)

func main() {
	num1 := 100
	fmt.Printf("%p\n", &num1)

	var arr1 [8]int
	fmt.Printf("%p\n", &arr1)
	fmt.Printf("%p\n", &arr1[0])
	fmt.Printf("%p\n", &arr1[1])
	fmt.Printf("%p\n", &arr1[2])
	fmt.Printf("%p\n", &arr1[3])
	fmt.Printf("%p\n", &arr1[4])
	fmt.Printf("%p\n", &arr1[5])
	fmt.Printf("%p\n", &arr1[6])
	fmt.Printf("%p\n", &arr1[7])
}
