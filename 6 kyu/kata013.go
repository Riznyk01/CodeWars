package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	str, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(CamelCase(str))
}

func CamelCase(s string) string {
	var combinedString string
	var modifiedWords []string
	arrStr := strings.Fields(s)

	for _, word := range arrStr {
		modifiedWord := strings.ToUpper(word[:1]) + word[1:]
		modifiedWords = append(modifiedWords, modifiedWord)
		combinedString = strings.Join(modifiedWords, "")
	}
	return combinedString
}
