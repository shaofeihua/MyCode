/*
	指针 pointer
		存储了另一个变量的内存地址的变量
*/
package main

import "fmt"

func main() {
	// 1、定义一个 int 类型的变量
	a := 10
	fmt.Println("a 的数值是：", a)    // 10
	fmt.Printf("a 的类型是：%T\n", a) // int
	fmt.Println("a 的地址是：", &a)   // 0xc00000a088

	//	2、创建一个指针变量，用于存储变量 a 的地址
	var p1 *int
	fmt.Println("p1: ", p1) // <nil>，空指针
	p1 = &a
	fmt.Println("p1: ", p1)             // 0xc00000a088
	fmt.Println("p1 的地址：", &p1)         // 0xc000006030
	fmt.Println("p1 指向的内存地址对应的值：", *p1) // 10

	//	3、操作变量，更改值，变量地址不变
	a = 100
	fmt.Println("a 的数值是：", a)  // 100
	fmt.Println("a 的地址是：", &a) // 0xc00000a088

	//	4、操作指针，更改变量的值
	*p1 = 200
	fmt.Println("a 的数值是：", a) // 200

	//	指针的指针
	p2 := &p1
	fmt.Printf("a 的类型：%T，p1 的类型：%T，p2 的类型：%T\n", a, p1, p2) // int, *int, **int

	fmt.Println("p2 的数值：", p2)                                              // 0xc000006030
	fmt.Printf("p2 的地址：%p\n", &p2)                                          // 0xc000006038
	fmt.Println("p2 存储的内存地址，对应的数值，就是 p1 的内存地址，对应的值，即 a 的内存地址：", *p2)        // 0xc00000a088
	fmt.Println("p2 存储的内存地址，对应的数值，就是 p1 的内存地址，对应的值，再获取其对应的值，即 a 的值：", **p2) // 200
}
