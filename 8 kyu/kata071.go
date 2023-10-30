package main

import "fmt"

func main() {
	var month int
	_, err := fmt.Scan(&month)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(QuarterOf(month))
}
func QuarterOf(month int) int {
	return (month + 2) / 3
}
