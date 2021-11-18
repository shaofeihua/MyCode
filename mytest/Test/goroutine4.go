package main

import (
	"fmt"
)

func main() {
	c1 := make(chan bool)
	go func() {
		fmt.Println("Go Go Go!!!")
		c1 <- true
		close(c1) //因为后面有 for 循环加 range 对 c1 不停地迭代，所以必须有 close 的操作，，否则会发生 goroutine 死锁
	}()

	for v := range c1 { //range 一直对 c1 进行迭代，直到前面 go func() 对 c1 赋值之后，这里才会将 c1 的值打印出来，但是此时仍然没有退出 main 函数，因为这里是无限循环，仍然继续对 c1 进行迭代。所以前面必有一个关闭 c1 的动作，通知 range 不需要再进行迭代了。此时 main 函数才执行完成，退出。
		fmt.Println(v)
	}
}
