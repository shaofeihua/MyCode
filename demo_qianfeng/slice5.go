package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3, 4}
	s2 := make([]int, 0, 0)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

	// 将 s1 的元素都追加到 s2
	for _, v := range s1 {
		s2 = append(s2, v)
	}
	fmt.Println("s2:", s2)

	s1[0] = 100
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
}
