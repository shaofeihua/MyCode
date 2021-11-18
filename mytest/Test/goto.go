package main

import "fmt"

func main() {
	for {
		for a := 1; a < 10; a++ {
			if a > 3 {
				fmt.Println(a)
				goto LABLE1 //因为标签的位置被放在了后面，此时使用 break 和 goto 的效果是一样的。注意：标签尽量放在 goto 后面，尽量避免形成死循环。
			}
		}
	}
LABLE1:
	fmt.Println("OK")
}
