package main

import "fmt"

func main() {
	s1 := make([]int, 3, 6)
	fmt.Printf("%p\n", s1) //%p 打印内存地址

	s1 = append(s1, 5, 1, 8)
	fmt.Printf("%v %p\n", s1, s1)

	s1 = append(s1, 5, 1, 8)
	fmt.Printf("%v %p\n", s1, s1)
}
