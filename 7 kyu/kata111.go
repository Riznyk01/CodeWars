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
	fmt.Println(FindShort(str))
}
func FindShort(s string) int {
	words := strings.Split(s, " ")
	minLen := len(words[0])
	for _, w := range words[1:] {
		if len(w) < minLen {
			minLen = len(w)
		}
	}
	return minLen
}
