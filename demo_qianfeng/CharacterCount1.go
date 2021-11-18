package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	//filename := "D:\\GoCode\\Mage\\111.txt"
	filename := "D:\\GoCode\\Mage\\I_have_a_dream.txt"
	count(filename)
}

func count(filename string) {
	// 遍历英文字母， for range
	// 过滤掉非英文字符（byte/rune）整数
	// xxx a--z    XXXX A--Z
	// 'a' <= xx <='z' || 'A'<=XX<='Z'

	// ioutil.ReadFile()，直接读取整个文件
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	stat := map[rune]int{}
	for _, ch := range string(bytes) {
		if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
			stat[ch]++
		}
	}

	for ch, v := range stat {
		fmt.Printf("%c: %d\n", ch, v)
	}
}
