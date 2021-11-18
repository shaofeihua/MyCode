package main

import "fmt"

const (
	B int = 1 << (iota * 10)
	KB
	MB
	GB
)

func main() {
	fmt.Println(B, KB, MB, GB)

	var a int = 10
	a++
	a--
	var p *int = &a
	fmt.Println(*p)
}
