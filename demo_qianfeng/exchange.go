package main

import (
	"fmt"
)

func main() {

	v1 := 5
	v2 := 10

	v1, v2 = v2, v1

	fmt.Println(v1, "\n", v2)
}
