package main

import (
	"fmt"
)

func main() {
	const a = 10
	fmt.Println("a = ", a)
	fmt.Printf("a is %T\n", a)

	const (
		e1 = iota
		e2
		e3
	)

	const (
		a1 = iota + 1
		a2
		a3
	)

	fmt.Println(e1, e2, e3)
	fmt.Println(a1, a2, a3)

	abc := 10
	fmt.Println("abc: ", abc)
	{
		// 这里的 abc := 30 和 abc = 30，语法上都没错误，但是意义完全不一样。
		// abc := 30
		abc = 30
		fmt.Println("abc: ", abc)
	}
	// 当上面大括号中的使用 abc := 30，这里的打印结果是 abc: 10
	// 当上面大括号中的使用 abc = 30，这里的打印结果是 abc: 30
	fmt.Println("abc: ", abc)
}
