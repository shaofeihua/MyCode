package main

import "fmt"

func main() {
	/*
		空接口（interface{}）
			不包含任何的方法，正因为如此，所有的类型都实现了空接口，因此空接口可以存储任意类型的数值
	*/
	var a1 A = Cat{"花猫"}
	var a2 A = Person{"王二狗", 30}
	var a3 A = "haha"
	var a4 A = 100
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

	fmt.Println("---------------------------------")

	test1(a1)
	test1(a2)
	test1(3.14)
	test1("Ruby")

	fmt.Println("---------------------------------")

	//map: key，字符串；value，任意类型
	map1 := make(map[string]interface{})
	map1["name"] = "李小花"
	map1["age"] = 30
	map1["friend"] = Person{"Jerry", 18}
	fmt.Println(map1)

	fmt.Println("---------------------------------")

	//切片，存储任意类型的数据
	slice1 := make([]interface{}, 0, 10)
	slice1 = append(slice1, a1, a2, a3, a4, 100, "abc")
	fmt.Println(slice1)
}

//接口A是空接口，理解为代表了任意类型
func test1(a A) {
	fmt.Println(a)
}

// 空接口
type A interface {
}
type Cat struct {
	color string
}
type Person struct {
	name string
	age  int
}
