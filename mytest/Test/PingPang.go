/*
思路：
1、先写一个函数 Pingpang() 作为 goroutine 启动。
2、定义一个全局变量的 channel 用于两个 goroutine 之间传递消息。
3、在 main 函数中启动 goroutine 准备接受消息。
4、在 Pingpang() 定义一个 for 无限循环，用于读取 main 传递过来的消息。
5、main 和 Pingpang 的消息传递次数在 main 函数中使用 for 循环来限制。
6、Pingpang() 中“先接收再发送”，也就是先打印（读取） channel 传递的消息，再给 channel 传入一个新消息。而主函数刚好相反，“先发送再接收”，也就是先给 channel 传入一个消息，再打印（读取） channel 的值。


*/

package main

import (
	"fmt"
)

var c chan string

func main() {
	c = make(chan string)
	go Pingpang()
	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("From main: Hello, #%d", i)
		fmt.Println(<-c)
	}
}

func Pingpang() {
	i := 0
	for {
		fmt.Println(<-c)
		c <- fmt.Sprintf("From Pingpang: Hi, #%d", i)
		i++
	}
}
