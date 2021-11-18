//map 嵌套时，每一级的 map 都需要单独的进行初始化
package main

import (
	"fmt"
)

var c = make(chan bool)

func main() {
	go Go()
	<-c //main 函数会等待，不会马上退出，直到读取到 c 的内容。

}

func Go() {
	fmt.Println("666，Go Go Go !!!")
	c <- true //Go() 执行完后，通过 c 向 main 函数传递信号，告诉主函数“我”已经执行完了。
}
