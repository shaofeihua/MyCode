package main

import "fmt"
import "unsafe"

var (
	a1 byte = 'G' // byte 单字节
	b1 rune = '景' // rune 多字节
)

func main() {
	fmt.Printf("%c占用%d个字节\n", 'G', unsafe.Sizeof(a1))
	fmt.Printf("%c占用%d个字节\n", '景', unsafe.Sizeof(b1))
}

// 在 Go 中，多字节字符使用标准的 Unicode 编码，而不是 utf8，因此会是 4 个字节。如果
