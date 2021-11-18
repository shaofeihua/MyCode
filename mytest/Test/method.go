package main

import (
	"fmt"
)

type A struct {
	Name string
}

type B struct {
	Name string
}

func main() {
	a := A{}
	a.Print()
	fmt.Println(a.Name)

	b := B{}
	b.Print()
}

func (a *A) Print() {
	a.Name = "xiaonan"
	fmt.Println("This is method A")
}

func (b B) Print() {
	fmt.Println("This is method B")
}
