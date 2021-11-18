//字符串的基本操作：包含、索引、切割、合并
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello world"

	// 是否包含
	fmt.Println(strings.Contains(s, "hello"), strings.Contains(s, "?"))

	// 索引
	fmt.Println(strings.Index(s, "o"))

	// 字符串切割
	ss := "1#2#345"
	splitedStr := strings.Split(ss, "#")
	fmt.Println(splitedStr)

	// 字符串合并
	fmt.Println(strings.Join(splitedStr, ","))

	// 判断前缀和后缀
	fmt.Println(strings.HasPrefix(s, "hel"), strings.HasSuffix(s, "rld"))
}
