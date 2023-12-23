package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(EquableTriangle(5, 12, 13))
}
func EquableTriangle(a, b, c int) bool {
	if a+b > c && a+c > b && b+c > a {
		semiperimeter := float64(a+b+c) / 2
		area := math.Sqrt(semiperimeter * (semiperimeter - float64(a)) * (semiperimeter - float64(b)) * (semiperimeter - float64(c)))
		if 2*semiperimeter == area {
			return true
		}
	}
	return false
}
