package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	str, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(ReverseLetters(str))
}
func ReverseLetters(s string) string {
	filtered := []rune(strings.Map(filter, s))
	result := make([]rune, len(filtered))
	l := len(filtered) - 1
	for i := l; i >= 0; i-- {
		result[l-i] = filtered[i]
	}
	return string(result)
}
func filter(r rune) rune {
	if unicode.IsLetter(r) {
		return r
	} else {
		return -1
	}
}
