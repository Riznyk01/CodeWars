package main

import "fmt"

func main() {
	fmt.Println(BubblesortOnce([]int{2, 1}))
}

func BubblesortOnce(numbers []int) []int {
	res := make([]int, len(numbers))
	copy(res, numbers)
	for i := 0; i < len(res)-1; i++ {
		if res[i] > res[i+1] {
			res[i], res[i+1] = res[i+1], res[i]
		}
	}
	return res
}
