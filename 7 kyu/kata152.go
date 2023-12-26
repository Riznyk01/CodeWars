package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(solve("code"))
}
func solve(str string) string {
	l, u := 0, 0
	for _, r := range str {
		if r >= 65 && r <= 90 {
			u++
		} else if r >= 97 && r <= 122 {
			l++
		}
	}
	if l == u || l > u {
		return strings.ToLower(str)
	}
	return strings.ToUpper(str)
}
