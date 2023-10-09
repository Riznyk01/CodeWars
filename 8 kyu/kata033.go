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
	fmt.Println(PositiveSum(nums))
}
func PositiveSum(numbers []int) int {
	sumPosit := 0
	for _, n := range numbers {
		if n > 0 {
			sumPosit += n
		}
	}
	return sumPosit
}
