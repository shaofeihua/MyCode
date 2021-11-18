// 闭包捕获外部变量
package main

import (
	"fmt"
)

func main() {
	a := 10
	b := "chunjing"

	func() {
		a = 56       // 注意不是 a:=56
		b = "yanyan" // 注意不是 b:="yanyan"
		fmt.Println("闭包内a =", a)
		fmt.Println("闭包内b =", b)
	}()

	fmt.Println("闭包外a =", a)
	fmt.Println("闭包外b =", b)
}
