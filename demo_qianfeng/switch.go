package main

import "fmt"

func main() {
	// var num int
	// fmt.Println("请选择您要去的楼层：")
	// fmt.Scan(&num)

	// switch num {
	// case 1:
	// 	fmt.Println("按下的是 1 楼")
	// case 2:
	// 	fmt.Println("按下的是 2 楼")
	// case 3:
	// 	fmt.Println("按下的是 3 楼")
	// case 4:
	// 	fmt.Println("按下的是 4 楼")
	// default:
	// 	fmt.Printf("按下的是 %d 楼", num)
	// }

	// score := 99
	var score int
	fmt.Println("请输入考生的分数：")
	fmt.Scan(&score)

	switch {
	case score >= 90:
		fmt.Println("优秀")

	case score >= 80:
		fmt.Println("良好")

	case score >= 70:
		fmt.Println("一般")

	case score >= 60:
		fmt.Println("及格")

	case score < 60:
		fmt.Println("不及格")
	}
}
