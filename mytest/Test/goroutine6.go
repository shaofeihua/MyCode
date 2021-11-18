package main

import (
	"fmt"
)

func main() {
	c1 := make(chan bool, 1) //有缓存是异步的，无缓存是同步阻塞的。
	go func() {
		fmt.Println("Go Go Go!!!")
		<-c1
	}()
	c1 <- true //由于前面设置了缓存，所以此处 c1 赋值结束后程序就退出了，不管前面的 goroutine 是否完成了。所以这种情况下程序就不会输出前面想打印的内容。
}
