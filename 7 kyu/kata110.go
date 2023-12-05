package main

import "fmt"

func main() {
	var blue int
	_, err := fmt.Scan(&blue)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(CountRedBeads(blue))
}

func CountRedBeads(n int) int {
	if n > 1 {
		return (n - 1) * 2
	}
	return 0
}
