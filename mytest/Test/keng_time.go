package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format("Mon Jan _2 15:04:05 2006"))
}
