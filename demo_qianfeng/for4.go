package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 2; i <= 100; i++ {
		flag := true
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			// for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(i)
		}
	}
}
