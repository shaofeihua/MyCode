package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i11 myint
	var i12 mystr
	fmt.Printf("i11 的类型是：%T\n", i11) // main.myint
	fmt.Printf("i12 的类型是：%T\n", i12) // main.mystring

	var i1 myint
	var i2 = 100
	i1 = 100
	fmt.Println(i1, i2)

	var name mystr
	name = "王二狗"
	var s1 string
	s1 = "李小花"
	fmt.Println(name, s1)

	// i1 = i2 // cannot use i2 (type int) as type myint in assignment
	// name = s1 // cannot use s1 (type string) as type mystr in assignment

	fmt.Printf("%T,%T,%T,%T\n", i1, i2, name, s1)

	fmt.Println("---------------------------------------")

	res1 := fun1()          // 错误的用法：fun1(1,3)
	fmt.Println(res1(1, 3)) // 13

	fmt.Println("---------------------------------------")
	var i3 myint2
	i3 = 1000
	fmt.Printf("i3 的类型是：%T，i3 的值是：%d\n", i3, i3)
	i3 = i2
	fmt.Printf("i3 的类型是：%T，i3 的值是：%d\n", i3, i3)
}

// 1、定义新的类型
type myint int
type mystr string

// 2、定义函数类型
type myfun func(int, int) string

func fun1() myfun {
	fun := func(a, b int) string {
		s := strconv.Itoa(a) + strconv.Itoa(b)
		return s
	}
	return fun
}

// 3、类型别名
type myint2 = int
