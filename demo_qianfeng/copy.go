package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	/*
	   拷贝文件：
	   使用 io 包下的 Read() 和 Write() 方法来实现
	*/
	srcFile := "D:\\GoCode\\mmexport1486475844019.gif"
	destFile := "D:\\GoCode\\mmexport1486475844019-1.gif"

	total, err := CopyFile1(srcFile, destFile)
	fmt.Println(total, err)
}

// 创建一个函数，用于通过 io 操作实现文件的拷贝，返回值是拷贝的总量（字节）和错误
func CopyFile1(srcFile, destFile string) (int, error) {
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

	// 读写
	bs := make([]byte, 1024, 1024)
	n := -1    // 读取的数据量
	total := 0 // 拷贝的数据总量
	// 当拷贝的文件比较大时，是边读边写。所以用循环来实现，每循环一次，就读一次，就统计一次，然后就写一次
	for {
		n, err = file1.Read(bs)
		fmt.Println("n: ", n)
		if err == io.EOF || n == 0 {
			fmt.Println("拷贝完毕。")
			break
		} else if err != nil {
			fmt.Println("报错了：")
			return total, err
		}
		total += n
		fmt.Println("total: ", total)
		file2.Write(bs[:n])
	}
	return total, nil
}
