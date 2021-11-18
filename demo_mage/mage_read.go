package main

import (
	"bufio"
	"io"
	"os"
)

func processLine(line []byte) {
	os.Stdout.Write(line)
}

func ReadLine(filePth string, hookfn func([]byte)) error {
	//f, err := os.Open("D:\\GoProject\\src\\I_have_a_dream.txt")
	f, err := os.Open("D:/GoProject/src/I_have_a_dream.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	bfRd := bufio.NewReader(f)
	for {
		line, err := bfRd.ReadBytes('\n')
		hookfn(line)    //放在错误处理前面，即使发生错误，也会处理已经读取到的数据。
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}

func main() {
	ReadLine("test.txt", processLine)
}
