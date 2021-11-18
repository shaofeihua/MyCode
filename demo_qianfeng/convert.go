package main

import "fmt"

func main() {
	var a int
	var b byte
	b = 'y'

	a = int(b)
	fmt.Println("a =", a)
}
