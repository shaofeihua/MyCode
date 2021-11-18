/*
	panic
		恐慌
		用于中断程序的执行
		内置函数，它可以接受任意类型的参数
		当发生恐慌时，已经被 defer 内容的仍然会执行，还未被 defer 的内容就不执行了
		当外围函数中的代码运行发生恐慌时，只有函数中所有的 defer 都执行完毕后，恐慌才会被传递到调用处
	recover
		恢复
		捕获 panic，用于恢复程序
		必须配合 defer 使用
		recover 函数不需要参数，但是它有返回值，它的返回值，就是发生恐慌时传递给 panic 的参数

	使用 panic 的场景（其他场景建议使用 error 处理）
		1、空指针引用
		2、下标越界
		3、除数为 0
		4、不应该出现的分支，比如 default
		5、输入不应该引起函数错误
*/
package main

import "fmt"

func main() {
	// recover() 放在 main() 中或者 funB() 中，代码最终的运行结果是不一样的
	defer func() {
		msg := recover()
		if msg != nil {
			fmt.Println(msg, "程序恢复了")
		}
	}()

	funA()
	defer myprint("defer main(): 3...")
	funB()
	defer myprint("defer main(): 4...")

	fmt.Println("main over")
}

func myprint(s string) {
	fmt.Println(s)
}

func funA() {
	fmt.Println("我是函数 A")
}

func funB() {
	//defer func() {
	//	msg := recover()
	//	if msg != nil {
	//		fmt.Println(msg, "程序恢复了")
	//	}
	//}()
	fmt.Println("我是函数 B")
	defer myprint("defer funB(): 1...")
	for i := 0; i <= 10; i++ {
		fmt.Println("i: ", i)
		if i == 5 {
			// 中断程序运行
			// 注意：i=5 仍然会打印出来！！！
			panic("funB() 发生恐慌了")
		}
	}
	defer myprint("defer funB(): 2...")
}
