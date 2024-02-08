package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(DontGiveMeFive(1, 9))
}
func DontGiveMeFive(start int, end int) int {
	cnt := 0
	for i := start; i <= end; i++ {
		if !consistDigit(i) {
			cnt++
		}
	}
	return cnt
}
func consistDigit(n int) bool {
	return strings.Contains(strconv.Itoa(n), strconv.Itoa(5))
}
