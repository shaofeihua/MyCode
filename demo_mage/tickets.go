package main

import (
	"fmt"
	"time"
)

/*
	并发场景当中的临界资源安全问题

	举例场景：
		4个窗口同时卖20张票
		4个goroutine，模拟4个窗口
*/

// 全局变量，表示票
var ticket = 20 // 20张票

func main() {
	go saleTickets("售票口1")
	go saleTickets("售票口2")
	go saleTickets("售票口3")
	go saleTickets("售票口4")

	time.Sleep(5 * time.Second)
}

func saleTickets(name string) {
	for {
		if ticket > 0 {
			fmt.Println(name, "售出：", ticket)
			ticket--
		} else {
			fmt.Println(name, "售罄")
			break
		}
	}
}
