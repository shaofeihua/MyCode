package main

import "fmt"

// 不确定参数的个数，使用...的方式。例如 ...int 或者 ...string
// 传递的实参可以是 0 个或者多个
func MyFunc(args ...int) {
	fmt.Println("len(args) = ", len(args)) // 获取参数的长度，即个数

	for i := 0; i < len(args); i++ {
		fmt.Printf("args[%d] = %d\n", i, args[i])
	}

	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++")
}

func main() {
	MyFunc()
	MyFunc(2, 5)
	MyFunc(2, 5, 7)
}
