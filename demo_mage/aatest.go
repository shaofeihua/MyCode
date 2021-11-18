package main

import (
	"fmt"
)

/*
选班长
    小红，小明，小李，小李，小李，小明
    小红： 1
    小明： 2
    小李： 3
*/
//统计"我有一个梦想"中每个英文（大小写区分）字母出现次数
func main() {
	//字符串
	article := `aasdfg,hjk.`
	// 遍历英文字母， for range
	// 过滤掉非英文字符（byte/rune）整数
	// xxx a--z    XXXX A--Z
	// 'a' <= xx <='z' || 'A'<=XX<='Z'
	stat := map[rune]int{}
	for _, ch := range article {
		if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
			fmt.Println(ch)
			stat[ch]++
			fmt.Println(stat)
		}
	}

	for ch, v := range stat {
		fmt.Printf("%c: %d\n", ch, v)
	}

	fmt.Print(stat)

}
