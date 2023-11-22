package main

import (
	"fmt"
	"strings"
)

func main() {

	var str, character string

	_, err := fmt.Scan(&str, &character)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(Contamination(str, character))
}

func Contamination(text, char string) string {
	return strings.Repeat(char, len(text))
}
