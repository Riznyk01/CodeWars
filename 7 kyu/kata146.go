package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(Capitalize("abcdef", []int{1, 2, 5}))
}
func Capitalize(st string, arr []int) string {
	str := []rune(st)
	for _, ind := range arr {
		if ind < len(str) {
			str[ind] = unicode.ToUpper(str[ind])
		}

	}
	return string(str)
}
