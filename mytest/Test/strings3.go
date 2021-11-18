// 使用 flag 解析命令行参数
package main

import (
	"flag"
	"fmt"
)

func main() {
	// 如果要读取的参数是字符串，则使用 flag.String。如果参数是整型，则使用 flag.Int
	// String 里面有三个参数，分别是名称、默认值、使用说明，均为可选项
	username := flag.String("name", "you", "Input one name")
	num := flag.Int("years", 10000, "Input number of years")

	// 解析
	flag.Parse()

	fmt.Println("I love", *username, *num, "years")
}
