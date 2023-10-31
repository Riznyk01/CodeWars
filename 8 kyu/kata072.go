package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(NearestSq(num))
}
func NearestSq(n int) int {
	s := int(math.Round(math.Sqrt(float64(n))))
	return s * s
}
