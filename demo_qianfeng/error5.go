/*
	error 错误
		内置的数据类型，内置的方法

	创建 error 有两种方法：
		1、使用 errors 下的 New() 函数，创建一个 error 对象
		2、fmt 包下的 Errorf() 函数：
			func Errorf(format string, a ...interface{}) error
*/
package main

import (
	"errors"
	"fmt"
)

func main() {
	// 1、第一种方法
	err1 := errors.New("相信你是我的错。")
	fmt.Println("err1:", err1)
	fmt.Printf("err1 的类型: %T\n", err1) // *errors.errorString

	// 2、第二种方法
	err2 := fmt.Errorf("错误的状态码: %d", 404)
	fmt.Println("err2:", err2)
	fmt.Printf("err2 的类型: %T\n", err2) // *errors.errorString

	res := checkAge(10)
	// 注意！！！如果这里不加 if 判断，直接打印 res，当 age 大于等于 0 ，也就是 checkAge() 的返回值为 nil 时，此时除了打印 checkAge() 正常的返回结果外，还会多打印一行，内容是：<nil>
	if res != nil {
		fmt.Println(res)
		return
	}
}

// 创建一个函数，验证年龄是否合法：年龄小于 0 则报年龄不合法，大于等于 0 即提示年龄合法。
func checkAge(age int) error {
	if age < 0 {
		//err := errors.New("Error：年龄不合法，请确认！")
		err := fmt.Errorf("您给定的年龄是：%d，不合法。", age)
		return err
	}
	fmt.Printf("年龄是：%d\n", age)
	return nil
}
