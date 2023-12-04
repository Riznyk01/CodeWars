package main

import (
	"fmt"
)

func main() {
	var num, div int
	_, err := fmt.Scan(&num, &div)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(Divisions(num, div))
}
func Divisions(n, divisor int) int {
	cnt := 0
	for {
		if n/divisor < 1 {
			return cnt
		}
		n /= divisor
		cnt++
	}
}
