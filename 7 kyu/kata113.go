package main

import "fmt"

func main() {
	var firstNum, secondNum int
	_, err := fmt.Scan(&firstNum, &secondNum)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(GetSum(firstNum, secondNum))
}
func GetSum(a, b int) int {
	sum := 0
	if a > b {
		a, b = b, a
	}
	for i := a; i <= b; i++ {
		sum += i
	}
	return sum
}
