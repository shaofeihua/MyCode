/*
	结构体指针

	new()
		使用 go 的内置函数 new() 可以创建任意类型的指针
		new() 创建出来的指针，不是 nil，不是空指针，它指向了新分配的类型的内存地址，这个地址对应的值是零值。
*/
package main

import "fmt"

func main() {
	p1 := Persion{"盛小楠", 33, "女", "沈阳市"}
	fmt.Println("p1: ", p1)           // {盛小楠 33 女 沈阳市}
	fmt.Printf("p1 的内存地址: %p\n", &p1) // 0xc000040040
	// 定义结构体指针
	var pp1 *Persion
	pp1 = &p1
	pp1.name = "傻小楠"                               // 这里省略 * ，等价于 *pp1.name，但不能用 *pp1.name
	fmt.Println("p1: ", p1)                        // {傻小楠 33 女 沈阳市}
	fmt.Printf("pp1 的类型：%T，pp1 的值：%p\n", pp1, pp1) // 0xc000040040

	// 使用内置函数 new() 创建指针
	pp2 := new(Persion)
	fmt.Println("pp2: ", pp2)                      // &{ 0  }，零值
	fmt.Printf("pp2 的类型：%T，pp2 的值：%p\n", pp2, pp2) // 0xc000040100
	pp2.name = "大王"
	pp2.age = 1000
	pp2.sex = "男"
	pp2.address = "皇宫"
	fmt.Println("pp2: ", pp2) // &{大王 1000 男 皇宫}
}

type Persion struct {
	name    string
	age     int
	sex     string
	address string
}
