package main

import (
	"fmt"
)

func main() {
	var a func()
	a = fun1
	a()

	// go 语言中，function 默认看做一个指针，所以它的值是一个地址

	/*
		注意！！！！！！！
		此时，arr1 := fun2() 和 arr1 := fun2 不是一回事
	*/
	arr1 := fun2() // arr1 是数组，arr1 = arr
	fmt.Printf("arr1 的类型是：%T，值是：%v，地址是：%p\n", arr1, arr1, &arr1)

	arr11 := fun2 // arr11 是函数
	fmt.Printf("arr11 的类型是：%T，值是：%v，地址是：%p\n", arr11, arr11, &arr11)

	arr2 := fun3() // arr2 是个指针，指向 arr 的地址
	fmt.Printf("arr2 的类型是：%T，值是：%v，地址是：%p\n", arr2, arr2, &arr2)
}

func fun1() {
	fmt.Println("fun1()...")
}
func fun2() [4]int {
	arr := [4]int{1, 2, 3, 4}
	fmt.Println("arr 在 fun2() 中的值是：", arr)
	fmt.Println("arr 在 fun2() 中的地址是：", &arr)
	return arr
}

func fun3() *[4]int {
	arr := [4]int{5, 6, 7, 8}
	fmt.Println("arr 在 fun3() 中的地址是：", &arr)
	return &arr
}
