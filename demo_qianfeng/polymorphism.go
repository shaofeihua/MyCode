/*
	接口
		接口是一组方法的签名。
		当某个类型实现了接口中的所有方法，它就实现了这个接口。
		不需要显式的声明该类型实现了该接口。

	1、当需要接口类型的对象时，可以使用任意实现类对象代替
	2、接口对象不能访问实现类中的属性。

	多态：
		go 语言通过接口模拟多态

		一个接口的实现：
			1、看成实现本身的类型，能够访问实现类中的属性和方法
			2、看成是对应的接口类型，那就只能够访问接口中的方法
		接口的用法：
			1、一个函数如果接受接口类型作为参数，那么实际上可以传入该接口的任意实现类型对象作为参数。
			2、定义一个类型为接口类型，实际上可以赋值为任意实现类的对象。

	鸭子类型：
*/
package main

import "fmt"

func main() {
	// 1、创建 Mouse 对象
	m1 := Mouse{"罗技小红"}
	fmt.Println("m1: ", m1)
	// 2、创建 Flashdisk 对象
	f1 := Flashdisk{"闪迪64GB"}
	fmt.Println("f1: ", f1)

	testInterface(m1)
	testInterface(f1)

	var usb USB
	usb = m1
	usb.start()
	usb.stop()
	//fmt.Println(usb.name) // 接口对象不能访问实现类中的属性

	f1.deleteData()
	usb = f1
	//usb.deleteData()  // 接口类型无法调用它的实现类型的方法

	var arr1 [3]USB
	arr1[0] = m1
	arr1[1] = f1
	fmt.Println("arr1: ", arr1) // [{罗技小红} {闪迪64GB} <nil>]
}

// 1、定义接口
type USB interface {
	start()
	stop()
}

// 2、实现类
type Mouse struct {
	name string
}

type Flashdisk struct {
	name string
}

func (m Mouse) start() {
	fmt.Println("鼠标，准备就绪，可以开始工作了...")
}

func (m Mouse) stop() {
	fmt.Println("鼠标，结束工作，可以拔掉了...")
}

func (f Flashdisk) start() {
	fmt.Println("U盘，准备就绪，可以开始读写数据了...")
}

func (f Flashdisk) stop() {
	fmt.Println("U盘，读写数据结束，可以安全拨出了...")
}

func (f Flashdisk) deleteData() {
	fmt.Println("删除 U 盘数据...")
}

// 3、测试方法
func testInterface(usb USB) {
	usb.start()
	usb.stop()
}
