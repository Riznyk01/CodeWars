package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var s string

	_, err := fmt.Scan(&n, &s)

	if err != nil {
		fmt.Println("Input error:", err)
		return
	}

	fmt.Println(RepeatStr(n, s))
}

func RepeatStr(repetitions int, value string) string {
	return strings.Repeat(value, repetitions)
}
