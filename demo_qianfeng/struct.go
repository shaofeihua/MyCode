package main

import (
	"fmt"
)

type name_struct struct {
	ID int
	iD int
}

var data name_struct

func main() {
	// 赋值方式1：把结构体中的变量单独拿出来赋值
	data.ID = 100
	data.iD = 10

	fmt.Println("data =", data)

	// 赋值方式2：把整个结构体赋值给一个变量的提同时对整个结构体进行赋值
	data2 := name_struct{
		ID: 88,
		iD: 66,
	}
	fmt.Println("data2 =", data2)
}
