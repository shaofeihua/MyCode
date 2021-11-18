package main

import (
	"fmt"
)

type humen struct {
	sex int
}

type teacher struct {
	humen //嵌套结构体
	name  string
	age   int
}

type student struct {
	humen //嵌套结构体
	name  string
	age   int
}

func main() {
	a := teacher{name: "wang", age: 50}
	b := student{name: "lilei", age: 20}
	a.sex = 0
	b.sex = 1

	//a.humen.sex = 0
	//b.humen.sex = 1
	a.sex = 0 //等同于a.humen.sex。因为前面的 humen 已经嵌套到 teacher 和 student 里面了，后两者可以直接饮用 humen 的字段。
	b.sex = 1 //等同于b.humen.sex。因为前面的 humen 已经嵌套到 teacher 和 student 里面了，后两者可以直接饮用 humen 的字段。
	fmt.Println(a, b)
}
