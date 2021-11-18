package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// 1、文件的路径
	fileName1 := "D:\\nihao.txt"
	fileName2 := "test.txt"
	fmt.Println(filepath.IsAbs(fileName1))      // true
	fmt.Println(filepath.IsAbs(fileName2))      // false
	fmt.Println(filepath.Abs(fileName1))        // D:\nihao.txt
	fmt.Println(filepath.Abs(fileName2))        // D:\GoCode\test.txt
	fmt.Println(filepath.Join(fileName1, "..")) // D:\

	// // 2、创建目录
	// err := os.Mkdir("D:\\nihao123.txt", os.ModePerm)
	// if err != nil {
	// 	fmt.Println("ERROR: ", err)
	// 	return
	// }
	// fmt.Println("文件夹创建成功")

	// // err1 := os.Mkdir("D:\\nihao123\\aa\\bb", os.ModePerm) // os.Mkdir 只能创建一层文件夹
	// err1 := os.MkdirAll("D:\\nihao123\\aa\\bb", os.ModePerm) // os.MkdirALL 可以创建多层文件夹
	// if err1 != nil {
	// 	fmt.Println("ERROR: ", err1)
	// 	return
	// }
	// fmt.Println("多层文件夹创建成功")

	// // 3、创建文件
	// // 如果文件存在，os.Create 会把文件清空
	// file, err2 := os.Create("D:\\nihao123\\aa\\bb\\readme.txt")
	// if err2 != nil {
	// 	fmt.Println("ERROR: ", err2)
	// 	return
	// }
	// fmt.Println(file)

	// os.Open 对文件的操作是只读，不能写入
	file3, err3 := os.Open("test.txt")
	if err3 != nil {
		fmt.Println("ERROR: ", err3)
		return
	}
	fmt.Println(file3)
}
