/*
	错误类型表示
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	//file, err := os.Open("123.txt") // 当 123.txt 存在，这里使用相对路径也导致找不到文件而发生 panic。即使 123.txt 跟当前 go 代码文件处于同一路径下。
	file, err := os.Open("D:\\GoCode\\Mage\\123.txt") // 当 123.txt 存在，并且使用绝对路径后，不再发生 panic
	if err != nil {
		fmt.Println(err)
		if ins, ok := err.(*os.PathError); ok {
			fmt.Println("Op:", ins.Op)     // open
			fmt.Println("Path:", ins.Path) // D:\GoCode\Mage\123.txt
			fmt.Println("Err:", ins.Err)   // The system cannot find the file specified.
		}
	}
	fmt.Println(file.Name(), "打开文件成功。")
}

/*
	上面产生的 panic 如下：
		panic: runtime error: invalid memory address or nil pointer dereference
		[signal 0xc0000005 code=0x0 addr=0x0 pc=0x8698ad]

		goroutine 1 [running]:
		os.(*File).Name(...)
			C:/Program Files/Go/src/os/file.go:55
		main.main()
			D:/GoCode/Mage/errortype.go:22 +0x10d
*/
