package main

import (
	"fmt"
)

func main() {
	var m map[int]string
	m = map[int]string{}

	var n map[int]string
	n = make(map[int]string)

	var p map[int]string = make(map[int]string)

	var x = make(map[int]string)

	y := make(map[int]string)

	y[2] = "xiaonan"

	fmt.Println(m, n, p, x, y)
}
