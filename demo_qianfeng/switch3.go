// 练习：switch: case 后面跟随多个值
package main

import (
	"fmt"
)

func main() {
	// var letter string
	letter := ""
	fmt.Println("请输入一个大写英文字母：")
	fmt.Scanln(&letter)

	switch letter {
	case "A", "E", "I", "O", "U":
		fmt.Println(letter, "是元音")
	case "M", "N":
		fmt.Println("M或N")
	default:
		fmt.Println("其他")
	}
}
