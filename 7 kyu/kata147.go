package main

import "fmt"

func main() {
	fmt.Println(RowSumOddNumbers(5))
}
func RowSumOddNumbers(n int) (sum int) {
	elem := -1
	for j := 1; j <= n; j++ {
		for i := 0; i < j; i++ {
			elem += 2
			if j == n {
				sum += elem
			}
		}
	}
	return sum
}
