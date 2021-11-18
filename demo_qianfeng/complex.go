package main

import (
	"fmt"
)

func main() {
	a := 3.2 + 5.8i
	fmt.Printf("%v type is %T\n", a, a)
	fmt.Printf("%v 的实部是 %v", a, real(a))
	fmt.Printf(", %v 的虚部是 %v\n", a, imag(a))
}
