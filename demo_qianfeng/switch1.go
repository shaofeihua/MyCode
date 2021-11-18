// 练习：模拟计算器：输入两个数字并选择一种算数运算，计算结果
package main

import (
	"fmt"
)

func main() {
	num1 := 0
	num2 := 0
	oper := ""
	fmt.Println("请输入一个整数：")
	fmt.Scanln(&num1)
	fmt.Println("请再次输入一个整数：")
	fmt.Scanln(&num2)
	fmt.Println("请输入一个操作：+,-,*,/")
	fmt.Scanln(&oper)

	switch oper {
	case "+":
		fmt.Printf("%d + %d = %d\n", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%d - %d = %d\n", num1, num2, num1-num2)
	case "*":
		fmt.Printf("%d + %d = %d\n", num1, num2, num1*num2)
	case "/":
		fmt.Printf("%d / %d = %d\n", num1, num2, num1/num2)
	}
}
