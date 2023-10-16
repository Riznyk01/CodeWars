package main

import (
	"fmt"
)

func main() {
	var a, b int
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(CountBy(a, b))
}
func CountBy(x, n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			slice[0] = x
		} else {
			slice[i] = slice[i-1] + x
		}
	}
	return slice
}
