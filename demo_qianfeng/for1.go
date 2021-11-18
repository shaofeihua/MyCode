package main

import "fmt"

func main() {
	// 用 for 循环的方式打印字符串中的每一个字符
	str := "abcde"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}
	fmt.Println(len(str))
}
