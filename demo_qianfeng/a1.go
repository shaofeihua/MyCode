package main

import "fmt"

/*
选班长
    小红，小明，小李，小李，小李，小明
    小红： 1
    小明： 2
    小李： 3
*/
//统计"我有一个梦想"中每个英文（大小写区分）字母出现次数
func main() {
	vs := []string{"abc", "nihao", "world"}
	delItem(vs, "world")
	//fmt.Println(vs)
}

func delItem(vs []string, s string) []string {
	for i := 0; i < len(vs); i++ {
		if s == vs[i] {
			vs = append(vs[:i], vs[i+1:]...)
			//i = i - 1
			fmt.Println(vs)
		}
	}
	return vs
}
