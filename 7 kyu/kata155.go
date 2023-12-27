package main

import "fmt"

func main() {
	fmt.Println(Collatz(20))
}
func Collatz(n int) int {
	cnt := 1
	for n > 1 {
		cnt++
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	return cnt
}
