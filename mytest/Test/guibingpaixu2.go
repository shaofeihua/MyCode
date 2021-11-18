// 完整的归并排序
package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sort"
)

func main() {
	const filename = "large.in"
	const n = 100000000 // 单位是 byte
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	// 使用 defer 延迟执行要好过把 Close 写到程序最执行，因为加入程序运行到一半出错，defer 仍然会执行
	defer file.Close()

	// 将随机生成数据写入新创建的文件
	p := RandomSource(n)
	// 使用 bufio.NewWriter() 增加写的速度
	WriterSink(bufio.NewWriter(file), p)
	// Flush 将缓存中的数据提交到底层的 io.Writer 中
	bufio.NewWriter(file).Flush()

	// 打开文件
	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 读取文件的内容
	// 使用 bufio.NewWriter() 增加写的速度
	p = ReaderSource(bufio.NewReader(file))
	for v := range p {
		fmt.Println(v)
	}
}

//func main() {
func mergeDemo() {
	p := Merge(InMemSort(ArraySource(32, 65, 6, 8, 9, 5, 59)), InMemSort(ArraySource(32, 65, 6, 8, 9, 5, 59, 23, 56)))

	for v := range p {
		fmt.Print(v, " ")
	}
}

func ArraySource(a ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

func InMemSort(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		// Read into memory
		a := []int{}
		for v := range in {
			a = append(a, v)
		}

		// Sort
		sort.Ints(a)

		// Output
		for _, v := range a {
			out <- v
		}

		close(out)
	}()
	return out
}

// 归并排序
func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
	}()
	return out
}

func ReaderSource(reader io.Reader) <-chan int {
	out := make(chan int)
	go func() {
		// 类型为 byte ，长度（容量未声明，所以容量等于长度，也是 8）为 8 的 slice，用于缓存读取到的数据
		// 64 位的操作系统，一个 int 的长度为 1 byte（字节），即 8 位。所以 slice 长度设置为 （64/8=）8
		buffer := make([]byte, 8)
		for {
			// n 代表读了多少字节，err 代表读取是否有错误。
			// 读到 EOF 表示有错误。或者 n 不到 8 字节，例如 n = 4 byte，那么还是存在 EOF
			// 将数据读进 buffer
			n, err := reader.Read(buffer)
			if n > 0 {
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}
			// 判断数据读取是否包含错误，如果错误不为空，则跳出循环
			if err != nil {
				break
			}
		}
		close(out)
	}()
	return out
}

func WriterSink(writer io.Writer, in <-chan int) {
	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))
		writer.Write(buffer)
	}
}

// 随机数据源
func RandomSource(count int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}

func MergeN(inputs ...<-chan int) <-chan int {
	// 先处理数据长度为 1 的情况，不需要做归并
	if len(inputs) == 1 {
		return inputs[0]
	}
	// 以多路两两归并为例（将一份数据来源拆分为两份（其中每一份可能再分为两份）分别进行归并，然后再将二者归并），先获取输入数据的长度
	m := len(inputs) / 2

	return Merge(MergeN(inputs[:m]...), MergeN(inputs[m:]))
}
