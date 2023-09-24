package main

import (
	"fmt"
	"strconv"
)

func main() {
	var inp int
	fmt.Scan(&inp)
	fmt.Println(Digitize(inp))
}
func Digitize(n int) []int {
	cnt := countDigits(n)
	numbers := make([]int, cnt, cnt)
	x := 1
	if n != 0 {
		for i := 0; i < cnt; i++ {
			numbers[i] = (n / x) % 10
			x *= 10
		}
		return numbers
	} else {
		numbers[0] = 0
		return numbers
	}
}
func countDigits(number int) int {
	numberStr := strconv.Itoa(number)
	return len(numberStr)
}
