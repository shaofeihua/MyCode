package main

import (
	"fmt"
)

func main() {
	var a1 [5]int = [5]int{1, 2, 3, 4, 5}
	a2 := a1[:]
	a3 := a2[:]
	fmt.Printf("The a1 is: %p\n", a1)
	fmt.Printf("The a2 is: %p %v\n", a2, a2)
	fmt.Printf("The a3 is: %p %v\n", a3, a3)

	s1 := []int{2, 3}
	s2 := s1[:]
	fmt.Printf("The s1 is: %p\n", s1)
	fmt.Printf("The s2 is: %p\n", s2)
}
