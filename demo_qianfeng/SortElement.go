package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{9, 8, 19, 10, 2, 8}

	sort.Ints(s)

	//fmt.Println(len(s)-1, s[len(s)-1])
	fmt.Printf("%d: %d\n", len(s)-1, s[len(s)-1])
}
