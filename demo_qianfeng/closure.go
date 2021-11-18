/*
	Go 语言支持函数式编程
		支持将一个函数作为另一个函数的参数
		也支持将一个函数作为另一个函数的返回值

	closure 闭包
		一个外层函数中，有内层函数，该内层函数中，会操作外层函数的局部变量（外层函数中的参数，或者外层函数中直接定义的变量），并且该外层函数的返回值就是这个局部变量。
		这个内层函数和外层函数的局部变量，统称为闭包结构。

		外层函数的局部变量的生命周期会发生改变：
			通常局部变量会随着函数调用而创建，随着函数的结束而销毁。
			但是闭包中的外层函数的局部变量，并不会随着外层函数的结束而销毁，因为内层函数还要继续使用。

*/
package main

import "fmt"

func main() {
	res1 := increment()                // res = fun
	fmt.Printf("res1 的类型： %T\n", res1) // func() int
	fmt.Println("res1: ", res1)

	v1 := res1()
	fmt.Println("v1: ", v1) // 1

	v2 := res1()
	fmt.Println("v2: ", v2) // 2

	fmt.Println(res1()) // 3
	fmt.Println(res1()) // 4
	fmt.Println(res1()) // 5
	fmt.Println(res1()) // 6

	res2 := increment()
	fmt.Println("res2: ", res2)
	v3 := res2()            // v3 = fun
	fmt.Println("v3: ", v3) // 1
	//	因为每调用一次 increment()，就会生成一个新的变量 i，所以 v3 的值是 1 ，而不是 7

	fmt.Println(res2()) // 2
	fmt.Println(res1()) // 7
}

func increment() func() int { // 外层函数
	// 1、定义了一个局部变量
	i := 0
	// 2、定义了一个匿名函数，给变量自增并返回
	fun := func() int { // 内层函数
		i++
		return i
	}
	// 3、返回该匿名函数
	return fun
}
