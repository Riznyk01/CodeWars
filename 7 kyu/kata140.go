package main

import "fmt"

func main() {
	fmt.Println(Factorial(0))
}
func Factorial(n int) int {
	res := 1
	for n > 1 {
		res *= n
		n--
	}
	return res
}
