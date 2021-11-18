package main

import (
	"fmt"
)

func main() {
	b1 := Book{}
	b1.bookName = "西游记"
	b1.price = 50.2

	s1 := Student{}
	s1.name = "王二狗"
	s1.age = 18
	s1.book = b1

	fmt.Printf("这个学生的姓名是：%s，年龄是：%d，正在读的书是：%s，书的价格是：%.2f\n", s1.name, s1.age, s1.book.bookName, s1.book.price)

	s2 := Student{name: "李小花", age: 18, book: Book{bookName: "三国", price: 56.8}}
	fmt.Println(s2.name, s2.age, s2.book.bookName, s2.book.price)

	s1.book.bookName = "水浒传"
	fmt.Println(s1)
	fmt.Println(b1)

	b4 := Book{bookName: "西厢记", price: 39.9}
	s4 := Student2{name: "崔莺莺", age: 19, book: &b4}
	fmt.Printf("b4 : %v, b4 的地址是：%p\n", b4, &b4) // b4 : {西厢记 39.9}, b4 的地址是：0xc000048460
	fmt.Println(s4)                              // {崔莺莺 19 0xc000048460}
	fmt.Println(s4.book)                         // &{西厢记 39.9}
	fmt.Println(*s4.book)                        // {西厢记 39.9}
	fmt.Println(&s4.book)

	s4.book.bookName = "东游记"
	fmt.Println(b4)
	fmt.Println(s4)
}

type Student struct {
	name string
	age  int
	book Book
}

type Student2 struct {
	name string
	age  int
	book *Book
}

type Book struct {
	bookName string
	price    float64
}
