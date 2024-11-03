package main

import (
	"fmt"
	"math"
)

func main() {
	var r float32
	var size float32
	fmt.Scan(&r)
	size = circlesize(r)
	fmt.Printf("%f", size)
}
func circlesize(r float32) float32 {
	var size, pi float32
	pi = math.Pi
	size = pi * r * r
	return size
}
