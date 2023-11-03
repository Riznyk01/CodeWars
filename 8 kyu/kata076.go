package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter the number of numbers in the array:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	count, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Input error:")
		return
	}

	numbers := make([]int, count)

	for i := 0; i < count; i++ {
		fmt.Printf("Enter the number %d: ", i+1)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error entering number. Try again.")
			i--
			continue
		}
		numbers[i] = num
	}
	fmt.Println(CountPositivesSumNegatives(numbers))
}
func CountPositivesSumNegatives(numbers []int) []int {
	res := make([]int, 2)
	for _, n := range numbers {
		if n > 0 {
			res[0]++
		} else if n < 0 {
			res[1] += n
		}
	}
	if res[0] > 0 || res[1] > 0 {
		return res
	}
	return []int{}
}
