package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(FreqSeq("hello world", "-"))
}

func FreqSeq(str string, sep string) string {
	runes := make(map[rune]int)
	var res string
	for _, r := range str {
		runes[r]++
	}
	for _, r := range str {
		res += strconv.Itoa(runes[r]) + sep
	}
	return res[:len(res)-1]
}
