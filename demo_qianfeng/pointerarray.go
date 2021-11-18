/*
	数组指针
		首先是一个指针，存储的是一个数组的内存地址
		*[4]Type

	指针数组
		首先是一个数组，存储的元素的数据类型是指针
		[4]*Type
*/

package main

import "fmt"

func main() {
	//	1、创建一个普通的数组
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println("arr1: ", arr1)           // [1 2 3 4]
	fmt.Printf("arr1 的内存地址: %p\n", &arr1) // 0xc0000101a0

	//	2、创建一个指针，存储数组的内存地址，即 数组指针
	var p1 *[4]int
	p1 = &arr1
	fmt.Println("p1: ", p1)         // &[1 2 3 4]
	fmt.Printf("p1: %p\n", p1)      // 0xc0000101a0，数据 a1 的内存地址
	fmt.Printf("p1 的地址: %p\n", &p1) // 0xc000006030，指针 p1 自己的内存地址

	//	根据数组指针来操作数组
	(*p1)[0] = 100
	fmt.Println("arr1: ", arr1) // [100 2 3 4]

	p1[0] = 200                 // 简化的写法
	fmt.Println("arr1: ", arr1) // [200 2 3 4]

	//	指针数组
	a, b, c, d := 1, 2, 3, 4
	arr2 := [4]int{a, b, c, d}
	arr3 := [4]*int{&a, &b, &c, &d}
	fmt.Println("arr2: ", arr2) // [1 2 3 4]
	fmt.Println("arr3: ", arr3) // [0xc00000a0e0 0xc00000a0e8 0xc00000a0f0 0xc00000a0f8]

	arr2[0] = 100
	fmt.Println("arr2: ", arr2) // [100 2 3 4]
	fmt.Println("a: ", a)       // 1

	*arr3[0] = 200
	fmt.Println("arr3: ", arr3) // [0xc00000a0e0 0xc00000a0e8 0xc00000a0f0 0xc00000a0f8]
	fmt.Println("a: ", a)       // 200

	b = 1000
	fmt.Println("arr2: ", arr2) // [100 2 3 4]
	fmt.Println("arr3: ", arr3) // [0xc00000a0e0 0xc00000a0e8 0xc00000a0f0 0xc00000a0f8]
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i]) // 100 2 3 4
	}
	for i := 0; i < len(arr3); i++ {
		fmt.Println(*arr3[i]) // 200 1000 3 4
	}
}
