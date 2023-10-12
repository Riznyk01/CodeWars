package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	str, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	numStrs := strings.Split(strings.ReplaceAll(strings.TrimSpace(str), " ", ""), ",")
	nums := make([]int, len(numStrs))

	for i, str := range numStrs {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Number conversion error:", err)
			return
		}
		nums[i] = num
	}
	fmt.Println("Array of numbers:", Invert)
}
func Invert(arr []int) []int {
	slice := make([]int, len(arr))
	for i, n := range arr {
		slice[i] = -n
	}
	return slice
}
