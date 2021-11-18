package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	num1 := rand.Int()
	fmt.Println(num1)

	rand.Seed(1000)
	num2 := rand.Intn(10)
	fmt.Println(num2)

	t1 := time.Now()
	fmt.Println(t1)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		fmt.Println("随机数：", rand.Intn(1000))
	}
}
