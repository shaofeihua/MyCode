/*
深拷贝和浅拷贝

深拷贝：
	拷贝的是数据本身。
	值类型的数据，默认都是深拷贝：array，int，float，string，bool，struct

浅拷贝：
	拷贝的数据的地址。
	引用类型的数据，默认都是浅拷贝：slice，map
	导致多个变量指向同一块内存。

切片是引用类型的数据，所以拷贝的是地址。
*/
package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3, 4}
	s2 := make([]int, 0) // len:0, cap:0
	for i := 0; i < len(s1); i++ {
		s2 = append(s2, s1[i])
	}
	fmt.Println(s1) // [1 2 3 4]
	fmt.Println(s2) // [1 2 3 4]

	fmt.Println("------------------------")

	s1[0] = 100
	fmt.Println(s1) // [100 2 3 4]
	fmt.Println(s2) // [1 2 3 4]

	fmt.Println("------------------------")

	// copy()
	s3 := []int{7, 8, 9}
	fmt.Println(s2) // [1 2 3 4]
	fmt.Println(s3) // [7 8 9]

	// 将 s2 的元素拷贝到 s3 中
	copy(s3, s2)
	fmt.Println(s2)
	fmt.Println(s3)

}
