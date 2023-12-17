package main

import "fmt"

func main() {
	arr := []int{1, 1, 1, 2, 2, 2}
	fmt.Println(Dominator(arr))
}
func Dominator(a []int) (d int) {
	nums := make(map[int]int)
	for _, n := range a {
		nums[n]++
	}
	for k, v := range nums {
		if v > len(a)/2 {
			return k
		}
	}
	return -1
}
