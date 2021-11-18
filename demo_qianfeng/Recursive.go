// 函数的递归调用
package main

import "fmt"

func Recursive(a int) {
	if a == 1 {
		fmt.Println("a =", a)
		return // 终止函数的调用
	}
	// 函数调用自身
	Recursive(a - 1)

	fmt.Println("a =", a)
}

func main() {
	var a int
	Recursive(3)
	fmt.Println("a =", a)
}
