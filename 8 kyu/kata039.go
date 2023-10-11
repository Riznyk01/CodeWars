package main

import (
	"fmt"
)

func main() {
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(OddCount(num))
}
func OddCount(n int) int {
	return n / 2
}
