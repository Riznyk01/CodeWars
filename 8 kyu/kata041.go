package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var str string
	str, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(ReplaceDots(str))
}
func ReplaceDots(str string) string {
	return strings.Replace(str, ".", "-", -1)
}
