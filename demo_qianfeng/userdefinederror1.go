/*
	自定义 error
*/
package main

import "fmt"

func main() {
	area, err := areaCectangular(3, 3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("矩形的面积是：", area)

	area1, err1 := areaCectangular(-5, -5)
	if err1 != nil {
		fmt.Println(err1)
		// 使用断言判断到底是长度出错还是宽度出错
		if err1, ok := err1.(*areaError); ok {
			if err1.lengthNegative() {
				fmt.Printf("Error: 长度 %.2f 小于 0\n", err1.length)
			}
			if err1.widthNegative() {
				fmt.Printf("Error：宽度 %.2f 小于 0\n", err1.width)
			}
		}
		return
	}
	fmt.Println("矩形的面积是：", area1)

}

type areaError struct {
	length float64
	width  float64
	msg    string // 用来描述错误
}

func (e *areaError) Error() string {
	return fmt.Sprintf("Error: %s", e.msg)
}

func (e *areaError) lengthNegative() bool {
	return e.length < 0
}

func (e *areaError) widthNegative() bool {
	return e.width < 0
}

func areaCectangular(length, width float64) (float64, error) {
	msg := ""
	if length < 0 {
		msg = "长度小于 0"
	}
	if width < 0 {
		if msg == "" {
			msg = "宽度小于 0"
		} else {
			msg += "，宽度也小于 0"
		}
	}

	if msg != "" {
		/*
			注意，这里必须使用 &areaError{} 而不是 areaError{}，否则报错：
			cannot use areaError{...} (type areaError) as type error in return argument:
				areaError does not implement error (Error method has pointer receiver)

		*/
		return 0, &areaError{length, width, msg}
	}
	return length * width, nil
}
