/*
	高阶函数
		根据 go 语言的数据类型的特点，可以把一个函数当做另一个函数的参数。

	fun1(), fun2()
		将函数 fun1 当做函数 fun2 的参数
			fun2 就叫做高阶函数：即接收了另一个函数作为它的参数
			fun1 就叫做回调函数：即它作为另一个函数的参数
*/
package main

import "fmt"

func main() {
	// 这里只是把 add 作为参数引用，但是还未调用，所以不能写成 add()，否则就变成了引用函数 add() 的计算结果
	res1 := oper(10, 20, add)
	fmt.Println("res1: ", res1)

	res2 := oper(5, 2, sub)
	fmt.Println("res2: ", res2)

	//	匿名函数作为参数：先把匿名函数赋值给一个变量
	fun1 := func(a, b int) int {
		return a * b
	}
	res3 := oper(2, 5, fun1)
	fmt.Println("res3: ", res3)

	//	匿名函数作为参数：直接作为函数参数
	res4 := oper(100, 20, func(a, b int) int {
		if b == 0 {
			fmt.Println("除数不能为 0")
			return 0
		}
		return a / b
	})
	fmt.Println("res4: ", res4)
}

// 加法
func add(a, b int) int {
	return a + b
}

// 减法
func sub(a, b int) int {
	return a - b
}

func oper(a, b int, fun func(a, b int) int) int {
	fmt.Println(a, b, fun) // 打印 3 个参数
	res := fun(a, b)
	return res
}
