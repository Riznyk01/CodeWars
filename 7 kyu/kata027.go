package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str string
	var n, m int

	str, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}

	_, err = fmt.Scan(&n, &m)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}

	fmt.Println(Solve(str, n, m))
}
func Solve(s string, a, b int) string {
	runes := []rune(s)
	if b > len(runes) {
		b = len(runes) - 1
	}
	for a < b {
		runes[a], runes[b] = runes[b], runes[a]
		a++
		b--
	}
	return string(runes)
}
