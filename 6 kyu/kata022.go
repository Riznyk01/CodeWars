package main

import (
	"fmt"
	"strconv"
)

func main() {

	var num uint
	_, err := fmt.Scan(&num)

	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(CountBits(num))
}

func CountBits(decNum uint) int {
	binaryString := strconv.FormatUint(uint64(decNum), 2)
	cnt := 0
	for _, r := range binaryString {
		if r == '1' {
			cnt++
		}
	}
	return cnt
}
