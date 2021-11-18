//冒泡排序
package main

import (
	"fmt"
)

func main() {
	a := [5]int{1, 9, 5, 8, 6}
	fmt.Println(a)

	num := len(a)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}

	fmt.Println(a)
}
