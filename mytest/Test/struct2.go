package main

import (
	"fmt"
)

//结构体包含的元素之一也是结构体
type person struct {
	name    string
	gender  string
	Contact struct {
		phone   string
		address string
	}
}

func main() {
	a := person{name: "xiaonan", gender: "female"}
	a.Contact.phone = "1351018"
	a.Contact.address = "beijing"
	fmt.Println(a)
}
