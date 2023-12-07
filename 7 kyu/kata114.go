package main

import "fmt"

func main() {
	var length, width int
	_, err := fmt.Scan(&length, &width)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(BreakChocolate(length, width))
}
func BreakChocolate(n, m int) int {
	if n*m <= 1 {
		return 0
	}
	return (n * m) - 1
}
