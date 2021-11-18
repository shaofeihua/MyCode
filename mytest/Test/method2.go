package main

import (
	"fmt"
)

type A int

func main() {
	var a A
	a = 0
	a.add()
	fmt.Println(a)
}

func (a *A) add() {
	*a += 100
}
