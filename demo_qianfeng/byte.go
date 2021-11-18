package main

import (
	"fmt"
)

func main() {
	str := "chunjing"
	fmt.Printf("%v 的类型是 %T, 长度是 %d\n", str, str, len(str))

	fmt.Printf("str[0] = %c, str[4] = %c\n", str[0], str[4])
}
