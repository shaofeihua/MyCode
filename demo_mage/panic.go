package main

import "fmt"

func main() {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg, "程序恢复啦")
		}
	}()
	
	funA()
	defer myprint("defer main: 3")
	funB()
	defer myprint("defer main: 4")
	fmt.Println("main ... over ...")
}

func myprint(s string) {
	fmt.Println(s)
}

func funA() {
	fmt.Println("我是函数funA")
}

func funB() {
	fmt.Println("我是函数funB")
	defer myprint("defer funB(): 1")
	for i := 1; i <= 10; i++ {
		fmt.Println("i:", i)
		if i == 5 {
			// 让程序中断
			panic("funB函数，恐慌了")
		}
	}
	defer myprint("defer funB(): 2")
}
