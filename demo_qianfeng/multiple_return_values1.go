package main

import "fmt"

// 多个返回值
// 对比两个数的大小
func MaxAndMin(a, b int) (max, min int) {
	if a > b {
		// 注意！！！应该将参数 a 和 b 赋值给返回值 max 和 min，即 max = a, min = b，而不是 a = max, b = min
		max = a
		min = b
	} else {
		min = a
		max = b
	}
	return
}

func main() {
	da, xiao := MaxAndMin(12, 15)
	fmt.Printf("da = %d\nxiao = %d\n", da, xiao)
}
