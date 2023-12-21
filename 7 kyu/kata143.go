package main

import "fmt"

func main() {
	fmt.Println(SumEvenFibonacci(1))
}
func SumEvenFibonacci(limit int) int {
	sum := 2
	x1, x2 := 1, 2
	for x2 <= limit {
		x1, x2 = x2, x1+x2
		if x2%2 == 0 {
			sum += x2
		}
	}
	return sum
}
