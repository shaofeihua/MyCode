package main

import "fmt"

func main() {
	a := [10]int{1, 0}
	fmt.Println(a)

	b := [8]int{6: 88}
	fmt.Println(b)

	c := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(c)

	d := [...]int{15: 86}
	fmt.Println(d)

}
