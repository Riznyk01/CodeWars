package main

import "fmt"

func main() {
	var num int

	_, err := fmt.Scan(&num)

	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(Summation(num))
}

func Summation(n int) int {
	out := 0
	for i := 1; i <= n; i++ {
		out += i
	}
	return out
}
