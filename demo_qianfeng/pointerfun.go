/*
	函数指针
		指向函数的指针
		go 语言中，function，默认看做一个指针，没有*

	指针函数
		一个函数，它的返回值是一个指针
*/
package main

import "fmt"

func main() {
	var a func()
	a = fun1
	a()                         // 我是 fun1()
	fmt.Println("fun1: ", fun1) // 0xc97b80
	fmt.Println("a: ", a)       // 0xc97b80
	//fmt.Println("fun1 的地址: ", &fun1) // cannot take the address of fun1
	fmt.Println("a 的地址: ", &a) // 0xc000006028

	arr1 := fun2()
	fmt.Printf("arr1 的类型: %T, 地址: %p, 数值: %v\n", arr1, &arr1, arr1) // [4]int, 0xc0000101c0, [1 2 3 4]

	arr2 := fun3()
	fmt.Printf("arr2 的类型: %T, 地址: %p, 数值: %v\n", arr2, &arr2, arr2) // *[4]int, 0xc000006038, &[5 6 7 8]
	fmt.Printf("指针 arr2 中存储的数组的地址: %p\n", arr2)                     // 指针 arr2 中存储的数组的地址: 0xc000010220
}

func fun3() *[4]int {
	arr := [4]int{5, 6, 7, 8}
	fmt.Printf("fun3 中 arr 的地址: %p\n", &arr) // fun3 中 arr 的地址: 0xc000010220
	return &arr
}

func fun2() [4]int {
	arr := [4]int{1, 2, 3, 4}
	return arr
}

func fun1() {
	fmt.Println("我是 fun1()")
}
