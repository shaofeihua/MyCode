package main

import "fmt"

func main() {
	var s1 []int
	fmt.Println("The slice1 is: ", s1)

	var a1 [6]int = [6]int{5, 1, 8, 8, 8, 6}
	fmt.Println("The array1 is: ", a1)

	s2 := a1[3:6]
	fmt.Println("The slice2 is: ", s2)

	s3 := make([]int, 3, 10)
	fmt.Println("The slice3 is: ", s3)
	fmt.Println("The lenth of slice3 is: ", len(s3))
	fmt.Println("The cap of slice3 is: ", cap(s3))

	s4 := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	s5 := s4[2:5]
	fmt.Println("The slice5 is: ", s5)
	fmt.Println("The value of slice5 is: ", string(s5))
}
