package main

import (
	"fmt"
	"math"
)

func main() {
	radius := -3.0
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		// return err
		return
	}
	fmt.Println("圆形的面积是：", area)
}

// 1、定义一个结构体，表示错误的类型
type areaError struct {
	msg    string
	radius float64
}

// 2、实现 error 接口，就是实现 Error() 方法
func (e *areaError) Error() string {
	// 格式化错误的输出
	// return fmt.Printf("error: 半径，%.2f，%s", e.radius, e.msg) // too many arguments to return，have (int, error)，want (string)
	// return e.msg
	return fmt.Sprintf("error: 半径，%.2f，%s", e.radius, e.msg)
}

// 求圆形的面积
func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"输入的半径值是非法的", radius}
	}
	return math.Pi * radius * radius, nil
}
