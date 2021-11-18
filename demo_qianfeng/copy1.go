package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	/*
	   拷贝文件：
	   使用 io 包下的 Copy() 方法来实现
	*/
	srcFile := "D:\\GoCode\\mmexport1486475844019.gif"
	destFile := "D:\\GoCode\\aaa.gif"

	total, err := CopyFile1(srcFile, destFile)
	fmt.Println(total, err)
}

// func CopyFile1(srcFile, destFile string) (int, error) {
func CopyFile1(srcFile, destFile string) (int64, error) { // io.Copy 的返回值是 int64
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}

	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}

	defer file1.Close()
	defer file2.Close()

	return io.Copy(file2, file1)
}
