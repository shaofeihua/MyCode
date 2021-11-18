// 打印乘法表
package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d X %d = %d\t", j, i, i*j)
		}
		fmt.Println() // 在这个位置放一个打印功能很重要：用来实现阶梯式地换行
	}
}
