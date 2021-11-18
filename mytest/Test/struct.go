package main

import (
	"fmt"
)

type female struct {
	Name   string
	Age    int
	Height int
}

func main() {
	a := female{}
	a.Name = "xiaonan"
	a.Age = 30
	a.Height = 160
	fmt.Println(a)
	X(a)
	fmt.Println(a) //Age 是 int 类型，执行的是“值拷贝”，所以这里的打印，Age 的值仍然是 30，而不是 31
}

func X(fe female) {
	fe.Age = 31
	fmt.Println("X: ", fe)
}
