package main

import "fmt"
import "runtime"
import "unsafe"

func main() {
	// 处理器架构
	fmt.Println("CPU: ", runtime.GOARCH)
	//整型变量占用的空间，单位是字节Byte
	fmt.Println("IntSize: ", unsafe.Sizeof(9999999999999))
}
