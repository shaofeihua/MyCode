package main

import (
	"fmt"
)

func main() {
	// 只给结构体赋值但不使用方法
	w1 := Worker{name: "王二狗", age: 20, sex: "男"}
	fmt.Println(w1)

	// 使用方法
	w1.work()

	w2 := &Worker{name: "王小花", age: 20, sex: "女"}
	fmt.Printf("w2 的类型是：%T\n", w2)
	w2.work()
	w2.rest()

	w3 := &Worker{name: "诸葛大力", age: 23, sex: "女"}
	c1 := &Cat{color: "黑色", age: 2}
	w3.printInfo()
	c1.printInfo()
}

// 1、定义一个工人结构体
type Worker struct {
	name string
	age  int
	sex  string
}

type Cat struct {
	color string
	age   int
}

// 2、定义方法
func (w Worker) work() {
	fmt.Println(w.name, "正在工作")
}

// 3、定义方法，调用它的参数的类型是指针
func (w *Worker) rest() {
	fmt.Println(w.name, "正在休息")
}

func (p *Worker) printInfo() {
	fmt.Printf("工人姓名：%s，工人年龄：%d，工人性别：%s\n", p.name, p.age, p.sex)
}

func (p *Cat) printInfo() {
	fmt.Printf("猫的颜色：%s，猫的年龄：%d\n", p.color, p.age)
}
