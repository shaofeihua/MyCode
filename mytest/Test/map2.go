package main

import "fmt"

func main() {
	m := make(map[int]string)
	m[1] = "ok"

	a := m[1]
	fmt.Println("The value of a is: ", a)

	delete(m, 1)
	b := m[1]
	fmt.Println("The value of b is: ", b)

	n := make(map[int]map[int]string) //定义一个 map ，它的值也是一个 map
	n[1] = make(map[int]string)       //由于第一层 map 的值也是 map，所以需要将这个值也进行初始化
	n[1][1] = "xiaonan"               //如果没有上一个语句的初始化，本赋值是不成立的
	x := n[1][1]
	fmt.Println(x)
}
