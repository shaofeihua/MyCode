package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i1 myint
	var i2 = 100
	i1 = 200
	fmt.Println(i1, i2)

	var name mystr
	name = "王二狗"
	var s1 string
	s1 = "李小花"
	fmt.Println(name, s1)

	fmt.Printf("%T,%T,%T,%T\n", i1, i2, name, s1) // main.myint,int,main.mystr,string

	f1 := fun1()
	fmt.Println(f1(10, 20)) // 1020
	fmt.Printf("%T,%T,%T\n",f1,fun1,fun1()) // main.myfun,func() main.myfun,main.myfun
}

// 自定义一个新类型
type myint int
type mystr string

// 自定义一个函数类型
type myfun func(int, int) string

func fun1() myfun {
	fun := func(a, b int) string {
		s := strconv.Itoa(a) + strconv.Itoa(b)
		return s
	}
	return fun
}
