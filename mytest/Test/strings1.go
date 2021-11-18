// 字符串转换：基本数值转化，解析，格式化。后两者为互逆的操作
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 基本数值转化
	// 整型转化为字符串
	fmt.Println(strconv.Itoa(10))
	// 字符串转化为整型。两个返回值。因为字符串在转化为整型的过程可能会报错。如果转化成功则第二个返回值则为 nil
	fmt.Println(strconv.Atoi("711"))
	// 字符串转化为整型。结果是：无效的语法
	fmt.Println(strconv.Atoi("a"))

	// 解析
	// 字符串转化为布尔型。两个返回值。
	fmt.Println(strconv.ParseBool("false"))
	// 字符串转化为浮点型。两个返回值。ParseFloat 的第二个参数是位数：32 或者 64 。二者的结果不一样
	fmt.Println(strconv.ParseFloat("3.1415926", 64))

	// 格式化
	// 将十进制数转化为二进制数。第二个参数是指进制
	fmt.Println(strconv.FormatInt(123, 2))
	// 将
	fmt.Println(strconv.FormatBool(true))
}
