package main

import (
	"fmt"
)

func main() {
	c1 := make(chan bool)
	go func() {
		fmt.Println("Go Go Go!!!")
		c1 <- true
	}()
	<-c1 //当 go 关键字启动一个 goroutine 之后，main 函数执行已经执行到此处。只要前面的函数未执行到 c1 <- true ，也就是说未给 c1 传入值，此时 c1 的读取就一直处于等待状态。
}
