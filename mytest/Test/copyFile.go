package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// 复制命令提供两个参数 -f 和 -v ，分别用来实现强制复制和显示过程
	var force，showProgress bool
	// 当目的文件存在时则强制执行复制
	flag.BoolVar(&force, "f", false, "force copy when existing")
	// 显示当前正在做什么（显示过程）
	flag.BoolVar(&showProgress, "v", false, "explain what is being done")

	// 解析
	flag.Parse()

	// 当参数少于 2 个，使用函数 Usage 提示用法  
	if flag.NArg() < 2 {
		flag.Usage()
		return
	}

}

// 判断目的文件是否已存在（或者说将要复制出来的文件名是否有重名）
// 如果文件不存在则返回 err
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

// 文件拷贝
// 文件拷贝的实质是打开并读取一个文件的内容复制给另一个文件
// 返回值 w int64 是指拷贝了多少字节
func copyFile(src, dst string) (w int64, err error) {
	// 打开源文件并返回句柄
	srcFile, err := os.Open(src)
	// 如果有错误的话则打印错误并返回
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}

// 拷贝文件得行为。src 源文件名，dst 目的文件名
func copyFileAction(src, dst string, showProgress, force bool) {
	if !force {
		// 调用上面写的函数判断目的文件 dst 是否存在
		if fileExists(dst) {
			// 如果存在则给出选择是否覆盖
			fmt.Printf("%s exists override? y/n\n", dst)
			// 前面给出了选择，现在需要使用从操作系统的标准输入读取用户的选择
			reader := bufio.NewReader(os.Stdin)
			// 然后通过 bufio 的 reader 读取一行
			data, _, _ := reader.ReadLine()

			// 提取并判断用户的输入，考虑到可能存在非法字符例如空格，所以先去掉两边的空格再转化为字符串，看是否等于 y
			// 如果不等于 y,则说明用户不希望拷贝继续,直接返回，不进行任何操作
			if strings.TrimSpace(string(data)) != "y" {
				return
			}

		}
	}

	copyFile(src, dst)
}
