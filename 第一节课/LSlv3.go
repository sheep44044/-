package main

import "fmt"

func main() {
	var y int
	fmt.Scan(&y)
	result := sushu(y)
	fmt.Println(result)
}
func sushu(y int) string {
	var result int
	for i := 1; i <= y; i++ {
		result = y % i
		if result != 0 {
			return "它不是素数"
			break
		}
	}
	return "它是素数"
}
