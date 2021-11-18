package main

import "fmt"

// 多个返回值
func MyFunc() (int, int, int) {
	return 111, 222, 333
}

// go 官方推荐写法
func MyFunc01() (a, b, c int) {
	a, b, c = 111, 222, 333
	return
}

func main() {
	x, y, z := MyFunc()
	fmt.Printf("x = %d\ny = %d\nz = %d\n", x, y, z)

	a, b, c := MyFunc01()
	fmt.Printf("a = %d\nb = %d\nc = %d\n", a, b, c)
}
