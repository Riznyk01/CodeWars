package main

import "fmt"

func main() {
	var a, b, c int
	_, err := fmt.Scan(&a, &b, &c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(SequenceSum(a, b, c))
}
func SequenceSum(start, end, step int) int {
	sum := 0
	for i := start; i <= end; i += step {
		sum += i
	}
	return sum
}
