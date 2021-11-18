package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c1 := make(chan bool, 10) //因为接下来要循环10次，就会创建 10 个 goroutine，为了保证 10 个 goroutine 读取 c1 的动作同时执行，效率最高，所以此处设置缓存也是 10
	for i := 0; i < 10; i++ {
		go Go(c1, i)
	}

	for i := 0; i < 10; i++ {
		<-c1
	}
}

func Go(c1 chan bool, index int) {
	a := 1
	for i := 0; i < 1000; i++ {
		a += i
	}
	fmt.Println(index, a)

	c1 <- true
}
