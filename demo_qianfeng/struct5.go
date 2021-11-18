/*
	结构体
*/
package main

import "fmt"

func main() {
	// 新建一个结构体对象
	// 方法1：var p1 Persion
	// 方法2：
	p1 := Persion{}
	fmt.Println("p1: ", p1) // { 0  }
	p1.name = "王二狗"
	p1.age = 30
	p1.sex = "男"
	p1.address = "内蒙古"
	fmt.Println("p1: ", p1) // {王二狗 30 男 内蒙古}
	fmt.Printf("p1 的姓名：%s，年龄：%d，性别：%s，住址：%s\n", p1.name, p1.age, p1.sex, p1.address)

	//	方法3：
	p2 := Persion{name: "盛小楠", age: 33, sex: "女", address: "沈阳市"}
	fmt.Printf("p2 的姓名：%s，年龄：%d，性别：%s，住址：%s\n", p2.name, p2.age, p2.sex, p2.address)

	// 方法4：
	p3 := Persion{"房凌月", 33, "女", "赤峰市"}
	fmt.Printf("p3 的姓名：%s，年龄：%d，性别：%s，住址：%s\n", p3.name, p3.age, p3.sex, p3.address)
}

type Persion struct {
	name    string
	age     int
	sex     string
	address string
}
