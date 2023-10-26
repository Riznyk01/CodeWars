package main

import (
	"fmt"
)

func main() {
	var flowerOne, flowerTwo int
	_, err := fmt.Scan(&flowerOne, &flowerTwo)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(LoveFunc(flowerOne, flowerTwo))
}
func LoveFunc(flower1, flower2 int) bool {
	if flower1%2 != flower2%2 {
		return true
	}
	return false
}
