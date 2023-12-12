package main

import (
	"fmt"
)

func main() {
	var num uint32
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ClosestMultipleOf10(num))
}
func ClosestMultipleOf10(n uint32) uint32 {
	m := (n / 10) * 10
	if n-m > 4 {
		return m + 10
	}
	return m
}
