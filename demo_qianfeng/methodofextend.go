/*
	OOP 中的继承性：
		如果两个类（class）存在继承关系，其中一个是子类，另一个作为父类，那么：
		1、子类可以直接访问父类的属性和方法；
		2、子类可以新增自己的属性和方法；
		3、子类可以重写父类的方法（override，就是将父类已有的方法，重新实现）

	method 是可以继承的，如果匿名字段实现了一个 method，那么包含这个匿名字段的 struct 也可以调用这个 method
*/
package main

import "fmt"

func main() {
	p1 := Person{"陈老狗", 60}
	p1.eat()

	s1 := Student{Person{"陈小狗", 25}, "仙家山大学"}
	s1.eat() // 子类对象，访问父类的方法

	s1.study() // 子类对象，访问自己新增的方法

	s1.eat() // 子类对象，重写父类的方法
}

// 1、定义一个“父类”
type Person struct {
	name string
	age  int
}

// 2、定义一个“子类”
type Student struct {
	Person // 结构体嵌套，模拟面向对象语言的继承性
	school string
}

// 3、方法
func (p Person) eat() {
	fmt.Println("这是父类的方法，吃窝窝头...")
}

// 子类新增的方法
func (s Student) study() {
	fmt.Println("子类新增的方法，开始学习啦")
}

// 子类重新 eat 方法
func (s Student) eat() {
	fmt.Println("子类改写的方法，不吃窝窝头，改吃炸鸡喝啤酒了...")
}
