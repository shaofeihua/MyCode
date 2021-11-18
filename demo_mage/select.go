package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 100
	}()

	go func() {
		ch2 <- 200
	}()

	select {
	case num1 := <-ch1:
		fmt.Println("ch1中获取的数据：", num1)
	case num2, ok := <-ch2:
		if ok {
			fmt.Println("ch2中获取的数据：", num2)
		} else {
			fmt.Println("ch2通道已经关闭")
		}
	}
	fmt.Println("main over")
}
