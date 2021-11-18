package main

import (
	"fmt"
)

func main() {
	res1 := increament()               // res1 = fun
	fmt.Printf("res1 的类型是：%T\n", res1) // func() int
	fmt.Println("res1 的值是：", res1)

	v1 := res1()
	fmt.Println("v1 的值是：", v1) // 1

	v2 := res1()
	/*
		因为在 increament() 中，局部变量 i 和内层函数 fun 形成了闭包，所以 increamnet() 调用结束后，变量 i 并没有随着销毁。第一次 increament() 被调用时 i 的值从 0 自增为 1 ，所以 increament() 第二次被调用时 i 就从 1 开始自增，而不是从 0 开始自增，所以 i 的值这次就变成了 2 。
	*/
	fmt.Println("v2 的值是：", v2) // 2

	fmt.Println(res1()) // 3
	fmt.Println(res1()) // 4
	fmt.Println(res1()) // 5
	fmt.Println(res1()) // 6
}

func increament() func() int { // 外层函数
	// 1. 定义了一个局部变量
	i := 0
	// 2. 定义了一个匿名函数，给变量自增，并返回
	fun := func() int { // 内层函数
		i++
		return i
	}
	// 3. 返回该匿名函数
	return fun
}
