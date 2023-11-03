package main

import (
	"fmt"
	"strings"
)

func main() {
	var nums string
	_, err := fmt.Scan(&nums)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(FakeBin(nums))
}
func FakeBin(x string) string {
	fc := func(r rune) rune {
		if r < '5' {
			return '0'
		}
		return '1'
	}
	return strings.Map(fc, x)
}
