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
			break // break 的作用是跳出循环。如果存在多层嵌套的循环，那么只跳出 break 所在的那一层循环，外层的循环仍然执行。
		}

		fmt.Println(i)
	}
}
