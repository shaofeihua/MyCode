package main

import "fmt"
import "time"

func main() {
	i := 0
	// for 后面不写任何东西，则这个循环条件永远为真，这个循环就是死循环
	for {
		i++
		// 延时1秒
		time.Sleep(time.Second)

		if i == 5 {
			continue // continue 的作用是跳过本次循环，整个循环不结束，继续执行
		}

		fmt.Println(i)
	}
}
