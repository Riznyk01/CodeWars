package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter an array of numbers separated by commas:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, " ", "")

	strSlice := strings.Split(input, ",")
	numbersSlice := make([]int, len(strSlice))

	for i := 0; i < len(strSlice); i++ {
		num, err := strconv.Atoi(strSlice[i])
		if err != nil {
			fmt.Println(err)
		}
		numbersSlice[i] = num
	}
	fmt.Println(multipleOfIndex(numbersSlice))
}
func multipleOfIndex(ints []int) []int {
	outSlice := make([]int, 0)
	for i, n := range ints {
		if i != 0 && n%i == 0 {
			outSlice = append(outSlice, ints[i])
		}
	}
	return outSlice
}
