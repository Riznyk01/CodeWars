package main

import "fmt"

func main() {
	arr := []int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5}
	fmt.Println(FindOdd(arr))
}

func FindOdd(seq []int) int {
	cnts := make(map[int]int)
	for _, num := range seq {
		cnts[num]++
	}
	for num, cnt := range cnts {
		if cnt%2 != 0 {
			return num
		}
	}
	return 0
}
