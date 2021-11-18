/*
	自定义 error
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	radius := -3.0
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println("err:", err)
		return // 这里的 return 的作用是，当报错（err 不为 nil）时，结束整个函数
	}
	fmt.Println("圆形的面积：", area) // 圆形的面积： 28.274333882308138
}

// 1、定义一个结构体，表示错误的类型
type areError struct {
	radius float64
	msg    string
}

// 2、实现 error 接口，就是实现 Error() 方法
func (e *areError) Error() string {
	return fmt.Sprintf("error: 半径，%.2f，%s", e.radius, e.msg)
}

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areError{radius, "半径是非法的"}
	}
	return math.Pi * radius * radius, nil
}
