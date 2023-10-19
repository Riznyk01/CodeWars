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
	fmt.Println(monkeyCount(num))
}
func monkeyCount(n int) []int {
	monkeySlice := make([]int, n)
	for i := 0; i < n; i++ {
		monkeySlice[i] = i + 1
	}
	return monkeySlice
}
