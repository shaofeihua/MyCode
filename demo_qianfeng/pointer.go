package main

import (
	"fmt"
)

func main() {
	a := 10
	var p *int
	fmt.Println("p 的值是：", p) // <nil>，空指针
	p = &a
	fmt.Println("a 的值是：", a)
	fmt.Println("a 的地址是：", &a)
	fmt.Println("p 的值是：", p)
	fmt.Println("p 的地址是：", &p)
	fmt.Println("p 的值对应的数值是：", *p)

	a = 100
	fmt.Println("a 的值是：", a)
	fmt.Println("a 的地址是：", &a) // 地址不变

	// 通过指针改变变量的值
	*p = 200
	fmt.Println("a 的值是：", a)
	fmt.Println("a 的地址是：", &a)

	// 指针的指针
	var p1 **int
	fmt.Println("p1 的值是：", p1)
	fmt.Println("p1 的地址是：", &p1)

	// p1 存储的是指针 p 的地址
	p1 = &p
	fmt.Println("通过 p1 获取到 p 的地址是：", p1)
	fmt.Println("通过 p1 获取到 p 的值是：", *p1)
	fmt.Println("通过 p1 获取到 p 的值对应的值是：", **p1)
}
