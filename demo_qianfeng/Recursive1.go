// 函数的递归调用
package main

import "fmt"

func test(n int8) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("n = ", n)
}

func main() {
	test(4)
}
