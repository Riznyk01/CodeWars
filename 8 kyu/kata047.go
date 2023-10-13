package main

import (
	"fmt"
)

func main() {
	var word string

	_, err := fmt.Scan(&word)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(RemoveChar(word))
}
func RemoveChar(word string) string {
	return word[1 : len(word)-1]
}
