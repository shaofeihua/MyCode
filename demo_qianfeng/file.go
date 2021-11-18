package main

import (
	"fmt"
	"os"
)

func main() {
	fileInfo, err := os.Stat("D:\\nihao.txt")
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println(fileInfo)
	fmt.Printf("%T\n", fileInfo) // *os.fileStat

	// 文件名
	fmt.Println(fileInfo.Name())
	// 文件大小
	fmt.Println(fileInfo.Size())
	// 是否是目录
	fmt.Println(fileInfo.IsDir())
	// 修改时间
	fmt.Println(fileInfo.ModTime())
	// 权限
	fmt.Println(fileInfo.Mode())
}
