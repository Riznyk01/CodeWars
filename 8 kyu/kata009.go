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
	fmt.Println(StringToArray(str))
}

func StringToArray(str string) []string {
	out := strings.Split(str, " ")
	return out
}
