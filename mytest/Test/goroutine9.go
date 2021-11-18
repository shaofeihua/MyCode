//使用 select 处理多个 channel

package main

import (
	"fmt"
)

func main() {
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool, 2)

	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c1", v)
			case v, ok := <-c2:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c2", v)
			}
		}
	}()

	c1 <- 1
	c2 <- "xiaonan"
	c1 <- 586
	c2 <- "meng"
	c1 <- 8611
	c2 <- "yes"

	close(c1)
	close(c2)

	for i := 0; i < 2; i++ {
		<-o
	}
}
