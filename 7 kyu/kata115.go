package main

import "fmt"

func main() {
	num := 0
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(SumCubes(num))
}

func SumCubes(n int) int {
	x := (n * (n + 1)) / 2
	return x * x
}
