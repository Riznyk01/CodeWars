package main

import (
	"fmt"
)

func main() {
	var h, d float64
	_, err := fmt.Scan(&h, &d)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(combat(h, d))
}
func combat(health, damage float64) float64 {
	if health-damage > 0 {
		return health - damage
	}
	return 0
}
