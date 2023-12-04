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
	fmt.Println(Capitalize(strings.TrimSpace(str)))
}

func Capitalize(st string) []string {
	f, s := []rune(st), []rune(st)
	for i, r := range []rune(st) {
		if i%2 == 0 {
			f[i] = unicode.ToUpper(r)
		} else {
			s[i] = unicode.ToUpper(r)
		}
	}
	return []string{string(f), string(s)}
}
