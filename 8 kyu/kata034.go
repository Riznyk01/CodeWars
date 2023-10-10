package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(NumberToString(num))
}
func NumberToString(n int) string {
	return strconv.Itoa(n)
}
