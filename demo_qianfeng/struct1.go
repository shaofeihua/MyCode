package main

import (
	"fmt"
)

func main() {
	// 定义结构体
	type Person struct {
		name    string
		age     int
		sex     string
		address string
	}

	// 方法一：
	var p1 Person
	p1.name = "小楠"
	p1.age = 30
	p1.sex = "女"
	p1.address = "沈阳"
	fmt.Printf("姓名：%s，年龄：%d，性别：%s，住址：%s\n", p1.name, p1.age, p1.sex, p1.address) //姓名：小楠，年龄：30，性别：女，住址：沈阳
	fmt.Println(p1)                                                              // {小楠 30 女 沈阳}

	// 方法二：
	p2 := Person{}
	p2.name = "ruby"
	p2.age = 28
	p2.sex = "女"
	p2.address = "北京"
	fmt.Printf("姓名：%s，年龄：%d，性别：%s，住址：%s\n", p2.name, p2.age, p2.sex, p2.address)

	// 方法三：
	p3 := Person{name: "如花", age: 16, sex: "未知", address: "广东"}
	fmt.Println(p3)
	p4 := Person{
		name:    "王二狗",
		age:     30,
		sex:     "男",
		address: "河南", // 最后一行的结尾必须有逗号
	}
	fmt.Println(p4)

	// 方法四：
	p5 := Person{"满满", 6, "男", "赤峰"} // 省略字段名时，初始化的顺序必须符合 struct 的顺序。例如第二字段 age 必须是整型
	fmt.Println(p5)
}
