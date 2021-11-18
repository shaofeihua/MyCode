package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	dirname := "C:\\Users\\lusibo\\Documents"
	listFiles(dirname)
}

func listFiles(dirname string) {
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfos)
	fmt.Printf("fileInfos 的数据类型是：%T\n", fileInfos)

	for i, fi := range fileInfos {
		filename := dirname + "\\" + fi.Name()
		fmt.Printf("%d, %s\n", i, filename)
		//	递归遍历子目录
		if fi.IsDir() {
			listFiles(filename)
		}
	}
}
