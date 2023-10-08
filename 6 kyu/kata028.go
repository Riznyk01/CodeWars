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
	fmt.Println(FindEvenIndex(toArr(str)))
}
func toArr(str string) []int {
	numStrs := strings.Split(strings.ReplaceAll(strings.TrimSpace(str), " ", ""), ",")
	nums := make([]int, len(numStrs))
	for i, str := range numStrs {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Number conversion error:", err)
			return []int{}
		}
		nums[i] = num
	}
	return nums
}
func FindEvenIndex(arr []int) int {
	for n := 0; n < len(arr); n++ {
		var sumLeft, sumRight int
		sumR := func(n int) int {
			for i := n + 1; i < len(arr); i++ {
				sumRight = sumRight + arr[i]
			}
			return sumRight
		}
		sumL := func(n int) int {
			for i := 0; i < n; i++ {
				sumLeft = sumLeft + arr[i]
			}
			return sumLeft
		}
		if n == 0 {
			if sumR(n) == 0 {
				return n
			}
		} else if n == len(arr) {
			if sumL(n) == 0 {
				return n
			}
		} else {
			if sumR(n) == sumL(n) {
				return n
			}
		}
	}
	return -1
}
