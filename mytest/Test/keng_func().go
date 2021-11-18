package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}

	for _, v := range s {
		//go func() {
		go func(v int) {
			fmt.Println(v)
			//}()
		}(v)
	}
	select {}
}
