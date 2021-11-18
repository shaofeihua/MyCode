package main

import (
	"fmt"
)

func main() {
	a := []int{2, 1, 2, 5, 6, 3, 4, 5, 2, 3, 9}
	z := Rm_duplicate(&a)
	fmt.Println(a, z)
}

func Rm_duplicate(list *[]int) []int {
	var x []int = []int{}
	for _, i := range *list {
		if len(x) == 0 {
			x = append(x, i)
		} else {
			for k, v := range x {
				if i == v {
					break
				}
				if k == len(x)-1 {
					x = append(x, i)
				}
			}
		}
	}
	return x
}
