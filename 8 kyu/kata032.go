package main

import (
	"fmt"
	"strconv"
)

func main() {
	var binStr string
	_, err := fmt.Scan(&binStr)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(BinToDec(binStr))
}
func BinToDec(bin string) int {
	decimal, err := strconv.ParseInt(bin, 2, 0)
	if err != nil {
		fmt.Println("Error during conversion:", err)
		return 0
	}
	return int(decimal)
}
