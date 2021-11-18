package main

import (
	"fmt"
)

func main() {
	var s1 []int
	fmt.Println(s1)

	s2 := []int{1, 2, 3, 4}
	fmt.Println(s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
	fmt.Printf("%T\n", s2)

	// 使用 make 函数创建切片
	s3 := make([]int, 3, 8)
	fmt.Println(s3, len(s3), cap(s3))
	// s3[4] = 4
	// fmt.Println(s3)
	// 切片的长度已经规定了，就不能使用 s3[4] 这种方式追加元素，而要使用 append 函数

	// append
	s3 = append(s3, 1, 2, 3)
	fmt.Println(s3)
	s3 = append(s3, 4, 5, 6, 7, 8, 9)
	fmt.Println(s3)
	// 把一个切片中的元素加入到另一个切片中
	s4 := []int{55, 66, 77}
	s3 = append(s3, s4...)
	fmt.Println(s3)
}
