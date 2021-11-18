package main

import (
	"fmt"
)

func main() {
	var v int
	// 匿名变量（占位符）
	v, _ = getData()
	fmt.Println(v)

	_, v = getData()
	fmt.Println(v)
}

func getData() (int, int) {
	return 22, 55
}
