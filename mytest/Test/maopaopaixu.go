package main

import "fmt"

func main() {
	a := [...]int{6, 1, 9, 3, 8}
	fmt.Println("The array a is: ", a)

	var len_a = len(a)

	//降序排列
	for i := 0; i < len_a; i++ {
		for j := i + 1; j < len_a; j++ {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	fmt.Println(a)
}
