package main

import "fmt"

func main() {
	fmt.Println("春景")

	// goto 的作用是无条件跳转：goto 后面是工程师自定义名字的标签，goto 的下一行代码直至自定义标签中间的代码全部不执行。
	goto Next

	fmt.Println("小楠")

Next:
	fmt.Println("张妍")
}
