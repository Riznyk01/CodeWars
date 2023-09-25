package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Enter the size of the array: ")

	_, err := fmt.Scan(&n)

	if err != nil {
		fmt.Println("Input error:", err)
		return
	}

	if n < 1 {
		fmt.Println("The size of the array must be greater than 0")
		return
	}

	arr := make([]int, n)

	fmt.Println("Enter the elements of the array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Element %d: ", i+1)
		_, err := fmt.Scan(&arr[i])

		if err != nil {
			fmt.Println("Input error:", err)
			return
		}
	}
	fmt.Println("Square(n) Sum:", SquareSum(arr))
}

func SquareSum(numbers []int) int {
	out := 0
	for i := 0; i < len(numbers); i++ {
		out += numbers[i] * numbers[i]
	}
	return out
}
