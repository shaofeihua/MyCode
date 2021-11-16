package main

import (
	"fmt"
	"reflect"
)

type Person struct{
	Name string
	Age int
	Sex string
}

func (p Person) Say(msg string){
	fmt.Println("hello,",msg)
}

func (p Person)PrintInfo(){
	fmt.Println("姓名：%s，年龄：%d，性别：%s",p.Name,p.Age,p.Sex)
}

func main(){
	p1:=Person{"王二狗",30,"男"}
	GetMessage(p1)
}

func GetMessage(input interface{}) {
	getType:=reflect.TypeOf(input)
	fmt.Println("get Type is: ",getType.Name())
	fmt.Println("get Kind is: ",getType.Kind())

	getValue:=reflect.ValueOf(input)
	fmt.Println("get all Field is: ",getValue)
}