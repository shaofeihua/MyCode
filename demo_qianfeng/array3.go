package main

import (
	"fmt"
)

func main() {
	var arr1 [4]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4
	// arr1[4] = 5
	fmt.Println(arr1)
	fmt.Println(arr1[0])
	fmt.Println(len(arr1))
	fmt.Println(cap(arr1))

	var a = [4]int{2, 4, 6, 8}
	fmt.Println(a)

	var a1 = [6]int{1, 3, 5}
	fmt.Println(a1)
	fmt.Println(len(a1))
	fmt.Println(cap(a1))

	var a2 = [5]int{0: 3, 3: 6, 2: 2}
	fmt.Println(a2)

	var a3 = [3]string{}
	fmt.Println("a3", a3)

	var a4 = [3]string{"rose", "王二狗", "哪吒"}
	fmt.Println("a4", a4)

}
