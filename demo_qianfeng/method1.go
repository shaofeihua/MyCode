/*
	方法
		方法就是包含了接收者的函数。
		接收者可以是命名类型或者结构体类型的一个指针或者一个值。
		所有给定类型的方法，属于该类型的方法集。

		函数名不能是相同的，而方法名可以相同，但是方法的接受者不能相同

		对比函数：
			1、意义不同
				方法：某个类别的行为功能，只能被指定的接收者调用
				函数：一段独立功能的代码，可以直接调用
			2、语法不同
				方法：方法名可以相同，只要接收者不同就行
				函数：函数名不能相同
*/
package main

import "fmt"

func main() {
	w1 := Worker{"王二狗", 20}
	fmt.Println("w1: ", w1) // {王二狗 20}
	w1.work()               // 王二狗 在工作...
	w1.name = "王老狗"
	fmt.Println("w1: ", w1) // {王老狗 20}
	w1.work()               // 王老狗 正在工作...

	w2 := &Worker{"老卢", 33}
	fmt.Println("w2: ", w2)                    // &{老卢 33}
	fmt.Printf("w2 的值：%p,w2 的类型：%T\n", w2, w2) // w2 的值：0xc0000040a8,w2 的类型：*main.Worker
	w2.work()                                  // 老卢 在工作...
	w2.name = "小卢"                             // &{小卢 33}
	fmt.Println("w2: ", w2)
	w2.work() // 小卢 正在工作...

	w3 := Worker{"老白", 28}
	fmt.Println("w3: ", w3) // {老白 28}
	w3.rest()               // 老白 正在休息...
	w3.name = "小白"
	fmt.Println("w3: ", w3) // {小白 28}
	w3.rest()               // 小白 正在休息...

	w4 := Worker{"邢育森", 40}
	w4.printinfo() // 工人的姓名：邢育森，工人的年龄：40

	c1 := Cat{"黑色", 2}
	c1.printinfo() // 猫的颜色：黑色，猫的年龄：2
}

// 1、定义一个工人结构体
type Worker struct {
	name string
	age  int
}

type Cat struct {
	color string
	age   int
}

// 2、定义行为方法
func (w Worker) work() {
	fmt.Println(w.name, "正在工作...")
}

func (w *Worker) rest() {
	//fmt.Println(*w.name, "正在休息...") // invalid indirect of w.name (type string)
	fmt.Println(w.name, "正在休息...") // *w.name 必须简写为 w.name，否则报错
}

func (w Worker) printinfo() {
	fmt.Printf("工人的姓名：%s，工人的年龄：%d\n", w.name, w.age)
}

func (c Cat) printinfo() {
	fmt.Printf("猫的颜色：%s，猫的年龄：%d\n", c.color, c.age)
}
