package main

import (
	_ "embed"
	"fmt"
)

//go:embed version
var version string

func main() {
	fmt.Println("version: ", version)
}
