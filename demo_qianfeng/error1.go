package main

import (
	"errors"
	"fmt"
)

func main() {
	// 1、创建一个 error 数据
	err1 := errors.New("自己创建玩的。。。")
	fmt.Println(err1)
	fmt.Printf("%T\n", err1) // *errors.errorString

	// 2、另一个创建 error 的方法
	err2 := fmt.Errorf("错误的信息码：%d", 100)
	fmt.Println(err2)
	fmt.Printf("%T\n", err2)

	Age(-10)
}

// 设计一个函数，如果年龄不合法，则返回错误信息
func Age(a int) error {
	if a < 0 {
		// return errors.New("您输入的年龄小于 0，不合法。")
		err := fmt.Errorf("您输入的年龄小于 0，不合法。")
		fmt.Println(err)
		return err
	}
	fmt.Printf("您输入的年龄是：%d\n", a)
	return nil // 如果年龄合法，也要有返回值，error 可以为空
}
