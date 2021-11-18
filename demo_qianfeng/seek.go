package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	/*
	   断点续传
	   Seek(offset int64,whence int)(int64,error)，设置光标偏移量
	   		第一个参数：偏移量
	   		第二个参数：如何设置
	   			0: seekstart, 表示相对于文件开始
	   			1: seekcurrent, 表示相对于当前位置的偏移量
	   			2: seekend, 表示相对于文件末尾
	*/
	fileName := "D:\\GoCode\\aaa.txt" // abcdefghij
	file, err := os.OpenFile(fileName, os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 读写
	bs := []byte{0}
	file.Read(bs)
	fmt.Println(string(bs)) // a。此时光标位于 a 后面

	file.Seek(4, io.SeekStart)
	file.Read(bs)
	fmt.Println(string(bs)) // e。此时光标位于 e 后面

	file.Seek(2, 0) // 这里的 0 相当于 seekstart
	file.Read(bs)
	fmt.Println(string(bs)) // c。此时光标位于 c 后面

	file.Seek(3, 1) // 这里的 1 相当于 seekcurrent
	file.Read(bs)
	fmt.Println(string(bs)) // g。此时光标位于 g 后面

	file.Seek(0, 2)         // 或者 file.Seek(0,io.SeekEnd)
	file.WriteString("ABC") // 在文件末位写入 ABC

}
