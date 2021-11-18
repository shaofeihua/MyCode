package main()

import (
	"fmt"
)

func main() {
	a:=[5]int{3,1,8,6,5}

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
