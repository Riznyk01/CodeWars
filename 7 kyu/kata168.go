package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(HighAndLow("1 9 3 4 -5"))
}
func HighAndLow(in string) string {
	arr := strings.Split(in, " ")
	maxNum, _ := strconv.Atoi(arr[0])
	minNum, _ := strconv.Atoi(arr[0])
	for _, numStr := range arr {
		n, _ := strconv.Atoi(numStr)
		if n < minNum {
			minNum = n
		} else if n > maxNum {
			maxNum = n
		}
	}
	return strconv.Itoa(maxNum) + " " + strconv.Itoa(minNum)
}
