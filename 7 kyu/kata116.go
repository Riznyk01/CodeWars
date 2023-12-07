package main

import (
	"fmt"
	"sort"
)

func main() {
	numSlice := []int{1, 2, 10, 50, 5}
	fmt.Println(SortNumbers(numSlice))
}

func SortNumbers(numbers []int) []int {
	sort.Ints(numbers)
	return numbers
}
