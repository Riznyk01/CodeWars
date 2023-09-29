package main

import "fmt"

func main() {
	var num int
	_, err := fmt.Scan(&num)

	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(DigitalRoot(num))
}
func DigitalRoot(n int) int {
	return (n-1)%9 + 1
}
