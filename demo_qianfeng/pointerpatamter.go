/*
	指针作为参数
*/
package main

import "fmt"

func main() {
	a := 10
	fmt.Println("fun1 调用前 a 的值：", a) // 10
	fun1(a)
	fmt.Println("fun1 调用后 a 的值：", a) // 10

	fun2(&a)
	fmt.Println("fun2 调用后 a 的值：", a) // 200

	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println("fun3 调用前 arr1 的值：", arr1) // [1 2 3 4]
	fun3(arr1)
	fmt.Println("fun3 调用后 arr1 的值：", arr1) // [1 2 3 4]

	fun4(&arr1)
	fmt.Println("fun4 调用后 arr1 的值：", arr1) // [200 2 3 4]
}

func fun4(p1 *[4]int) {
	fmt.Println("fun4 函数中的数组指针：", p1) // &[1 2 3 4]
	p1[0] = 200
	fmt.Println("fun4 函数中的数组指针：", p1) // &[200 2 3 4]
}

func fun3(arr [4]int) {
	fmt.Println("fun3 函数中的数组：", arr) // [1 2 3 4]
	arr[0] = 100
	fmt.Println("fun3 函数中修改后的数组：", arr) // [100 2 3 4]
}

func fun2(p1 *int) {
	fmt.Println("fun2 函数中 p1 的值：", *p1) // 10
	*p1 = 200
	fmt.Println("fun2 函数修改 p1 后 p1 的值：", *p1) // 200
}

func fun1(num int) {
	fmt.Println("fun1 函数中 num 的值：", num) // 10
	num = 100
	fmt.Println("fun1 函数修改 num 后 num 的值为：", num) // 100
}
