package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(Solve("Codewars@codewars123.com"))
}
func Solve(s string) []int {
	outArr := make([]int, 4)
	for _, r := range s {
		if unicode.IsUpper(r) {
			outArr[0]++
		} else if unicode.IsLower(r) {
			outArr[1]++
		} else if unicode.IsNumber(r) {
			outArr[2]++
		} else if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			outArr[3]++
		}
	}
	return outArr
}
