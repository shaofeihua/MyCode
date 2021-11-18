package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("type.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f.Name(), "打开文件成功")
}
