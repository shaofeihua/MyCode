package main

import (
	"fmt"
)

func main() {
	// 声明变量，未初始化（按照默认值初始化）
	var name string
	// 声明变量，特定值初始化
	var user string = "Jing"
	// int 类型的默认值为 0
	var age int
	// float64类型的默认值为 0
	var money float64
	// 定义变量时给了初始值，那么也可以省略类型的声明，如下：变量a将自动被识别为浮点型
	var a = 0.1
	// go 与其他语言的区别之一，go 将 bool 型数据的初始值定义为 false
	var b bool

	fmt.Printf("%T(%v)\n", name, name)
	fmt.Printf("%T(%v)\n", user, user)
	fmt.Printf("%T(%v)\n", age, age)
	fmt.Printf("%T(%v)\n", money, money)
	fmt.Printf("%T(%v)\n", a, a)
	fmt.Printf("%T(%v)\n", b, b)
}
