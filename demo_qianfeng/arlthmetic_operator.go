package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var a int
	// var b float64
	var x int
	var y float64
	fmt.Println("请输入一个整数和一个浮点类型：")
	fmt.Scanln(&x, &y)
	fmt.Printf("a 的值是：%d，b 的值是：%f\n", x, y)

	fmt.Println("请输入一个字符串：")
	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('\n')
	fmt.Println("读到的字符串是：", s1)
}
