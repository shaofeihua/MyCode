package main

import (
	"fmt"
)

func main() {
	/*
		数组指针：
			1、首先是一个指针
			2、存贮的是一个数组的地址
	*/
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println("arr1: ", arr1)

	// 创建一个指针，存储这个数组的地址
	var p1 *[4]int
	p1 = &arr1
	fmt.Println(p1)  // &[1 2 3 4] 。打印出来的不是 arr1 的内存地址
	fmt.Println(&p1) // p1 本身的地址
	fmt.Println(*p1)

	// 根据指针来操作驻数组
	(*p1)[0] = 100
	fmt.Println("arr1: ", arr1)

	// 简化写法
	p1[1] = 200
	fmt.Println("arr1: ", arr1)

	/*
		指针数组：
			1、首先是一个数组
			2、数组里的元素都是指针
	*/
	a := 1
	b := 2
	c := 3
	d := 4
	arr2 := [4]int{a, b, c, d}
	arr3 := [4]*int{&a, &b, &c, &d}

	fmt.Println("a 的地址是：", &a)
	fmt.Println("b 的地址是：", &b)
	fmt.Println("c 的地址是：", &c)
	fmt.Println("d 的地址是：", &d)
	fmt.Println("arr2: ", arr2)
	fmt.Println("arr3: ", arr3)

	arr2[0] = 100
	fmt.Println("a: ", a)       // a 的值不变，还是 1
	fmt.Println("a 的地址是：", &a)  // a 的地址不变
	fmt.Println("arr2: ", arr2) // [100,2,3,4]

	*arr3[0] = 200
	fmt.Println("arr3: ", arr3)
	fmt.Println("arr2: ", arr2)
	fmt.Println("a: ", a)
}
