package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("请输入变量 a 的值：")

	// 阻塞等待用户的输入
	fmt.Scan(&a)
	fmt.Println("a =", a)
}
