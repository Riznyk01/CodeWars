package main

import "fmt"

func main() {
	var num, lim int
	_, err := fmt.Scan(&num, &lim)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(FindMultiples(num, lim))
}
func FindMultiples(integer, limit int) []int {
	resSlice := make([]int, 0)
	for i := integer; i <= limit; i = i + integer {
		resSlice = append(resSlice, i)
	}
	return resSlice
}
