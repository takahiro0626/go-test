package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func main() {
	result := add(1, 2)
	fmt.Println(result)
}