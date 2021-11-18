package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(reflect.TypeOf(a))

	//s1 和 s2 的长度都未超过数组的长度，也就是说内存也未超过数组的边界
	s1 := a[2:5]
	fmt.Println(reflect.TypeOf(s1))
	s2 := a[1:3]
	fmt.Printf("%p %v\n", a, a)
	fmt.Printf("%p %v\n", s1, s1)
	fmt.Printf("%p %v\n", s2, s2)
	fmt.Printf("-------------------------------\n")

	//在 s1 和 s2 的长度都未超过数组长度的情况下，修改公共元素的值
	s1[0] = 9 //s1 的索引号为 0 的元素是数组 a 的元素 3 ，是 s1 和 s2 的公共元素
	fmt.Printf("The first modified s1 is: %p %v\n", s1, s1)
	fmt.Printf("The latest value of s2 is: %p %v\n", s2, s2) //s1 的索引号为 0 的元素是数组 a 的元素 3 ，是 s1 和 s2 的公共元素。所以当在 s1 中修改了这个公共元素的值之后，s2 中的该元素也跟着一起改变了
	fmt.Printf("%p %v\n", a, a)                              //同理，数组中的这个元素也跟着改变了
	fmt.Printf("-------------------------------\n")

	//在切片 s2 长度超过数组 a 的长度的情况下，修改 s1 中的公共元素的值，s2 的对应元素不再改变了。因为 s2 已经变成了一个全新的切片，内存位置也已经发生改变
	s2 = append(s2, 5, 1, 8)
	s1[0] = 888 //在执行此语句之前，元素 “9” 仍然是数组 a 和 切片 s1、s2 的公共元素。上一个语句给 s2 追加了几个元素后，s2 的长度已经超出了数组 a 的长度，此时再来通过 s1 修改“公共元素 9”
	fmt.Printf("The second modified s1 is: %p %v\n", s1, s1)
	fmt.Printf("The append value of s2 is: %p %v\n", s2, s2) //s2 中的“公共元素 9 ”并没有跟着 s1 发生改变
	fmt.Printf("%p %v\n", a, a)                              //数组 a 中的“公共元素 9” 跟着 s1 发生了改变
}
