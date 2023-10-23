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
	fmt.Println(NoSpace(str))
}
func NoSpace(word string) string {
	return strings.Replace(word, " ", "", -1)
}
