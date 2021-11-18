package main

import (
	"fmt"
)

var c chan int

func main() {
	c <- 10
	go MR()
}

func MR() {
	for i := 0; i < 10; i++ {
		v := make(chan int)
		select {
		case v := <-c:
			if i%2 == 0 {
				fmt.Println("G1 执行第%d读取完毕。开始发送。", i+1, v)
				c += 1
			} else {
				fmt.Println("G2 执行第%d读取完毕。开始发送。", i+1, v)
				c += 1
			}
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout")
		}
	}
}
