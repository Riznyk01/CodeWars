package main

import (
	"fmt"
)

func main() {
	var n, x, y int
	_, err := fmt.Scan(&n, &x, &y)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(IsDivisible(n, x, y))
}
func IsDivisible(n, x, y int) bool {
	return n%x == 0 && n%y == 0
}
