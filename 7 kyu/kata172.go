package main

import (
	"fmt"
)

func main() {
	fmt.Println(Number([][2]int{{3, 0}, {9, 1}, {4, 10}, {12, 2}, {6, 1}, {7, 10}}))
}
func Number(busStops [][2]int) (res int) {
	for _, stop := range busStops {
		res += stop[0] - stop[1]
	}
	return res
}
