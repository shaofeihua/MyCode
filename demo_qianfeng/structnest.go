/*
	结构体的嵌套
*/
package main

import "fmt"

func main() {
	b1 := Book{}
	b1.bookname = "西游记"
	b1.price = 35.8
	fmt.Println("b1: ", b1) // {西游记 35.8}

	s1 := Student{}
	s1.name = "王二狗"
	s1.age = 10
	s1.book = b1
	fmt.Println("s1: ", s1) // {王二狗 10 {西游记 35.8}}

	fmt.Printf("学生姓名：%s，学生年龄：%d，看的书是：%s，书的价格是：%.2f\n", s1.name, s1.age, s1.book.bookname, s1.book.price) // 学生姓名：王二狗，学生年龄：10，看的书是：西游记，书的价格是：35.80

	s2 := Student{name: "盛小楠", age: 18, book: Book{bookname: "沉默的羔羊", price: 25.5}}
	fmt.Println("s2: ", s2) // {盛小楠 18 {沉默的羔羊 25.5}}

	s1.book.bookname = "红楼梦"
	fmt.Println("s1: ", s1) // {王二狗 10 {红楼梦 35.8}} ，变了
	fmt.Println("b1: ", b1) // {西游记 35.8} ，没变
}

type Book struct {
	bookname string
	price    float64
}

type Student struct {
	name string
	age  int
	book Book
}
