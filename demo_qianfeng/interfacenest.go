/*
	接口嵌套
*/
package main

import "fmt"

func main() {
	var cat Cat = Cat{}
	cat.test1()
	cat.test2()
	cat.test3()
	fmt.Printf("cat 的类型是：%T\n", cat) // main.Cat

	fmt.Println("-------------------------------")

	var a1 A = cat
	a1.test1() // a1 只能调用 test1()
	var b1 B = cat
	b1.test2() // b1 只能调用 test2()
	var c1 C = cat
	c1.test1() // c1 可以调用 test1()、test2()、test3()
	c1.test2()
	c1.test3()

	fmt.Println("-------------------------------")

	//var c2 C=a1 // 不能这样赋值。因为 a1 没有实现 C 的所有方法
	var a2 A = c1
	a2.test1()

}

type A interface {
	test1()
}

type B interface {
	test2()
}

type C interface {
	A
	B
	test3()
}

type Cat struct { // 如果 Cat 想实现接口 C，那不只要实现接口 C 的方法，还要实现接口 A 和 B 的方法
}

func (c Cat) test1() {
	fmt.Println("test1()...")
}

func (c Cat) test2() {
	fmt.Println("test2()...")
}

func (c Cat) test3() {
	fmt.Println("test3()...")
}
