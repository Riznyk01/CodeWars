package main

import "fmt"

func main() {
	fmt.Println(Fib(-5))
}
func Fib(n int) int {
	var negative bool
	cnt := 1
	if n == 1 {
		return 1
	}
	if n < 0 {
		n = -n
		negative = true
	}
	for x1, x2 := 0, 1; cnt <= n; x1, x2 = x2, x1+x2 {
		if cnt++; cnt == n {
			if negative {
				return -(x1 + x2)
			}
			return x1 + x2
		}
	}
	return 0
}
