package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(DeadAntCount("...ant...ant..nat.ant.t..ant...ant..ant..ant.anant..t"))
}
func DeadAntCount(ants string) int {
	replacer := strings.NewReplacer("ant", "")
	parts := replacer.Replace(ants)
	return findMax(countParts(parts))
}
func countParts(input string) map[rune]int {
	numOfParts := make(map[rune]int)
	for _, r := range input {
		if r == 'a' || r == 'n' || r == 't' {
			numOfParts[r]++
		}
	}
	return numOfParts
}
func findMax(totalParts map[rune]int) (maxCount int) {
	for _, n := range totalParts {
		if n > maxCount {
			maxCount = n
		}
	}
	return maxCount
}
