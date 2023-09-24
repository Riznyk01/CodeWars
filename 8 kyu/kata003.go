package main

import "fmt"

func main() {
	var firstNum, secondNum int

	_, err := fmt.Scan(&firstNum, &secondNum)

	if err != nil {
		fmt.Println("Input error:", err)
		return
	}

	fmt.Println(Between(firstNum, secondNum))
}
func Between(a, b int) []int {
	quantity := b - a + 1
	numbers := make([]int, quantity)
	for i := 0; i <= quantity-1; i++ {
		numbers[i] = a
		a += 1
	}
	return numbers
}
