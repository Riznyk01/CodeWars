package main

import (
	"fmt"
	"strings"
)

func main() {
	var inpStr string
	_, err := fmt.Scan(&inpStr)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(MakeUpperCase(inpStr))
}
func MakeUpperCase(str string) string {
	return strings.ToUpper(str)
}
