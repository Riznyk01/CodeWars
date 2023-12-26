package main

import (
	"fmt"
)

func main() {
	fmt.Println(Gcd(30, 12))
}

func Gcd(x, y uint32) uint32 {
	for x != 0 && y != 0 {
		if x > y {
			x = x % y
		} else {
			y = y % x
		}
	}
	return x + y
}
