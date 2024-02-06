package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(RangeBitCount(2, 7))
}

func RangeBitCount(a, b int) int {
	cnt := 0
	for i := a; i <= b; i++ {
		for _, d := range strconv.FormatInt(int64(i), 2) {
			if d == '1' {
				cnt++
			}
		}
	}
	return cnt
}
