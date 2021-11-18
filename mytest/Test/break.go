package main

import "fmt"

func main() {
LABLE1:
	for {
		for a := 1; a < 10; a++ {
			if a > 3 {
				fmt.Println(a)
				break LABLE1
			}
		}
	}
	fmt.Println("OK")
}
