package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Step1：打开文件
	fileName := "D:\\GoCode\\abc.txt"
	/*
		Open
			为只读权限
		OpenFile
			可以通过 os.O_CREATE 实现在文件不存在时创建
			通过 os.O_WRONLY 确保写权限
			通过 os.O_APPEND 实现在文件中追加内容。当没有 os.O_APPEND 时，默认从文件起始处开始写
	*/
	// file, err := os.Open(fileName)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	// Step3：关闭文件
	defer file.Close()
	// Step2：写入数据
	// bs := []byte{65, 66, 67, 68, 69, 70} // ABCDEF
	bs := []byte{97, 98, 99, 100} // abcd
	// n, err := file.Write(bs)
	n, err := file.Write(bs[:2]) // 只将切片中的前两个字符写入到文件中
	fmt.Println(n)
	HandleErr(err)
	file.WriteString("\n")

	// 直接写入字符串
	n, err = file.WriteString("HelloWorld")
	fmt.Println(n)
	HandleErr(err)
	file.WriteString("\n")

	n, err = file.Write([]byte("today"))
	fmt.Println(n)
	HandleErr(err)
}

// 写一个专门用来处理错误的函数，直接调用即可，不用每次都写对于错误的判断
func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
