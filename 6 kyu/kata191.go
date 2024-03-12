package main

import "fmt"

func main() {
	fmt.Println(NumberOfPairs([]string{"gray", "black", "purple", "purple", "gray", "black"}))
}
func NumberOfPairs(gloves []string) (pairs int) {
	colorCount := make(map[string]int)
	for _, g := range gloves {
		colorCount[g]++
	}

	for _, v := range colorCount {
		pairs += v / 2
	}

	return pairs
}
