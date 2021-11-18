package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Println("Hello World!")
}

func Info(o interface{}) {
	t := reflect.TypeOf(o) //获取接口的类型
}

func main() {

}
