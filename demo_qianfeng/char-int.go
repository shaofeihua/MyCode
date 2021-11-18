package main

import "fmt"

var (
	b1 byte = 'G'
	r1 rune = '景'
)

func main() {
	fmt.Println(b1)
	fmt.Println(r1)
	fmt.Printf("%c: %T\n", b1, b1)
	fmt.Printf("%c: %T\n", r1, r1)
}
