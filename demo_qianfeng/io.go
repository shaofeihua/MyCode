package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// step1：打开文件
	file, err := os.Open("D:\\GoCode\\test.txt") // test.txt 的内容是：abcdefghij
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("file 的类型是：%T\n", file) // *os.File
	fmt.Println(file)                  // &{0xc000080780}
	fmt.Println(*file)                 // {0xc000080780}

	// step3：关闭文件
	defer file.Close()

	// step2：读取文件
	// 读取文件的返回结果是个切片，所以先定义一个切片用来接收读取的数据
	bs := make([]byte, 4, 4)
	/*
		// 第一次读取
		n, err := file.Read(bs)
		fmt.Println(err)        // <nil>
		fmt.Println(n)          // 4
		fmt.Println(bs)         // [97 98 99 100]
		fmt.Println(string(bs)) // abcd

		// 第二次读取
		n, err = file.Read(bs)
		fmt.Println(err)        // <nil>
		fmt.Println(n)          // 4
		fmt.Println(bs)         // [101 102 103 104]
		fmt.Println(string(bs)) // efgh

		// 第三次读取
		n, err = file.Read(bs)
		fmt.Println(err)        // <nil>
		fmt.Println(n)          // 2
		fmt.Println(bs)         // [105 106 103 104]
		fmt.Println(string(bs)) // ijgh

		// 第四次读取
		n, err = file.Read(bs)
		fmt.Println(err) // EOF
		fmt.Println(n)   // 0
	*/
	// 以上四次读取用循环来实现
	n := 1
	for {
		n, err = file.Read(bs)
		if n == 0 || err == io.EOF {
			fmt.Println("已经读取到了文件的末尾，读取文件结束。")
			break
		}
		// fmt.Println(string(bs))     // ijgh
		fmt.Println(string(bs[:n])) // ij
	}
}
