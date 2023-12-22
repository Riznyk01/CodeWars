package main

import "fmt"

func main() {
	fmt.Println(WordsToMarks("love"))
}
func WordsToMarks(s string) int {
	sum := 0
	for _, r := range s {
		sum += int(r) - 96
	}
	return sum
}
