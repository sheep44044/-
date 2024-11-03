package main

import (
	"fmt"
)

func main() {
	var x, y, result int
	fmt.Scan(&x, &y)
	result = sum(x, y)
	fmt.Printf("%d + %d = %d\n", x, y, result)
}

func sum(x, y int) int {
	var result int
	result = x + y
	return result
}
