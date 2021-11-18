// 不用作用域的同名变量
// 不同作用域，允许定义同名变量
// 使用变量的原则，就近原则
package main

import (
	"fmt"
)

var a byte //全局变量

func main() {
	var a int // 局部变量
	fmt.Printf("a1: %T\n", a)

	{
		var a float64
		fmt.Printf("a2: %T\n", a)
	}

	test()
}

func test() {
	fmt.Printf("a3: %T\n", a)
}
