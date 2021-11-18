// 接口断言
package main

import (
	"fmt"
	"math"
)

func main() {
	var t1 Triangle = Triangle{3, 4, 5}
	fmt.Println(t1.peri())
	fmt.Println(t1.area())
	fmt.Println(t1.a, t1.b, t1.c)

	var c1 Circle = Circle{4}
	fmt.Println(c1.peri())
	fmt.Println(c1.area())
	fmt.Println(c1.radius)

	// 定义一个接口
	var s1 Shape
	s1 = t1
	fmt.Println(s1.peri())
	fmt.Println(s1.area())

	var s2 Shape
	s2 = c1
	fmt.Println(s2.peri())
	fmt.Println(s2.area())

	testShape(t1)
	testShape(c1)
	testShape(s1)
	testShape(s2)

	getType(t1)
	getType(c1)
	getType(s1)
	getType(s2)

	var t2 *Triangle = &Triangle{a: 3, b: 4, c: 2}
	fmt.Printf("t2 的类型是：%T，地址是：%p，数据地址是：%p\n", t2, &t2, t2)
	getType(t2)

	getType2(t2)
	getType2(t1)
}

func getType2(s Shape) {
	switch ins := s.(type) {
	case Triangle:
		fmt.Println("三角形", ins.a, ins.b, ins.c)
	case Circle:
		fmt.Println("圆形", ins.radius)
	case *Triangle:
		fmt.Println("三角形结构的指针", ins.a, ins.b, ins.c)
	}
}
func getType(s Shape) {
	// 断言
	if ins, ok := s.(Triangle); ok {
		// fmt.Printf("是三角形，三边是：%.2f，%.2f，%.2f\n", ins.a, ins.b, ins.c)
		fmt.Println("是三角形，三边是：", ins.a, ins.b, ins.c)
	} else if ins, ok := s.(Circle); ok {
		fmt.Println("是圆形，半径是：", ins.radius)
	} else if ins, ok := s.(*Triangle); ok {
		fmt.Printf("ins 的类型是：%T，地址是：%p，数据的地址是：%p\n", ins, &ins, ins)
		fmt.Printf("s 的类型是：%T，地址是：%p，数据的地址是：%p\n", s, &s, s)
	} else {
		fmt.Println("我也不知道是什么图形了")
	}
}

func testShape(s Shape) {
	fmt.Printf("周长是：%.2f，面积是：%.2f\n", s.peri(), s.area())
}

// 1、定义一个接口，包含两个方法，分别是求图形的周长和面积。
type Shape interface {
	peri() float64 // 周长，有返回值
	area() float64 // 面积，有返回值
}

// 2、定义实现接口的类：三角形，圆形
type Triangle struct {
	// 三角形，有三个边长
	a, b, c float64
}

type Circle struct {
	// 圆形，半径
	radius float64
}

// 3、定义一个方法求三角形或圆形的周长及面积。有 float64 的返回值
// 三角形的周长
func (t Triangle) peri() float64 {
	// return a + b + c // 错误！！！
	return t.a + t.b + t.c
}

// 三角形的面积
func (t Triangle) area() float64 {
	// 根据海伦公式，先求三角形周长的一半
	// P := (t.a + t.b + t.c) / 2
	p := t.peri() / 2
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}

// 圆形的周长
func (c Circle) peri() float64 {
	// return c.radius * 2 * 3.14
	return c.radius * 2 * math.Pi
}

// 圆的面积
func (c Circle) area() float64 {
	// return c.radius * *2 * math.Pi
	return math.Pow(c.radius, 2) * math.Pi
}
