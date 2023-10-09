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
	fmt.Println(MakeNegative(num))
}
func MakeNegative(x int) int {
	if x < 0 {
		return x
	}
	return x * -1
}
