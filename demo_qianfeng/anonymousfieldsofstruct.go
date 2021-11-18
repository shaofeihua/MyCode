/*
	匿名结构体和匿名字段
		匿名结构体：
			没有名字的结构体
			在创建匿名结构体时，同时创建对象
		匿名字段：
			结构体的字段没有字段名
			以类型名作为字段的名称，例如 string
			不能存在两个相同的匿名字段（即 匿名字段的类型不能重复），例如不允许存在两个 string
*/
package main

import "fmt"

func main() {
	// 创建一个匿名结构体
	stu1 := struct {
		name string
		sex  string
	}{
		name: "张三",
		sex:  "男",
	}
	fmt.Println("stu1: ", stu1) // {张三 男}

	w1 := Worker{"王五", 30}
	fmt.Println("w1: ", w1)        // {王五 30}
	fmt.Println(w1.string, w1.int) // 王五 30
	w2 := Worker{"北京市", 5}
	fmt.Println("w2: ", w2) // {北京市 5}
}

type Worker struct {
	string // 匿名字段
	int    // 匿名字段
	// string。前面已经有一个 string，这里不能再有一个 string ，因为不允许存在两个相同的匿名字段
}
