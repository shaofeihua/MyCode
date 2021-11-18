package main

import "fmt"

func main() {
	/*
		接口的嵌套
	*/
	var cat Cat = Cat{}
	cat.test1()
	cat.test2()
	cat.test3()

	fmt.Println("-------------------------------")

	// a1只能调用test1()
	var a1 A = cat
	a1.test1()

	fmt.Println("-------------------------------")

	// c1可以调用三个方法
	var c1 C = cat
	c1.test1()
	c1.test2()
	c1.test3()
}

type A interface {
	test1()
}

type B interface {
	test2()
}

//接口C继承了接口A和接口B
type C interface {
	A
	B
	test3()
}

type Cat struct {
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
