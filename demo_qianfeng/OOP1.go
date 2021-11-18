package main

import "fmt"

func main() {
	// 1、创建父类的对象
	p1 := Person{name: "王五", age: 30}
	fmt.Println("p1: ", p1) // {王五 30}

	// 2、创建子类的对象
	//s1 := Student{Person{"王无忌", 10}, "精武门"} // 省略字段名，直接写值
	s1 := Student{Person: Person{name: "王无忌", age: 10}, school: "精武门"} // 不省略字段名
	fmt.Println("s1: ", s1)                                            // {{王无忌 10} 精武门}

	var s2 Student
	s2.Person.name = "盛小楠"
	s2.Person.age = 33
	s2.school = "蒙中"
	fmt.Println("s2: ", s2) // {{盛小楠 33} 蒙中}

	var s3 Student
	s3.name = "满满" // 省略 Person
	s3.age = 6     // 省略 Person
	s3.school = "红黄蓝"
	fmt.Println("s3: ", s3) // {{满满 6} 红黄蓝}
	/*
		s3.Person.name 等价于 s3.name
		Student 结构体把 Person 结构体作为一个匿名字段了
		那么 Person 中的字段，对于 Student 来说就是提升字段
		Student 可以直接访问 Person 中的字段（父类对象的提升字段可被子类对象直接访问）
	*/
}

// 1、定义父类
type Person struct {
	name string
	age  int
}

// 2、定义子类
type Student struct {
	Person // 匿名字段是另一个结构体的名称
	school string
}
