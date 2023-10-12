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
	fmt.Println(OtherAngle(a, b))
}
func OtherAngle(a int, b int) int {
	return 180 - a - b
}
