package main

import "fmt"

func main() {
	s := []int{9, 8, 19, 10, 2, 8}
	fmt.Println("切割前的 s: ", s)
	delOfValue(s, 10)
}

// 根据元素的值，删除切片中的元素
// 切片的元素类型为 int
func delOfValue(s []int, element int) []int {
	for i := 0; i < len(s); i++ {
		//for i, _ := range s {
		//这里使用
		if element == s[i] {
			fmt.Println(s[:i], s[i+1:])
			s = append(s[:i], s[i+1:]...)
			fmt.Println("切割后的 s: ", s)
		}
		//break
	}
	fmt.Println("返回的 s: ", s)
	fmt.Println()
	return s
}

// 根据元素的索引，删除切片中的元素
// 切片的元素类型为 int
