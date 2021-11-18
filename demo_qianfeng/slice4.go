package main

import (
	"fmt"
)

func main() {
	// 使用数组创建切片
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := a[:5] // 1-5
	fmt.Println("s1：", s1)
	s2 := a[5:9] //6-9
	fmt.Println("s2：", s2)
	s3 := a[6:] // 7-10
	fmt.Println("s3：", s3)
	s4 := a[:] //1-10
	fmt.Println("s4：", s4)

	fmt.Println("--------------------------------")

	fmt.Printf(" a 的内存地址：%p\n", &a)
	fmt.Printf("s1 的内存地址：%p\n", s1)
	fmt.Printf("s2 的内存地址：%p\n", s2)
	fmt.Printf("s3 的内存地址：%p\n", s3)
	fmt.Printf("s4 的内存地址：%p\n", s4)
}
