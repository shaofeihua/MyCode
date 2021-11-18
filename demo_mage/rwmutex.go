package main

import (
	"fmt"
	"sync"
	"time"
)

var rwMutex *sync.RWMutex
var wg sync.WaitGroup

func main() {
	rwMutex = new(sync.RWMutex)

	wg.Add(2)
	go readData(1)
	go readData(2)

	wg.Wait()
	fmt.Println("main over")
}

func readData(i int) {
	defer wg.Done()

	fmt.Println(i, "开始读：read start")
	rwMutex.RLock() // 读操作上锁
	fmt.Println(i, "正在读取数据：reading。。。")
	time.Sleep(1 * time.Second)
	rwMutex.RUnlock() // 读操作解锁
	fmt.Println(i, "读结束：read over")
}
