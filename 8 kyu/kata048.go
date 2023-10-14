package main

import (
	"fmt"
)

func main() {
	var x, y int

	_, err := fmt.Scan(&x, &y)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(Multiply(x, y))
}

func Multiply(a, b int) int {
	return a * b
}
