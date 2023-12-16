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
		fmt.Println(err)
	}
	fmt.Println(RakeGarden(str))
}
func RakeGarden(garden string) string {
	arr := make([]string, len(strings.Split(garden, " ")))
	for i, w := range strings.Split(garden, " ") {
		if w != "rock" && w != "gravel" {
			arr[i] = "gravel"
		} else {
			arr[i] = w
		}
	}
	return strings.Join(arr, " ")
}
