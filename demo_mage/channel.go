package main

import "fmt"

/*
	channel，通道
		是引用类型的数据
*/

func main() {
	var a chan int
	fmt.Printf("%T,%v\n", a, a)

	if a == nil {
		fmt.Println("channel a 为 nil，不能使用，即将创建通道")
		a = make(chan int)
		fmt.Println(a) // 0xc00003e060
	}

	// 内存地址跟前面一样，说明 channel 是引用类型的数据
	test1(a) // 0xc00003e060
}

func test1(ch chan int){
	fmt.Printf("%T,%v\n",ch,ch)
}