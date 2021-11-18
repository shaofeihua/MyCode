/*
	type 关键字
		用于定义新类型和定义别名
		1、定义新类型：type 类型名 Type
		2、定义别名：type 类型名=Type
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	type myint int // 定义新类型 myint，它和 int 是两个不同的类型
	var i1 myint
	i1 = 100
	var i2 = 200
	fmt.Printf("i1 的类型：%T\n", i1) // main.myint
	fmt.Printf("i2 的类型：%T\n", i2) // int
	//i1 = i2 // cannot use i2 (type int) as type myint in assignment

	type mystring = string // 定义别名，mystring 就是 string
	var str1 mystring
	str1 = "盛小楠"
	var str2 = "白展堂"
	fmt.Printf("str1 的类型：%T\n", str1) // string
	fmt.Printf("str2 的类型：%T\n", str2) // string

	res := funXiaoNan()
	fmt.Println(res(10, 20)) // 1020

	var stu Student
	//stu.name = "李大嘴" // ambiguous selector stu.name 。因为 Person 和 People 中都有 name
	stu.Person.name = "李大嘴"
	stu.Person.show()                              // 李大嘴
	fmt.Printf("s.Person() 的类型是：%T\n", stu.Person) // main.Person
	fmt.Printf("s.People() 的类型是：%T\n", stu.People) // main.Person
}

// 1、定义一个新的类型
//type myint int
//type mystr string

// 2、定义函数类型
type myfun func(int, int) string // type myfun func(a,b int) string 这样写是错的，定义函数类型时，不用写形参

func funXiaoNan() myfun {
	funXN := func(a, b int) string {
		s := strconv.Itoa(a) + strconv.Itoa(b)
		return s
	}
	return funXN
}

// 3、定义 struct
type Person struct {
	name string
}

type People = Person // 是别名。不是新类型！

type Student struct {
	Person
	People
}

func (p Person) show() {
	fmt.Println("Person: ", p.name)
}

// func (p People) show() {} // Person.show redeclared in this block previous declaration at .\type.go:68:6
// Person 和 People 本质上是一回事，所以前面 Person 已经定义过 show() 方法，这里不能再次定义
