package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(WrapPresent(17, 32, 11))
}
func WrapPresent(height, width, length int) int {
	arr := []int{height, width, length}
	sort.Ints(arr)
	return 2*arr[2] + 2*arr[1] + 4*arr[0] + 20
}
