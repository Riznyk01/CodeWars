package main

import "fmt"

func main() {
	var b, f int
	_, err := fmt.Scan(&b, &f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(CheckForFactor(b, f))
}
func CheckForFactor(base int, factor int) bool {
	return base%factor == 0
}
