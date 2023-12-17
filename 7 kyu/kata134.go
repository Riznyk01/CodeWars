package main

import "fmt"

func main() {
	fmt.Println(GrowingPlant(10, 7, 652))
}
func GrowingPlant(upSpeed, downSpeed, desiredHeight int) (days int) {
	g := 0
	for g < desiredHeight {
		days++
		g = g + upSpeed
		if g >= desiredHeight {
			return days
		}
		g = g - downSpeed
	}
	return days
}
