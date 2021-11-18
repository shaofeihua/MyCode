package main

import (
	"fmt"
	"pipeline/pipe"
)

func main() {
	p := pipeline.ArraySource(5, 8, 3, 7, 9, 0)
	//for {
	//	if num, ok := <-p; ok {
	//		fmt.Println(num)
	//	} else {
	//		break
	//	}
	//}
	for v := range p {
		fmt.Println(v)
	}
}
