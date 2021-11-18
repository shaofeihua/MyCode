package main

import (
	"fmt"
)

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

// 函数也是一种数据类型，通过 type 关键字给一个函数类型起名为 FuncType
// 函数的类型为 func(int,int) int ，注意这里没有括号
type FuncType func(int, int) int

func main() {
	var result1, result2 int
	result1 = Add(1, 1)
	result2 = Minus(1, 1)
	fmt.Println("result1 =", result1)
	fmt.Println("result2 =", result2)

	// 声明一个类型为函数类型的变量，变量名为 fTest
	var fTest FuncType
	var result int
	fTest = Add
	result = fTest(10, 20)
	fmt.Println("result3 =", result)

	fTest = Minus
	result = fTest(10, 20)
	fmt.Println("result4 =", result)
}
