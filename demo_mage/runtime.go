package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GOROOT: ", runtime.GOROOT())
	fmt.Println("os: ", runtime.GOOS)
	fmt.Println("逻辑CPU的数量：",runtime.NumCPU())
}
