package main

import "fmt"

func main() {
	a := [5]int{}
	a[2] = 8
	fmt.Println(a)

	b := new([5]int)
	b[3] = 6
	fmt.Println(b)

	c := [2][5]int{ //多维数组。数组有2个元素，其中每个元素都是包含5和元素的数组
		{1, 2, 3, 4, 5},
		{8, 6, 8, 6}}
	fmt.Println(c)
}
