package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var body, tailStr string
	var tail rune

	reader := bufio.NewReader(os.Stdin)
	body, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}

	tailStr, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}

	tail = []rune(tailStr)[0]
	fmt.Println(CorrectTail(body, tail))
}
func CorrectTail(body string, tail rune) bool {
	return strings.HasSuffix(body, string(tail))
}
