package main

import "fmt"

func main() {
	fmt.Println(Arithmetic(8, 2, "add"))
}
func Arithmetic(a int, b int, operator string) int {
	switch operator {
	case "add":
		return a + b
	case "subtract":
		return a - b
	case "multiply":
		return a * b
	case "divide":
		return a / b
	}
	return 0
}
