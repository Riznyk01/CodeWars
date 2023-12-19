package main

import "fmt"

func main() {
	fmt.Println(TwoOldestAges([]int{69, 67, 59, 45, 59, 37, 65, 39, 85, 21}))
}
func TwoOldestAges(ages []int) [2]int {
	res := [2]int{}
	idx := 0
	for i, n := range ages {
		if n > res[1] {
			res[1] = n
			idx = i
		}
	}
	for i, num := range ages {
		if i != idx {
			if num > res[0] {
				res[0] = num
			}
		}
	}
	return res
}
