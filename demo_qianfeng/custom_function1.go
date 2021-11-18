package main

import "fmt"

// 有参数无返回值的自定义函数
// 括号内的“a int”相当于在函数体进行变量声明的语句“var a int”
// 定义函数时，小括号内的参数叫形参
func MyFunc01(a int) {
	fmt.Println("a =", a)
}

// 多个参数
func MyFunc02(a int, b int) {
	fmt.Printf("a = %d, b = %d\n", a, b)
}

// 多个参数为相同类型时的简单写法
func MyFunc03(a, b int) {
	fmt.Printf("a = %d, b = %d\n", a, b)
}

// 多个不同类型的参数
func MyFunc04(a int, b string, c float64) {
	fmt.Printf("a = %d, b = %v, c = %v\n", a, b, c)
}

// 调用有参数的函数时，括号内不能为空，必须为被调用的函数的参数赋值
// 调用函数时传递的参数叫实参
// 参数传递只能由实参传递给形参，不能由形参传递给实参，是单向传递
func main() {
	MyFunc01(222)
	MyFunc02(222, 333)
	MyFunc03(222, 333)
	MyFunc04(222, "333", 444.2)
}
