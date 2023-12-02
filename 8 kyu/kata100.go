package main

import (
	"fmt"
	"strings"
)

func main() {
	var word string
	_, err := fmt.Scan(&word)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(IsPalindrome(word))
}

func IsPalindrome(str string) bool {
	strUp := strings.ToUpper(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if strUp[i] != strUp[j] {
			return false
		}
	}
	return true
}
