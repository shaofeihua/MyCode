package main

import "fmt"

func main() {
	// 与，&&
	// 或，||
	// 非，！
	fmt.Println("4 > 3 is:", 4 > 3)
	fmt.Println("4 != 3 is:", 4 != 3)
	fmt.Println("!(4 != 3) is:", !(4 != 3))

	fmt.Println("true && true 的结果是：", true && true)
	fmt.Println("true && false 的结果是：", true && false)
	fmt.Println("false && false 的结果是：", false && false)

	fmt.Println("true || true 的结果是：", true || true)
	fmt.Println("true || false 的结果是：", true || false)
	fmt.Println("false || false 的结果是：", false || false)
}
