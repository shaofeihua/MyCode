package main

import "fmt"

func main() {
	// 类型别名
	type chang int64
	type chuan string

	var (
		a chang = 12
		b chuan = "yan"
	)

	fmt.Printf("a type is %T\n", a)
	fmt.Printf("b type is %T\n", b)

}
