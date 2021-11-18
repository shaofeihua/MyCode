package main

import (
	"fmt"
)

func main() {
	var arr1 [4]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4

	fmt.Println(arr1)
	// 由于数组是定长的（即包含元素的个数固定不变），所以数组的长度和容量是相等的
	fmt.Println("数组的长度：", len(arr1)) // 长度是指容器中实际存储的数据量
	fmt.Println("数组的容量：", cap(arr1)) // 容量是指容器中能够存储的最大数据量

	// 修改数组中的数据
	arr1[0] = 5
	fmt.Println(arr1, arr1[0])

	// 数组的其他创建方式
	var a [4]int // 等同于：var a = [4] int
	// 默认存储 0 值
	fmt.Println(a)

	// 直接赋值
	var b = [4]int{1, 2, 3, 4}
	fmt.Println(b)

	// 长度不足，自动补齐 0 值
	var c = [5]int{1, 2, 3, 4}
	fmt.Println(c)

	// 指定下标位置对应的元素的值
	var d = [5]int{1: 11, 3: 22}
	fmt.Println(d)

	// 数据类型为字符串
	var e = [5]string{"taotao", "xiaonan", "yanyan"}
	fmt.Println(e)

	// 自动推导数组的长度
	f := [...]int{1, 2, 3, 4, 5}
	fmt.Println(f)
	fmt.Println(len(f))

	// 既不指定数组的长度，也不给全所有元素的值
	g := [...]int{1: 23, 8: 88}
	fmt.Println(g)
	fmt.Println(len(g))
}
