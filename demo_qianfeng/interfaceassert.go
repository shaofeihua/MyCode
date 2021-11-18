/*
	接口断言
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var t1 Triangle = Triangle{3, 4, 5}
	fmt.Println("t1 的三条边长为：", t1.a, t1.b, t1.c) // 三条边长为： 3 4 5
	fmt.Println("t1 的周长为：", t1.perimeter())     // 12
	fmt.Println("t1 的面积为：", t1.area())          // 6

	var c1 Cricle = Cricle{4}
	fmt.Println("c1 的半径为：", c1.radius)      // 4
	fmt.Println("c1 的周长为：", c1.perimeter()) // 25.132741228718345
	fmt.Println("c1 的面积为：", c1.area())      // 50.26548245743669

	var s1 Shape
	s1 = t1
	//fmt.Println(s1.perimeter) // 0x9e8ec0
	fmt.Println("s1 的周长：", s1.perimeter()) // 12
	//fmt.Println(s1.area) // 0x9e8f40
	fmt.Println("s1 的周长：", s1.area()) // 6

	var s2 Shape
	s2 = c1
	fmt.Println("s2 的周长：", s2.perimeter()) // 25.132741228718345
	fmt.Println("s2 的面积：", s2.area())      // 50.26548245743669

	testShape(t1) // 周长是：12.00，面积是：6.00
	testShape(c1) // 周长是：25.13，面积是：50.27
	testShape(s1) // 周长是：12.00，面积是：6.00

	getType(t1) // 是三角形，三条边分别为： 3 4 5
	getType(c1) // 是圆形，半径是： 4

	var t2 *Triangle = &Triangle{3, 4, 2}
	fmt.Printf("t2 的类型是：%T，自己的内存地址：%p，值是：%p\n", t2, &t2, t2) // t2 的类型是：*main.Triangle，自己的内存地址：0xc000006030，值是：0xc0000121f8
	getType(t2)

	getType2(t1) // 是三角形，三条边是： 3 4 5
	getType2(c1) // 是圆形，半径是： 4
	getType2(t2) // 是三角形结构体指针，三条边是： 3 4 2
}

func getType2(s Shape) {
	// 断言
	switch instance := s.(type) {
	case Triangle:
		fmt.Println("是三角形，三条边是：", instance.a, instance.b, instance.c)
	case Cricle:
		fmt.Println("是圆形，半径是：", instance.radius)
	case *Triangle:
		fmt.Println("是三角形结构体指针，三条边是：", instance.a, instance.b, instance.c)
	case *Cricle:
		fmt.Println("是圆形结构体指针，半径是：", instance.radius)
	}
}

func getType(s Shape) {
	//	断言
	if instance, ok := s.(Triangle); ok {
		fmt.Println("是三角形，三条边分别为：", instance.a, instance.b, instance.c)
	} else if instance, ok := s.(Cricle); ok {
		fmt.Println("是圆形，半径是：", instance.radius)
	} else if instance, ok := s.(*Triangle); ok {
		fmt.Printf("instance 的类型是：%T，自己的内存地址：%p，值是：%p\n", instance, &instance, instance) // instance 的类型是：*main.Triangle，自己的内存地址：0xc000006038，值是：0xc0000121f8
		fmt.Printf("s 的类型是：%T，自己的内存地址：%p，值是：%p\n", s, &s, s)                             // s 的类型是：*main.Triangle，自己的内存地址：0xc000034260，值是：0xc0000121f8
	} else {
		fmt.Println("我也不知道是什么形状...")
	}
}

func testShape(s Shape) {
	fmt.Printf("周长是：%.2f，面积是：%.2f\n", s.perimeter(), s.area())
}

// 1、定义一个接口：形状
type Shape interface {
	// 如果方法有返回值，这里也不能忘了写，例如 float64
	perimeter() float64 // 周长
	area() float64      // 面积
}

// 2、定义实现类：三角形
type Triangle struct {
	a, b, c float64 // 三角形的三条边
}

// 3、方法
// 计算三角形的周长
func (t Triangle) perimeter() float64 {
	return t.a + t.b + t.c
}

// 计算三角形的面积
func (t Triangle) area() float64 {
	// 海伦公式：根据三角形边长计算面积
	p := t.perimeter() / 2 // 调用前面的方法，先求三角形周长的一半
	s := math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
	return s
}

// 定义实现类：圆形
type Cricle struct {
	radius float64 // 半径
}

func (c Cricle) perimeter() float64 {
	return math.Pi * 2 * c.radius
}

func (c Cricle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}
