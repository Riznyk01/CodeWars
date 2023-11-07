package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter an array:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	inputSlice := strings.Split(input, ",")

	fmt.Println(Well(inputSlice))
}
func Well(x []string) string {
	goodCount := 0
	for _, idea := range x {
		if idea == "good" {
			goodCount++
		}
	}

	if goodCount == 1 || goodCount == 2 {
		return "Publish!"
	} else if goodCount > 2 {
		return "I smell a series!"
	}
	return "Fail!"
}
