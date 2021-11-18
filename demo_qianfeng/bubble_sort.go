// 数组的排序：冒泡排序
// 展示冒泡排序的具体过程
package main

import (
	"fmt"
)

func main() {
	arr := [5]int{15, 23, 8, 10, 7}
	// 升序：[7,8,10,15,23]
	// 降序：[23,15,10,8,7]
	// 第一轮排序
	for j := 0; j < 4; j++ {
		if arr[j] > arr[j+1] {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
	fmt.Println(arr) // [15 8 10 7 23]

	// 第二轮排序
	for j := 0; j < 3; j++ {
		if arr[j] > arr[j+1] {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
	fmt.Println(arr) // [8 10 7 15 23]

	// 第三轮排序
	for j := 0; j < 2; j++ {
		if arr[j] > arr[j+1] {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
	fmt.Println(arr) // [8 7 10 15 23]

	// 第四轮排序
	for j := 0; j < 2; j++ {
		if arr[j] > arr[j+1] {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
	fmt.Println(arr) // [7 8 10 15 23]
}
