package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	srcFile := "D:\\GoProject\\src\\aa.txt"
	destFile := "D:\\GoProject\\src\\ab.txt"

	total, err := CopyFile1(srcFile, destFile)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("已复制：", total)
}

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

	bs := make([]byte, 1024, 1024)
	n := -1
	total := 0
	for {
		n, err = file1.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("拷贝完毕。。。")
			break
		} else if err != nil {
			fmt.Println("出错了")
			return total, err
		}
		total += n
		file2.Write(bs[:n])
	}
	return total, nil
}
