/*
	匿名函数

	定义一个匿名函数，可以直接调用，通常只能调用一次。
	可以将匿名函数赋值给某个变量，就可以调用多次。

	匿名函数的用途：
		Go支持函数式变成：
		1、将匿名函数作为另一个函数的参数，即 回调函数。
		2、将匿名函数作为另一个函数的返回值，可以形成闭包结构。
*/
package main

import "fmt"

func main() {
	fun1()
	fun2 := fun1
	fun2()
	fmt.Printf("fun1的类型是： %T\n", fun1)

	// 匿名函数，大括号的结尾跟一对小括号，直接调用
	func() {
		fmt.Println("我是一个匿名函数")
	}()

	fun3 := func() {
		fmt.Println("我也是一个匿名函数")
	}
	/*
		注意：
			上面 fun3 这部分代码，因为这里是把匿名函数赋值给变量，并且该匿名函数没有参数，所以就没法赋值，
			也就无法在赋值时直接调用，所以大括号后面不加小括号()，通过变量 fun3 调用时才加小括号
	*/
	fun3()
	fun3()

	//	带参数的匿名函数
	func(a, b int) {
		fmt.Println(a, b)
	}(1, 2)

	//	带返回值的匿名函数：赋值给变量时时直接调用
	sum := func(a, b int) int {
		return a + b
	}(1, 2)
	fmt.Println("sum: ", sum)

	//	带返回值的匿名函数：赋值给变量时不调用，后面通过变量调用
	sum1 := func(a, b int) int {
		return a + b
	}
	fmt.Println(sum1) // 0x3f99e0
	fmt.Println(sum1(10, 20))
	//sum1(10, 20)，这样调用，不会输出任何结果。程序也不报错。
}

func fun1() {
	fmt.Println("我是fun1()")
}
