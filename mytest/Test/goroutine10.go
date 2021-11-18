package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		for v := range c {
			fmt.Println(v)
		}
	}()

	for i := 0; i < 10; i++ {
		select {
		case c <- 1:
		case c <- 0:
		}
	}
}
