package main

import (
	"fmt"
)

func main() {
	// go 是静态语言，要求变量在声明时候的类型和赋值的类型必须一致
	// 变量的值改变后，地址并没有变化
	var num int
	num = 100
	fmt.Printf("num 的数值是：%d，地址是：%p\n", num, &num)

	num = 200
	fmt.Printf("num 的数值是：%d，地址是：%p\n", num, &num)
}
