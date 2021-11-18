package main

import (
	"fmt"
)

// 准备一个会报错的函数：让分母为 0
func test(x int) {
	result := 10 / x
	fmt.Println("result =", result)
}

func main() {
	defer fmt.Println("a")
	defer fmt.Println("b")

	defer test(0) // 调用会报错的函数，看看 defer 的影响

	defer fmt.Println("c")
}
