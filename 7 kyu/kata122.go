package main

import (
	"fmt"
	"strings"
)

func main() {
	var start, end string
	_, err := fmt.Scan(&start, &end)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(solution(start, end))
}
func solution(str, ending string) bool {
	return strings.HasSuffix(str, ending)
}
