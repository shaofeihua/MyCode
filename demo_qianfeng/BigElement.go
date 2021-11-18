/*
	1、找到切片中最大的元素索引和值(打印) []int{9, 8, 19, 10, 2, 8}

	2、找到切片第二大(10)的元素索引和值(打印) []int{9, 8, 19, 10, 2, 8}

	3、将最大的元素移动到切片最末尾 []int{9, 8, 19, 10, 2, 8}

	4、将第二大的元素移动到切片倒数第二位 []int{9, 8, 19, 10, 2, 8}

	5、将切片进行排序，按小到大(冒泡排序)

	6、在切片中查找元素(二分查找) []int{2, 8, 9, 10, 19}
*/
package main

import "fmt"

func main() {
	s := []int{9, 8, 19, 10, 2, 8}

	// 重新创建一个切片，避免查找最大值时对原切片进行修改
	s1 := []int{}
	for _, v := range s {
		s1 = append(s1, v)
	}

	fmt.Println("s1 排序前：")
	fmt.Println("s: ", s)
	fmt.Println("s1:", s1)
	fmt.Println()
	//for i, v := range s {
	//	fmt.Println(i, v)
	//}

	// 对 s1 进行排序，通过排序后的顺序，找到最大值和第二大的值
	// 冒泡排序：降序
	for i := 0; i < len(s1)-1; i++ {
		for j := i + 1; j < len(s1); j++ {
			if s1[i] < s1[j] {
				s1[i], s1[j] = s1[j], s1[i]
			}
		}
	}

	fmt.Println("s1 排序后：")
	fmt.Println("s: ", s)
	fmt.Println("s1:", s1)
	fmt.Println()

	fmt.Printf("s 的元素的最大值是：%d\n", s1[0])
	fmt.Println()

	for i, _ := range s {
		if s[i] == s1[0] {
			fmt.Printf("s 中最大的元素索引和值分别是：%d, %d\n", i, s[i])
		}
	}

	//	从切片中，将最大值 19 移动到切片的末尾
}
