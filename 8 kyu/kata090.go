package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//var firstStr, secondStr string
	firstStr, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	secondStr, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	firstStr = strings.TrimSpace(firstStr)
	secondStr = strings.TrimSpace(secondStr)
	fmt.Println(Sol(firstStr, secondStr))
}
func Sol(a, b string) string {
	var short, long string
	if len(a) < len(b) {
		short = a
		long = b
	} else {
		short = b
		long = a
	}
	return fmt.Sprintf("%s%s%s", short, long, short)
}
