package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(countSheep(n))
}
func countSheep(num int) string {
	var str string
	for i := 1; i <= num; i++ {
		str = str + strconv.Itoa(i) + " sheep..."
	}
	return str
}
