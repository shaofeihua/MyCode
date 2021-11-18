package main

import (
	"fmt"
)

func main() {
	p1 := Person{"王二狗", 30, "男", "北京市"}
	fmt.Printf("p1 的类型是：%T，值是：%v，地址是：%p\n", p1, p1, &p1)

	fmt.Println()

	p2 := p1
	fmt.Printf("p2 的类型是：%T，值是：%v，地址是：%p\n", p2, p2, &p2)
	p2.name = "王小花"
	fmt.Printf("p2 的类型是：%T，值是：%v，地址是：%p\n", p2, p2, &p2) // p2 改变
	fmt.Printf("p1 的类型是：%T，值是：%v，地址是：%p\n", p1, p1, &p1) // p1 不变

	// 定义结构体指针
	pp1 := &p1
	fmt.Printf("pp1 的类型是：%T，值是：%v，存储的地址是：%p，本身的地址是：%p\n", pp1, pp1, pp1, &pp1)
	fmt.Println(pp1)
	fmt.Println(*pp1)

	fmt.Println()

	pp1.name = "大刀王五"
	fmt.Println(p1)

	// 使用专门用于创建某种类型的指针的函数：new()

	p3 := new(Person)
	fmt.Println(p3)
	fmt.Printf("p3 的类型是：%T，值是：%v，存储的地址是：%p，本身的地址是：%p\n", p3, p3, p3, &p3)

	p4 := new(int)
	fmt.Println(p4)
	fmt.Printf("p4 的类型是：%T，值是：%v，存储的地址是：%p，本身的地址是：%p\n", p4, p4, p4, &p4)
}

type Person struct {
	name    string
	age     int
	sex     string
	address string
}
