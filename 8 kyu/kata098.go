package main

import "fmt"

func main() {
	fmt.Println(NthEven(100))
}
func NthEven(n int) int {
	return (n - 1) * 2
}
