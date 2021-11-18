// 切片
package main

import (
	"fmt"
)

func main() {
	// 1.数组
	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr) // [1 2 3 4]

	fmt.Println("----------------------")

	// 2.切片
	var s1 []int
	fmt.Println(s1) // []

	s2 := []int{1, 2, 3, 4}
	fmt.Println(s2) // [1 2 3 4]

	fmt.Printf("%T,%T\n", arr, s2) // [4]int,[]int

	s3 := make([]int, 3, 8)
	fmt.Println(s3)
	fmt.Printf("s3 的长度是：%d, 容量是：%d\n", len(s3), cap(s3))

	s3[0] = 1
	s3[1] = 2
	s3[2] = 3
	fmt.Println(s3)
	// fmt.Println(s3[3]) // panic: runtime error: index out of range

	fmt.Println("----------------------")

	// append()，专门用于向切片的尾部追加元素
	s4 := make([]int, 0, 5)
	fmt.Println(s4) // []
	s4 = append(s4, 1, 2)
	fmt.Println(s4) // [1 2]
	s4 = append(s4, 3, 4, 5, 6, 7)
	fmt.Println(s4)                                      // [1 2 3 4 5 6 7]
	fmt.Printf("s4 的长度是：%d, 容量是：%d\n", len(s4), cap(s4)) // s4 的长度是：7, 容量是：10
	// 将一个切片中的元素添加给另一个切片
	s4 = append(s4, s3...)
	fmt.Println(s4)                                      // [1 2 3 4 5 6 7 1 2 3]
	fmt.Printf("s4 的长度是：%d, 容量是：%d\n", len(s4), cap(s4)) // s4 的长度是：10, 容量是：10
}
