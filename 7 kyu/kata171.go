package main

import "fmt"

func main() {
	fmt.Println(RoundToNext5(2))
}
func RoundToNext5(n int) int {
	for n%5 != 0 {
		n++
	}
	return n
}
