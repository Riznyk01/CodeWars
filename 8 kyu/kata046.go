package main

import (
	"fmt"
)

func main() {
	var t float64
	_, err := fmt.Scan(&t)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(Litres(t))
}
func Litres(time float64) int {
	return int(time * 0.5)
}
