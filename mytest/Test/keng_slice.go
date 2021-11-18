//切片是引用类型的数据
package main

import (
	"fmt"
)

func Pingpong(s []int) []int {
	s = append(s, 3)
	return s
}

func main() {
	s := make([]int, 0)
	fmt.Println(s)
	s = Pingpong(s)
	fmt.Println(s)
}
