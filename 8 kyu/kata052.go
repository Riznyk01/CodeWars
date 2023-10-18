package main

import (
	"fmt"
)

func main() {
	var val bool
	_, err := fmt.Scan(&val)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(BoolToWord(val))
}
func BoolToWord(word bool) string {
	if word {
		return "Yes"
	}
	return "No"
}
