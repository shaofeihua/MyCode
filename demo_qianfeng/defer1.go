// defer 和匿名函数结合使用
package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 20
	defer func() {
		fmt.Printf("内部 a = %d, b = %d\n", a, b)
	}() // ()表示调用次函数

	a = 111
	b = 222
	fmt.Printf("外部 a = %d, b = %d\n", a, b)
}
