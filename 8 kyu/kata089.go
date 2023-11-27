package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var num int
	str, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	numsStr := strings.Split(strings.ReplaceAll(strings.TrimSpace(str), " ", ""), ",")
	nums := make([]int, len(numsStr))

	for i, str := range numsStr {
		n, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Number conversion error:", err)
			return
		}
		nums[i] = n
	}
	_, err = fmt.Scan(&num)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(EachCons(nums, num))
}
func EachCons(arr []int, n int) [][]int {
	var result [][]int
	for i := 0; i <= len(arr)-n; i++ {
		result = append(result, arr[i:i+n])
	}
	return result
}
