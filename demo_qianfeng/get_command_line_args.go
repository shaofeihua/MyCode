// 获取命令行参数
package main

import (
	"fmt"
	"os"
)

func main() {
	// 接收用户传递的参数，都是以字符串的方式传递的
	list := os.Args
	n := len(list)
	fmt.Println(list)
	fmt.Println(n)

	for i := 0; i < n; i++ {
		fmt.Printf("list[%d] = %s\n", i, list[i])
	}

	for i, data := range list {
		fmt.Printf("list[%d] = %s\n", i, data)
	}
}
