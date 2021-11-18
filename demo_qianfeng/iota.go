package main

import (
	"fmt"
)

func main() {
	// iota，常量自动生成器，每隔一行，自动累加 1
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Printf("a = %d\nb = %d\nc = %d\n", a, b, c)

	// iota 遇到 const，则重置为 0
	const d = iota
	fmt.Println("d =", d)

	// 可以只写一个 iota，后面每一行的常量即使不赋值，也会自动依次加 1
	const (
		a1 = iota
		b1
		c1
	)
	fmt.Printf("a1 = %d\nb1 = %d\nc1 = %d\n", a1, b1, c1)

	// 如果同一行的常量都赋值为 iota，则值都一样
	const (
		l          = 12
		m1, m2, m3 = iota, iota, iota
		n          = iota
	)
	fmt.Printf("l = %d\nm1 = %d\nm2 = %d\nm3 = %d\nn = %d\n", l, m1, m2, m3, n)

	// iota 对应的常量的值，与 iota 对应的常量属于第几行有关，与上一行常量的值无关。例如，总共 3 行常量，iota 出现在第 3 行，则其对应的常量的值为 2；总共 5 行常量，iota 出现在第 5 行，则其对应的常量的值为 4
	const (
		i          = 3
		j1, j2, j3 = 2, 1, 5
		k          = 88
		x          = 99
		y          = iota
	)
	fmt.Printf("i = %d\nj1 = %d\nj2 = %d\nj3 = %d\nk = %d\nx = %d\ny = %d\n", i, j1, j2, j3, k, x, y)
}
