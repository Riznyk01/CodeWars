package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numStr string
	_, err := fmt.Scan(&numStr)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(StringToNumber(numStr))
}
func StringToNumber(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}
	return num
}
