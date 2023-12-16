package main

import "fmt"

func main() {
	fmt.Println(Strong(1))
}

func Strong(n int) string {
	if toFactorial(n) == n {
		return "STRONG!!!!"
	}
	return "Not Strong !!"
}
func toFactorial(num int) (res int) {
	for num > 0 {
		res += factorial(num % 10)
		num = num / 10
	}
	return res
}
func factorial(m int) int {
	f := 1
	if m == 0 {
		return 1
	}
	for m > 0 {
		f *= m
		m--
	}
	return f
}
