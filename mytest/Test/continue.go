package main

import "fmt"

func main() {
LABLE1:
	for a := 1; a < 10; a++ {
		for {
			fmt.Println(a)
			continue LABLE1 //第二级的 for 循环是无限循环，当它执行到 continue ，马上就跳转到标签位置，所有continue 后面的那个 fmt.Println(a) 就没有执行的机会了
			fmt.Println(a)
		}
	}
	fmt.Println("OK")
}
