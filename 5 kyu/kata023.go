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
	fmt.Println(MoveZeros(arr))
}

func MoveZeros(integers []int) []int {
	outArr := make([]int, len(integers))
	ind := 0
	for _, num := range integers {
		if num != 0 {
			outArr[ind] = num
			ind++
		}
	}
	return outArr
}
