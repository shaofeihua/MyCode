package main

import (
	"fmt"
)

func main() {
	// 1. 数据类型
	num := 10
	fmt.Printf("%T\n", num)

	fmt.Println("-----------------------")

	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [3]float64{2.15, 3.18, 6.19}
	arr3 := [4]int{5, 6, 7, 8}
	arr4 := [4]string{"hello", "world"}

	fmt.Printf("%T\n", arr1) // [4]int
	fmt.Printf("%T\n", arr2) // [3]float64
	fmt.Printf("%T\n", arr3) // [4]int
	fmt.Printf("%T\n", arr4) // [4]string
	// 说明：每个数组都有自己的数据类型
	// 数据的数据类型是：[size]type

	fmt.Println("-----------------------")

	// 2. 赋值
	num2 := num            // 值传递。传递的是 num 的值的副本（备份）
	fmt.Println(num, num2) // 运行结果：10 10
	num2 = 20
	fmt.Println(num, num2) // 运行结果：10 20 。修改 num2 的值对 num 没影响

	fmt.Println("-----------------------")

	// 数组的赋值
	arr5 := arr1
	fmt.Println(arr1) // [1 2 3 4]
	fmt.Println(arr5) // [1 2 3 4]
	// 说明：数据也是值传递

	arr5[0] = 100
	fmt.Println(arr1) // [1 2 3 4]
	fmt.Println(arr5) // [100 2 3 4]
	// 类似 int 类型的变量，改变数组 arr5 的其中一个数据，对数组 arr1 没有影响

	fmt.Println("-----------------------")

	a := 3
	b := 4
	fmt.Println(a == b)       // false
	fmt.Println(arr1 == arr5) // false
	fmt.Println(arr1 == arr2) // 53:19: invalid operation: arr1 == arr2 (mismatched types [4]int and [3]float64)
	// 说明：不同类型的数组不能进行比较。数组的比较必须长度和里面数值的类型都一样
}
