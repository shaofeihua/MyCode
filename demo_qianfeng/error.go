package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("D:\\GoCode\\test.txt")
	if err != nil {
		// fmt.Println(err) // open /test.txt: The system cannot find the file specified.
		log.Fatal(err) // 2020/08/10 19:53:21 open /test.txt: The system cannot find the file specified.\nexit status 1
		return         // 如果存在错误，必须使用 return 进行返回
	}
	fmt.Println(f.Name(), "opened successfully")
}
