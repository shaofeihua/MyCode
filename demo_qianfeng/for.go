package main

import "fmt"

func main() {
	// 求1到100之间的所有整数之和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("1到100之间的整数相加之和等于：", sum)
}
