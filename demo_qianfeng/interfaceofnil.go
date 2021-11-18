/*
	空接口
		不包含任何方法，因此，所有的类型都实现了空接口，所以，空接口可以存储任意类型的数值
*/
package main

import "fmt"

func main() {
	var a1 A = Cat{"花猫"}
	var a2 A = Person{"白展堂", 22}
	var a3 A = "hello"
	var a4 A = 100
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

	test1(a1)
	test1(a2)
	test1(a3)
	test1(a4)
	test1(3.14)
	test1("邢育森")

	test2(a1)
	test2(a2)
	test2(a3)
	test2(a4)
	test2("8")
	test2("燕小六")

	//	map，key 为（固定的）字符串类型，value 为任意类型
	map1 := make(map[string]interface{})
	map1["name"] = "佟湘玉"
	map1["age"] = 30
	map1["friend"] = Person{"郭芙蓉", 20}
	fmt.Println("map1: ", map1) // map[age:30 friend:{郭芙蓉 20} name:佟湘玉]

	//	切片，存储任意类型的数据
	slice1 := make([]interface{}, 0, 10)
	slice1 = append(slice1, a1, a2, a3, a4, 200, "abc")
	fmt.Println("slice1: ", slice1) // [{花猫} {白展堂 22} hello 100 200 abc]

	test3(slice1)
}

// 遍历元素类型为任意类型的切片
func test3(slice2 []interface{}) {
	for i := 0; i < len(slice2); i++ {
		fmt.Printf("切片的第%d个元素是：%v\n", i+1, slice2[i])
	}
}

// 接口 A 是空接口，可以是任意类型
func test1(a A) {
	fmt.Println(a)
}

// 参数类型为匿名空接口，也可以接受任意类型的参数
func test2(a interface{}) {
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
