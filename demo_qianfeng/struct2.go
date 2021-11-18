package main

import (
	"fmt"
)

func main() {
	// 匿名结构体
	s2 := struct {
		name string
		age  int
	}{
		name: "李四",
		age:  30,
	}
	fmt.Println(s2.name, s2.age)

	w1 := Worker{
		"李小花", 30,
	}
	fmt.Println(w1)
	fmt.Println(w1.string, w1.int)
}

// 匿名字段：只有类型，没有字段名。此时字段类型不可重复，例如只能有一个 string
type Worker struct {
	string
	int
}
