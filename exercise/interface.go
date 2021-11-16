package main

import (
	"fmt"
)

func main() {
	// 1、创建 Mouse 类型
	m1 := Mouse{"罗技小红"}
	fmt.Println(m1.name)
	//2、创建 FlashDisk
	f1 := FlashDisk{"闪迪64G"}
	fmt.Println(f1.name)

	fmt.Println("---------------------------------------")

	testInterface(m1)
	testInterface(f1)

	fmt.Println("---------------------------------------")

	var usb USB
	usb=m1
	usb.start()
	usb.stop()
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

type FlashDisk struct {
	name string
}

func (m Mouse) start() {
	fmt.Println(m.name, "鼠标就绪，可以开始工作了")
}

func (m Mouse) stop() {
	fmt.Println(m.name, "结束工作，鼠标可以安全退出了")
}

func (f FlashDisk) start() {
	fmt.Println(f.name, "U盘就绪，可以开始存储数据了")
}

func (f FlashDisk) stop() {
	fmt.Println(f.name, "结束工作，U盘可以安全弹出了")
}

// 3、测试方法
func testInterface(usb USB) { // usb=m1 usb=f1
	usb.start()
	usb.stop()
}
