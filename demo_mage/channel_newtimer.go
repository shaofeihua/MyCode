package main

import (
	"fmt"
	"time"
)

func main() {
	//timer := time.NewTimer(3 * time.Second)
	//fmt.Println(time.Now())
	//
	//ch2 := timer.C
	//fmt.Println(<-ch2)

	timer2 := time.NewTimer(5 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer2 结束了")
	}()

	time.Sleep(3 * time.Second)
	flag := timer2.Stop()
	if flag {
		fmt.Println("Timer2 停止了")
	}
}
