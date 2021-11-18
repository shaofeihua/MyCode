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

	//声明一个接口类型
	var s1 Shape
	s1 = t1
	fmt.Println(s1.peri())
	fmt.Println(s1.area())
	//fmt.Println(s1.a, s1.b, s1.c) // 不能这样使用的原因是：接口无法调用实现类中的字段

	fmt.Println("-----------------------------------------")

	testShape(t1)
	testShape(c1)
	testShape(s1)

	fmt.Println("-----------------------------------------")

	getType(t1)
	getType(c1)
	getType(s1)
}

func getType(s Shape) {
	//断言
	if ins, ok := s.(Triangle); ok {
		fmt.Println("是三角形，三条边分别是：", ins.a, ins.b, ins.c)
	} else if ins, ok := s.(Circle); ok {
		fmt.Println("是圆形，半径是：", ins.radius)
	}
}
func testShape(s Shape) {
	fmt.Printf("周长：%.2f, 面积：%.2f\n", s.peri(), s.area())
}

//1、定义一个接口
type Shape interface {
	peri() float64 // 形状的周长
	area() float64 // 形状的面积
}

// 2、定义实现类: 三角形
type Triangle struct {
	a, b, c float64
}

// 求三角形的周长
func (t Triangle) peri() float64 {
	return t.a + t.b + t.c
}

// 求三角形的面积
func (t Triangle) area() float64 {
	p := t.peri() / 2
	s := math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
	return s
}

// 圆形
type Circle struct {
	radius float64 // 半径
}

func (c Circle) peri() float64 {
	return c.radius * 2 * math.Pi
}

func (c Circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}
