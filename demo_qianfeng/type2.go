package main

import "fmt"

func main() {
	var s Student
	// s.name = "王二狗" // ambiguous selector s.name
	s.Person.name = "李小花"
	fmt.Println(s)
	s.People.name = "王大锤"
	fmt.Println(s)
	// s.show() // ambiguous selector s.name
	s.Person.show()
	s.People.show()
}

type Person struct {
	name string
}

func (p Person) show() {
	fmt.Printf(p.name)
	fmt.Println()
}

type People = Person

type Student struct {
	// 嵌入两个结构体
	Person
	People
}
