package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type MyString string

func main() {
	phrase, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	phrase = strings.TrimSpace(phrase)
	fmt.Println(MyString(phrase).IsUpperCase())
}

func (s MyString) IsUpperCase() bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && r != 32 {
			return false
		}
	}
	return true
}
