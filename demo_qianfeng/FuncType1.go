// 回调函数：函数的参数之一为函数
// 示例：实现计算器 calc
package main

import (
	"fmt"
)

type FuncType func(int, int) int

func Calc(a, b int, fTest FuncType) (result int) {
	fmt.Println("Calc")
	result = fTest(a, b)
	return
}

// 实现加法
func Add(a, b int) int {
	return a + b
}

// 实现乘法
func Multiply(a, b int) int {
	return a * b
}

func main() {
	a := Calc(2, 5, Add)
	fmt.Println("a =", a)

	b := Calc(2, 5, Multiply)
	fmt.Println("b =", b)
}
