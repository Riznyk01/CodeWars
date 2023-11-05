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
	fmt.Println(Grow(numbersSlice))
}
func Grow(arr []int) int {
	m := 1
	for _, n := range arr {
		m *= n
	}
	return m
}
