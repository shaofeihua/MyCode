package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "D:\\GoProject\\src\\aa.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	defer file.Close()

	fmt.Printf("%T\n", file) // *os.File
	fmt.Println(file)        // &{0xc000074780}

	bs := make([]byte, 4, 4)

	n, err := file.Read(bs)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println(n)
	fmt.Println(bs)
}
