package main

import "fmt"

func MyFunc() int {
	return 222
}

// 给返回值一个变量名，例如 result，这是 go 的推荐写法
func MyFunc01() (result int) {
	result = 333
	return
}

func main() {
	var a int
	a = MyFunc()
	fmt.Println("a =", a)

	// 当变量被赋值为函数时，也可以使用“:=”这种简便写法
	b := MyFunc()
	fmt.Println("b =", b)

	c := MyFunc01()
	fmt.Println("c =", c)
}
