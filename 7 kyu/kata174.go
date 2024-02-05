package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinMax([]int{1, 2, 3, 4, 5}))
}
func MinMax(arr []int) [2]int {
	sort.Ints(arr)
	return [2]int{arr[0], arr[len(arr)-1]}
}

//func MinMax(arr []int) [2]int {
//	a := [2]int{arr[0], arr[0]}
//	for _, num := range arr {
//		if num < a[0] {
//			a[0] = num
//		} else if num > a[1] {
//			a[1] = num
//		}
//	}
//	return a
//}
