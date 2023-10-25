package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	b, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	b = strings.TrimRight(b, "\n")
	var d string
	d, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	d = strings.TrimRight(b, "\n")

	fmt.Println(Feast(b, d))
}

func Feast(beast string, dish string) bool {
	if beast[0] == dish[0] && beast[(len(beast))-1] == dish[(len(dish))-1] {
		return true
	}
	return false
}
