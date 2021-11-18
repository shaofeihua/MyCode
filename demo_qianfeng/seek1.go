package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file_name := flag.String("f", "", "文件名")

	flag.Parse()
	if *file_name != "" {
		is_file(*file_name)
	} else {
		fmt.Printf("test\n")
	}
}

func is_file(file_name string) {
	file, err := os.Open(file_name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	input, _ := ioutil.ReadAll(file)
	fmt.Println(input)
}
