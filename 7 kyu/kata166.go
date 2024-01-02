package main

import "fmt"

func main() {
	var addOne = Add(1)
	fmt.Println(addOne(3)) // 4
}
func Add(n int) func(m int) int {
	fc := func(m int) int {
		return n + m
	}
	return fc
}
