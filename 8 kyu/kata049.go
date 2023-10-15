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
	fmt.Println(CalculateYears(y))
}
func CalculateYears(years int) (result [3]int) {
	if years == 1 || years == 2 {
		result[1], result[2] = 15+9*(years-1), 15+9*(years-1)
	} else {
		result[1], result[2] = 24+4*(years-2), 24+5*(years-2)
	}
	result[0] = years
	return result
}
