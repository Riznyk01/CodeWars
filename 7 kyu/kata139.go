package main

import "fmt"

func main() {
	fmt.Println(Angle(6))
}
func Angle(n int) int {
	return 180 * (n - 2)
}
