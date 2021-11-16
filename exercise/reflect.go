package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)

	fmt.Printf("v 的类型是：%T, v 的值是：%v\n", v)
	fmt.Println("kind: ", v.Kind())
	fmt.Println("type: ", v.Type())
	fmt.Println("value: ", v.Float())

	convertValue := v.Interface().(float64)
	fmt.Println(convertValue)

	pointer := reflect.ValueOf(&x)
	convertPointer := pointer.Interface().(*float64)
	fmt.Println(convertPointer)
}
