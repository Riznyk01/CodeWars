package main

import (
	"fmt"
)

func main() {
	fmt.Println(Derive(7, 8))
}

func Derive(coefficient, exponent int) string {
	result := fmt.Sprintf("%d%s%d", coefficient*exponent, "x^", exponent-1)
	return result
}
