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
	fmt.Println(toWeirdCase(str))

}
func toWeirdCase(s string) string {
	var combinedString string
	var modifiedWords []string
	arrStr := strings.Fields(s)

	for _, word := range arrStr {
		modifiedWord := []rune(word)

		if len(modifiedWord) == 1 {
			modifiedWord[0] = unicode.ToUpper(modifiedWord[0])
		} else {
			for i := 0; i < len(modifiedWord); i += 2 {
				modifiedWord[i] = unicode.ToUpper(modifiedWord[i])
			}
			for i := 1; i < len(modifiedWord); i += 2 {
				modifiedWord[i] = unicode.ToLower(modifiedWord[i])
			}

		}

		modifiedWords = append(modifiedWords, string(modifiedWord))
		combinedString = strings.Join(modifiedWords, " ")
	}
	return combinedString
}
