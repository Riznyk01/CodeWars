package main

import (
	"bufio"
	"fmt"
	"math"
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
	fmt.Println(SquareOrSquareRoot(numbersSlice))
}
func SquareOrSquareRoot(arr []int) []int {
	outArr := make([]int, len(arr))
	for i, n := range arr {
		nSq := math.Sqrt(float64(n))
		if math.Floor(nSq) == nSq {
			outArr[i] = int(nSq)
		} else {
			outArr[i] = n * n
		}
	}
	return outArr
}
