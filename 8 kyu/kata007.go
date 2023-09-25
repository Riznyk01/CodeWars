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
	fmt.Println(ToAlternatingCase(str))
}

func ToAlternatingCase(str string) string {
	filter := func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		} else if r >= 'a' && r <= 'z' {
			return unicode.ToUpper(r)
		} else if r >= 'A' && r <= 'Z' {
			return unicode.ToLower(r)
		} else {
			return r
		}
	}
	return strings.Map(filter, str)
}
