package main

import "fmt"

func main() {
	fmt.Println(ReverseList([]int{1, 2, 3, 4}))
}
func ReverseList(lst []int) []int {
	reversedArr := make([]int, len(lst))
	for i, n := range lst {
		reversedArr[len(lst)-i-1] = n
	}
	return reversedArr
}
