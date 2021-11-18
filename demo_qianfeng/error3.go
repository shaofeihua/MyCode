package main

import (
	"fmt"
)

func main() {
	length, width := 3.5, 2.6
	area, err := rectArea(length, width)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("矩形的面积是：", area)
}

type areaError struct {
	msg    string
	length float64
	width  float64
}

func (e *areaError) Error() string {
	return e.msg
}

func (e *areaError) lengthNegative() bool {
	return e.length < 0
}

func (e *areaError) widthNegative() bool {
	return e.width < 0
}

func rectArea(length, width float64) (float64, error) {
	// 定义一个空的字符串变量用来存储错误信息
	msg := ""
	if length < 0 {
		msg = "矩形的长度小于零"
	}
	if width < 0 {
		if msg == "" {
			msg = "矩形的宽度小于零"
		} else {
			msg += "，矩形的宽度也小于零"
		}
	}

	if msg != "" {
		return 0, &areaError{msg, length, width}
	}
	return length * width, nil
}
