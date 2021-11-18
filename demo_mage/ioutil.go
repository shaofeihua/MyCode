package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fileName := "D:\\GoProject\\src\\aa.txt"
	data, err := ioutil.ReadFile(fileName)
	fmt.Println(err)
	fmt.Println(data)
	fmt.Println(string(data))

	fileName1 := "D:\\GoProject\\src\\ab.txt"
	s1 := "床前明月光"
	err = ioutil.WriteFile(fileName1, []byte(s1), os.ModePerm)
	fmt.Println(err)
	data1, err1 := ioutil.ReadFile(fileName1)
	fmt.Printf("data1 的数据类型是：%T\n", data1)
	fmt.Println(err1)
	fmt.Println(data1)
	fmt.Println(string(data1))

	s2 := "小楠和萌萌是好朋友"
	r1 := strings.NewReader(s2)
	fmt.Printf("r1 的数据类型是：%T\n", r1)
	fmt.Println(err)
	fmt.Println(data)
	fmt.Println(string(data))
}
