package main

import "fmt"

func main() {
	s := []int{5, 4, 3, 2, 1}
	fmt.Println(Smaller(s))
}
func Smaller(arr []int) []int {
	result := make([]int, len(arr))
	for j, n := range arr {
		cnt := 0
		for i := j + 1; i < len(arr); i++ {
			if n > arr[i] {
				cnt++
			}
		}
		result[j] = cnt
	}
	return result
}
