package main

import (
	"fmt"
)

func main() {
	var y int
	_, err := fmt.Scan(&y)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(century(y))
}

func century(year int) int {
	if year%100 == 0 {
		return year / 100
	} else {
		return year/100 + 1
	}
}
