package main

import "fmt"

func main() {
	a2 := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(a2)
	fmt.Println(a2[0])
	fmt.Println(a2[1][1])
	fmt.Println("a2 的长度是：", len(a2))
	fmt.Println("a2[0] 的长度是：", len(a2[0]))
	fmt.Printf("a2 的内存地址是：%p\n", &a2)
	fmt.Printf("a2[0] 的内存地址是：%p\n", &a2[0])
	fmt.Printf("a2[1] 的内存地址是：%p\n", &a2[1])
	fmt.Printf("a2[2] 的内存地址是：%p\n", &a2[2])

	// 遍历二维数组
	for i := 0; i < len(a2); i++ {
		for j := 0; j < len(a2[i]); j++ {
			fmt.Print(a2[i][j], "\t")
		}
		fmt.Println()
	}

	fmt.Println("-------------------------------------")

	// 遍历二维数组：使用 range
	for _, v1 := range a2 {
		for _, v2 := range v1 {
			fmt.Print(v2, "\t")
		}
		fmt.Println()
	}
}
