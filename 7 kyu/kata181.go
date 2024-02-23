package main

import (
	"fmt"
)

func main() {
	fmt.Println(ArrowArea(4, 2))
}

func ArrowArea(a, b int) float64 {
	return 0.25 * float64(a) * float64(b)
}
