// 归并排序
package main

import (
	"fmt"
	"sort"
)

func main() {
	p := Merge(InMemSort(ArraySource(3, 8, 6, 9, 2, 6, 2, 1, 7, 4, 5)), InMemSort(ArraySource(6, 12, 33, 4, 10, 8, 88)))
	for v := range p {
		fmt.Print(v, " ")
	}
}

// 读取数据源
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

// 内部排序
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
		for _, v1 := range a {
			out <- v1
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
				// 每循环一次，in1 里面的元素被读取一个，in1 里面的元素就减少一个，读取之后要读 v1 重新赋值，避免读取过的元素被重复读取
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
