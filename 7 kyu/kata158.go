package main

import "fmt"

func main() {
	fmt.Println(InAscOrder([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}))
}
func InAscOrder(numbers []int) bool {
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < numbers[i-1] {
			return false
		}
	}
	return true
}
